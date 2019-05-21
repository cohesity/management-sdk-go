// Copyright 2019 Cohesity Inc.
package restoretasks

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the RESTORETASKS_IMPL
 */
type RESTORETASKS interface {
    CreateApplicationsCloneTask (*models.ApplicationsRestoreTaskRequest) (*models.RestoreTask, error)

    CreateApplicationsRecoverTask (*models.ApplicationsRestoreTaskRequest) (*models.RestoreTask, error)

    CreateCloneTask (*models.CloneTaskRequest) (*models.RestoreTask, error)

    DeletePublicDestroyCloneTask (int64) (error)

    CreateDeployTask (*models.DeployTaskRequest) (*models.RestoreTask, error)

    CreateDownloadFilesAndFolders (*models.DownloadFilesAndFoldersParams) (*models.RestoreTask, error)

    SearchRestoredFiles ([]int64, *bool, []int64, []int64, *int64, *int64, *int64, *string, []int64, []models.EnvironmentSearchRestoredFilesEnum, *int64) (*models.FileSearchResults, error)

    CreateRestoreFilesTask (*models.RestoreFilesTaskRequest) (*models.RestoreTask, error)

    GetFileSnapshotsInformation (int64, int64, int64, int64, string) ([]*models.FileSnapshotInformation, error)

    SearchObjects (*int64, []string, *string, *int64, []int64, []int64, []models.EnvironmentSearchObjectsEnum, *int64, *int64, *string, []int64, *int64) (*models.ObjectSearchResults, error)

    GetOutlookEmails ([]int64, []int64, []string, []string, *int64, *int64, *string, []int64, []string, *bool, *int64, *string, *string, *int64, *bool) (*models.FileSearchResults, error)

    CreateRecoverTask (*models.RecoverTaskRequest) (*models.RestoreTask, error)

    UpdateRestoreTask (*models.UpdateRestoreTaskParams) (*models.RestoreTask, error)

    GetRestoreTasks ([]int64, *int64, *int64, *int64, []string, models.EnvironmentGetRestoreTasksEnum) ([]*models.RestoreTask, error)

    UpdateCancelRestoreTask (int64) (error)

    GetRestoreTaskById (int64) (*models.RestoreTask, error)

    GetVirtualDiskInformation (int64, int64, int64, int64, int64, int64) ([]*models.VirtualDiskInformation, error)

    GetVmVolumesInformation (*int64, *int64, *int64, *int64, *int64, *int64, *int64, *int64, *bool) ([]*models.VmVolumesInformation, error)
}

/*
 * Factory for the RESTORETASKS interaface returning RESTORETASKS_IMPL
 */
func NewRESTORETASKS(config configuration.CONFIGURATION) *RESTORETASKS_IMPL {
    client := new(RESTORETASKS_IMPL)
    client.config = config
    return client
}