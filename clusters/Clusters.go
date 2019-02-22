package clusters

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the CLUSTERS_IMPL
 */
type CLUSTERS interface {
    GetExternalClientSubnets () (*models.SpecifiesTheExternalClientSubnetsThatCanCommunicateWithThisCluster, error)

    GetClusterKeys () (*models.ClusterPublicKeys, error)

    UpdateExternalClientSubnets (*models.SpecifiesTheExternalClientSubnetsThatCanCommunicateWithThisCluster) (*models.SpecifiesTheExternalClientSubnetsThatCanCommunicateWithThisCluster, error)
}

/*
 * Factory for the CLUSTERS interaface returning CLUSTERS_IMPL
 */
func NewCLUSTERS(config configuration.CONFIGURATION) *CLUSTERS_IMPL {
    client := new(CLUSTERS_IMPL)
    client.config = config
    return client
}
