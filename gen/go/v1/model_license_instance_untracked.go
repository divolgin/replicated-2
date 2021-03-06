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

type LicenseInstanceUntracked struct {
	CurrentSequence int64 `json:"CurrentSequence,omitempty"`
	CurrentVersion string `json:"CurrentVersion,omitempty"`
	LicenseId string `json:"LicenseId,omitempty"`
	ReplicatedAgentVersion string `json:"ReplicatedAgentVersion,omitempty"`
	ReplicatedOperatorVersion string `json:"ReplicatedOperatorVersion,omitempty"`
	ReplicatedSyncTime time.Time `json:"ReplicatedSyncTime,omitempty"`
	ReplicatedUiVersion string `json:"ReplicatedUiVersion,omitempty"`
	ReplicatedUpdaterVersion string `json:"ReplicatedUpdaterVersion,omitempty"`
	ReplicatedVersion string `json:"ReplicatedVersion,omitempty"`
	SyncTime time.Time `json:"SyncTime,omitempty"`
}
