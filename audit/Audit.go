package audit

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the AUDIT_IMPL
 */
type AUDIT interface {
    SearchClusterAuditLogs (*bool, []string, *string, *int64, *string, *string, []string, []string, []string, *int64, *int64, *int64) (*models.ClusterAuditLogFilterResult, error)
}

/*
 * Factory for the AUDIT interaface returning AUDIT_IMPL
 */
func NewAUDIT(config configuration.CONFIGURATION) *AUDIT_IMPL {
    client := new(AUDIT_IMPL)
    client.config = config
    return client
}
