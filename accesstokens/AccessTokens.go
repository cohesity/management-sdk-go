package accesstokens

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the ACCESSTOKENS_IMPL
 */
type ACCESSTOKENS interface {
    CreateGenerateAccessToken (*models.CreateAccessTokenCredentialRequest) (*models.AccessToken, error)
}

/*
 * Factory for the ACCESSTOKENS interaface returning ACCESSTOKENS_IMPL
 */
func NewACCESSTOKENS(config configuration.CONFIGURATION) *ACCESSTOKENS_IMPL {
    client := new(ACCESSTOKENS_IMPL)
    client.config = config
    return client
}
