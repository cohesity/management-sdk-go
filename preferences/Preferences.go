package preferences

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the PREFERENCES_IMPL
 */
type PREFERENCES interface {
    GetUserPreferences () ([]*models.UserPreferencesProtoUserPreferencesPreference, error)

    UpdateUserPreferences ([]*models.UserPreferencesProtoUserPreferencesPreference) ([]*models.UserPreferencesProtoUserPreferencesPreference, error)
}

/*
 * Factory for the PREFERENCES interaface returning PREFERENCES_IMPL
 */
func NewPREFERENCES(config configuration.CONFIGURATION) *PREFERENCES_IMPL {
    client := new(PREFERENCES_IMPL)
    client.config = config
    return client
}
