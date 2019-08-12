// Copyright 2019 Cohesity Inc.
package stats

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the STATS_IMPL
 */
type STATS interface {
    GetActiveAlertsStats (int64, int64) (*models.ActiveAlertsStats, error)

    GetConsumerStats (models.ConsumerTypeGetConsumerStatsEnum, *int64, *string, []int64, []int64, []string) (*models.GetConsumerStatsResult, error)

    GetFileDistributionStats (models.EntityTypeGetFileDistributionStatsEnum) ([]*models.FileDistributionStats, error)

    GetProtectionRunsStats (models.StatusGetProtectionRunsStatsEnum, int64, int64) (*models.ProtectionRunsStats, error)

    GetLastProtectionRunStats () (*models.LastProtectionRunStats, error)

    GetProtectedObjectsSummary ([]models.ExcludeTypeGetProtectedObjectsSummaryEnum) (*models.ProtectedObjectsSummary, error)

    GetRestoreStats (int64, int64) (*models.RestoreStats, error)

    GetStorageStats () (*models.StorageStats, error)

    GetTenantStats (models.ConsumerTypeGetTenantStatsEnum, *int64, *string, []int64, []string) (*models.GetTenantStatsResult, error)

    GetVaultStats () (*models.VaultStats, error)

    GetVaultProviderStats (models.RunTypeGetVaultProviderStatsEnum) ([]*models.VaultProviderStatsInfo, error)

    GetVaultRunStats (models.RunTypeGetVaultRunStatsEnum, int64, int64, models.IntervalEnum) (*models.VaultRunStatsSummary, error)

    GetViewBoxStats ([]int64, []string) (*models.GetViewBoxStatsResult, error)

    GetViewStats (models.MetricEnum, *int64) (*models.ViewStatsSnapshot, error)

    GetViewProtocolStats () ([]*models.ViewProtocolStats, error)
}

/*
 * Factory for the STATS interaface returning STATS_IMPL
 */
func NewSTATS(config configuration.CONFIGURATION) *STATS_IMPL {
    client := new(STATS_IMPL)
    client.config = config
    return client
}