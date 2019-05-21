// Copyright 2019 Cohesity Inc.
package protectionruns

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the PROTECTIONRUNS_IMPL
 */
type PROTECTIONRUNS interface {
    GetProtectionRuns (*bool, *int64, *int64, *bool, []string, *bool, *int64, *int64, *int64, *int64, *bool) ([]*models.ProtectionRunInstance, error)

    UpdateProtectionRuns (*models.UpdateProtectionJobRunsParam) (error)

    CreateCancelProtectionJobRun (int64, *models.CancelProtectionJobRunParam) (error)
}

/*
 * Factory for the PROTECTIONRUNS interaface returning PROTECTIONRUNS_IMPL
 */
func NewPROTECTIONRUNS(config configuration.CONFIGURATION) *PROTECTIONRUNS_IMPL {
    client := new(PROTECTIONRUNS_IMPL)
    client.config = config
    return client
}