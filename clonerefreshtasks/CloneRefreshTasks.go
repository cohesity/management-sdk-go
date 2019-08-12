// Copyright 2019 Cohesity Inc.
package clonerefreshtasks

import "github.com/cohesity/management-sdk-go/configuration"
import "github.com/cohesity/management-sdk-go/models"

/*
 * Interface for the CLONEREFRESHTASKS_IMPL
 */
type CLONEREFRESHTASKS interface {
    CreateCloneRefreshTask (*models.CloneRefreshRequest) (*models.RestoreTaskWrapper, error)
}

/*
 * Factory for the CLONEREFRESHTASKS interaface returning CLONEREFRESHTASKS_IMPL
 */
func NewCLONEREFRESHTASKS(config configuration.CONFIGURATION) *CLONEREFRESHTASKS_IMPL {
    client := new(CLONEREFRESHTASKS_IMPL)
    client.config = config
    return client
}