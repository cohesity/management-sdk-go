package cluster

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the CLUSTER_IMPL
 */
type CLUSTER interface {
    UpdateCluster (*models.UpdateCluster) (*models.CohesityCluster, error)

    GetCluster (*bool) (*models.CohesityCluster, error)

    GetBasicClusterInfo () (*models.BasicCohesityClusterInformation, error)
}

/*
 * Factory for the CLUSTER interaface returning CLUSTER_IMPL
 */
func NewCLUSTER(config configuration.CONFIGURATION) *CLUSTER_IMPL {
    client := new(CLUSTER_IMPL)
    client.config = config
    return client
}
