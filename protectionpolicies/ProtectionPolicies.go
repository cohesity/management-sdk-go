package protectionpolicies

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the PROTECTIONPOLICIES_IMPL
 */
type PROTECTIONPOLICIES interface {
    GetProtectionPolicyById (string) (*models.ProtectionPolicy, error)

    UpdateProtectionPolicy (*models.ProtectionPolicyRequest, string) (*models.ProtectionPolicy, error)

    DeleteProtectionPolicy (string) (error)

    GetProtectionPolicies ([]models.Environments1Enum, []int64, []string, *bool, []string, []string) ([]*models.ProtectionPolicy, error)

    CreateProtectionPolicy (*models.ProtectionPolicyRequest) (*models.ProtectionPolicy, error)
}

/*
 * Factory for the PROTECTIONPOLICIES interaface returning PROTECTIONPOLICIES_IMPL
 */
func NewPROTECTIONPOLICIES(config configuration.CONFIGURATION) *PROTECTIONPOLICIES_IMPL {
    client := new(PROTECTIONPOLICIES_IMPL)
    client.config = config
    return client
}
