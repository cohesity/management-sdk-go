package groups

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the GROUPS_IMPL
 */
type GROUPS interface {
    GetGroups ([]string, *bool, *string, *string) ([]*models.GroupDetails, error)

    DeleteGroups (*models.DeleteGroupsRequest) (error)

    CreateGroup (*models.GroupRequest) (*models.GroupDetails, error)

    UpdateGroup (*models.GroupRequest) (*models.GroupDetails, error)
}

/*
 * Factory for the GROUPS interaface returning GROUPS_IMPL
 */
func NewGROUPS(config configuration.CONFIGURATION) *GROUPS_IMPL {
    client := new(GROUPS_IMPL)
    client.config = config
    return client
}
