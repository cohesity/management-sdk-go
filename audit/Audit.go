// Copyright 2019 Cohesity Inc.
package audit

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the AUDIT_IMPL
 */
type AUDIT interface {
    SearchClusterAuditLogs (*int64, *string, []string, []string, *string, *int64, *int64, *bool, []string, []string, *int64, *string) (*models.ClusterAuditLogsSearchResult, error)
}

/*
 * Factory for the AUDIT interaface returning AUDIT_IMPL
 */
func NewAUDIT(config configuration.CONFIGURATION) *AUDIT_IMPL {
    client := new(AUDIT_IMPL)
    client.config = config
    return client
}