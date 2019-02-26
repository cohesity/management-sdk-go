package interfacegroup

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the INTERFACEGROUP_IMPL
 */
type INTERFACEGROUP interface {
    GetInterfaceGroups () ([]*models.InterfaceGroup, error)

    DeleteInterfaceGroup (string) (error)

    CreateInterfaceGroup (*models.InterfaceGroup) (*models.InterfaceGroup, error)

    UpdateInterfaceGroup (*models.InterfaceGroup) (*models.InterfaceGroup, error)
}

/*
 * Factory for the INTERFACEGROUP interaface returning INTERFACEGROUP_IMPL
 */
func NewINTERFACEGROUP(config configuration.CONFIGURATION) *INTERFACEGROUP_IMPL {
    client := new(INTERFACEGROUP_IMPL)
    client.config = config
    return client
}
