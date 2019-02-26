package appinstance

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the APPINSTANCE_IMPL
 */
type APPINSTANCE interface {
    GetAppInstances () ([]*models.AppInstance, error)

    UpdateAppInstanceState (int64, *models.UpdateAppInstanceStateParameters) (error)

    CreateLaunchAppInstance (*models.LaunchAppInstance) (*models.AppInstanceIdParameterSpecifiesAppInstanceIdInPathParameter, error)
}

/*
 * Factory for the APPINSTANCE interaface returning APPINSTANCE_IMPL
 */
func NewAPPINSTANCE(config configuration.CONFIGURATION) *APPINSTANCE_IMPL {
    client := new(APPINSTANCE_IMPL)
    client.config = config
    return client
}
