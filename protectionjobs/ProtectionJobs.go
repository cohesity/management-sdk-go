package protectionjobs

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the PROTECTIONJOBS_IMPL
 */
type PROTECTIONJOBS interface {
    DeleteProtectionJob (int64, *bool) (error)

    CreateRunProtectionJob (int64, *models.ProtectionRunParameters) (error)

    GetProtectionJobById (int64) (*models.ProtectionJob, error)

    UpdateProtectionJob (*models.ProtectionJobRequest, int64) (*models.ProtectionJob, error)

    GetProtectionJobs (*bool, []string, *bool, *bool, *bool, []models.EnvironmentsEnum, []string, *bool, []int64, []string) ([]*models.ProtectionJob, error)

    ChangeProtectionJobState (int64, *models.ChangeProtectionJobStateParameters) (error)

    CreateProtectionJob (*models.ProtectionJobRequest) (*models.ProtectionJob, error)

    UpdateProtectionJobsState (*models.UpdateStateParametersOfProtectionJobs) (*models.SpecifiesTheResponseOfUpdationOfStateOfMultipleProtectionJobs, error)
}

/*
 * Factory for the PROTECTIONJOBS interaface returning PROTECTIONJOBS_IMPL
 */
func NewPROTECTIONJOBS(config configuration.CONFIGURATION) *PROTECTIONJOBS_IMPL {
    client := new(PROTECTIONJOBS_IMPL)
    client.config = config
    return client
}
