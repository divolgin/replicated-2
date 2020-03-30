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

type LicensesV2 struct {
}
