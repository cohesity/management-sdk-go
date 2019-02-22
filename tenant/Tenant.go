package tenant

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the TENANT_IMPL
 */
type TENANT interface {
    UpdateTenantGroups () ([]*models.GroupDetails, error)

    UpdateTenantViewBox (*models.TenantViewBoxUpdateDetails) (*models.TenantViewBoxUpdate, error)

    UpdateTenantVlan (*models.TenantVlanUpdateDetails) (*models.TenantVlanUpdate, error)

    GetTenantsProxyConfigRequest () ([]int64, error)

    UpdateTenantUsers (*models.TenantUserUpdateDetails) ([]*models.UserDetails, error)

    UpdateTenantView (*models.TenantViewUpdateDetails) (*models.TenantViewUpdate, error)

    UpdateTenantLdapProvider (*models.TenantLdapProviderUpdateDetails) (*models.TenantLdapProviderUpdate, error)

    UpdateTenantProtectionPolicy (*models.TenantProtectionPolicyUpdateDetails) (*models.TenantProtectionPolicyUpdate, error)

    UpdateTenantProtectionJob (*models.TenantProtectionJobUpdateDetails) (*models.TenantProtectionJobUpdate, error)

    GetTenantsProxies ([]string) ([]*models.TenantProxy, error)

    CreateTenant (*models.TenantCreateRequest) (*models.TenantDetails, error)

    DeleteTenant (*string) ([]*models.TenantDetails, error)

    UpdateTenantActiveDirectory (*models.TenantActiveDirectoryUpdateDetails) (*models.TenantActiveDirectoryUpdate, error)

    UpdateTenantEntity (*models.TenantEntityUpdateDetails) (*models.TenantEntityUpdate, error)

    GetTenants ([]string, []models.PropertiesEnum, *bool, *bool, []models.Status7Enum) ([]*models.TenantDetails, error)

    UpdateTenant (*models.TenantUpdateDetails) (*models.TenantDetails, error)
}

/*
 * Factory for the TENANT interaface returning TENANT_IMPL
 */
func NewTENANT(config configuration.CONFIGURATION) *TENANT_IMPL {
    client := new(TENANT_IMPL)
    client.config = config
    return client
}
