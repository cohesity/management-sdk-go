package ldapprovider

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the LDAPPROVIDER_IMPL
 */
type LDAPPROVIDER interface {
    DeleteLdapProvider (int64) (error)

    GetLdapProviderStatus (int64) (error)

    UpdateLdapProvider (*models.UpdateLDAPProviderParameters) (*models.LDAPProviderResponse, error)

    CreateLdapProvider (*models.LDAP) (*models.LDAPProviderResponse, error)

    GetLdapProvider ([]string, *bool, []int64) ([]*models.LDAPProviderResponse, error)
}

/*
 * Factory for the LDAPPROVIDER interaface returning LDAPPROVIDER_IMPL
 */
func NewLDAPPROVIDER(config configuration.CONFIGURATION) *LDAPPROVIDER_IMPL {
    client := new(LDAPPROVIDER_IMPL)
    client.config = config
    return client
}
