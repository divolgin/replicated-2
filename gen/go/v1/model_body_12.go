/*
 * Vendor API V1
 *
 * Apps documentation
 *
 * API version: 1.0.0
 * Contact: info@replicated.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Body12 struct {
	// If activation is required this is the email the code will be sent to.
	ActivationEmail       string `json:"activation_email"`
	AirgapDownloadEnabled bool   `json:"airgap_download_enabled"`
	// License Label name, ie name of customer who is using license.
	Assignee             string `json:"assignee"`
	AssistedSetupEnabled bool   `json:"assisted_setup_enabled"`
	// Channel id that the license will be associated with.
	ChannelId string `json:"channel_id"`
	// Date that the license will expire, can be null for no expiration or formated by year-month-day ex. 2016-02-02.
	ExpirationDate string `json:"expiration_date"`
	// Defines expiration policy for this license.  Values: ignore: replicated will take no action on a expired license noupdate-norestart: application updates will not be downloaded, and once the application is stopped, it will not be started again noupdate-stop: application updates will not be downloaded and the application will be stopped
	ExpirationPolicy string `json:"expiration_policy"`
	// Defines whether this license should use the external support bundle generator.
	ExternalSupportBundle bool `json:"external_support_bundle,omitempty"`
	// An array of field values for custom fields of a given app
	FieldValues []LicenseFieldValue `json:"field_values"`
	// A license can be optionally locked to a specific release
	IsAppVersionLocked bool `json:"is_app_version_locked,omitempty"`
	// LicenseType can be set to \"dev\", \"trial\", or \"prod\"
	LicenseType string `json:"license_type"`
	// If app version is locked, this is the version to lock it to (sequence)
	LockedAppVersion int64 `json:"locked_app_version,omitempty"`
	// If this license requires activation set to true, make sure to set activation email as well.
	RequireActivation bool `json:"require_activation"`
	// If set to automatic will auto update remote license installation with every release. If set to manual will update only when on-premise admin clicks the install update button.
	UpdatePolicy string `json:"update_policy"`
	// Defines whether this license should use support bundle specs from console.
	UseConsoleSupportSpec bool `json:"use_console_support_spec,omitempty"`
}
