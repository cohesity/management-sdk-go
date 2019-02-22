package smbfileopens

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the SMBFILEOPENS_IMPL
 */
type SMBFILEOPENS interface {
    GetSmbFileOpens (*string, *string, *string, *int64) (*models.SMBActiveFileOpenResponse, error)

    CreateCloseSmbFileOpen (*models.CloseSMBFileOpenParameters) (error)
}

/*
 * Factory for the SMBFILEOPENS interaface returning SMBFILEOPENS_IMPL
 */
func NewSMBFILEOPENS(config configuration.CONFIGURATION) *SMBFILEOPENS_IMPL {
    client := new(SMBFILEOPENS_IMPL)
    client.config = config
    return client
}
