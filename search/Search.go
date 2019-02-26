package search

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the SEARCH_IMPL
 */
type SEARCH interface {
    SearchProtectionSources ([]int64, []models.PhysicalServerHostTypesEnum, []string, *int64, *int64, *string, []string, []models.Environments7Enum) ([]*models.ProtectionSourceResponse, error)

    SearchProtectionRuns (string) (*models.ProtectionRunResponse, error)
}

/*
 * Factory for the SEARCH interaface returning SEARCH_IMPL
 */
func NewSEARCH(config configuration.CONFIGURATION) *SEARCH_IMPL {
    client := new(SEARCH_IMPL)
    client.config = config
    return client
}
