package remoterestore

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the REMOTERESTORE_IMPL
 */
type REMOTERESTORE interface {
    ListRemoteVaultSearchJobs () ([]*models.RemoteVaultSearchInformation, error)

    CreateRemoteVaultSearchJob (*models.CreateRemoteVaultSearchRequest) (*models.RemoteVaultSearchJobUid, error)

    DeleteStopRemoteVaultSearchJob (*models.SearchJobStopRequest) (error)

    ListRemoteVaultSearchJobById (int64) (*models.RemoteVaultSearchInformation, error)

    GetRemoteVaultSearchJobResults (int64, int64, int64, *int64, *string, *string) (*models.RemoteVaultSearchJobResult, error)

    UploadVaultEncryptionKeys (int64, []*models.VaultEncryptionKey) (error)

    ListRemoteVaultRestoreTasks () ([]*models.RemoteVaultRestoreTaskStatus, error)

    CreateRemoteVaultRestoreTask (*models.CreateRemoteVaultRestoreTaskRequest) (*models.UniqueGlobalId, error)
}

/*
 * Factory for the REMOTERESTORE interaface returning REMOTERESTORE_IMPL
 */
func NewREMOTERESTORE(config configuration.CONFIGURATION) *REMOTERESTORE_IMPL {
    client := new(REMOTERESTORE_IMPL)
    client.config = config
    return client
}
