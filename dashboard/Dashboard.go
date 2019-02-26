package dashboard

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the DASHBOARD_IMPL
 */
type DASHBOARD interface {
    GetDashboard (*int64, *bool, *bool, []models.TileTypesEnum) (*models.DashboardResponse, error)
}

/*
 * Factory for the DASHBOARD interaface returning DASHBOARD_IMPL
 */
func NewDASHBOARD(config configuration.CONFIGURATION) *DASHBOARD_IMPL {
    client := new(DASHBOARD_IMPL)
    client.config = config
    return client
}
