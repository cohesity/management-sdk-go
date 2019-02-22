package app

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the APP_IMPL
 */
type APP interface {
    UploadApp () (*models.AppInformation, error)

    DeleteUninstallApp (int64, int64) (error)

    CreateInstallApp (int64, int64) (*models.AppInformation, error)

    GetApps () ([]*models.AppInformation, error)
}

/*
 * Factory for the APP interaface returning APP_IMPL
 */
func NewAPP(config configuration.CONFIGURATION) *APP_IMPL {
    client := new(APP_IMPL)
    client.config = config
    return client
}
