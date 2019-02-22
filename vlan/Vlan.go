package vlan

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the VLAN_IMPL
 */
type VLAN interface {
    UpdateVlan (int64, *models.VLAN) (*models.VLAN, error)

    RemoveVlan (int64) (error)

    GetVlans (*bool, []string) ([]*models.VLAN, error)

    CreateVlan () (*models.VLAN, error)

    GetVlanById (int64) (*models.VLAN, error)
}

/*
 * Factory for the VLAN interaface returning VLAN_IMPL
 */
func NewVLAN(config configuration.CONFIGURATION) *VLAN_IMPL {
    client := new(VLAN_IMPL)
    client.config = config
    return client
}
