package statistics

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the STATISTICS_IMPL
 */
type STATISTICS interface {
    GetTimeSeriesStats (string, string, string, int64, *string, *int64, *int64) (*models.MetricDataBlock, error)

    GetEntitySchemaByName (string) ([]*models.EntitySchema, error)

    GetEntities (string, *bool, []string) ([]*models.Entity, error)

    GetEntitiesSchema ([]string, []string) ([]*models.EntitySchema, error)
}

/*
 * Factory for the STATISTICS interaface returning STATISTICS_IMPL
 */
func NewSTATISTICS(config configuration.CONFIGURATION) *STATISTICS_IMPL {
    client := new(STATISTICS_IMPL)
    client.config = config
    return client
}
