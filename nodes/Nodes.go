package nodes

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the NODES_IMPL
 */
type NODES interface {
    GetNodes () ([]*models.Node, error)

    GetNodeById (int64) ([]*models.Node, error)
}

/*
 * Factory for the NODES interaface returning NODES_IMPL
 */
func NewNODES(config configuration.CONFIGURATION) *NODES_IMPL {
    client := new(NODES_IMPL)
    client.config = config
    return client
}
