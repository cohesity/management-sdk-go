package mimport

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the MIMPORT_IMPL
 */
type MIMPORT interface {
    UpdateImportConfig (*models.ImportConfigurations) (error)
}

/*
 * Factory for the MIMPORT interaface returning MIMPORT_IMPL
 */
func NewMIMPORT(config configuration.CONFIGURATION) *MIMPORT_IMPL {
    client := new(MIMPORT_IMPL)
    client.config = config
    return client
}
