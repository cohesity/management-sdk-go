package kmsconfiguration

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the KMSCONFIGURATION_IMPL
 */
type KMSCONFIGURATION interface {
    GetKmsConfig (*string) ([]*models.GetKMSConfigurationResponseParameters, error)

    UpdateKmsConfig (*models.KMSConfiguration) (*models.GetKMSConfigurationResponseParameters, error)

    CreateKmsConfig (*models.KMSConfiguration) (*models.GetKMSConfigurationResponseParameters, error)
}

/*
 * Factory for the KMSCONFIGURATION interaface returning KMSCONFIGURATION_IMPL
 */
func NewKMSCONFIGURATION(config configuration.CONFIGURATION) *KMSCONFIGURATION_IMPL {
    client := new(KMSCONFIGURATION_IMPL)
    client.config = config
    return client
}
