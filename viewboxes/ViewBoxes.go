package viewboxes

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the VIEWBOXES_IMPL
 */
type VIEWBOXES interface {
    GetViewBoxById (int64) (*models.DomainViewBox, error)

    UpdateViewBox (int64, *models.StorageDomainViewBoxRequest) (*models.DomainViewBox, error)

    GetViewBoxes (*bool, []string, *bool, []int64, []string, []int64) ([]*models.DomainViewBox, error)

    CreateViewBox (*models.StorageDomainViewBoxRequest) (*models.DomainViewBox, error)
}

/*
 * Factory for the VIEWBOXES interaface returning VIEWBOXES_IMPL
 */
func NewVIEWBOXES(config configuration.CONFIGURATION) *VIEWBOXES_IMPL {
    client := new(VIEWBOXES_IMPL)
    client.config = config
    return client
}
