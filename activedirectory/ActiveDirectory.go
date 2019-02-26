package activedirectory

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the ACTIVEDIRECTORY_IMPL
 */
type ACTIVEDIRECTORY interface {
    GetActiveDirectoryDomainControllers () (*models.DomainControllers, error)

    AddActiveDirectoryPrincipals ([]*models.AddGroupsOrUsersRequest) ([]*models.NonLOCALGroupOrUser, error)

    SearchActiveDirectoryPrincipals (*string, models.ObjectClass4Enum, *string, []string, *bool) ([]*models.ActiveDirectoryPrincipal, error)

    GetActiveDirectoryEntry ([]string, []string, *bool) ([]*models.ActiveDirectory, error)

    UpdatePreferredDomainControllers ([]*models.UpdatePreferredDomainControllerRequest, string) (*models.ActiveDirectory, error)

    UpdateActiveDirectoryMachineAccounts (*models.UpdateMachineAccountsRequest, string) (*models.ActiveDirectory, error)

    UpdateActiveDirectoryLdapProvider (*models.UpdateLDAPProviderRequest, string) (*models.ActiveDirectory, error)

    UpdateActiveDirectoryIgnoredTrustedDomains (*models.UpdateBlacklistedTrustedDomainRequest, string) (*models.ActiveDirectory, error)

    UpdateActiveDirectoryIdMapping (*models.UpdateIDMappingInformationRequest, string) (*models.ActiveDirectory, error)

    ListCentrifyZones (*string) ([]*models.ListCentrifyZone, error)

    DeleteActiveDirectoryEntry (*models.ActiveDirectory) (error)

    CreateActiveDirectoryEntry (*models.ActiveDirectory) (*models.ActiveDirectory, error)
}

/*
 * Factory for the ACTIVEDIRECTORY interaface returning ACTIVEDIRECTORY_IMPL
 */
func NewACTIVEDIRECTORY(config configuration.CONFIGURATION) *ACTIVEDIRECTORY_IMPL {
    client := new(ACTIVEDIRECTORY_IMPL)
    client.config = config
    return client
}
