package routes

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the ROUTES_IMPL
 */
type ROUTES interface {
    GetRoutes () ([]*models.Route, error)

    AddRoute (*models.Route) (*models.Route, error)

    DeleteRoute (*models.DeleteRouteParameters) (error)
}

/*
 * Factory for the ROUTES interaface returning ROUTES_IMPL
 */
func NewROUTES(config configuration.CONFIGURATION) *ROUTES_IMPL {
    client := new(ROUTES_IMPL)
    client.config = config
    return client
}
