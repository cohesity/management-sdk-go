package idps

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the IDPS_IMPL
 */
type IDPS interface {
    DeleteIdp (int64) (error)

    GetIdps ([]string, []int64, []string) ([]*models.IdPServiceConfiguration, error)

    UpdateIdp (int64, *models.UpdateIdPConfigurationRequest) (*models.IdPServiceConfiguration, error)

    GetIdpLogin (*string) (error)

    CreateIdp (*models.CreateIdPConfigurationRequest) (*models.IdPServiceConfiguration, error)
}

/*
 * Factory for the IDPS interaface returning IDPS_IMPL
 */
func NewIDPS(config configuration.CONFIGURATION) *IDPS_IMPL {
    client := new(IDPS_IMPL)
    client.config = config
    return client
}
