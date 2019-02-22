package roles

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the ROLES_IMPL
 */
type ROLES interface {
    UpdateRole (string, *models.RoleUpdate) (*models.RoleInformation, error)

    GetRoles ([]string, *bool, *string) ([]*models.RoleInformation, error)

    CreateRole (*models.RoleCreate) (*models.RoleInformation, error)

    DeleteRoles (*models.DeleteRole) (error)
}

/*
 * Factory for the ROLES interaface returning ROLES_IMPL
 */
func NewROLES(config configuration.CONFIGURATION) *ROLES_IMPL {
    client := new(ROLES_IMPL)
    client.config = config
    return client
}
