package CohesityManagementSdk

import (
	"github.com/cohesity/management-sdk-go/accesstokens"
	"github.com/cohesity/management-sdk-go/activedirectory"
	"github.com/cohesity/management-sdk-go/alerts"
	"github.com/cohesity/management-sdk-go/app"
	"github.com/cohesity/management-sdk-go/appinstance"
	"github.com/cohesity/management-sdk-go/audit"
	"github.com/cohesity/management-sdk-go/certificates"
	"github.com/cohesity/management-sdk-go/cluster"
	"github.com/cohesity/management-sdk-go/clusterpartitions"
	"github.com/cohesity/management-sdk-go/clusters"
	"github.com/cohesity/management-sdk-go/configuration"
	"github.com/cohesity/management-sdk-go/dashboard"
	"github.com/cohesity/management-sdk-go/export"
	"github.com/cohesity/management-sdk-go/groups"
	"github.com/cohesity/management-sdk-go/idps"
	"github.com/cohesity/management-sdk-go/interfacegroup"
	"github.com/cohesity/management-sdk-go/kmsconfiguration"
	"github.com/cohesity/management-sdk-go/ldapprovider"
	"github.com/cohesity/management-sdk-go/mimport"
	"github.com/cohesity/management-sdk-go/models"
	"github.com/cohesity/management-sdk-go/nodes"
	"github.com/cohesity/management-sdk-go/notifications"
	"github.com/cohesity/management-sdk-go/preferences"
	"github.com/cohesity/management-sdk-go/principals"
	"github.com/cohesity/management-sdk-go/privileges"
	"github.com/cohesity/management-sdk-go/protectionjobs"
	"github.com/cohesity/management-sdk-go/protectionpolicies"
	"github.com/cohesity/management-sdk-go/protectionruns"
	"github.com/cohesity/management-sdk-go/protectionsources"
	"github.com/cohesity/management-sdk-go/remotecluster"
	"github.com/cohesity/management-sdk-go/remoterestore"
	"github.com/cohesity/management-sdk-go/restoretasks"
	"github.com/cohesity/management-sdk-go/roles"
	"github.com/cohesity/management-sdk-go/routes"
	"github.com/cohesity/management-sdk-go/search"
	"github.com/cohesity/management-sdk-go/smbfileopens"
	"github.com/cohesity/management-sdk-go/staticroute"
	"github.com/cohesity/management-sdk-go/statistics"
	"github.com/cohesity/management-sdk-go/tenant"
	"github.com/cohesity/management-sdk-go/tenants"
	"github.com/cohesity/management-sdk-go/vaults"
	"github.com/cohesity/management-sdk-go/viewboxes"
	"github.com/cohesity/management-sdk-go/views"
	"github.com/cohesity/management-sdk-go/vlan"
)

/*
 * Interface for the COHESITYMANAGEMENTSDK_IMPL
 */
type COHESITYMANAGEMENTSDK interface {
	Alerts() alerts.ALERTS
	ActiveDirectory() activedirectory.ACTIVEDIRECTORY
	Tenant() tenant.TENANT
	StaticRoute() staticroute.STATICROUTE
	Preferences() preferences.PREFERENCES
	Notifications() notifications.NOTIFICATIONS
	Principals() principals.PRINCIPALS
	Routes() routes.ROUTES
	RemoteCluster() remotecluster.REMOTECLUSTER
	Nodes() nodes.NODES
	InterfaceGroup() interfacegroup.INTERFACEGROUP
	Clusters() clusters.CLUSTERS
	Certificates() certificates.CERTIFICATES
	App() app.APP
	AppInstance() appinstance.APPINSTANCE
	Vlan() vlan.VLAN
	Views() views.VIEWS
	ViewBoxes() viewboxes.VIEWBOXES
	RestoreTasks() restoretasks.RESTORETASKS
	Vaults() vaults.VAULTS
	Tenants() tenants.TENANTS
	Statistics() statistics.STATISTICS
	SMBFileOpens() smbfileopens.SMBFILEOPENS
	Search() search.SEARCH
	Roles() roles.ROLES
	RemoteRestore() remoterestore.REMOTERESTORE
	ProtectionSources() protectionsources.PROTECTIONSOURCES
	ProtectionRuns() protectionruns.PROTECTIONRUNS
	ProtectionPolicies() protectionpolicies.PROTECTIONPOLICIES
	ProtectionJobs() protectionjobs.PROTECTIONJOBS
	Audit() audit.AUDIT
	KmsConfiguration() kmsconfiguration.KMSCONFIGURATION
	Privileges() privileges.PRIVILEGES
	LdapProvider() ldapprovider.LDAPPROVIDER
	Import() mimport.MIMPORT
	Idps() idps.IDPS
	Groups() groups.GROUPS
	Dashboard() dashboard.DASHBOARD
	ClusterPartitions() clusterpartitions.CLUSTERPARTITIONS
	Export() export.EXPORT
	Cluster() cluster.CLUSTER
	AccessTokens() accesstokens.ACCESSTOKENS
	Configuration() configuration.CONFIGURATION
	Authorize() (*models.AccessToken, error)
}

/*
 * Factory for the COHESITYMANAGEMENTSDK interface returning COHESITYMANAGEMENTSDK_IMPL
 */
func NewCohesitySdkClient(clustervip string, username string, password string, domain string, skipssl bool) COHESITYMANAGEMENTSDK {
	cohesityManagementSdkClient := new(COHESITYMANAGEMENTSDK_IMPL)
	cohesityManagementSdkClient.config = configuration.NewCONFIGURATION()

	cohesityManagementSdkClient.config.SetUsername(username)
	cohesityManagementSdkClient.config.SetPassword(password)
	cohesityManagementSdkClient.config.SetDomain(domain)
	cohesityManagementSdkClient.config.SetClusterVip(&clustervip)
	cohesityManagementSdkClient.config.SetSkipSSL(skipssl)
	cohesityManagementSdkClient.Authorize()

	return cohesityManagementSdkClient
}

//CohesityPatch
func NewCohesityClientWithToken(clustervip string, mgmtauthtoken *models.AccessToken) COHESITYMANAGEMENTSDK {
	cohesityManagementSdkClient := new(COHESITYMANAGEMENTSDK_IMPL)
	cohesityManagementSdkClient.config = configuration.NewCONFIGURATION()
	cohesityManagementSdkClient.config.SetClusterVip(&clustervip)
	cohesityManagementSdkClient.config.SetSkipSSL(true)
	cohesityManagementSdkClient.config.SetAccessToken(mgmtauthtoken)
	return cohesityManagementSdkClient
}
