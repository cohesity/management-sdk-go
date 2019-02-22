package notifications

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the NOTIFICATIONS_IMPL
 */
type NOTIFICATIONS interface {
    UpdateNotifications () (error)

    GetNotifications () (*models.Notifications, error)
}

/*
 * Factory for the NOTIFICATIONS interaface returning NOTIFICATIONS_IMPL
 */
func NewNOTIFICATIONS(config configuration.CONFIGURATION) *NOTIFICATIONS_IMPL {
    client := new(NOTIFICATIONS_IMPL)
    client.config = config
    return client
}
