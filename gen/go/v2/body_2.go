/*
 * Vendor API V2
 *
 * The Vendor API is available as an alternative to configuring your application via the Vendor website. This API will allow you as the Vendor to manage your apps and app configuration, manage customer licenses, view your private container images.  The API can be accessed at the URL https://api.replicated.com/vendor/v2.  == Authentication  Authentication is achieved via a token that can be generated via the Vendor website. You can create a new token by clicking the New API Token button in the API Tokens table. You must give this token a nickname before saving.  After you have generated your token, you will need to send this token in the Authorization header of your requests to the Vendor API.  The following Error response codes can be expected:  [options=\"header\", cols=\".^1,.^3,.^10\"] |=== |Code|Response|Description |400|Bad Request|We were unable to parse your request. |401|Unauthorized|Confirm that your token or username/password is valid. |403|Forbidden|The server is refusing to alloo you to complete this request. |404|Not found|The requested resource was not found. |409|Conflict|The action would result in a conflict and is being rejected. |500|Internal Server Error|Something unexpected happened. |===  You download JSON schemas for this API in our github repo.  Check out our Developer Help Center for sample recipes on using our Vendor API
 *
 * OpenAPI spec version: 1.0.0
 * Contact: info@replicated.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type Body2 struct {

	// If activation is required this is the email the code will be sent to.
	ActivationEmail string `json:"activation_email"`

	// License Label name, ie name of customer who is using license.
	Assignee string `json:"assignee"`

	// Channel id that the license will be associated with.
	ChannelId string `json:"channel_id"`

	Channels *LicenseChannels `json:"channels,omitempty"`

	ConsoleAuthOptions []string `json:"console_auth_options,omitempty"`

	EnabledFeatures map[string]interface{} `json:"enabled_features,omitempty"`

	// Date that the license will expire, can be null for no expiration or formated by year-month-day ex. 2016-02-02.
	ExpirationDate string `json:"expiration_date"`

	// Defines expiration policy for this license.  Values: ignore: replicated will take no action on a expired license noupdate-norestart: application updates will not be downloaded, and once the application is stopped, it will not be started again noupdate-stop: application updates will not be downloaded and the application will be stopped
	ExpirationPolicy string `json:"expiration_policy"`

	FieldValues *LicenseFieldValues `json:"field_values"`

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
}
