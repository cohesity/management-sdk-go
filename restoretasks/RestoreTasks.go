package restoretasks

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the RESTORETASKS_IMPL
 */
type RESTORETASKS interface {
    SearchObjects (*string, []int64, []int64, []models.Environments1Enum, *int64, *string, *int64, []int64, *int64, *int64, *int64, []string) (*models.ObjectSearchResults, error)

    GetVmVolumesInformation (*int64, *int64, *bool, *int64, *int64, *int64, *int64, *int64, *int64) ([]*models.VMVolumeInformation, error)

    UpdateCancelRestoreTask (int64) (error)

    GetRestoreTaskById (int64) (*models.RestoreTask1, error)

    GetVirtualDiskInformation (int64, int64, int64, int64, int64, int64) ([]*models.SpecifiesInformationAboutTheVirtualDisk, error)

    GetRestoreTasks ([]int64, *int64, *int64, *int64, []string, models.Environment2Enum) ([]*models.RestoreTask1, error)

    GetOutlookEmails (*int64, *bool, *int64, *string, []int64, []int64, []int64, *string, *int64, *string, []string, *int64, []string, *bool, []string) (*models.FileFolderSearchResult1, error)

    GetFileSnapshotsInformation (int64, int64, int64, string, int64) ([]*models.FileFolderSnapshotInformation, error)

    CreateRecoverTask (*models.CreateRestoreTaskRequest) (*models.RestoreTask1, error)

    SearchRestoredFiles ([]int64, *int64, *int64, []int64, *int64, *int64, *bool, *string, []int64, []int64, []models.Environments1Enum) (*models.FileFolderSearchResult1, error)

    CreateDeployTask (*models.DeployTaskRequest) (*models.RestoreTask1, error)

    CreateDownloadFilesAndFolders (*models.DownloadFilesAndFoldersParameters) (*models.RestoreTask1, error)

    CreateRestoreFilesTask (*models.RestoreTask) (*models.RestoreTask1, error)

    UpdateRestoreTask (*models.UpdateRestoreTaskParams) (*models.RestoreTask1, error)

    CreateApplicationsCloneTask (*models.CreateApplicationsRestoreTaskRequest) (*models.RestoreTask1, error)

    CreateApplicationsRecoverTask (*models.CreateApplicationsRestoreTaskRequest) (*models.RestoreTask1, error)

    CreateCloneTask (*models.CloneRestoreTaskRequest) (*models.RestoreTask1, error)

    DestroyCloneTask (int64) (error)
}

/*
 * Factory for the RESTORETASKS interaface returning RESTORETASKS_IMPL
 */
func NewRESTORETASKS(config configuration.CONFIGURATION) *RESTORETASKS_IMPL {
    client := new(RESTORETASKS_IMPL)
    client.config = config
    return client
}
