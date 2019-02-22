package alerts

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the ALERTS_IMPL
 */
type ALERTS interface {
    GetNotificationRules () ([]*models.NotificationRule, error)

    GetAlertCategories () ([]*models.AlertCategoryName, error)

    UpdateNotificationRule () (*models.NotificationRule, error)

    GetAlerts (int64, *bool, []string, *string, []models.AlertStateListEnum, []models.AlertSeverityListEnum, []int64, []string, []int64, []models.AlertCategoryList1Enum, *string, *int64, *int64) ([]*models.Alert, error)

    GetResolutions (int64, []string, *int64, *int64, []string, *bool, []int64) ([]*models.AlertResolution1, error)

    GetAlertById (string) (*models.Alert, error)

    GetAlertTypes () ([]*models.AlertMetadata, error)

    UpdateResolution (int64, *models.UpdateAlertResolutionRequest) (*models.AlertResolution1, error)

    GetResolutionById (int64) (*models.AlertResolution1, error)

    CreateResolution (*models.CreateAlertResolutionRequest) (*models.AlertResolution1, error)

    DeleteNotificationRule (int64) (error)

    CreateNotificationRule (*models.NotificationRule) (*models.NotificationRule, error)
}

/*
 * Factory for the ALERTS interaface returning ALERTS_IMPL
 */
func NewALERTS(config configuration.CONFIGURATION) *ALERTS_IMPL {
    client := new(ALERTS_IMPL)
    client.config = config
    return client
}
