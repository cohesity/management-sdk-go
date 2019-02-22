package certificates

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the CERTIFICATES_IMPL
 */
type CERTIFICATES interface {
    DeleteWebServerCertificate () (error)

    GetWebServerCertificate () (*models.SSLCertificateConfiguration, error)

    UpdateWebServerCertificate (*models.SSLCertificateConfiguration) (*models.SSLCertificateConfiguration, error)
}

/*
 * Factory for the CERTIFICATES interaface returning CERTIFICATES_IMPL
 */
func NewCERTIFICATES(config configuration.CONFIGURATION) *CERTIFICATES_IMPL {
    client := new(CERTIFICATES_IMPL)
    client.config = config
    return client
}
