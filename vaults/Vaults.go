package vaults

import "github.com/cohesity/management-sdk-go/models"
import "github.com/cohesity/management-sdk-go/configuration"

/*
 * Interface for the VAULTS_IMPL
 */
type VAULTS interface {
    GetArchiveMediaInfo (int64, int64, int64, *int64, []int64) ([]*models.TapeMediaInformation, error)

    UpdateBandwidthSettings (*models.VaultBandwidthLimits) (*models.VaultBandwidthLimits, error)

    GetVaultEncryptionKey (int64) (*models.VaultEncryptionKey, error)

    GetVaults (*bool, *int64, *string) ([]*models.Vault, error)

    CreateVault (*models.Vault) (*models.Vault, error)

    GetBandwidthSettings () (*models.VaultBandwidthLimits, error)

    GetVaultById (int64) (*models.Vault, error)

    UpdateVault (int64, *models.Vault) (*models.Vault, error)
}

/*
 * Factory for the VAULTS interaface returning VAULTS_IMPL
 */
func NewVAULTS(config configuration.CONFIGURATION) *VAULTS_IMPL {
    client := new(VAULTS_IMPL)
    client.config = config
    return client
}
