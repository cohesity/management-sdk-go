package clusterpartitions

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the CLUSTERPARTITIONS_IMPL
 */
type CLUSTERPARTITIONS interface {
    GetClusterPartitions ([]int64, []string) ([]*models.ClusterPartition, error)

    GetClusterPartitionById (int64) (*models.ClusterPartition, error)
}

/*
 * Factory for the CLUSTERPARTITIONS interaface returning CLUSTERPARTITIONS_IMPL
 */
func NewCLUSTERPARTITIONS(config configuration.CONFIGURATION) *CLUSTERPARTITIONS_IMPL {
    client := new(CLUSTERPARTITIONS_IMPL)
    client.config = config
    return client
}
