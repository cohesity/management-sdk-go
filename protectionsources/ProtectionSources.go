package protectionsources

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the PROTECTIONSOURCES_IMPL
 */
type PROTECTIONSOURCES interface {
    ListVirtualMachines (*int64, []string, []string, *bool) ([]*models.ProtectionSource, error)

    DeleteUnregisterProtectionSource (int64) (error)

    ListProtectionSourcesRegistrationInfo ([]string, *bool, []models.Environments2Enum, []int64, *bool, []string) (*models.GetRegistrationInformationResponse, error)

    CreateRegisterProtectionSource (*models.RegisterProtectionSourceParameters) (*models.ProtectionSource, error)

    ListProtectionSourcesRootNodes ([]models.Environments2Enum, *string, *int64) ([]*models.NodeInAProtectionSourcesTree, error)

    ListSqlAagHostsAndDatabases ([]int64) ([]*models.SQLAAGHostAndDatabases, error)

    ListProtectedObjects (models.Environment2Enum, int64, *bool) ([]*models.ProtectedVMInformation, error)

    ListDataStoreInformation (int64) ([]*models.ProtectionSource, error)

    GetProtectionSourcesObjects ([]int64) ([]*models.ProtectionSource, error)

    GetProtectionSourcesObjectById (int64) (*models.ProtectionSource, error)

    CreateRefreshProtectionSourceById (int64) (error)

    ListApplicationServers (models.Environment2Enum, *int64, models.ApplicationEnum, *int64) ([]*models.RegisteredApplicationServer, error)

    ListProtectionSources (*bool, *int64, []models.ExcludeOffice365TypesEnum, *bool, *bool, *bool, []string, []string, []models.ExcludeTypesEnum, *bool, []models.Environments2Enum, *string) ([]*models.NodeInAProtectionSourcesTree, error)

    UpdateApplicationServers (*models.UpdateApplicationServerParameters) (*models.ProtectionSource, error)

    CreateRegisterApplicationServers (*models.RegisterApplicationServersParameters) (*models.ProtectionSource, error)

    DeleteUnregisterApplicationServers (int64) (*models.ProtectionSource, error)

    CreateUpgradePhysicalAgents (*models.UpgradePhysicalServerAgentsRequest) (*models.StatusOfPhysicalUpgradeRequest, error)

    GetDownloadPhysicalAgent (models.HostType8Enum, models.PkgTypeEnum) ([]int64, error)
}

/*
 * Factory for the PROTECTIONSOURCES interaface returning PROTECTIONSOURCES_IMPL
 */
func NewPROTECTIONSOURCES(config configuration.CONFIGURATION) *PROTECTIONSOURCES_IMPL {
    client := new(PROTECTIONSOURCES_IMPL)
    client.config = config
    return client
}
