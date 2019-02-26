package remotecluster

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the REMOTECLUSTER_IMPL
 */
type REMOTECLUSTER interface {
    GetReplicationEncryptionKey () (*models.ReplicationEncryptionKey, error)

    DeleteRemoteCluster (int64) (error)

    GetRemoteClusters (*bool, *bool, []int64, []string) ([]*models.RemoteCluster, error)

    GetRemoteClusterById (int64) ([]*models.RemoteCluster, error)

    UpdateRemoteCluster (int64, *models.RegisterRemoteCluster) (*models.RemoteCluster, error)

    CreateRemoteCluster (*models.RegisterRemoteCluster) (*models.RemoteCluster, error)
}

/*
 * Factory for the REMOTECLUSTER interaface returning REMOTECLUSTER_IMPL
 */
func NewREMOTECLUSTER(config configuration.CONFIGURATION) *REMOTECLUSTER_IMPL {
    client := new(REMOTECLUSTER_IMPL)
    client.config = config
    return client
}
