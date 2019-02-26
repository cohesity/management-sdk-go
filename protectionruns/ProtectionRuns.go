package protectionruns

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the PROTECTIONRUNS_IMPL
 */
type PROTECTIONRUNS interface {
    CreateCancelProtectionJobRun (int64, *models.CancelAProtectionJobRun) (error)

    GetProtectionRuns (*int64, *bool, *int64, *int64, *int64, *int64, []string, *bool, *bool, *int64) ([]*models.ProtectionJobRunInstance, error)

    UpdateProtectionRuns (*models.UpdateProtectionJobRunsParameters) (error)
}

/*
 * Factory for the PROTECTIONRUNS interaface returning PROTECTIONRUNS_IMPL
 */
func NewPROTECTIONRUNS(config configuration.CONFIGURATION) *PROTECTIONRUNS_IMPL {
    client := new(PROTECTIONRUNS_IMPL)
    client.config = config
    return client
}
