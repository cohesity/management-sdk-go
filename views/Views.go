package views

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the VIEWS_IMPL
 */
type VIEWS interface {
    GetViewByName (string) (*models.View, error)

    UpdateViewByName (string, *models.View1) (*models.View, error)

    DeleteView (string) (error)

    GetFileLockStatus (string) (*models.FileLockStatus, error)

    CreateLockFile (string, *models.LockFileParameters) (*models.FileLockStatus, error)

    GetViews ([]string, []string, *bool, *int64, *int64, *bool, []string, *bool, []int64, []int64, *bool, *bool) (*models.GetViewsResult, error)

    CreateCloneView (*models.CloneViewRequest) (*models.View, error)

    CreateCloneDirectory (*models.CloneDirectoryRequestParams) (error)

    CreateOverwriteView (*models.OverwriteViewParameters) (*models.View, error)

    CreateRenameView (*models.RenameViewParameters, string) (*models.View, error)

    GetViewUserQuotas (*bool, *string, *bool, *int64, *int64, *string, *string, *bool, *int64, []int64, []string, []string, *string) (*models.ViewUserQuotas, error)

    UpdateViewUserQuota (*models.ViewUserQuotaParameters) (*models.UserQuotaAndUsage, error)

    CreateViewUserQuota (*models.ViewUserQuotaParameters) (*models.UserQuotaAndUsage, error)

    DeleteViewUsersQuota (*models.DeleteViewUsersQuotaParameters) (error)

    UpdateUserQuotaSettings (*models.UpdateUserQuotaSettingsForView) (*models.UserQuotaSettings, error)

    UpdateView (*models.View1) (*models.View, error)

    CreateView (*models.CreateViewRequest) (*models.View, error)

    GetViewsByShareName ([]string, *bool, *string, *int64, *string) (*models.GetViewsAndAliasesByShareResult, error)

    CreateActivateViewAliases (string) (*models.ActivateViewAliasesResult, error)

    CreateViewAlias (*models.ViewAlias) (*models.ViewAlias, error)

    DeleteViewAlias (string) (error)
}

/*
 * Factory for the VIEWS interaface returning VIEWS_IMPL
 */
func NewVIEWS(config configuration.CONFIGURATION) *VIEWS_IMPL {
    client := new(VIEWS_IMPL)
    client.config = config
    return client
}
