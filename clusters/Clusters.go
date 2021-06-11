// Copyright 2019 Cohesity Inc.
package clusters

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the CLUSTERS_IMPL
 */
type CLUSTERS interface {
    GetClusterKeys () (*models.ClusterPublicKeys, error)

    DestroyCluster () (error)

    CreateCloudCluster (*models.CreateCloudClusterParameters) (*models.CreateClusterResult, error)

    CreateExpandCloudCluster (*models.ExpandCloudClusterParameters) (*models.CreateClusterResult, error)

    GetClusterCreationProgress () (*models.ClusterCreationProgressResult, error)

    CreatePhysicalCluster (*models.CreatePhysicalClusterParameters) (*models.CreateClusterResult, error)

    CreateExpandPhysicalCluster (*models.ExpandPhysicalClusterParameters) (*models.CreateClusterResult, error)

    ChangeServiceState (*models.ChangeServiceStateParameters) (*models.ChangeServiceStateResult, error)

    UpdateUpgradeCluster (*models.UpgradeClusterParameters) (*models.UpgradeClusterResult, error)

    CreateVirtualCluster (*models.CreateVirtualClusterParameters) (*models.CreateClusterResult, error)

    GetExternalClientSubnets () (*models.ExternalClientSubnets, error)

    UpdateExternalClientSubnets (*models.ExternalClientSubnets) (*models.ExternalClientSubnets, error)
}

/*
 * Factory for the CLUSTERS interaface returning CLUSTERS_IMPL
 */
func NewCLUSTERS(config configuration.CONFIGURATION) *CLUSTERS_IMPL {
    client := new(CLUSTERS_IMPL)
    client.config = config
    return client
}