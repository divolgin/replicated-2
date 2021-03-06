package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/replicatedhq/replicated/cli/print"
	"github.com/replicatedhq/replicated/pkg/platformclient"
	"github.com/spf13/cobra"
)

// releaseInspectCmd represents the inspect command
var releaseInspectCmd = &cobra.Command{
	Use:   "inspect SEQUENCE",
	Short: "Print the YAML config for a release",
	Long:  "Print the YAML config for a release",
}

func init() {
	releaseCmd.AddCommand(releaseInspectCmd)
}

func (r *runners) releaseInspect(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("release sequence is required")
	}
	seq, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return fmt.Errorf("Failed to parse sequence argument %s", args[0])
	}

	release, err := r.platformAPI.GetRelease(r.appID, seq)
	if err != nil {
		if err == platformclient.ErrNotFound {
			return fmt.Errorf("No such release %d", seq)
		}
		return err
	}

	return print.Release(r.w, release)
}
