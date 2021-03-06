package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/url"

	"os"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/mitchellh/cli"
	"github.com/pkg/errors"
	"github.com/replicatedhq/ship/pkg/e2e"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	shipReleaseCommand.AddCommand(Cmd())
}

func Cmd() *cobra.Command {
	vip := viper.New()

	var stdoutLogger log.Logger
	stdoutLogger = log.NewLogfmtLogger(os.Stdout)
	stdoutLogger = log.With(stdoutLogger, "ts", log.DefaultTimestampUTC)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "API client for creating ship releases",
		Long: `
`,
		RunE: func(cmd *cobra.Command, args []string) error {

			releaser := &Releaser{
				viper:  vip,
				logger: stdoutLogger,
				ui: &cli.ColoredUi{
					OutputColor: cli.UiColorNone,
					ErrorColor:  cli.UiColorRed,
					WarnColor:   cli.UiColorYellow,
					InfoColor:   cli.UiColorGreen,
					Ui: &cli.BasicUi{
						Reader:      os.Stdin,
						Writer:      os.Stdout,
						ErrorWriter: os.Stderr,
					},
				},
			}

			err := releaser.Release(context.Background())

			if err != nil {
				return errors.Wrap(err, "promote release")
			}

			return nil
		},
	}

	cmd.Flags().String("vendor-token", "", "Token to use to communicate with https://g.replicated.com")
	cmd.Flags().String("graphql-api-address", "https://g.replicated.com/graphql", "upstream g. address")
	cmd.Flags().String("spec-file", "", "spec file to promote")
	cmd.Flags().String("channel-id", "", "channel id to promote")
	cmd.Flags().String("semver", "", "semver of the release")
	cmd.Flags().String("release-notes", "", "release notes")
	cmd.Flags().String("log-level", "off", "log level")

	vip.BindPFlags(cmd.Flags())
	vip.BindPFlags(cmd.PersistentFlags())
	vip.AutomaticEnv()
	vip.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	return cmd
}

type Releaser struct {
	viper  *viper.Viper
	logger log.Logger
	ui     cli.Ui
}

func (r *Releaser) getParams() (token, specContents, semver, channelID, gqlAddr string, err error) {
	token = r.viper.GetString("vendor-token")
	if token == "" {
		err = errors.New("param vendor-token is required")
		return
	}

	specFile := r.viper.GetString("spec-file")
	if specFile == "" {
		err = errors.New("param spec-file is required")
		return
	}

	specBytes, err := ioutil.ReadFile(specFile)
	if err != nil {
		err = errors.Wrapf(err, "read file %s", specFile)
		return
	}
	specContents = string(specBytes)

	semver = r.viper.GetString("semver")
	if semver == "" {
		err = errors.New("param semver is required")
		return
	}

	channelID = r.viper.GetString("channel-id")
	if channelID == "" {
		err = errors.New("param channel-id is required")
		return
	}

	gqlAddr = r.viper.GetString("graphql-api-address")
	return
}

func (r *Releaser) Release(ctx context.Context) error {
	token, specContents, semver, channelID, gqlAddr, err := r.getParams()
	if err != nil {
		return errors.Wrap(err, "load params")
	}

	gqlServer, err := url.Parse(gqlAddr)
	if err != nil {
		return errors.Wrapf(err, "parse graphql-api-address URL \"%s\"", gqlAddr)
	}
	client := &e2e.GraphQLClient{
		GQLServer: gqlServer,
		Token:     token,
		Logger:    r.logger,
	}

	data, err := client.PromoteRelease(
		string(specContents),
		channelID,
		semver,
		r.viper.GetString("release-notes"),
	)

	r.ui.Info(fmt.Sprintf("received data %+v", data))

	if err != nil {
		return errors.Wrapf(err, "promote release")
	}

	return nil
}
