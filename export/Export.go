package export

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the EXPORT_IMPL
 */
type EXPORT interface {
    CreateExportConfig (*models.ExportParameters) (error)
}

/*
 * Factory for the EXPORT interaface returning EXPORT_IMPL
 */
func NewEXPORT(config configuration.CONFIGURATION) *EXPORT_IMPL {
    client := new(EXPORT_IMPL)
    client.config = config
    return client
}
