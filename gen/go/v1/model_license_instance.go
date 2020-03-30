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

import (
	"time"
)

type LicenseInstance struct {
	Active                    bool                            `json:"Active,omitempty"`
	AppStatus                 string                          `json:"AppStatus,omitempty"`
	AssistSessionId           string                          `json:"AssistSessionId,omitempty"`
	Cloud                     string                          `json:"Cloud,omitempty"`
	Created                   time.Time                       `json:"Created,omitempty"`
	InstanceId                string                          `json:"InstanceId,omitempty"`
	LastActive                time.Time                       `json:"LastActive,omitempty"`
	LicenseId                 string                          `json:"LicenseId,omitempty"`
	ReplicatedAgentVersion    string                          `json:"ReplicatedAgentVersion,omitempty"`
	ReplicatedOperatorVersion string                          `json:"ReplicatedOperatorVersion,omitempty"`
	ReplicatedSyncTime        time.Time                       `json:"ReplicatedSyncTime,omitempty"`
	ReplicatedUiVersion       string                          `json:"ReplicatedUiVersion,omitempty"`
	ReplicatedUpdaterVersion  string                          `json:"ReplicatedUpdaterVersion,omitempty"`
	ReplicatedVersion         string                          `json:"ReplicatedVersion,omitempty"`
	VersionHistory            []LicenseInstanceVersionHistory `json:"VersionHistory,omitempty"`
}
