package principals

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the PRINCIPALS_IMPL
 */
type PRINCIPALS interface {
    GetSessionUser () (*models.UserDetails, error)

    CreateResetS3SecretKey (*models.ResetS3SecretAccessKey) (*models.NewS3SecretAccessKey, error)

    UpdateUser (*models.UserRequest) (*models.UserDetails, error)

    CreateUser (*models.UserRequest) (*models.UserDetails, error)

    DeleteUsers (*models.DeleteUsersRequest) (error)

    GetUserPrivileges () ([]string, error)

    GetUsers (*bool, []string, []string, *string, []string) ([]*models.UserDetails, error)

    SearchPrincipals (models.ObjectClass4Enum, *string, []string, *bool, *string) ([]*models.Principal, error)

    UpdateSourcesForPrincipals (*models.UpdateSourcesForPrincipalParameters) (error)

    ListSourcesForPrincipals ([]string) ([]*models.SourcesForSid, error)
}

/*
 * Factory for the PRINCIPALS interaface returning PRINCIPALS_IMPL
 */
func NewPRINCIPALS(config configuration.CONFIGURATION) *PRINCIPALS_IMPL {
    client := new(PRINCIPALS_IMPL)
    client.config = config
    return client
}
