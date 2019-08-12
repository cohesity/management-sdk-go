// Copyright 2019 Cohesity Inc.
package monitoring

import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the MONITORING_IMPL
 */
type MONITORING interface {
    GetAllJobRuns (int64, int64, []string, *int64, *int64) (error)
}

/*
 * Factory for the MONITORING interaface returning MONITORING_IMPL
 */
func NewMONITORING(config configuration.CONFIGURATION) *MONITORING_IMPL {
    client := new(MONITORING_IMPL)
    client.config = config
    return client
}