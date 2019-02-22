package models

import "time"


/*
 * Structure for the custom type ApplicationsWrapper
 */
type ApplicationsWrapper struct {
    Applications    []*MapReduceInformation `json:"applications,omitempty" form:"applications,omitempty"` //Applications specifies the list of available map-reduce applications
}

/*
 * Structure for the custom type ApplicationSpecificParameters
 */
type ApplicationSpecificParameters struct {
    TruncateExchangeLog *bool           `json:"truncateExchangeLog,omitempty" form:"truncateExchangeLog,omitempty"` //If true, after the Cohesity Cluster successfully captures a Snapshot
}

/*
 * Structure for the custom type VMVolumeInformation
 */
type VMVolumeInformation struct {
    FilesystemVolumes []*FilesystemVolume `json:"filesystemVolumes,omitempty" form:"filesystemVolumes,omitempty"` //Array of Filesystem Volumes.
}

/*
 * Structure for the custom type ClusterNetworkingResourceInformation
 */
type ClusterNetworkingResourceInformation struct {
    Endpoints       []*ClusterNetworkingEndpoint `json:"endpoints,omitempty" form:"endpoints,omitempty"` //The endpoints by which the resource is accessible.
    Type            *string         `json:"type,omitempty" form:"type,omitempty"` //The type of the resource.
}

/*
 * Structure for the custom type ContinuousSchedule
 */
type ContinuousSchedule struct {
    BackupIntervalMins *int64          `json:"backupIntervalMins,omitempty" form:"backupIntervalMins,omitempty"` //If specified, this field defines the time interval in minutes when
}

/*
 * Structure for the custom type ClusterPublicKeys
 */
type ClusterPublicKeys struct {
    SshPublicKey    *string         `json:"sshPublicKey,omitempty" form:"sshPublicKey,omitempty"` //Specifies the SSH public key used to login to Cluster nodes.
}

/*
 * Structure for the custom type CloudParameters
 */
type CloudParameters struct {
    FailoverToCloud *bool           `json:"failoverToCloud,omitempty" form:"failoverToCloud,omitempty"` //Specifies whether the Protection Sources in this Protection Job
}

/*
 * Structure for the custom type BackupPolicyProtoOneOffSchedule
 */
type BackupPolicyProtoOneOffSchedule struct {
    Time            *time.Time      `json:"time,omitempty" form:"time,omitempty"` //A message to encapusulate time of a day. Users of this proto will have to
}

/*
 * Structure for the custom type BackupJobProtoExcludeSource
 */
type BackupJobProtoExcludeSource struct {
    Entities        []*Entity       `json:"entities,omitempty" form:"entities,omitempty"` //An intersection of leaf-level entities will be obtained after expanding
}

/*
 * Structure for the custom type BackupJobProtoBackupSource
 */
type BackupJobProtoBackupSource struct {
    Entities        []*Entity       `json:"entities,omitempty" form:"entities,omitempty"` //Source entities.
}

/*
 * Structure for the custom type AppInstanceIdParameterSpecifiesAppInstanceIdInPathParameter
 */
type AppInstanceIdParameterSpecifiesAppInstanceIdInPathParameter struct {
    AppInstanceId   int64           `json:"appInstanceId" form:"appInstanceId"` //Specifies the app instance Id.
}

/*
 * Structure for the custom type ActivateViewAliasesResult
 */
type ActivateViewAliasesResult struct {
    Aliases         []*ViewAliasInformation `json:"aliases,omitempty" form:"aliases,omitempty"` //Aliases created for the view. A view alias allows a directory path inside
}

/*
 * Structure for the custom type AcropolisProtectionSource
 */
type AcropolisProtectionSource struct {
    ClusterUuid     *string         `json:"clusterUuid,omitempty" form:"clusterUuid,omitempty"` //Specifies the UUID of the Acropolis cluster instance to which this
    Description     *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description about the Protection Source.
    MountPath       *bool           `json:"mountPath,omitempty" form:"mountPath,omitempty"` //Specifies whether the the VM is an agent VM. This is applicable to
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the Acropolis Object.
    Type            TypeEnum        `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of an Acropolis Protection Source Object such as
    Uuid            *string         `json:"uuid,omitempty" form:"uuid,omitempty"` //Specifies the UUID of the Acropolis Object. This is unique within the
}

/*
 * Structure for the custom type AppInformation
 */
type AppInformation struct {
    AppId              *int64          `json:"appId,omitempty" form:"appId,omitempty"` //TODO: Bhargava - Change it to AppUid
    Clusters           []*ClusterInfo  `json:"clusters,omitempty" form:"clusters,omitempty"` //Specifies the list of clusters on which the app is installed for a
    InstallState       InstallStateEnum `json:"installState,omitempty" form:"installState,omitempty"` //Specifies app installation status.
    InstallTime        *int64          `json:"installTime,omitempty" form:"installTime,omitempty"` //Specifies timestamp when the app was installed.
    IsLatest           *bool           `json:"isLatest,omitempty" form:"isLatest,omitempty"` //Specifies whether the app currently installed on all clusters is the
    Metadata           AppMetadataInformation `json:"metadata,omitempty" form:"metadata,omitempty"` //AppMetadata provides metadata information about an application.
    RequiredPrivileges *[]RequiredPrivilegeEnum `json:"requiredPrivileges,omitempty" form:"requiredPrivileges,omitempty"` //Specifies privileges that are required for this app.
    UninstallTime      *int64          `json:"uninstallTime,omitempty" form:"uninstallTime,omitempty"` //Specifies timestamp when the app was uninstalled.
    Version            *int64          `json:"version,omitempty" form:"version,omitempty"` //Specifies application version assigned by the AppStore.
}

/*
 * Structure for the custom type AmazonCloudCredentials
 */
type AmazonCloudCredentials struct {
    AccessKeyId      *string         `json:"accessKeyId,omitempty" form:"accessKeyId,omitempty"` //Specifies the access key for Amazon service account.
    C2sAccessPortal  C2SAccessPortalCAP `json:"c2sAccessPortal,omitempty" form:"c2sAccessPortal,omitempty"` //Specifies information required to connect to CAP to get AWS credentials.
    Region           *string         `json:"region,omitempty" form:"region,omitempty"` //Specifies the region to use for the Amazon service account.
    SecretAccessKey  *string         `json:"secretAccessKey,omitempty" form:"secretAccessKey,omitempty"` //Specifies the secret access key for Amazon service account.
    ServiceUrl       *string         `json:"serviceUrl,omitempty" form:"serviceUrl,omitempty"` //Specifies the URL (Endpoint) for the service such as s3like.notamazon.com.
    SignatureVersion *int64          `json:"signatureVersion,omitempty" form:"signatureVersion,omitempty"` //Specifies the version of the S3 Compliance.
    TierType         TierTypeEnum    `json:"tierType,omitempty" form:"tierType,omitempty"` //Specifies the storage class of AWS.
    UseHttps         *bool           `json:"useHttps,omitempty" form:"useHttps,omitempty"` //Specifies whether to use http or https to connect to the service.
}

/*
 * Structure for the custom type ClusterNetworkingEndpoint
 */
type ClusterNetworkingEndpoint struct {
    Fqdn            *string         `json:"fqdn,omitempty" form:"fqdn,omitempty"` //The Fully Qualified Domain Name.
    Ipv4Addr        *string         `json:"ipv4Addr,omitempty" form:"ipv4Addr,omitempty"` //The IPv4 address.
    Ipv6Addr        *string         `json:"ipv6Addr,omitempty" form:"ipv6Addr,omitempty"` //The IPv6 address.
}

/*
 * Structure for the custom type ClusterInfo
 */
type ClusterInfo struct {
    ClusterId       *int64          `json:"clusterId,omitempty" form:"clusterId,omitempty"` //Specifies the id of the cluster.
    IncarnationId   *int64          `json:"incarnationId,omitempty" form:"incarnationId,omitempty"` //Specifies the incarnation id of the cluster.
}

/*
 * Structure for the custom type UpdateAlertResolutionRequest
 */
type UpdateAlertResolutionRequest struct {
    AlertIdList     *[]string       `json:"alertIdList,omitempty" form:"alertIdList,omitempty"` //Specifies the Alerts to resolve, which are specified by Alert Ids.
}

/*
 * Structure for the custom type UpdateProtectionJobRunsParameters
 */
type UpdateProtectionJobRunsParameters struct {
    JobRuns         []*UpdateProtectionJobRun `json:"jobRuns,omitempty" form:"jobRuns,omitempty"` //Array of Job Runs.
}

/*
 * Structure for the custom type UpdateLDAPProviderRequest
 */
type UpdateLDAPProviderRequest struct {
    LdapProviderId  *int64          `json:"ldapProviderId,omitempty" form:"ldapProviderId,omitempty"` //Specifies the LDAP provider id which is mapped to an Active Directory
}

/*
 * Structure for the custom type UpdateBlacklistedTrustedDomainRequest
 */
type UpdateBlacklistedTrustedDomainRequest struct {
    IgnoredTrustedDomains *[]string       `json:"ignoredTrustedDomains,omitempty" form:"ignoredTrustedDomains,omitempty"` //Specifies the list of trusted domains that were set by the user to be
}

/*
 * Structure for the custom type StoragePolicyOverride
 */
type StoragePolicyOverride struct {
    DisableInlineDedupAndCompression *bool           `json:"disableInlineDedupAndCompression,omitempty" form:"disableInlineDedupAndCompression,omitempty"` //If false, the inline deduplication and compression settings inherited
}

/*
 * Structure for the custom type SearchJobStopRequest
 */
type SearchJobStopRequest struct {
    SearchJobUid    UniqueGlobalId  `json:"searchJobUid,omitempty" form:"searchJobUid,omitempty"` //Specifies the unique id of the Remote Vault search job in progress.
}

/*
 * Structure for the custom type RPOSchedule
 */
type RPOSchedule struct {
    RpoIntevalMinutes *int64          `json:"rpoIntevalMinutes,omitempty" form:"rpoIntevalMinutes,omitempty"` //If this field is set, then at any point, a recovery point should be
}

/*
 * Structure for the custom type CohesityClusterHardware
 */
type CohesityClusterHardware struct {
    HardwareModels  *[]string       `json:"hardwareModels,omitempty" form:"hardwareModels,omitempty"` //TODO: Write general description for this field
    HardwareVendors *[]string       `json:"hardwareVendors,omitempty" form:"hardwareVendors,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type ClusterAuditLogFilterResult
 */
type ClusterAuditLogFilterResult struct {
    ClusterAuditLogs []*ClusterAuditLog `json:"clusterAuditLogs,omitempty" form:"clusterAuditLogs,omitempty"` //Array of Cluster Audit Logs.
    TotalCount       *int64          `json:"totalCount,omitempty" form:"totalCount,omitempty"` //Specifies the total number of logs that match the specified
}

/*
 * Structure for the custom type ClusterAuditLogConfiguration
 */
type ClusterAuditLogConfiguration struct {
    Enabled             bool            `json:"enabled" form:"enabled"` //Specifies if the Cluster audit logging is enabled on the
    RetentionPeriodDays int64           `json:"retentionPeriodDays" form:"retentionPeriodDays"` //Specifies the number of days to keep (retain) the Cluster audit logs.
}

/*
 * Structure for the custom type CloudDeployTarget
 */
type CloudDeployTarget struct {
    DeployVmsToCloudParams DeployVMsToCloudParams `json:"deployVmsToCloudParams,omitempty" form:"deployVmsToCloudParams,omitempty"` //Contains Cloud specific information needed to identify various resources
    TargetEntity           Entity          `json:"targetEntity,omitempty" form:"targetEntity,omitempty"` //Specifies the attributes and the latest statistics about an entity.
    Type                   *int64          `json:"type,omitempty" form:"type,omitempty"` //The type of the CloudDeploy target.
}

/*
 * Structure for the custom type CloseSMBFileOpenParameters
 */
type CloseSMBFileOpenParameters struct {
    FilePath        *string         `json:"filePath,omitempty" form:"filePath,omitempty"` //Specifies the filepath in the view relative to the root filesystem.
    OpenId          *int64          `json:"openId,omitempty" form:"openId,omitempty"` //Specifies the id of the active open.
    ViewName        *string         `json:"viewName,omitempty" form:"viewName,omitempty"` //Specifies the name of the View in which to search. If a view name is not
}

/*
 * Structure for the custom type CloneTaskInformation
 */
type CloneTaskInformation struct {
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Name of the recovery task.
    TaskId          *string         `json:"taskId,omitempty" form:"taskId,omitempty"` //Id of the recovery task.
}

/*
 * Structure for the custom type AlertMetadata
 */
type AlertMetadata struct {
    AlertDocumentList          []*AlertDocument `json:"alertDocumentList,omitempty" form:"alertDocumentList,omitempty"` //Specifies alert documentation one per each language supported.
    AlertTypeId                *int64          `json:"alertTypeId,omitempty" form:"alertTypeId,omitempty"` //Specifies unique id for the alert type.
    Category                   Category1Enum   `json:"category,omitempty" form:"category,omitempty"` //Specifies category of the alert type.
    DedupIntervalSeconds       *int64          `json:"dedupIntervalSeconds,omitempty" form:"dedupIntervalSeconds,omitempty"` //Specifies dedup interval in seconds. If the same alert is raised multiple
    DedupUntilResolved         *bool           `json:"dedupUntilResolved,omitempty" form:"dedupUntilResolved,omitempty"` //Specifies if the alerts are to be deduped until the current one (if
    HideAlertFromUser          *bool           `json:"hideAlertFromUser,omitempty" form:"hideAlertFromUser,omitempty"` //Specifies whether to show the alert in the iris UI and CLI.
    IgnoreDuplicateOccurrences *bool           `json:"ignoreDuplicateOccurrences,omitempty" form:"ignoreDuplicateOccurrences,omitempty"` //Specifies whether to ignore duplicate occurrences completely.
    PrimaryKeyList             *[]string       `json:"primaryKeyList,omitempty" form:"primaryKeyList,omitempty"` //Specifies properties that serve as primary keys.
    PropertyList               *[]string       `json:"propertyList,omitempty" form:"propertyList,omitempty"` //Specifies list of properties that the client is supposed to provide when
    SendSupportNotification    *bool           `json:"sendSupportNotification,omitempty" form:"sendSupportNotification,omitempty"` //Specifies whether to send support notification for the alert.
    SnmpNotification           *bool           `json:"snmpNotification,omitempty" form:"snmpNotification,omitempty"` //Specifies whether an SNMP notification is sent when an alert is raised.
    Version                    *int64          `json:"version,omitempty" form:"version,omitempty"` //Specifies version of the metadata.
}

/*
 * Structure for the custom type CloneDirectoryRequestParams
 */
type CloneDirectoryRequestParams struct {
    DestinationDirectoryName       *string         `json:"destinationDirectoryName,omitempty" form:"destinationDirectoryName,omitempty"` //Name of the new directory which will contain the clone contents.
    DestinationParentDirectoryPath *string         `json:"destinationParentDirectoryPath,omitempty" form:"destinationParentDirectoryPath,omitempty"` //Specifies the path of the destination parent directory. The source dir
    SourceDirectoryPath            *string         `json:"sourceDirectoryPath,omitempty" form:"sourceDirectoryPath,omitempty"` //Specifies the path of the source directory
}

/*
 * Structure for the custom type AgentInformation
 */
type AgentInformation struct {
    CbmrVersion            *string         `json:"cbmrVersion,omitempty" form:"cbmrVersion,omitempty"` //Specifies the version if Cristie BMR product is installed on the host.
    HostType               HostType3Enum   `json:"hostType,omitempty" form:"hostType,omitempty"` //Specifies the host type where the agent is running. This is only set for
    Id                     *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies the agent's id.
    Name                   *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the agent's name.
    RegistrationInfo       RegisteredSourceInfo `json:"registrationInfo,omitempty" form:"registrationInfo,omitempty"` //Specifies information about a registered Source.
    SourceSideDedupEnabled *bool           `json:"sourceSideDedupEnabled,omitempty" form:"sourceSideDedupEnabled,omitempty"` //Specifies whether source side dedup is enabled or not.
    Status                 StatusEnum      `json:"status,omitempty" form:"status,omitempty"` //Specifies the agent status.
    StatusMessage          *string         `json:"statusMessage,omitempty" form:"statusMessage,omitempty"` //Specifies additional details about the agent status.
    Upgradability          UpgradabilityEnum `json:"upgradability,omitempty" form:"upgradability,omitempty"` //Specifies the upgradability of the agent running on the physical server.
    UpgradeStatus          UpgradeStatusEnum `json:"upgradeStatus,omitempty" form:"upgradeStatus,omitempty"` //Specifies the status of the upgrade of the agent on a physical server.
    UpgradeStatusMessage   *string         `json:"upgradeStatusMessage,omitempty" form:"upgradeStatusMessage,omitempty"` //Specifies detailed message about the agent upgrade failure. This field
    Version                *string         `json:"version,omitempty" form:"version,omitempty"` //Specifies the version of the Agent software.
}

/*
 * Structure for the custom type ChangeProtectionJobStateParameters
 */
type ChangeProtectionJobStateParameters struct {
    Pause           *bool           `json:"pause,omitempty" form:"pause,omitempty"` //If true, the specified Protection Job is paused and no new Runs
    PauseReason     *int64          `json:"pauseReason,omitempty" form:"pauseReason,omitempty"` //Specifies the reason of pausing the job so that depending on the pause
}

/*
 * Structure for the custom type AgentDeploymentStatusResponse
 */
type AgentDeploymentStatusResponse struct {
    CompactVersion       *string         `json:"compactVersion,omitempty" form:"compactVersion,omitempty"` //Specifies the compact version of Cohesity agent. For example, 6.0.1.
    HealthStatus         HealthStatusEnum `json:"healthStatus,omitempty" form:"healthStatus,omitempty"` //Specifies the health status of the Cohesity agent.
    HostIp               *string         `json:"hostIp,omitempty" form:"hostIp,omitempty"` //Specifies the IP of the host on which the agent is installed.
    HostOsType           HostOsTypeEnum  `json:"hostOsType,omitempty" form:"hostOsType,omitempty"` //Specifies the host type on which the agent is installed.
    LastUpgradeStatus    LastUpgradeStatusEnum `json:"lastUpgradeStatus,omitempty" form:"lastUpgradeStatus,omitempty"` //Specifies the status of the last upgrade attempt.
    Upgradability        Upgradability1Enum `json:"upgradability,omitempty" form:"upgradability,omitempty"` //Specfies the upgradability of the agent running on the server.
    UpgradeStatusMessage *string         `json:"upgradeStatusMessage,omitempty" form:"upgradeStatusMessage,omitempty"` //Specifies detailed message about the agent upgrade failure. This field
    Version              *string         `json:"version,omitempty" form:"version,omitempty"` //Specifies the Cohesity software version of the agent. For example:
}

/*
 * Structure for the custom type NonLOCALGroupOrUser
 */
type NonLOCALGroupOrUser struct {
    CreatedTimeMsecs     *int64          `json:"createdTimeMsecs,omitempty" form:"createdTimeMsecs,omitempty"` //Specifies the epoch time in milliseconds when the group or user
    Description          *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description about the user or group.
    Domain               *string         `json:"domain,omitempty" form:"domain,omitempty"` //Specifies the domain of the Active Directory where the
    LastUpdatedTimeMsecs *int64          `json:"lastUpdatedTimeMsecs,omitempty" form:"lastUpdatedTimeMsecs,omitempty"` //Specifies the epoch time in milliseconds when the group or user
    ObjectClass          ObjectClass1Enum `json:"objectClass,omitempty" form:"objectClass,omitempty"` //Specifies the type of the referenced Active Directory principal.
    PrincipalName        *string         `json:"principalName,omitempty" form:"principalName,omitempty"` //Specifies the name of the Active Directory principal,
    Restricted           *bool           `json:"restricted,omitempty" form:"restricted,omitempty"` //Whether the principal is a restricted principal. A restricted principal
    Roles                *[]string       `json:"roles,omitempty" form:"roles,omitempty"` //Array of Roles.
    Sid                  *string         `json:"sid,omitempty" form:"sid,omitempty"` //Specifies the unique Security ID (SID) of the Active Directory principal
}

/*
 * Structure for the custom type CapacityByTier
 */
type CapacityByTier struct {
    StorageTier                  StorageTierEnum `json:"storageTier,omitempty" form:"storageTier,omitempty"` //StorageTier is the type of StorageTier.
    TierMaxPhysicalCapacityBytes *int64          `json:"tierMaxPhysicalCapacityBytes,omitempty" form:"tierMaxPhysicalCapacityBytes,omitempty"` //TierMaxPhysicalCapacityBytes is the maximum physical capacity in bytes of
}

/*
 * Structure for the custom type CancelAProtectionJobRun
 */
type CancelAProtectionJobRun struct {
    CopyTaskUid     UniqueGlobalId  `json:"copyTaskUid,omitempty" form:"copyTaskUid,omitempty"` //Specifies an id for an object that is unique across Cohesity Clusters.
    JobRunId        *int64          `json:"jobRunId,omitempty" form:"jobRunId,omitempty"` //Run Id of a Protection Job Run that needs to be cancelled. If this Run
}

/*
 * Structure for the custom type BlackoutPeriod
 */
type BlackoutPeriod struct {
    Day             Day1Enum        `json:"day,omitempty" form:"day,omitempty"` //Blackout Day.
    EndTime         TimeOfDay       `json:"endTime,omitempty" form:"endTime,omitempty"` //Specifies the end time of the blackout time range.
    StartTime       TimeOfDay       `json:"startTime,omitempty" form:"startTime,omitempty"` //Specifies the start time of the blackout time range.
}

/*
 * Structure for the custom type ActiveDirectory
 */
type ActiveDirectory struct {
    DomainName                 *string         `json:"domainName,omitempty" form:"domainName,omitempty"` //Specifies the fully qualified domain name (FQDN) of an Active Directory.
    FallbackUserIdMappingInfo  UserIDMapping   `json:"fallbackUserIdMappingInfo,omitempty" form:"fallbackUserIdMappingInfo,omitempty"` //Specifies how the Unix and Windows users are mapped in an Active Directory.
    IgnoredTrustedDomains      *[]string       `json:"ignoredTrustedDomains,omitempty" form:"ignoredTrustedDomains,omitempty"` //Specifies the list of trusted domains that were set by the user to be
    LdapProviderId             *int64          `json:"ldapProviderId,omitempty" form:"ldapProviderId,omitempty"` //Specifies the LDAP provider id which is map to this Active Directory
    MachineAccounts            *[]string       `json:"machineAccounts,omitempty" form:"machineAccounts,omitempty"` //Array of Machine Accounts.
    OuName                     *string         `json:"ouName,omitempty" form:"ouName,omitempty"` //Specifies an optional Organizational Unit name.
    Password                   *string         `json:"password,omitempty" form:"password,omitempty"` //Specifies the password for the specified userName.
    PreferredDomainControllers []*UpdatePreferredDomainControllerRequest `json:"preferredDomainControllers,omitempty" form:"preferredDomainControllers,omitempty"` //Specifies Map of Active Directory domain names to its preferred domain
    TenantId                   *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
    TrustedDomains             *[]string       `json:"trustedDomains,omitempty" form:"trustedDomains,omitempty"` //Specifies the trusted domains of the Active Directory domain.
    UnixRootSid                *string         `json:"unixRootSid,omitempty" form:"unixRootSid,omitempty"` //Specifies the SID of the Active Directory domain user to be mapped to
    UserIdMappingInfo          UserIDMapping   `json:"userIdMappingInfo,omitempty" form:"userIdMappingInfo,omitempty"` //Specifies how the Unix and Windows users are mapped in an Active Directory.
    UserName                   *string         `json:"userName,omitempty" form:"userName,omitempty"` //Specifies a userName that has administrative privileges in the domain.
    Workgroup                  *string         `json:"workgroup,omitempty" form:"workgroup,omitempty"` //Specifies an optional Workgroup name.
}

/*
 * Structure for the custom type AcropolisRestoreParameters
 */
type AcropolisRestoreParameters struct {
    DisableNetwork     *bool           `json:"disableNetwork,omitempty" form:"disableNetwork,omitempty"` //Specifies whether the network should be left in disabled state.
    NetworkId          *int64          `json:"networkId,omitempty" form:"networkId,omitempty"` //Specifies a network configuration to be attached to the cloned or
    PoweredOn          *bool           `json:"poweredOn,omitempty" form:"poweredOn,omitempty"` //Specifies the power state of the cloned or recovered objects.
    Prefix             *string         `json:"prefix,omitempty" form:"prefix,omitempty"` //Specifies a prefix to prepended to the source object name to derive a
    StorageContainerId *int64          `json:"storageContainerId,omitempty" form:"storageContainerId,omitempty"` //A storage container where the VM's files should be restored to. This
    Suffix             *string         `json:"suffix,omitempty" form:"suffix,omitempty"` //Specifies a suffix to appended to the original source object name
}

/*
 * Structure for the custom type BasicTaskInfoHandlesTheBasicElementsOfTheNotificationTask
 */
type BasicTaskInfoHandlesTheBasicElementsOfTheNotificationTask struct {
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Name of the recovery task.
    TaskId          *string         `json:"taskId,omitempty" form:"taskId,omitempty"` //Id of the recovery task.
}

/*
 * Structure for the custom type BandwidthLimitOverride
 */
type BandwidthLimitOverride struct {
    BytesPerSecond  *int64          `json:"bytesPerSecond,omitempty" form:"bytesPerSecond,omitempty"` //Specifies the value to override the regular maximum bandwidth rate
    TimePeriods     TimePeriod      `json:"timePeriods,omitempty" form:"timePeriods,omitempty"` //Specifies a time period by specifying a single daily time period
}

/*
 * Structure for the custom type DeleteRole
 */
type DeleteRole struct {
    Names           []string        `json:"names" form:"names"` //Array of Role Names.
}

/*
 * Structure for the custom type ReplicationEncryptionKey
 */
type ReplicationEncryptionKey struct {
    EncryptionKey   *string         `json:"encryptionKey,omitempty" form:"encryptionKey,omitempty"` //Specifies a replication encryption key.
}

/*
 * Structure for the custom type RenameViewParameters
 */
type RenameViewParameters struct {
    NewViewName     string          `json:"newViewName" form:"newViewName"` //Specifies the new name of the View.
}

/*
 * Structure for the custom type ReducersWrapper
 */
type ReducersWrapper struct {
    Reducers        []*InformationAboutAReducer `json:"reducers,omitempty" form:"reducers,omitempty"` //Reducers specifies the list of available reducers in analytics workbench.
}

/*
 * Structure for the custom type QoS
 */
type QoS struct {
    PrincipalName   *string         `json:"principalName,omitempty" form:"principalName,omitempty"` //Specifies the name of the QoS Policy used for the View such as
}

/*
 * Structure for the custom type PureEnvironmentJobParameters
 */
type PureEnvironmentJobParameters struct {
    MaxSnapshotsOnPrimary *int64          `json:"maxSnapshotsOnPrimary,omitempty" form:"maxSnapshotsOnPrimary,omitempty"` //Specifies how many recent snapshots of each backed up entity to retain on
}

/*
 * Structure for the custom type PhysicalFileBackupParameters
 */
type PhysicalFileBackupParameters struct {
    BackupPathInfoVec []*BackupPathInformation `json:"backupPathInfoVec,omitempty" form:"backupPathInfoVec,omitempty"` //Specifies the paths to backup on the Physical source.
}

/*
 * Structure for the custom type NtpSettingsConfig
 */
type NtpSettingsConfig struct {
    NtpServersInternal *bool           `json:"ntpServersInternal,omitempty" form:"ntpServersInternal,omitempty"` //Flag to specify if the NTP servers are on internal network or not.
}

/*
 * Structure for the custom type ParametersForABackupOp
 */
type ParametersForABackupOp struct {
    InstanceId      *string         `json:"instanceId,omitempty" form:"instanceId,omitempty"` //Id of that particular instance
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Name of the recovery task.
    StartTimeUsecs  *string         `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"` //Denotes the start time of the backuptask, needed for deeplinking.
    TaskId          *string         `json:"taskId,omitempty" form:"taskId,omitempty"` //Id of the recovery task.
}

/*
 * Structure for the custom type NewS3SecretAccessKey
 */
type NewS3SecretAccessKey struct {
    NewKey          *string         `json:"newKey,omitempty" form:"newKey,omitempty"` //Specifies the new S3 Secret Access key.
}

/*
 * Structure for the custom type OnlyOneOfTheFollowingFieldsShouldBeSet
 */
type OnlyOneOfTheFollowingFieldsShouldBeSet struct {
    EndAfterNumBackups *int64          `json:"endAfterNumBackups,omitempty" form:"endAfterNumBackups,omitempty"` //The following field has been deprecated.
    EndTimeUsecs       *int64          `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"` //If specified, the backup job will no longer be run after this time.
}

/*
 * Structure for the custom type HypervCloneParameters
 */
type HypervCloneParameters struct {
    DisableNetwork  *bool           `json:"disableNetwork,omitempty" form:"disableNetwork,omitempty"` //Specifies whether the network should be left in disabled state.
    NetworkId       *int64          `json:"networkId,omitempty" form:"networkId,omitempty"` //Specifies a network configuration to be attached to the cloned or
    PoweredOn       *bool           `json:"poweredOn,omitempty" form:"poweredOn,omitempty"` //Specifies the power state of the cloned or recovered objects.
    Prefix          *string         `json:"prefix,omitempty" form:"prefix,omitempty"` //Specifies a prefix to prepended to the source object name to derive a
    ResourceId      *int64          `json:"resourceId,omitempty" form:"resourceId,omitempty"` //The resource (HyperV host) to which the restored VM will be attached.
    Suffix          *string         `json:"suffix,omitempty" form:"suffix,omitempty"` //Specifies a suffix to appended to the original source object name
}

/*
 * Structure for the custom type HealthTile
 */
type HealthTile struct {
    CapacityBytes          *int64          `json:"capacityBytes,omitempty" form:"capacityBytes,omitempty"` //Raw Cluster Capacity in Bytes. This is not usable capacity and does not
    ClusterCloudUsageBytes *int64          `json:"clusterCloudUsageBytes,omitempty" form:"clusterCloudUsageBytes,omitempty"` //Usage in Bytes on the cloud.
    LastDayAlerts          []*Alert        `json:"lastDayAlerts,omitempty" form:"lastDayAlerts,omitempty"` //Alerts in last 24 hours.
    LastDayNumCriticals    *int64          `json:"lastDayNumCriticals,omitempty" form:"lastDayNumCriticals,omitempty"` //Number of Critical Alerts.
    LastDayNumWarnings     *int64          `json:"lastDayNumWarnings,omitempty" form:"lastDayNumWarnings,omitempty"` //Number of Warning Alerts.
    NumNodes               *int64          `json:"numNodes,omitempty" form:"numNodes,omitempty"` //Number of nodes in the cluster.
    NumNodesWithIssues     *int64          `json:"numNodesWithIssues,omitempty" form:"numNodesWithIssues,omitempty"` //Number of nodes in the cluster that are unhealthy.
    PercentFull            *float64        `json:"percentFull,omitempty" form:"percentFull,omitempty"` //Percent the cluster is full.
    RawUsedBytes           *int64          `json:"rawUsedBytes,omitempty" form:"rawUsedBytes,omitempty"` //Raw Bytes used in the cluster.
}

/*
 * Structure for the custom type HardwareInformation
 */
type HardwareInformation struct {
    ChassisModel          *string         `json:"chassisModel,omitempty" form:"chassisModel,omitempty"` //TODO: Write general description for this field
    ChassisSerial         *string         `json:"chassisSerial,omitempty" form:"chassisSerial,omitempty"` //TODO: Write general description for this field
    ChassisType           *string         `json:"chassisType,omitempty" form:"chassisType,omitempty"` //TODO: Write general description for this field
    CohesityChassisSerial *string         `json:"cohesityChassisSerial,omitempty" form:"cohesityChassisSerial,omitempty"` //TODO: Write general description for this field
    CohesityNodeSerial    *string         `json:"cohesityNodeSerial,omitempty" form:"cohesityNodeSerial,omitempty"` //TODO: Write general description for this field
    HbaModel              *string         `json:"hbaModel,omitempty" form:"hbaModel,omitempty"` //TODO: Write general description for this field
    IpmiLanChannel        *string         `json:"ipmiLanChannel,omitempty" form:"ipmiLanChannel,omitempty"` //TODO: Write general description for this field
    MaxSlots              *string         `json:"maxSlots,omitempty" form:"maxSlots,omitempty"` //TODO: Write general description for this field
    NodeModel             *string         `json:"nodeModel,omitempty" form:"nodeModel,omitempty"` //TODO: Write general description for this field
    NodeSerial            *string         `json:"nodeSerial,omitempty" form:"nodeSerial,omitempty"` //TODO: Write general description for this field
    ProductModel          *string         `json:"productModel,omitempty" form:"productModel,omitempty"` //TODO: Write general description for this field
    SlotNumber            *string         `json:"slotNumber,omitempty" form:"slotNumber,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type GroupRequest
 */
type GroupRequest struct {
    Description        *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description of the group.
    Domain             *string         `json:"domain,omitempty" form:"domain,omitempty"` //Specifies the domain of the group.
    IsSmbPrincipalOnly *bool           `json:"isSmbPrincipalOnly,omitempty" form:"isSmbPrincipalOnly,omitempty"` //Specify that this group can have only SMB principals as members if this is
    Name               *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the group.
    Restricted         *bool           `json:"restricted,omitempty" form:"restricted,omitempty"` //Whether the group is a restricted group. Users belonging to a restricted
    Roles              *[]string       `json:"roles,omitempty" form:"roles,omitempty"` //Array of Roles.
    SmbPrincipals      []*SpecifiesStructWithSMBPrincipalDetails `json:"smbPrincipals,omitempty" form:"smbPrincipals,omitempty"` //Specifies the SMB principals. Principals will be added to this group only
    TenantId           *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
    Users              *[]string       `json:"users,omitempty" form:"users,omitempty"` //Specifies the users who are members of the group.
}

/*
 * Structure for the custom type GroupDetails
 */
type GroupDetails struct {
    CreatedTimeMsecs     *int64          `json:"createdTimeMsecs,omitempty" form:"createdTimeMsecs,omitempty"` //Specifies the epoch time in milliseconds when the group was created/added.
    Description          *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description of the group.
    Domain               *string         `json:"domain,omitempty" form:"domain,omitempty"` //Specifies the domain of the group.
    IsSmbPrincipalOnly   *bool           `json:"isSmbPrincipalOnly,omitempty" form:"isSmbPrincipalOnly,omitempty"` //Specify that this group can have only SMB principals as members if this is
    LastUpdatedTimeMsecs *int64          `json:"lastUpdatedTimeMsecs,omitempty" form:"lastUpdatedTimeMsecs,omitempty"` //Specifies the epoch time in milliseconds when the group was last modified.
    Name                 *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the group.
    Restricted           *bool           `json:"restricted,omitempty" form:"restricted,omitempty"` //Whether the group is a restricted group. Users belonging to a restricted
    Roles                *[]string       `json:"roles,omitempty" form:"roles,omitempty"` //Array of Roles.
    Sid                  *string         `json:"sid,omitempty" form:"sid,omitempty"` //Specifies the unique Security ID (SID) of the group.
    SmbPrincipals        []*SpecifiesStructWithSMBPrincipalDetails `json:"smbPrincipals,omitempty" form:"smbPrincipals,omitempty"` //Specifies the SMB principals. Principals will be added to this group only
    TenantId             *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
    Users                *[]string       `json:"users,omitempty" form:"users,omitempty"` //Specifies the users who are members of the group.
}

/*
 * Structure for the custom type GetMapReduceAppRunParameters
 */
type GetMapReduceAppRunParameters struct {
    AppId                 *int64          `json:"appId,omitempty" form:"appId,omitempty"` //ApplicationId is the Id of the map reduce application.
    AppInstanceId         *int64          `json:"appInstanceId,omitempty" form:"appInstanceId,omitempty"` //ApplicationInstanceId is the Id of the map reduce application instance.
    IncludeDetails        *bool           `json:"includeDetails,omitempty" form:"includeDetails,omitempty"` //If this flag is true, then send details of instance, else send only
    LastNumInstances      *int64          `json:"lastNumInstances,omitempty" form:"lastNumInstances,omitempty"` //Give last N instance of an app based on end time.
    MaxRunEndTimeInSecs   *int64          `json:"maxRunEndTimeInSecs,omitempty" form:"maxRunEndTimeInSecs,omitempty"` //MaxRunEndTimestampInSecs specifies the maximum job run end timestamp
    MaxRunStartTimeInSecs *int64          `json:"maxRunStartTimeInSecs,omitempty" form:"maxRunStartTimeInSecs,omitempty"` //MaxRunStartTimestampInSecs specifies the maximum job run start timestamp
    MinRunEndTimeInSecs   *int64          `json:"minRunEndTimeInSecs,omitempty" form:"minRunEndTimeInSecs,omitempty"` //MinRunEndTimestampInSecs specifies the minimum job run end timestamp
    MinRunStartTimeInSecs *int64          `json:"minRunStartTimeInSecs,omitempty" form:"minRunStartTimeInSecs,omitempty"` //MinRunStartTimestampInSecs specifies the minimum job run start timestamp
    PageSize              *int64          `json:"pageSize,omitempty" form:"pageSize,omitempty"` //Number of results to be displayed on a page.
    RunStatus             *string         `json:"runStatus,omitempty" form:"runStatus,omitempty"` //Filter instances based on the map reduce application run status.
    StartOffset           *int64          `json:"startOffset,omitempty" form:"startOffset,omitempty"` //Start offset for pagination from where result needs to be fetched.
}

/*
 * Structure for the custom type FlashBladeFileSystem
 */
type FlashBladeFileSystem struct {
    BackupEnabled        *bool           `json:"backupEnabled,omitempty" form:"backupEnabled,omitempty"` //Specifies whether the .snapshot directory exists on the file system.
    CreatedTimeMsecs     *int64          `json:"createdTimeMsecs,omitempty" form:"createdTimeMsecs,omitempty"` //Specifies the time when the filesystem was created in Unix epoch time
    LogicalCapacityBytes *int64          `json:"logicalCapacityBytes,omitempty" form:"logicalCapacityBytes,omitempty"` //Specifies the total capacity in bytes of the file system.
    LogicalUsedBytes     *int64          `json:"logicalUsedBytes,omitempty" form:"logicalUsedBytes,omitempty"` //Specifies the size of logical data currently represented on the
    NfsInfo              FlashBladeNFSInformation `json:"nfsInfo,omitempty" form:"nfsInfo,omitempty"` //Specifies information specific to NFS protocol exposed by Pure Flash Blade
    PhysicalUsedBytes    *int64          `json:"physicalUsedBytes,omitempty" form:"physicalUsedBytes,omitempty"` //Specifies the size of physical data currently consumed by the file
    Protocols            *[]ProtocolEnum `json:"protocols,omitempty" form:"protocols,omitempty"` //List of Protocols.
    SmbInfo              FlashBladeSMBInformation `json:"smbInfo,omitempty" form:"smbInfo,omitempty"` //Specifies information specific to SMB shares exposed by Pure Flash Blade
    UniqueUsedBytes      *int64          `json:"uniqueUsedBytes,omitempty" form:"uniqueUsedBytes,omitempty"` //Specifies the size of physical data cconsumed by the file system
}

/*
 * Structure for the custom type FilesystemVolume
 */
type FilesystemVolume struct {
    Disks             []*DiskInAVolume `json:"disks,omitempty" form:"disks,omitempty"` //Array of Disks and Partitions.
    DisplayName       *string         `json:"displayName,omitempty" form:"displayName,omitempty"` //Specifies a description about the filesystem.
    FilesystemType    *string         `json:"filesystemType,omitempty" form:"filesystemType,omitempty"` //Specifies type of the filesystem on this volume.
    FilesystemUuid    *string         `json:"filesystemUuid,omitempty" form:"filesystemUuid,omitempty"` //Specifies the uuid of the filesystem.
    IsSupported       *bool           `json:"isSupported,omitempty" form:"isSupported,omitempty"` //If true, this is a supported filesystem volume type.
    LogicalVolume     LogicalVolume   `json:"logicalVolume,omitempty" form:"logicalVolume,omitempty"` //Specify attributes for a kLMV (Linux) or kLDM (Windows) filesystem.
    LogicalVolumeType LogicalVolumeTypeEnum `json:"logicalVolumeType,omitempty" form:"logicalVolumeType,omitempty"` //Specifies the type of logical volume such as kSimpleVolume, kLVM or kLDM.
    Name              *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the volume such as /C.
    VolumeGuid        *string         `json:"volumeGuid,omitempty" form:"volumeGuid,omitempty"` //VolumeGuid is the Volume guid.
}

/*
 * Structure for the custom type EnvBackupParams
 */
type EnvBackupParams struct {
    HypervBackupParams   HypervBackupEnvironmentParameters `json:"hypervBackupParams,omitempty" form:"hypervBackupParams,omitempty"` //Message to capture any additional backup params for a HyperV environment.
    NasBackupParams      NASBackupParameters `json:"nasBackupParams,omitempty" form:"nasBackupParams,omitempty"` //Message to capture any additional backup params for a NAS environment.
    OutlookBackupParams  OutlookBackupEnvironmentParameters `json:"outlookBackupParams,omitempty" form:"outlookBackupParams,omitempty"` //Message to capture any additional backup params for an Outlook environment.
    PhysicalBackupParams PhysicalBackupEnvironmentParameters `json:"physicalBackupParams,omitempty" form:"physicalBackupParams,omitempty"` //Message to capture any additional backup params for a Physical environment.
    SqlBackupJobParams   SQLBackupJobParameters `json:"sqlBackupJobParams,omitempty" form:"sqlBackupJobParams,omitempty"` //Message to capture additional backup job params specific to SQL.
    VmwareBackupParams   VmwareBackupEnvironmentParameters `json:"vmwareBackupParams,omitempty" form:"vmwareBackupParams,omitempty"` //Message to capture any additional backup params for a VMware environment.
}

/*
 * Structure for the custom type TimeSeries
 */
type TimeSeries struct {
    MetricDescriptiveName            *string         `json:"metricDescriptiveName,omitempty" form:"metricDescriptiveName,omitempty"` //Specifies a descriptive name for the metric that is displayed in the
    MetricName                       *string         `json:"metricName,omitempty" form:"metricName,omitempty"` //Specifies the name of the metric such as 'kUnmorphedUsageBytes'.
    MetricUnit                       MetricUnit      `json:"metricUnit,omitempty" form:"metricUnit,omitempty"` //Specifies the unit of measure for the metric.
    RawMetricPublishIntervalHintSecs *int64          `json:"rawMetricPublishIntervalHintSecs,omitempty" form:"rawMetricPublishIntervalHintSecs,omitempty"` //Specifies a suggestion for the interval to collect raw data points.
    TimeToLiveSecs                   *int64          `json:"timeToLiveSecs,omitempty" form:"timeToLiveSecs,omitempty"` //Specifies how long the data point will be stored.
    ValueType                        *int64          `json:"valueType,omitempty" form:"valueType,omitempty"` //Specifies the value type for this metric.
}

/*
 * Structure for the custom type EntitySchema
 */
type EntitySchema struct {
    AttributesDescriptor    AttributesDescriptor `json:"attributesDescriptor,omitempty" form:"attributesDescriptor,omitempty"` //Specifies a list of attributes about an entity.
    IsInternalSchema        *bool           `json:"isInternalSchema,omitempty" form:"isInternalSchema,omitempty"` //Specifies if this schema should be displayed in Advanced Diagnostics
    Name                    *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies a name that uniquely identifies an entity schema such as
    SchemaDescriptiveName   *string         `json:"schemaDescriptiveName,omitempty" form:"schemaDescriptiveName,omitempty"` //Specifies the name of the Schema as displayed in Advanced Diagnostics
    SchemaHelpText          *string         `json:"schemaHelpText,omitempty" form:"schemaHelpText,omitempty"` //Specifies an optional informational description about the schema.
    TimeSeriesDescriptorVec []*TimeSeries   `json:"timeSeriesDescriptorVec,omitempty" form:"timeSeriesDescriptorVec,omitempty"` //Array of Time Series.
    Version                 *int64          `json:"version,omitempty" form:"version,omitempty"` //Specifies the version of the entity schema.
}

/*
 * Structure for the custom type DiskPartition
 */
type DiskPartition struct {
    LengthBytes     *int64          `json:"lengthBytes,omitempty" form:"lengthBytes,omitempty"` //Specifies the length of the block in bytes.
    Number          *int64          `json:"number,omitempty" form:"number,omitempty"` //Specifies a unique number of the partition within the linear disk file.
    OffsetBytes     *int64          `json:"offsetBytes,omitempty" form:"offsetBytes,omitempty"` //Specifies the offset of the block (in bytes) from the beginning
    TypeUuid        *string         `json:"typeUuid,omitempty" form:"typeUuid,omitempty"` //Specifies the partition type uuid.
    Uuid            *string         `json:"uuid,omitempty" form:"uuid,omitempty"` //Specifies the partition uuid.
}

/*
 * Structure for the custom type DeviceTree
 */
type DeviceTree struct {
    CombineMethod   CombineMethodEnum `json:"combineMethod,omitempty" form:"combineMethod,omitempty"` //Specifies how to combine the children of this node.
    DeviceLength    *int64          `json:"deviceLength,omitempty" form:"deviceLength,omitempty"` //Specifies the length of this device. This number should match the
    DeviceNodes     []*DeviceNode   `json:"deviceNodes,omitempty" form:"deviceNodes,omitempty"` //Specifies the children of this node in the device tree.
    StripeSize      *int64          `json:"stripeSize,omitempty" form:"stripeSize,omitempty"` //Specifies the size of the striped data if the data is striped.
}

/*
 * Structure for the custom type DeployVMsToAWSParams
 */
type DeployVMsToAWSParams struct {
    InstanceType          Entity          `json:"instanceType,omitempty" form:"instanceType,omitempty"` //Specifies the attributes and the latest statistics about an entity.
    KeyPairName           Entity          `json:"keyPairName,omitempty" form:"keyPairName,omitempty"` //Specifies the attributes and the latest statistics about an entity.
    NetworkSecurityGroups []*Entity       `json:"networkSecurityGroups,omitempty" form:"networkSecurityGroups,omitempty"` //Names of the network security groups within the above VPC. At least
    Region                Entity          `json:"region,omitempty" form:"region,omitempty"` //Specifies the attributes and the latest statistics about an entity.
    Subnet                Entity          `json:"subnet,omitempty" form:"subnet,omitempty"` //Specifies the attributes and the latest statistics about an entity.
    Vpc                   Entity          `json:"vpc,omitempty" form:"vpc,omitempty"` //Specifies the attributes and the latest statistics about an entity.
}

/*
 * Structure for the custom type DeployTaskRequest
 */
type DeployTaskRequest struct {
    Name            string          `json:"name" form:"name"` //Specifies the name of the Deploy Task. This field must be set and must be
    NewParentId     *int64          `json:"newParentId,omitempty" form:"newParentId,omitempty"` //Specifies a new registered parent Protection Source. If specified
    Objects         []*RestoreObject `json:"objects,omitempty" form:"objects,omitempty"` //Array of Objects.
    Target          CloudDeployTargetDetails `json:"target,omitempty" form:"target,omitempty"` //Message that specifies the details about CloudDeploy target where backup
}

/*
 * Structure for the custom type DataTransferFromVaultSummary
 */
type DataTransferFromVaultSummary struct {
    DataTransferPerTask                         []*DataTransferFromVaultPerTask `json:"dataTransferPerTask,omitempty" form:"dataTransferPerTask,omitempty"` //Array of Data Transferred Per Task.
    NumLogicalBytesTransferred                  *int64          `json:"numLogicalBytesTransferred,omitempty" form:"numLogicalBytesTransferred,omitempty"` //Specifies the total number of logical bytes that have been transferred
    NumPhysicalBytesTransferred                 *int64          `json:"numPhysicalBytesTransferred,omitempty" form:"numPhysicalBytesTransferred,omitempty"` //Specifies the total number of physical bytes that have been transferred
    NumTasks                                    *int64          `json:"numTasks,omitempty" form:"numTasks,omitempty"` //Specifies the number of recover or clone tasks that have transferred data
    PhysicalDataTransferredBytesDuringTimeRange *[]int64        `json:"physicalDataTransferredBytesDuringTimeRange,omitempty" form:"physicalDataTransferredBytesDuringTimeRange,omitempty"` //Array of Physical Data Transferred Per Day.
    VaultName                                   *string         `json:"vaultName,omitempty" form:"vaultName,omitempty"` //Specifies the name of the Vault (External Target).
}

/*
 * Structure for the custom type DataTransferFromVaultPerTask
 */
type DataTransferFromVaultPerTask struct {
    NumLogicalBytesTransferred  *int64          `json:"numLogicalBytesTransferred,omitempty" form:"numLogicalBytesTransferred,omitempty"` //Specifies the total number of logical bytes that are transferred from
    NumPhysicalBytesTransferred *int64          `json:"numPhysicalBytesTransferred,omitempty" form:"numPhysicalBytesTransferred,omitempty"` //Specifies the total number of physical bytes that are transferred
    TaskName                    *string         `json:"taskName,omitempty" form:"taskName,omitempty"` //Specifies the task name.
    TaskType                    *string         `json:"taskType,omitempty" form:"taskType,omitempty"` //Specifies the task type.
}

/*
 * Structure for the custom type Dashboard
 */
type Dashboard struct {
    AuditLogs         AuditLogsTile   `json:"auditLogs,omitempty" form:"auditLogs,omitempty"` //Audit logs for Dashboard.
    ClusterId         *int64          `json:"clusterId,omitempty" form:"clusterId,omitempty"` //Id of the cluster for which dashboard is given.
    Health            HealthTile      `json:"health,omitempty" form:"health,omitempty"` //Health for Dashboard.
    Iops              IOPSTile        `json:"iops,omitempty" form:"iops,omitempty"` //IOPs information for dashboard.
    JobRuns           JobRunsTile     `json:"jobRuns,omitempty" form:"jobRuns,omitempty"` //Jon Runs information.
    ProtectedObjects  ProtectedObjectsTile `json:"protectedObjects,omitempty" form:"protectedObjects,omitempty"` //Protected Objects information.
    Protection        ProtectionTile  `json:"protection,omitempty" form:"protection,omitempty"` //Protection information and statistics.
    Recoveries        RecoveriesTile  `json:"recoveries,omitempty" form:"recoveries,omitempty"` //Recoveries information.
    StorageEfficiency StorageEfficiencyTile `json:"storageEfficiency,omitempty" form:"storageEfficiency,omitempty"` //StorageEfficiencyTile gives tile information for the storage saved
    Throughput        ThroughputTile  `json:"throughput,omitempty" form:"throughput,omitempty"` //Throughput information for dashboard.
}

/*
 * Structure for the custom type CreateRemoteVaultRestoreTaskRequest
 */
type CreateRemoteVaultRestoreTaskRequest struct {
    GlacierRetrievalType GlacierRetrievalTypeEnum `json:"glacierRetrievalType,omitempty" form:"glacierRetrievalType,omitempty"` //Specifies the way data needs to be retrieved from the external target.
    RestoreObjects       []*RestoreProtectionJobIndexAndSnapshots `json:"restoreObjects,omitempty" form:"restoreObjects,omitempty"` //Array of Restore Objects.
    SearchJobUid         UniqueGlobalId  `json:"searchJobUid" form:"searchJobUid"` //Specifies the unique id of the remote Vault search Job.
    TaskName             string          `json:"taskName" form:"taskName"` //Specifies a name of the restore task.
    VaultId              int64           `json:"vaultId" form:"vaultId"` //Specifies the id of the Vault that contains the index and
}

/*
 * Structure for the custom type CopySnapshotTaskStatus
 */
type CopySnapshotTaskStatus struct {
    Error              *string         `json:"error,omitempty" form:"error,omitempty"` //Specifies if an error occurred (if any) while running this task.
    Source             ProtectionSource `json:"source,omitempty" form:"source,omitempty"` //Specifies a generic structure that represents a node
    Stats              CopyRunStats    `json:"stats,omitempty" form:"stats,omitempty"` //Stats for one copy task or aggregated stats of a Copy Run in a
    Status             Status1Enum     `json:"status,omitempty" form:"status,omitempty"` //Specifies the status of the source object being protected.
    TaskEndTimeUsecs   *int64          `json:"taskEndTimeUsecs,omitempty" form:"taskEndTimeUsecs,omitempty"` //Specifies the end time of the copy task. The end time
    TaskStartTimeUsecs *int64          `json:"taskStartTimeUsecs,omitempty" form:"taskStartTimeUsecs,omitempty"` //Specifies the start time of the copy task. The start time
}

/*
 * Structure for the custom type CopyRunStats
 */
type CopyRunStats struct {
    EndTimeUsecs             *int64          `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"` //Specifies the time when this replication ended. If not set, then the
    IsIncremental            *bool           `json:"isIncremental,omitempty" form:"isIncremental,omitempty"` //Specifies whether this archival is incremental for archival targets.
    LogicalBytesTransferred  *int64          `json:"logicalBytesTransferred,omitempty" form:"logicalBytesTransferred,omitempty"` //Specifies the number of logical bytes transferred for this replication
    LogicalSizeBytes         *int64          `json:"logicalSizeBytes,omitempty" form:"logicalSizeBytes,omitempty"` //Specifies the total amount of logical data to be transferred for this
    LogicalTransferRateBps   *int64          `json:"logicalTransferRateBps,omitempty" form:"logicalTransferRateBps,omitempty"` //Specifies average logical bytes transfer rate in bytes per second for
    PhysicalBytesTransferred *int64          `json:"physicalBytesTransferred,omitempty" form:"physicalBytesTransferred,omitempty"` //Specifies the number of physical bytes sent over the wire for
    StartTimeUsecs           *int64          `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"` //Specifies the time when this replication was started. If not set, then
}

/*
 * Structure for the custom type CohesityClusterStatistics
 */
type CohesityClusterStatistics struct {
    CloudUsagePerfStats UsageAndPerformanceStatistics `json:"cloudUsagePerfStats,omitempty" form:"cloudUsagePerfStats,omitempty"` //Provides usage and performance statistics for the remote data stored on
    DataReductionRatio  *float64        `json:"dataReductionRatio,omitempty" form:"dataReductionRatio,omitempty"` //Specifies the ratio of logical bytes (not reduced by
    Id                  *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies the id of the Cohesity Cluster.
    LocalUsagePerfStats UsageAndPerformanceStatistics `json:"localUsagePerfStats,omitempty" form:"localUsagePerfStats,omitempty"` //Provides usage and performance statistics for local data stored directly
    LogicalStats        LogicalStatistics `json:"logicalStats,omitempty" form:"logicalStats,omitempty"` //Specifies the total logical data size of all the local and
    UsagePerfStats      UsageAndPerformanceStatistics `json:"usagePerfStats,omitempty" form:"usagePerfStats,omitempty"` //Provides usage and performance statistics about the local data
}

/*
 * Structure for the custom type ClusterPartition
 */
type ClusterPartition struct {
    HostName        *string         `json:"hostName,omitempty" form:"hostName,omitempty"` //Specifies that hostname that resolves to one or more Virtual IP
    Id              *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies a unique identifier for the Cluster Partition.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the Cluster Partition.
    NodeIds         *[]int64        `json:"nodeIds,omitempty" form:"nodeIds,omitempty"` //Array of Node Ids.
    Vips            *[]string       `json:"vips,omitempty" form:"vips,omitempty"` //Array of VIPs.
    VlanIps         *[]string       `json:"vlanIps,omitempty" form:"vlanIps,omitempty"` //Array of VLAN IPs.
    Vlans           []*VLAN         `json:"vlans,omitempty" form:"vlans,omitempty"` //Array of VLANs.
}

/*
 * Structure for the custom type ClusterConfigProtoSID
 */
type ClusterConfigProtoSID struct {
    IdentifierAuthority *[]int64        `json:"identifierAuthority,omitempty" form:"identifierAuthority,omitempty"` //The authority under which the SID was created. This is always 6 bytes
    RevisionLevel       *int64          `json:"revisionLevel,omitempty" form:"revisionLevel,omitempty"` //The revision level of the SID.
    SubAuthority        *[]int64        `json:"subAuthority,omitempty" form:"subAuthority,omitempty"` //List of ids relative to the identifier_authority that uniquely
}

/*
 * Structure for the custom type CloudDeployTargetDetails
 */
type CloudDeployTargetDetails struct {
    AwsParams       AWSParameters   `json:"awsParams,omitempty" form:"awsParams,omitempty"` //Specifies various resources when converting and deploying a VM to AWS.
    AzureParams     AzureParameters `json:"azureParams,omitempty" form:"azureParams,omitempty"` //Specifies various resources when converting and deploying a VM to Azure.
    Id              *int64          `json:"id,omitempty" form:"id,omitempty"` //Entity corresponding to the cloud deploy target.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the inner object's name or a human-readable string made off the
    Type            Type21Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of the CloudDeploy target.
}

/*
 * Structure for the custom type CIFSShareInformation
 */
type CIFSShareInformation struct {
    Acls            *[]string       `json:"acls,omitempty" form:"acls,omitempty"` //Array of Access Control Lists.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the CIFS share.
    Path            *string         `json:"path,omitempty" form:"path,omitempty"` //Specifies the path of this share under the Vserver's root.
    ServerName      *string         `json:"serverName,omitempty" form:"serverName,omitempty"` //Specifies the CIFS server name (such as 'NETAPP-01') specified by the
}

/*
 * Structure for the custom type ChassisInformation
 */
type ChassisInformation struct {
    ChassisId       *int64          `json:"chassisId,omitempty" form:"chassisId,omitempty"` //ChassisId is a unique id assigned to the chassis.
    ChassisName     *string         `json:"chassisName,omitempty" form:"chassisName,omitempty"` //ChassisName is the name of the chassis. This could be the chassis serial
    Location        *string         `json:"location,omitempty" form:"location,omitempty"` //Location is the location of the chassis within the rack.
    RackId          *int64          `json:"rackId,omitempty" form:"rackId,omitempty"` //Rack is the rack within which this chassis lives.
}

/*
 * Structure for the custom type CentrifyZone
 */
type CentrifyZone struct {
    CentrifySchema    CentrifySchemaEnum `json:"centrifySchema,omitempty" form:"centrifySchema,omitempty"` //Specifies the schema of this Centrify zone.
    Description       *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description of the Centrify zone.
    DistinguishedName *string         `json:"distinguishedName,omitempty" form:"distinguishedName,omitempty"` //Specifies the distinguished name of the Centrify zone.
}

/*
 * Structure for the custom type C2SAccessPortalCAP
 */
type C2SAccessPortalCAP struct {
    Agency                    *string         `json:"agency,omitempty" form:"agency,omitempty"` //Name of the agency.
    BaseUrl                   *string         `json:"baseUrl,omitempty" form:"baseUrl,omitempty"` //The base url of C2S CAP server.
    ClientCertificatePassword *string         `json:"clientCertificatePassword,omitempty" form:"clientCertificatePassword,omitempty"` //Encrypted password of the client private key.
    Mission                   *string         `json:"mission,omitempty" form:"mission,omitempty"` //Name of the mission.
    Role                      *string         `json:"role,omitempty" form:"role,omitempty"` //Role type.
}

/*
 * Structure for the custom type BandwidthLimit
 */
type BandwidthLimit struct {
    BandwidthLimitOverrides []*BandwidthLimitOverride `json:"bandwidthLimitOverrides,omitempty" form:"bandwidthLimitOverrides,omitempty"` //Array of Override Bandwidth Limits.
    RateLimitBytesPerSec    *int64          `json:"rateLimitBytesPerSec,omitempty" form:"rateLimitBytesPerSec,omitempty"` //Specifies the maximum allowed data transfer rate between the local Cluster
    Timezone                *string         `json:"timezone,omitempty" form:"timezone,omitempty"` //Specifies a time zone for the specified time period.
}

/*
 * Structure for the custom type BackupSourceParameters
 */
type BackupSourceParameters struct {
    AppEntityIdVec  *[]int64        `json:"appEntityIdVec,omitempty" form:"appEntityIdVec,omitempty"` //If we are backing up an application (such as SQL), this contains
    PhysicalParams  PhysicalBackupSourceParameters `json:"physicalParams,omitempty" form:"physicalParams,omitempty"` //Message to capture additional backup params for a Physical type source.
    SkipIndexing    *bool           `json:"skipIndexing,omitempty" form:"skipIndexing,omitempty"` //Set to true, if indexing is not required for given source.
    SourceId        *int64          `json:"sourceId,omitempty" form:"sourceId,omitempty"` //Source entity id.
    VmwareParams    VmwareBackupSourceParameters `json:"vmwareParams,omitempty" form:"vmwareParams,omitempty"` //Message to capture additional backup params for a VMware type source.
}

/*
 * Structure for the custom type BackupScript
 */
type BackupScript struct {
    FullBackupScript        RemoteScript    `json:"fullBackupScript,omitempty" form:"fullBackupScript,omitempty"` //Specifies the script that should run for the Full (no CBT) backup schedule
    IncrementalBackupScript RemoteScript    `json:"incrementalBackupScript,omitempty" form:"incrementalBackupScript,omitempty"` //Specifies the script that should run for the CBT-based backup
    LogBackupScript         RemoteScript    `json:"logBackupScript,omitempty" form:"logBackupScript,omitempty"` //Specifies the script that should run for the Log backup schedule
    RemoteHost              RemoteHost      `json:"remoteHost,omitempty" form:"remoteHost,omitempty"` //Specifies the remote host where the remote scripts are executed.
    Username                *string         `json:"username,omitempty" form:"username,omitempty"` //Specifies the username that will be used to login to the remote host.
}

/*
 * Structure for the custom type BackupJobPreOrPostScript
 */
type BackupJobPreOrPostScript struct {
    BackupScript     ScriptPathAndParams `json:"backupScript,omitempty" form:"backupScript,omitempty"` //A message to encapsulate pre or post script associated with a backup job
    FullBackupScript ScriptPathAndParams `json:"fullBackupScript,omitempty" form:"fullBackupScript,omitempty"` //A message to encapsulate pre or post script associated with a backup job
    LogBackupScript  ScriptPathAndParams `json:"logBackupScript,omitempty" form:"logBackupScript,omitempty"` //A message to encapsulate pre or post script associated with a backup job
    RemoteHostParams ContainsParametersToConnectToARemoteHost `json:"remoteHostParams,omitempty" form:"remoteHostParams,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type AzureSourceCredentials
 */
type AzureSourceCredentials struct {
    ApplicationId   *string         `json:"applicationId,omitempty" form:"applicationId,omitempty"` //Specifies Application Id of the active directory of Azure account.
    ApplicationKey  *string         `json:"applicationKey,omitempty" form:"applicationKey,omitempty"` //Specifies Application key of the active directory of Azure account.
    AzureType       AzureTypeEnum   `json:"azureType,omitempty" form:"azureType,omitempty"` //Specifies the entity type such as 'kSubscription' if the environment is
    SubscriptionId  *string         `json:"subscriptionId,omitempty" form:"subscriptionId,omitempty"` //Specifies Subscription id inside a customer's Azure account. It represents
    TenantId        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies Tenant Id of the active directory of Azure account.
}

/*
 * Structure for the custom type AWSParameters
 */
type AWSParameters struct {
    InstanceId              *int64          `json:"instanceId,omitempty" form:"instanceId,omitempty"` //Specfies id of the AWS instance type in which to deploy the VM.
    NetworkSecurityGroupIds *[]int64        `json:"networkSecurityGroupIds,omitempty" form:"networkSecurityGroupIds,omitempty"` //Specifies ids of the netwrok security groups within above VPC.
    Region                  *int64          `json:"region,omitempty" form:"region,omitempty"` //Specifies id of the AWS region in which to deploy the VM.
    SubnetId                *int64          `json:"subnetId,omitempty" form:"subnetId,omitempty"` //Specifies id of the subnet within above VPC.
    VirtualPrivateCloudId   *int64          `json:"virtualPrivateCloudId,omitempty" form:"virtualPrivateCloudId,omitempty"` //Specifies id of the Virtual Private Cloud to chose for the instance type.
}

/*
 * Structure for the custom type AWSSourceCredentials
 */
type AWSSourceCredentials struct {
    AccessKey          *string         `json:"accessKey,omitempty" form:"accessKey,omitempty"` //Specifies Access key of the AWS account.
    AmazonResourceName *string         `json:"amazonResourceName,omitempty" form:"amazonResourceName,omitempty"` //Specifies Amazon Resource Name (owner ID) of the IAM user, act as an
    AwsType            AwsTypeEnum     `json:"awsType,omitempty" form:"awsType,omitempty"` //Specifies the entity type such as 'kIAMUser' if the environment is kAWS.
    SecretAccessKey    *string         `json:"secretAccessKey,omitempty" form:"secretAccessKey,omitempty"` //Specifies Secret Access key of the AWS account.
}

/*
 * Structure for the custom type ApplicationServerObjectToRestore
 */
type ApplicationServerObjectToRestore struct {
    ApplicationServerId  *int64          `json:"applicationServerId,omitempty" form:"applicationServerId,omitempty"` //Specifies the Application Server to restore (for example, kSQL).
    SqlRestoreParameters SQLApplicationServerRestoreParameters `json:"sqlRestoreParameters,omitempty" form:"sqlRestoreParameters,omitempty"` //Specifies the parameters specific the Application Server instance.
    TargetHostId         *int64          `json:"targetHostId,omitempty" form:"targetHostId,omitempty"` //Specifies the target host if the application is to be restored to a
    TargetRootNodeId     *int64          `json:"targetRootNodeId,omitempty" form:"targetRootNodeId,omitempty"` //Specifies the registered root node, like vCenter, of targetHost.
}

/*
 * Structure for the custom type ApplicationInformation
 */
type ApplicationInformation struct {
    ApplicationTreeInfo []*NodeInAProtectionSourcesTree `json:"applicationTreeInfo,omitempty" form:"applicationTreeInfo,omitempty"` //Application Server and the subtrees below them.
    Environment         Environment4Enum `json:"environment,omitempty" form:"environment,omitempty"` //Specifies the environment type of the application such as 'kSQL',
}

/*
 * Structure for the custom type AppMetadataInformation
 */
type AppMetadataInformation struct {
    Author           *string         `json:"author,omitempty" form:"author,omitempty"` //Specifies author of the app.
    CreatedDate      *string         `json:"createdDate,omitempty" form:"createdDate,omitempty"` //Specifies date when the first version of the app was created.
    Description      *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies description about what app does.
    DevVersion       *string         `json:"devVersion,omitempty" form:"devVersion,omitempty"` //Specifies version of the app provided by the developer.
    IconImage        *string         `json:"iconImage,omitempty" form:"iconImage,omitempty"` //Specifies application icon.
    LastModifiedDate *string         `json:"lastModifiedDate,omitempty" form:"lastModifiedDate,omitempty"` //Specifies date when the app was last modified.
    Name             *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies name of the app.
}

/*
 * Structure for the custom type AlertResolution
 */
type AlertResolution struct {
    ResolutionDetails *string         `json:"resolutionDetails,omitempty" form:"resolutionDetails,omitempty"` //Specifies detailed notes about the Resolution.
    ResolutionId      *int64          `json:"resolutionId,omitempty" form:"resolutionId,omitempty"` //Specifies Unique id assigned by the Cohesity Cluster for this Resolution.
    ResolutionSummary *string         `json:"resolutionSummary,omitempty" form:"resolutionSummary,omitempty"` //Specifies short description about the Resolution.
    TimestampUsecs    *int64          `json:"timestampUsecs,omitempty" form:"timestampUsecs,omitempty"` //Specifies unix epoch timestamp (in microseconds) when the Alerts were
    UserName          *string         `json:"userName,omitempty" form:"userName,omitempty"` //Specifies name of the Cohesity Cluster user who resolved the Alerts.
}

/*
 * Structure for the custom type AlertDocument
 */
type AlertDocument struct {
    AlertCause       *string         `json:"alertCause,omitempty" form:"alertCause,omitempty"` //Specifies cause of the Alert that is included in the body of the email
    AlertDescription *string         `json:"alertDescription,omitempty" form:"alertDescription,omitempty"` //Specifies brief description about the Alert that is used in the subject
    AlertHelpText    *string         `json:"alertHelpText,omitempty" form:"alertHelpText,omitempty"` //Specifies instructions describing how to resolve the Alert that is
    AlertName        *string         `json:"alertName,omitempty" form:"alertName,omitempty"` //Specifies short name that describes the Alert type such as DiskBad,
}

/*
 * Structure for the custom type AggregatedSubtreeInfo
 */
type AggregatedSubtreeInfo struct {
    Environment      Environment3Enum `json:"environment,omitempty" form:"environment,omitempty"` //Specifies the environment such as 'kSQL' or 'kVMware', where the
    LeavesCount      *int64          `json:"leavesCount,omitempty" form:"leavesCount,omitempty"` //Specifies the number of leaf nodes under the subtree of this node.
    TotalLogicalSize *int64          `json:"totalLogicalSize,omitempty" form:"totalLogicalSize,omitempty"` //Specifies the total logical size of the data under the subtree
}

/*
 * Structure for the custom type AddGroupsOrUsersRequest
 */
type AddGroupsOrUsersRequest struct {
    Description     *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description about the user or group.
    Domain          *string         `json:"domain,omitempty" form:"domain,omitempty"` //Specifies the domain of the Active Directory where the
    ObjectClass     ObjectClass1Enum `json:"objectClass,omitempty" form:"objectClass,omitempty"` //Specifies the type of the referenced Active Directory principal.
    PrincipalName   *string         `json:"principalName,omitempty" form:"principalName,omitempty"` //Specifies the name of the Active Directory principal,
    Restricted      *bool           `json:"restricted,omitempty" form:"restricted,omitempty"` //Whether the principal is a restricted principal. A restricted principal
    Roles           *[]string       `json:"roles,omitempty" form:"roles,omitempty"` //Array of Roles.
}

/*
 * Structure for the custom type ActiveDirectoryPrincipal
 */
type ActiveDirectoryPrincipal struct {
    Domain          *string         `json:"domain,omitempty" form:"domain,omitempty"` //Specifies the domain name of the where the principal' account is
    FullName        *string         `json:"fullName,omitempty" form:"fullName,omitempty"` //Specifies the full name (first and last names) of the principal.
    ObjectClass     ObjectClassEnum `json:"objectClass,omitempty" form:"objectClass,omitempty"` //Specifies the object class of the principal (either 'kGroup' or 'kUser').
    PrincipalName   *string         `json:"principalName,omitempty" form:"principalName,omitempty"` //Specifies the name of the principal.
    Sid             *string         `json:"sid,omitempty" form:"sid,omitempty"` //Specifies the unique Security id (SID) of the principal.
}

/*
 * Structure for the custom type NetworkingInformation
 */
type NetworkingInformation struct {
    ResourceVec     []*ClusterNetworkingResourceInformation `json:"resourceVec,omitempty" form:"resourceVec,omitempty"` //The list of resources on the the system that are accessible by an
}

/*
 * Structure for the custom type MappersWrapper
 */
type MappersWrapper struct {
    Mappers         []*InformationAboutAMapper `json:"mappers,omitempty" form:"mappers,omitempty"` //Mappers specifies the list of available mappers in analytics workbench.
}

/*
 * Structure for the custom type MapReduceInfoAppProperty
 */
type MapReduceInfoAppProperty struct {
    CsvHeader       *string         `json:"csvHeader,omitempty" form:"csvHeader,omitempty"` //csv_header should be tab separated column names.
}

/*
 * Structure for the custom type MapReduceFileFormats
 */
type MapReduceFileFormats struct {
    SupportedFormats *[]string       `json:"supportedFormats,omitempty" form:"supportedFormats,omitempty"` //Specifies the list of formats supported with integer enum mapping to file
}

/*
 * Structure for the custom type MapReduceAuxData
 */
type MapReduceAuxData struct {
    PatternVec      []*ProtoToRepresentPatternAuxiliaryDataType `json:"patternVec,omitempty" form:"patternVec,omitempty"` //Pattern auxiliary data for a MapReduce.
}

/*
 * Structure for the custom type MSExchangeParameters
 */
type MSExchangeParameters struct {
    PerformLogTruncation *bool           `json:"performLogTruncation,omitempty" form:"performLogTruncation,omitempty"` //If this is set to true, then an attempt will be made to truncate
}

/*
 * Structure for the custom type LDAPProviderStatus
 */
type LDAPProviderStatus struct {
    StatusMessage   *string         `json:"statusMessage,omitempty" form:"statusMessage,omitempty"` //Specifies the connection status message of an LDAP provider.
}

/*
 * Structure for the custom type KillMapReduceInstanceResult
 */
type KillMapReduceInstanceResult struct {
    Error           ErrorProto      `json:"error,omitempty" form:"error,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type IdPReachabilityTestResult
 */
type IdPReachabilityTestResult struct {
    Reachable       *bool           `json:"reachable,omitempty" form:"reachable,omitempty"` //Specifies the flag for Idp reachability.
}

/*
 * Structure for the custom type HypervEnvironmentJobParameters
 */
type HypervEnvironmentJobParameters struct {
    FallbackToCrashConsistent *bool           `json:"fallbackToCrashConsistent,omitempty" form:"fallbackToCrashConsistent,omitempty"` //If true, takes a crash-consistent snapshot when app-consistent snapshot
}

/*
 * Structure for the custom type HypervBackupEnvironmentParameters
 */
type HypervBackupEnvironmentParameters struct {
    AllowCrashConsistentSnapshot *bool           `json:"allowCrashConsistentSnapshot,omitempty" form:"allowCrashConsistentSnapshot,omitempty"` //Whether to fallback to take a crash-consistent snapshot incase taking
}

/*
 * Structure for the custom type FileLockStatusParameters
 */
type FileLockStatusParameters struct {
    Path            *string         `json:"path,omitempty" form:"path,omitempty"` //Specifies the file path that needs to be locked.
}

/*
 * Structure for the custom type SpecifiesTheExternalClientSubnetsThatCanCommunicateWithThisCluster
 */
type SpecifiesTheExternalClientSubnetsThatCanCommunicateWithThisCluster struct {
    ClientSubnets   []*Subnet       `json:"clientSubnets,omitempty" form:"clientSubnets,omitempty"` //Specifies the Client Subnets for the cluster.
}

/*
 * Structure for the custom type ExportParameters
 */
type ExportParameters struct {
    Path            *string         `json:"path,omitempty" form:"path,omitempty"` //Specifies the directory path where to create a configuration files.
}

/*
 * Structure for the custom type EntityIdentifier
 */
type EntityIdentifier struct {
    EntityId        Value           `json:"entityId,omitempty" form:"entityId,omitempty"` //Specifies a data type and data field used to store data.
}

/*
 * Structure for the custom type DomainControllers
 */
type DomainControllers struct {
    DomainControllers *[]string       `json:"domainControllers,omitempty" form:"domainControllers,omitempty"` //Domain Controllers of a domain of an Active Directory domain.
}

/*
 * Structure for the custom type DataTransferToVaultsSummaryResponse
 */
type DataTransferToVaultsSummaryResponse struct {
    DataTransferSummary []*DataTransferToVaultSummary `json:"dataTransferSummary,omitempty" form:"dataTransferSummary,omitempty"` //Array of Summary Data Transfer Statistics.
}

/*
 * Structure for the custom type DataTransferFromVaultsSummaryResponse
 */
type DataTransferFromVaultsSummaryResponse struct {
    DataTransferSummary []*DataTransferFromVaultSummary `json:"dataTransferSummary,omitempty" form:"dataTransferSummary,omitempty"` //Array of Summary Data Transfer Statistics.
}

/*
 * Structure for the custom type DailyWeeklySchedule
 */
type DailyWeeklySchedule struct {
    Days            *[]DayEnum      `json:"days,omitempty" form:"days,omitempty"` //Array of Days.
}

/*
 * Structure for the custom type RemoteVaultSearchJobUid
 */
type RemoteVaultSearchJobUid struct {
    SearchJobUid    UniqueGlobalId  `json:"searchJobUid,omitempty" form:"searchJobUid,omitempty"` //Specifies the unique id assigned for the search Job on the Cluster.
}

/*
 * Structure for the custom type FileFolderSearchResult
 */
type FileFolderSearchResult struct {
    DocumentType       *string         `json:"documentType,omitempty" form:"documentType,omitempty"` //Specifies the inferred document type.
    EmailMetaData      EmailMetaData   `json:"emailMetaData,omitempty" form:"emailMetaData,omitempty"` //Specifies details about the emails and the folder containing emails.
    FileVersions       []*FileFolderVersion `json:"fileVersions,omitempty" form:"fileVersions,omitempty"` //Array of File Versions.
    Filename           *string         `json:"filename,omitempty" form:"filename,omitempty"` //Specifies the name of the found file or folder.
    IsFolder           *bool           `json:"isFolder,omitempty" form:"isFolder,omitempty"` //Specifies if the found item is a folder.
    JobId              *int64          `json:"jobId,omitempty" form:"jobId,omitempty"` //Specifies the Job id for the Protection Job that is currently
    JobUid             UniqueGlobalId  `json:"jobUid,omitempty" form:"jobUid,omitempty"` //Specifies the universal id of the Protection Job that backed up
    ProtectionSource   ProtectionSource `json:"protectionSource,omitempty" form:"protectionSource,omitempty"` //Specifies a generic structure that represents a node
    RegisteredSourceId *int64          `json:"registeredSourceId,omitempty" form:"registeredSourceId,omitempty"` //Specifies the id of the top-level registered source (such as a
    SourceId           *int64          `json:"sourceId,omitempty" form:"sourceId,omitempty"` //Specifies the source id of the object that contains the file or folder.
    Type               Type27Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of the file document such as KDirectory, kFile, etc.
    ViewBoxId          *int64          `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"` //Specifies the id of the Domain (View Box) where the source object that
}

/*
 * Structure for the custom type FileLevelDataLockConfigurations
 */
type FileLevelDataLockConfigurations struct {
    AutoLockAfterDurationIdle         *int64          `json:"autoLockAfterDurationIdle,omitempty" form:"autoLockAfterDurationIdle,omitempty"` //Specifies the duration to lock a file that has not been accessed or
    DefaultFileRetentionDurationMsecs *int64          `json:"defaultFileRetentionDurationMsecs,omitempty" form:"defaultFileRetentionDurationMsecs,omitempty"` //Specifies a global default retention duration for files in this view, if
    ExpiryTimestampMsecs              *int64          `json:"expiryTimestampMsecs,omitempty" form:"expiryTimestampMsecs,omitempty"` //Specifies a definite timestamp in milliseconds for retaining the file.
    LockingProtocol                   LockingProtocolEnum `json:"lockingProtocol,omitempty" form:"lockingProtocol,omitempty"` //Specifies the supported mechanisms to explicity lock a file from NFS/SMB
    MaxRetentionDurationMsecs         *int64          `json:"maxRetentionDurationMsecs,omitempty" form:"maxRetentionDurationMsecs,omitempty"` //Specifies a maximum duration in milliseconds for which any file in this
    MinRetentionDurationMsecs         *int64          `json:"minRetentionDurationMsecs,omitempty" form:"minRetentionDurationMsecs,omitempty"` //Specifies a minimum retention duration in milliseconds after a file gets
}

/*
 * Structure for the custom type EmailMetaData
 */
type EmailMetaData struct {
    BccRecipientAddresses *[]string       `json:"bccRecipientAddresses,omitempty" form:"bccRecipientAddresses,omitempty"` //Specifies the email addresses of the bcc recipients.
    CcRecipientAddresses  *[]string       `json:"ccRecipientAddresses,omitempty" form:"ccRecipientAddresses,omitempty"` //Specifies the email addresses of the cc recipients.
    DomainIds             *[]int64        `json:"domainIds,omitempty" form:"domainIds,omitempty"` //Specifies the domain Ids in which mailboxes are registered.
    EmailSubject          *string         `json:"emailSubject,omitempty" form:"emailSubject,omitempty"` //Specifies the subject of the email.
    FolderKey             *int64          `json:"folderKey,omitempty" form:"folderKey,omitempty"` //Specifes the Parent Folder key.
    FolderName            *string         `json:"folderName,omitempty" form:"folderName,omitempty"` //Specifies the parent folder name of the email.
    HasAttachments        *bool           `json:"hasAttachments,omitempty" form:"hasAttachments,omitempty"` //Specifies whether the emails have any attachments.
    ItemKey               *string         `json:"itemKey,omitempty" form:"itemKey,omitempty"` //Specifies the Key(unique within mailbox) for Outlook item such as Email.
    MailboxIds            *[]int64        `json:"mailboxIds,omitempty" form:"mailboxIds,omitempty"` //Specifies the mailbox Ids which contains the emails/folders.
    ProtectionJobIds      *[]int64        `json:"protectionJobIds,omitempty" form:"protectionJobIds,omitempty"` //Specifies the protection job Ids which have backed up mailbox(es)
    ReceivedEndTime       *int64          `json:"receivedEndTime,omitempty" form:"receivedEndTime,omitempty"` //Specifies the unix end time for querying on email's received time.
    ReceivedStartTime     *int64          `json:"receivedStartTime,omitempty" form:"receivedStartTime,omitempty"` //Specifies the unix start time for querying on email's received time.
    ReceivedTimeSeconds   *int64          `json:"receivedTimeSeconds,omitempty" form:"receivedTimeSeconds,omitempty"` //Specifies the unix time when the email was received.
    RecipientAddresses    *[]string       `json:"recipientAddresses,omitempty" form:"recipientAddresses,omitempty"` //Specifies the email addresses of the recipients.
    SenderAddress         *string         `json:"senderAddress,omitempty" form:"senderAddress,omitempty"` //Specifies the email address of the sender.
    SentTimeSeconds       *int64          `json:"sentTimeSeconds,omitempty" form:"sentTimeSeconds,omitempty"` //Specifies the unix time when the email was sent.
    ShowOnlyEmailFolders  *bool           `json:"showOnlyEmailFolders,omitempty" form:"showOnlyEmailFolders,omitempty"` //Specifies whether the query result should include only Email folders.
}

/*
 * Structure for the custom type DiskInAVolume
 */
type DiskInAVolume struct {
    DiskBlocks           []*DiskBlock    `json:"diskBlocks,omitempty" form:"diskBlocks,omitempty"` //Array of Disk Blocks.
    DiskFormat           DiskFormatEnum  `json:"diskFormat,omitempty" form:"diskFormat,omitempty"` //Specifies the format of the virtual disk.
    DiskPartitions       []*DiskPartition `json:"diskPartitions,omitempty" form:"diskPartitions,omitempty"` //Array of Partitions.
    PartitionTableFormat PartitionTableFormatEnum `json:"partitionTableFormat,omitempty" form:"partitionTableFormat,omitempty"` //Specifies partition table format on a disk.
    SectorSizeBytes      *int64          `json:"sectorSizeBytes,omitempty" form:"sectorSizeBytes,omitempty"` //Specifies the sector size of hard disk. It is used for mapping the disk
    Uuid                 *string         `json:"uuid,omitempty" form:"uuid,omitempty"` //Specifies the disk uuid.
    VmdkFileName         *string         `json:"vmdkFileName,omitempty" form:"vmdkFileName,omitempty"` //Specifies the disk file name. This is the VMDK name and not the
    VmdkSizeBytes        *int64          `json:"vmdkSizeBytes,omitempty" form:"vmdkSizeBytes,omitempty"` //Specifies the disk size in bytes.
}

/*
 * Structure for the custom type DeployVMsToAzureParams
 */
type DeployVMsToAzureParams struct {
    AzureManagedDiskParams AzureManagedDiskParams `json:"azureManagedDiskParams,omitempty" form:"azureManagedDiskParams,omitempty"` //Contains managed disk parameters needed to deploy to Azure using managed
    ComputeOptions         Entity          `json:"computeOptions,omitempty" form:"computeOptions,omitempty"` //Specifies the attributes and the latest statistics about an entity.
    NetworkResourceGroup   Entity          `json:"networkResourceGroup,omitempty" form:"networkResourceGroup,omitempty"` //Specifies the attributes and the latest statistics about an entity.
    NetworkSecurityGroup   Entity          `json:"networkSecurityGroup,omitempty" form:"networkSecurityGroup,omitempty"` //Specifies the attributes and the latest statistics about an entity.
    ResourceGroup          Entity          `json:"resourceGroup,omitempty" form:"resourceGroup,omitempty"` //Specifies the attributes and the latest statistics about an entity.
    StorageAccount         Entity          `json:"storageAccount,omitempty" form:"storageAccount,omitempty"` //Specifies the attributes and the latest statistics about an entity.
    StorageContainer       Entity          `json:"storageContainer,omitempty" form:"storageContainer,omitempty"` //Specifies the attributes and the latest statistics about an entity.
    StorageKey             Entity          `json:"storageKey,omitempty" form:"storageKey,omitempty"` //Specifies the attributes and the latest statistics about an entity.
    StorageResourceGroup   Entity          `json:"storageResourceGroup,omitempty" form:"storageResourceGroup,omitempty"` //Specifies the attributes and the latest statistics about an entity.
    Subnet                 Entity          `json:"subnet,omitempty" form:"subnet,omitempty"` //Specifies the attributes and the latest statistics about an entity.
    VirtualNetwork         Entity          `json:"virtualNetwork,omitempty" form:"virtualNetwork,omitempty"` //Specifies the attributes and the latest statistics about an entity.
}

/*
 * Structure for the custom type DataTransferToVaultSummary
 */
type DataTransferToVaultSummary struct {
    DataTransferPerProtectionJob                []*DataTransferToVaultPerProtectionJob `json:"dataTransferPerProtectionJob,omitempty" form:"dataTransferPerProtectionJob,omitempty"` //Array of Data Transfer Statistics Per Protection Jobs.
    LogicalDataTransferredBytesDuringTimeRange  *[]int64        `json:"logicalDataTransferredBytesDuringTimeRange,omitempty" form:"logicalDataTransferredBytesDuringTimeRange,omitempty"` //Array of Logical Data Transferred Per Day.
    NumLogicalBytesTransferred                  *int64          `json:"numLogicalBytesTransferred,omitempty" form:"numLogicalBytesTransferred,omitempty"` //Specifies the total number of logical bytes that are transferred
    NumPhysicalBytesTransferred                 *int64          `json:"numPhysicalBytesTransferred,omitempty" form:"numPhysicalBytesTransferred,omitempty"` //Specifies the total number of physical bytes that are transferred
    NumProtectionJobs                           *int64          `json:"numProtectionJobs,omitempty" form:"numProtectionJobs,omitempty"` //Specifies the number of Protection Jobs that transfer data to
    PhysicalDataTransferredBytesDuringTimeRange *[]int64        `json:"physicalDataTransferredBytesDuringTimeRange,omitempty" form:"physicalDataTransferredBytesDuringTimeRange,omitempty"` //Array of Physical Data Transferred Per Day.
    StorageConsumedBytes                        *int64          `json:"storageConsumedBytes,omitempty" form:"storageConsumedBytes,omitempty"` //Specifies the storage consumed on the Vault as of last day in the
    VaultId                                     *int64          `json:"vaultId,omitempty" form:"vaultId,omitempty"` //The vault Id associated with the vault.
    VaultName                                   *string         `json:"vaultName,omitempty" form:"vaultName,omitempty"` //Specifies the name of the Vault (External Target).
    VaultType                                   VaultType1Enum  `json:"vaultType,omitempty" form:"vaultType,omitempty"` //Specifies the type of Vault.
}

/*
 * Structure for the custom type PhysicalEnvironmentJobParameters
 */
type PhysicalEnvironmentJobParameters struct {
    FilePathFilters                FileOrDirectoryPathFilter `json:"filePathFilters,omitempty" form:"filePathFilters,omitempty"` //Specifies filters to match files and directories on a Server.
    IncrementalSnapshotUponRestart *bool           `json:"incrementalSnapshotUponRestart,omitempty" form:"incrementalSnapshotUponRestart,omitempty"` //If true, performs an incremental backup after server restarts. Otherwise
}

/*
 * Structure for the custom type PhysicalBackupEnvironmentParameters
 */
type PhysicalBackupEnvironmentParameters struct {
    EnableIncrementalBackupAfterRestart *bool           `json:"enableIncrementalBackupAfterRestart,omitempty" form:"enableIncrementalBackupAfterRestart,omitempty"` //If this is set to true, then incremental backup will be performed
    FilteringPolicy                     FilteringPolicyProto `json:"filteringPolicy,omitempty" form:"filteringPolicy,omitempty"` //Proto to encapsulate the filtering policy for backup objects like files or
}

/*
 * Structure for the custom type PatternRequestBody
 */
type PatternRequestBody struct {
    ApplicationDataType int64           `json:"applicationDataType" form:"applicationDataType"` //Specifies the data type for which suppprted patterns can be fetched.
    ApplicationId       int64           `json:"applicationId" form:"applicationId"` //Specifies AWB Application ID.
    UserPattern         SupportedPattern `json:"userPattern" form:"userPattern"` //Specifies details of the pattern available for search available in an
}

/*
 * Structure for the custom type ProtoToRepresentPatternAuxiliaryDataType
 */
type ProtoToRepresentPatternAuxiliaryDataType struct {
    IsSystemDefined *bool           `json:"isSystemDefined,omitempty" form:"isSystemDefined,omitempty"` //Whether this pattern is system defined.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Name of the pattern. This is marked optional but is required.
    Type            *int64          `json:"type,omitempty" form:"type,omitempty"` //Pattern type.
    Value           *string         `json:"value,omitempty" form:"value,omitempty"` //Value of the pattern. This is marked optional but is required.
}

/*
 * Structure for the custom type OverwriteViewParameters
 */
type OverwriteViewParameters struct {
    SourceViewName  string          `json:"sourceViewName" form:"sourceViewName"` //Specifies the source view name.
    TargetViewName  string          `json:"targetViewName" form:"targetViewName"` //Specifies the target view name.
}

/*
 * Structure for the custom type OutlookRestoreParameters
 */
type OutlookRestoreParameters struct {
    OutlookMailboxList []*OutlookMailbox `json:"outlookMailboxList,omitempty" form:"outlookMailboxList,omitempty"` //Specifies the list of mailboxes to be restored.
    TargetFolderPath   *string         `json:"targetFolderPath,omitempty" form:"targetFolderPath,omitempty"` //Specifies the target folder path to restore the mailboxes. This will
    TargetMailbox      ProtectionSource `json:"targetMailbox,omitempty" form:"targetMailbox,omitempty"` //Specifies a generic structure that represents a node
}

/*
 * Structure for the custom type OutlookMailbox
 */
type OutlookMailbox struct {
    MailboxObject        RestoreObject   `json:"mailboxObject,omitempty" form:"mailboxObject,omitempty"` //Specifies an object to recover or clone or an object to restore files
    OutlookFolderList    []*OutlookFolder `json:"outlookFolderList,omitempty" form:"outlookFolderList,omitempty"` //Specifies the list of folders to be restored incase user wishes not to
    RestoreEntireMailbox *bool           `json:"restoreEntireMailbox,omitempty" form:"restoreEntireMailbox,omitempty"` //Specifies whether the enitre mailbox is to be restored.
}

/*
 * Structure for the custom type OutlookEnvironmentJobParameters
 */
type OutlookEnvironmentJobParameters struct {
    FilePathFilter  FileOrDirectoryPathFilter `json:"filePathFilter,omitempty" form:"filePathFilter,omitempty"` //Specifies filters to match files and directories on a Server.
}

/*
 * Structure for the custom type OutlookBackupEnvironmentParameters
 */
type OutlookBackupEnvironmentParameters struct {
    FilteringPolicy FilteringPolicyProto `json:"filteringPolicy,omitempty" form:"filteringPolicy,omitempty"` //Proto to encapsulate the filtering policy for backup objects like files or
}

/*
 * Structure for the custom type OracleSession
 */
type OracleSession struct {
    Location         *string         `json:"location,omitempty" form:"location,omitempty"` //Location is the path where Oracle is installed.
    SystemIdentifier *string         `json:"systemIdentifier,omitempty" form:"systemIdentifier,omitempty"` //SystemIdentifier is the unique Oracle System Identifier for the DB instance.
}

/*
 * Structure for the custom type OracleHost
 */
type OracleHost struct {
    CpuCount        *int64          `json:"cpuCount,omitempty" form:"cpuCount,omitempty"` //Specifies the count of CPU available on the host.
    IpAddresses     *[]string       `json:"ipAddresses,omitempty" form:"ipAddresses,omitempty"` //Specifies the IP address of the host.
    Sessions        []*OracleSession `json:"sessions,omitempty" form:"sessions,omitempty"` //Specifies multiple session configurations available for this host.
}

/*
 * Structure for the custom type ObjectsProtectedByPolicy
 */
type ObjectsProtectedByPolicy struct {
    ObjectsProtected []*ObjectsByEnv `json:"objectsProtected,omitempty" form:"objectsProtected,omitempty"` //Protected Objects.
    PolicyId         *string         `json:"policyId,omitempty" form:"policyId,omitempty"` //Id of the policy.
    PolicyName       *string         `json:"policyName,omitempty" form:"policyName,omitempty"` //Name of the policy.
}

/*
 * Structure for the custom type ObjectsByEnv
 */
type ObjectsByEnv struct {
    EnvType         *string         `json:"envType,omitempty" form:"envType,omitempty"` //Environment Type.
    NumObjects      *int64          `json:"numObjects,omitempty" form:"numObjects,omitempty"` //Number of Objects.
}

/*
 * Structure for the custom type ObjectSearchResults
 */
type ObjectSearchResults struct {
    ObjectSnapshotInfo []*ObjectSnapshot `json:"objectSnapshotInfo,omitempty" form:"objectSnapshotInfo,omitempty"` //Array of Snapshot Objects.
    TotalCount         *int64          `json:"totalCount,omitempty" form:"totalCount,omitempty"` //Specifies the total number of backup objects that match the filter and
}

/*
 * Structure for the custom type Notifications
 */
type Notifications struct {
    Count            *int64          `json:"count,omitempty" form:"count,omitempty"` //Notification Count.
    NotificationList []*TaskNotification `json:"notificationList,omitempty" form:"notificationList,omitempty"` //Notification list.
    UnreadCount      *int64          `json:"unreadCount,omitempty" form:"unreadCount,omitempty"` //Unread Notification Count.
}

/*
 * Structure for the custom type NodeSystemDiskInfo
 */
type NodeSystemDiskInfo struct {
    DevicePath      *string         `json:"devicePath,omitempty" form:"devicePath,omitempty"` //DevicePath is the device path of the disk.
    Id              *int64          `json:"id,omitempty" form:"id,omitempty"` //Id is the id of the disk.
    Offline         *bool           `json:"offline,omitempty" form:"offline,omitempty"` //Offline specifies whether a disk is marked offline.
}

/*
 * Structure for the custom type NodeStatistics
 */
type NodeStatistics struct {
    Id              *int64          `json:"id,omitempty" form:"id,omitempty"` //Id is the Id of the Node.
    UsagePerfStats  UsageAndPerformanceStatistics `json:"usagePerfStats,omitempty" form:"usagePerfStats,omitempty"` //Provides usage and performance statistics
}

/*
 * Structure for the custom type NodeHardwareInformation
 */
type NodeHardwareInformation struct {
    Cpu             *string         `json:"cpu,omitempty" form:"cpu,omitempty"` //Cpu provides the information regarding the CPU.
    MemorySizeBytes *int64          `json:"memorySizeBytes,omitempty" form:"memorySizeBytes,omitempty"` //MemorySizeBytes provides the memory size in bytes.
    Network         *string         `json:"network,omitempty" form:"network,omitempty"` //Network provides the information regarding the network cards.
}

/*
 * Structure for the custom type NetAppClusterInformation
 */
type NetAppClusterInformation struct {
    ContactInfo     *string         `json:"contactInfo,omitempty" form:"contactInfo,omitempty"` //Specifies information about the contact for the NetApp cluster
    Location        *string         `json:"location,omitempty" form:"location,omitempty"` //Specifies where this NetApp cluster is located.
    SerialNumber    *string         `json:"serialNumber,omitempty" form:"serialNumber,omitempty"` //Specifies the serial number of the NetApp cluster in the
}

/*
 * Structure for the custom type MountVolumeResultDetails
 */
type MountVolumeResultDetails struct {
    MountError      Error           `json:"mountError,omitempty" form:"mountError,omitempty"` //Specifies the cause of the mount failure if the mounting of a
    MountPoint      *string         `json:"mountPoint,omitempty" form:"mountPoint,omitempty"` //Specifies the mount point where the volume is mounted.
    VolumeName      *string         `json:"volumeName,omitempty" form:"volumeName,omitempty"` //Specifies the name of the original volume.
}

/*
 * Structure for the custom type MetricValue
 */
type MetricValue struct {
    MetricName      *string         `json:"metricName,omitempty" form:"metricName,omitempty"` //Specifies the metric name.
    TimestampMsecs  *int64          `json:"timestampMsecs,omitempty" form:"timestampMsecs,omitempty"` //Specifies the creation time of a data point as a Unix epoch Timestamp
    Value           Value           `json:"value,omitempty" form:"value,omitempty"` //Specifies a data type and data field used to store data.
}

/*
 * Structure for the custom type MetricDataPoint
 */
type MetricDataPoint struct {
    Data            Data            `json:"data,omitempty" form:"data,omitempty"` //Specifies the fields to store data of a given type.
    TimestampMsecs  *int64          `json:"timestampMsecs,omitempty" form:"timestampMsecs,omitempty"` //Specifies a timestamp when the metric data point was captured.
}

/*
 * Structure for the custom type MetricDataBlock
 */
type MetricDataBlock struct {
    DataPointVec    []*MetricDataPoint `json:"dataPointVec,omitempty" form:"dataPointVec,omitempty"` //Array of Data Points.
    MetricName      *string         `json:"metricName,omitempty" form:"metricName,omitempty"` //Specifies the name of a metric such as 'kDiskAwaitTimeMsecs'.
    Type            *int64          `json:"type,omitempty" form:"type,omitempty"` //Specifies the data type of the data points.
}

/*
 * Structure for the custom type MapReduceInstanceInputParam
 */
type MapReduceInstanceInputParam struct {
    Key             *string         `json:"key,omitempty" form:"key,omitempty"` //TODO: Write general description for this field
    Value           *string         `json:"value,omitempty" form:"value,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type MapReduceInstanceWrapper
 */
type MapReduceInstanceWrapper struct {
    LogPath            *string         `json:"logPath,omitempty" form:"logPath,omitempty"` //LogPath is the path of the log files for the MR instance run.
    MrInstance         MapReduceInstance `json:"mrInstance,omitempty" form:"mrInstance,omitempty"` //Information about a Map reduce instance. An instance can be run only once.
    OutputFilePathList *[]string       `json:"outputFilePathList,omitempty" form:"outputFilePathList,omitempty"` //OutputFilePathList is the list containing the output files path suffix
}

/*
 * Structure for the custom type MapReduceInfoRequiredProperty
 */
type MapReduceInfoRequiredProperty struct {
    DefaultValue    *string         `json:"defaultValue,omitempty" form:"defaultValue,omitempty"` //Default Value of the property.
    Description     *string         `json:"description,omitempty" form:"description,omitempty"` //Description of this property
    IsRequired      *bool           `json:"isRequired,omitempty" form:"isRequired,omitempty"` //Whether the property is required or optional.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Name of the property.
}

/*
 * Structure for the custom type LogicalStatistics
 */
type LogicalStatistics struct {
    TotalLogicalUsageBytes *int64          `json:"totalLogicalUsageBytes,omitempty" form:"totalLogicalUsageBytes,omitempty"` //Provides the logical usage as computed by the Cohesity Cluster.
}

/*
 * Structure for the custom type LockFileParameters
 */
type LockFileParameters struct {
    ExpiryTimestampMsecs *int64          `json:"expiryTimestampMsecs,omitempty" form:"expiryTimestampMsecs,omitempty"` //Specifies a definite timestamp in milliseconds for retaining the file, or
    Path                 *string         `json:"path,omitempty" form:"path,omitempty"` //Specifies the file path that needs to be locked.
}

/*
 * Structure for the custom type LegalHoldings
 */
type LegalHoldings struct {
    HoldForLegalPurpose *bool           `json:"holdForLegalPurpose,omitempty" form:"holdForLegalPurpose,omitempty"` //Specifies whether the source is put on legal hold or not.
    ProtectionSourceId  *int64          `json:"protectionSourceId,omitempty" form:"protectionSourceId,omitempty"` //Specifies an Protection Source Id in the snapshot.
}

/*
 * Structure for the custom type LatestProtectionJobRunInformation
 */
type LatestProtectionJobRunInformation struct {
    LatestSnapshotInfo LatestProtectionRun `json:"latestSnapshotInfo,omitempty" form:"latestSnapshotInfo,omitempty"` //Specifies the information about the latest Protection Run.
    LocationName       *string         `json:"locationName,omitempty" form:"locationName,omitempty"` //Specifies the name of location that the object is backedup to.
    NumSnapshots       *int64          `json:"numSnapshots,omitempty" form:"numSnapshots,omitempty"` //Specifies of number of successful snapshots.
}

/*
 * Structure for the custom type LatencyThresholds
 */
type LatencyThresholds struct {
    ActiveTaskMsecs *int64          `json:"activeTaskMsecs,omitempty" form:"activeTaskMsecs,omitempty"` //If the latency of a datastore is above this value, existing backup tasks
    NewTaskMsecs    *int64          `json:"newTaskMsecs,omitempty" form:"newTaskMsecs,omitempty"` //If the latency of a datastore is above this value, then new backup tasks
}

/*
 * Structure for the custom type KeyValuePair
 */
type KeyValuePair struct {
    Key             *string         `json:"key,omitempty" form:"key,omitempty"` //Specifies the name of the key.
    Value           Value           `json:"value,omitempty" form:"value,omitempty"` //Specifies a data type and data field used to store data.
}

/*
 * Structure for the custom type DEPRECATEDIn50
 */
type DEPRECATEDIn50 struct {
    BackupPolicy            DEPRECATEDIn501 `json:"backupPolicy,omitempty" form:"backupPolicy,omitempty"` //If a backup does not get a chance to when it's due (either due to the system
    SnapshotTargetPolicyVec []*SnapshotTargetPolicyProto `json:"snapshotTargetPolicyVec,omitempty" form:"snapshotTargetPolicyVec,omitempty"` //Specifies additional policies that can be used to copy snapshots created
}

/*
 * Structure for the custom type IsilonSMBMountPoint
 */
type IsilonSMBMountPoint struct {
    AccessZoneId    *int64          `json:"accessZoneId,omitempty" form:"accessZoneId,omitempty"` //Specifies the Access Zone Id.
    Description     *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies the description of the NFS mount point.
    ShareName       *string         `json:"shareName,omitempty" form:"shareName,omitempty"` //Specifies the name of the SMB/CIFS share.
}

/*
 * Structure for the custom type IsilonNFSMountPoint
 */
type IsilonNFSMountPoint struct {
    AccessZoneName  *string         `json:"accessZoneName,omitempty" form:"accessZoneName,omitempty"` //Specifies the Access Zone name.
    Description     *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies the description of the NFS mount point.
    Id              *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies the Id of the NFS export.
}

/*
 * Structure for the custom type IsilonCluster
 */
type IsilonCluster struct {
    Description     *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies the description of an Isilon Cluster.
    Guid            *string         `json:"guid,omitempty" form:"guid,omitempty"` //Specifies a globally unique id of an Isilon Cluster.
}

/*
 * Structure for the custom type IsilonAccessZone
 */
type IsilonAccessZone struct {
    Id              *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies the id of the access zone.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the access zone.
    Path            *string         `json:"path,omitempty" form:"path,omitempty"` //Specifies the path of the access zone in ifs. This should include the
}

/*
 * Structure for the custom type ISCSISANPort
 */
type ISCSISANPort struct {
    IpAddress       *string         `json:"ipAddress,omitempty" form:"ipAddress,omitempty"` //Specifies the IP address of the SAN port.
    Iqn             *string         `json:"iqn,omitempty" form:"iqn,omitempty"` //Specifies the qualified name of the iSCSI port (IQN).
    TcpPort         *int64          `json:"tcpPort,omitempty" form:"tcpPort,omitempty"` //Specifies the listening port(tcp) of the SAN port.
}

/*
 * Structure for the custom type InterfaceGroup
 */
type InterfaceGroup struct {
    Id                  *int64          `json:"id,omitempty" form:"id,omitempty"` //Interface group Id.
    ModelInterfaceLists []*ProductModelAndInterfaceTuple `json:"modelInterfaceLists,omitempty" form:"modelInterfaceLists,omitempty"` //Specifies the product model and interface lists.
    Name                *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the interface group.
}

/*
 * Structure for the custom type InputSpecFileTimeFilter
 */
type InputSpecFileTimeFilter struct {
    ChangeTimeRangeEndSecs   *int64          `json:"changeTimeRangeEndSecs,omitempty" form:"changeTimeRangeEndSecs,omitempty"` //End of file's change time range.
    ChangeTimeRangeStartSecs *int64          `json:"changeTimeRangeStartSecs,omitempty" form:"changeTimeRangeStartSecs,omitempty"` //Start of file's change time range.
}

/*
 * Structure for the custom type InputSelectorSelectsTheFilesToMapOver
 */
type InputSelectorSelectsTheFilesToMapOver struct {
    FilesSelector   InputSpecInputFilesSelector `json:"filesSelector,omitempty" form:"filesSelector,omitempty"` //If mapper is going to run over files on SnapFS, this selects the input
    OnNfsFiles      *bool           `json:"onNfsFiles,omitempty" form:"onNfsFiles,omitempty"` //Selects whether input is files inside vmdks or files on NFS. One of
    VmSelector      InputSpecInputVMsSelector `json:"vmSelector,omitempty" form:"vmSelector,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type IndexingPolicyProto
 */
type IndexingPolicyProto struct {
    AllowPrefixes   *[]string       `json:"allowPrefixes,omitempty" form:"allowPrefixes,omitempty"` //List of directory prefixes to allow for indexing.
    DenyPrefixes    *[]string       `json:"denyPrefixes,omitempty" form:"denyPrefixes,omitempty"` //List of directory prefixes to filter out.
    DisableIndexing *bool           `json:"disableIndexing,omitempty" form:"disableIndexing,omitempty"` //If this field is set all the files in the VM will be filtered.
}

/*
 * Structure for the custom type IndexingPolicy
 */
type IndexingPolicy struct {
    AllowPrefixes   *[]string       `json:"allowPrefixes,omitempty" form:"allowPrefixes,omitempty"` //Array of Indexed Directories.
    DenyPrefixes    *[]string       `json:"denyPrefixes,omitempty" form:"denyPrefixes,omitempty"` //Array of Excluded Directories.
    DisableIndexing *bool           `json:"disableIndexing,omitempty" form:"disableIndexing,omitempty"` //Specifies if the files found in an Object (such as a VM) should be
}

/*
 * Structure for the custom type UpdateIDMappingInformationRequest
 */
type UpdateIDMappingInformationRequest struct {
    FallbackUserIdMappingInfo UserIDMapping   `json:"fallbackUserIdMappingInfo,omitempty" form:"fallbackUserIdMappingInfo,omitempty"` //Specifies how the Unix and Windows users are mapped in an Active Directory.
    UnixRootSid               *string         `json:"unixRootSid,omitempty" form:"unixRootSid,omitempty"` //Specifies the SID of the Active Directory domain user to be mapped to
    UserIdMappingInfo         UserIDMapping   `json:"userIdMappingInfo,omitempty" form:"userIdMappingInfo,omitempty"` //Specifies how the Unix and Windows users are mapped in an Active Directory.
}

/*
 * Structure for the custom type GroupInformation
 */
type GroupInformation struct {
    Domain          *string         `json:"domain,omitempty" form:"domain,omitempty"` //Specifies domain name of the user.
    GroupName       *string         `json:"groupName,omitempty" form:"groupName,omitempty"` //Specifies group name of the group.
    Sid             *string         `json:"sid,omitempty" form:"sid,omitempty"` //Specifies unique Security ID (SID) of the user.
}

/*
 * Structure for the custom type DeleteGroupsRequest
 */
type DeleteGroupsRequest struct {
    Domain          *string         `json:"domain,omitempty" form:"domain,omitempty"` //Specifies the domain associated with the groups to delete.
    Names           *[]string       `json:"names,omitempty" form:"names,omitempty"` //Array of Groups.
}

/*
 * Structure for the custom type GranularityBucket
 */
type GranularityBucket struct {
    Granularity     *int64          `json:"granularity,omitempty" form:"granularity,omitempty"` //The base time period granularity that determines the frequency at which
    Multiplier      *int64          `json:"multiplier,omitempty" form:"multiplier,omitempty"` //A factor to multiply the granularity by.
}

/*
 * Structure for the custom type GoogleAccountInformation
 */
type GoogleAccountInformation struct {
    AccountId       *string         `json:"accountId,omitempty" form:"accountId,omitempty"` //Specifies the Account Id assigned by Google.
    UserId          *string         `json:"userId,omitempty" form:"userId,omitempty"` //Specifies the User Id assigned by Google.
}

/*
 * Structure for the custom type GetViewsResult
 */
type GetViewsResult struct {
    LastResult      *bool           `json:"lastResult,omitempty" form:"lastResult,omitempty"` //If false, more Views are available to return. If the number of
    Views           []*View         `json:"views,omitempty" form:"views,omitempty"` //Array of Views.
}

/*
 * Structure for the custom type GetViewsAndAliasesByShareResult
 */
type GetViewsAndAliasesByShareResult struct {
    PaginationCookie *string         `json:"paginationCookie,omitempty" form:"paginationCookie,omitempty"` //If set, i.e. there are more results to display, use this value to get
    SharesList       []*Share        `json:"sharesList,omitempty" form:"sharesList,omitempty"` //Array of Views and Aliases by Share name.
}

/*
 * Structure for the custom type GetRegistrationInformationResponse
 */
type GetRegistrationInformationResponse struct {
    RootNodes       []*RegistrationAndProtectionInformation `json:"rootNodes,omitempty" form:"rootNodes,omitempty"` //Specifies the registration, protection and permission information of either
    Stats           ProtectionSummary `json:"stats,omitempty" form:"stats,omitempty"` //Specifies the sum of all the stats of protection of Protection Sources
    StatsByEnv      []*ProtectionSummaryByEnvironment `json:"statsByEnv,omitempty" form:"statsByEnv,omitempty"` //Specifies the breakdown of the stats by environment
}

/*
 * Structure for the custom type GetMRJarUploadPathResult
 */
type GetMRJarUploadPathResult struct {
    Error           ErrorProto      `json:"error,omitempty" form:"error,omitempty"` //TODO: Write general description for this field
    JarUploadPath   *string         `json:"jarUploadPath,omitempty" form:"jarUploadPath,omitempty"` //Path where Jars can be uploaded by Iris.
}

/*
 * Structure for the custom type GetAlertTypesParams
 */
type GetAlertTypesParams struct {
    AlertCategoryList *[]AlertCategoryListEnum `json:"alertCategoryList,omitempty" form:"alertCategoryList,omitempty"` //Specifies a list of Alert Categories to filter alert types by.
}

/*
 * Structure for the custom type FullSnapshotInfo
 */
type FullSnapshotInfo struct {
    RestoreInfo     FullSnapshotInformation `json:"restoreInfo,omitempty" form:"restoreInfo,omitempty"` //Specifies the info regarding a full SQL snapshot.
    SnapshotTarget  []*SnapshotTarget1 `json:"snapshotTarget,omitempty" form:"snapshotTarget,omitempty"` //Specifies the location holding snapshot copies that may be used for
}

/*
 * Structure for the custom type FlashBladeSMBInformation
 */
type FlashBladeSMBInformation struct {
    AclMode         AclModeEnum     `json:"aclMode,omitempty" form:"aclMode,omitempty"` //ACL mode for this SMB share.
}

/*
 * Structure for the custom type FlashBladeNFSInformation
 */
type FlashBladeNFSInformation struct {
    ExportRules     *string         `json:"exportRules,omitempty" form:"exportRules,omitempty"` //Specifies NFS protocol export rules. Rules are in the form host(options).
}

/*
 * Structure for the custom type FlashBladeNetworkInterface
 */
type FlashBladeNetworkInterface struct {
    IpAddress       *string         `json:"ipAddress,omitempty" form:"ipAddress,omitempty"` //Specifies the IP address of the Pure Storage FlashBlade Array.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the network interface.
    Vlan            *int64          `json:"vlan,omitempty" form:"vlan,omitempty"` //Specifies the id of the VLAN network of the Pure Storage FlashBlade Array.
}

/*
 * Structure for the custom type FixedUnixIDMapping
 */
type FixedUnixIDMapping struct {
    Gid             *int64          `json:"gid,omitempty" form:"gid,omitempty"` //Specifies the fixed Unix GID, when mapping type is set to kFixed.
    Uid             *int64          `json:"uid,omitempty" form:"uid,omitempty"` //Specifies the fixed Unix UID, when mapping type is set to kFixed.
}

/*
 * Structure for the custom type FilteringPolicyProto
 */
type FilteringPolicyProto struct {
    AllowFilters    *[]string       `json:"allowFilters,omitempty" form:"allowFilters,omitempty"` //List of filters to allow matched objects for backup.
    DenyFilters     *[]string       `json:"denyFilters,omitempty" form:"denyFilters,omitempty"` //List of filters to deny matched objects for backup.
}

/*
 * Structure for the custom type FileAndFoldersInformation
 */
type FileAndFoldersInformation struct {
    AbsolutePath    *string         `json:"absolutePath,omitempty" form:"absolutePath,omitempty"` //AbsolutePath specifies the absolute path of the specified file or folder.
    IsDirectory     *bool           `json:"isDirectory,omitempty" form:"isDirectory,omitempty"` //IsDirectory specifies if specified object is a directory or not.
}

/*
 * Structure for the custom type FilerAuditLogConfiguration
 */
type FilerAuditLogConfiguration struct {
    Enabled             bool            `json:"enabled" form:"enabled"` //Specifies if filer audit logging is enabled on the Cohesity Cluster.
    RetentionPeriodDays int64           `json:"retentionPeriodDays" form:"retentionPeriodDays"` //Specifies the number of days to keep (retain) the filer audit logs.
}

/*
 * Structure for the custom type FilenamePatternToDirectory
 */
type FilenamePatternToDirectory struct {
    Directory       *string         `json:"directory,omitempty" form:"directory,omitempty"` //Specifies the directory where to keep the files matching the pattern.
    FilenamePattern *string         `json:"filenamePattern,omitempty" form:"filenamePattern,omitempty"` //Specifies a pattern to be matched with filenames. This can be a
}

/*
 * Structure for the custom type FileFolderVersion
 */
type FileFolderVersion struct {
    ModifiedTimeUsecs *int64          `json:"modifiedTimeUsecs,omitempty" form:"modifiedTimeUsecs,omitempty"` //Specifies the time when the file or folder was last modified.
    SizeBytes         *int64          `json:"sizeBytes,omitempty" form:"sizeBytes,omitempty"` //Specifies the size of the file or folder (in bytes)
    Snapshots         []*SnapshotAttempt `json:"snapshots,omitempty" form:"snapshots,omitempty"` //Array of Snapshots.
}

/*
 * Structure for the custom type FileFolderSearchResult1
 */
type FileFolderSearchResult1 struct {
    Files           []*FileFolderSearchResult `json:"files,omitempty" form:"files,omitempty"` //Array of Files and Folders.
    TotalCount      *int64          `json:"totalCount,omitempty" form:"totalCount,omitempty"` //Specifies the total number of files and folders that match the filter and
}

/*
 * Structure for the custom type FileRestoreInformation
 */
type FileRestoreInformation struct {
    Error            Error           `json:"error,omitempty" form:"error,omitempty"` //Details about the Error.
    Filename         *string         `json:"filename,omitempty" form:"filename,omitempty"` //Specifies the path of the file/directory.
    FilesystemVolume FilesystemVolume `json:"filesystemVolume,omitempty" form:"filesystemVolume,omitempty"` //Specifies information about a filesystem volume.
    IsFolder         *bool           `json:"isFolder,omitempty" form:"isFolder,omitempty"` //Specifies whether the file path is a folder.
}

/*
 * Structure for the custom type FileOrDirectoryToProtect
 */
type FileOrDirectoryToProtect struct {
    BackupFilePath    *string         `json:"backupFilePath,omitempty" form:"backupFilePath,omitempty"` //Specifies absolute path to a file or a directory in a Physical Server
    ExcludedFilePaths *[]string       `json:"excludedFilePaths,omitempty" form:"excludedFilePaths,omitempty"` //Array of Excluded File Paths.
    SkipNestedVolumes *bool           `json:"skipNestedVolumes,omitempty" form:"skipNestedVolumes,omitempty"` //Specifies if any subdirectories under backupFilePath, where local or
}

/*
 * Structure for the custom type FileOrDirectoryPathFilter
 */
type FileOrDirectoryPathFilter struct {
    ExcludeFilters  *[]string       `json:"excludeFilters,omitempty" form:"excludeFilters,omitempty"` //Array of Excluded File Path Filters.
    ProtectFilters  *[]string       `json:"protectFilters,omitempty" form:"protectFilters,omitempty"` //Array of Protected File Path Filters.
}

/*
 * Structure for the custom type FileExtensionFilter
 */
type FileExtensionFilter struct {
    FileExtensionsList *[]string       `json:"fileExtensionsList,omitempty" form:"fileExtensionsList,omitempty"` //The list of file extensions to apply
    IsEnabled          *bool           `json:"isEnabled,omitempty" form:"isEnabled,omitempty"` //If set, it enables the file extension filter
    Mode               ModeEnum        `json:"mode,omitempty" form:"mode,omitempty"` //The mode applied to the list of file extensions
}

/*
 * Structure for the custom type ErrorProto
 */
type ErrorProto struct {
    ErrorMsg        *string         `json:"errorMsg,omitempty" form:"errorMsg,omitempty"` //An optional detail.
    Type            *int64          `json:"type,omitempty" form:"type,omitempty"` //Error.
}

/*
 * Structure for the custom type MetricUnit
 */
type MetricUnit struct {
    Type            *int64          `json:"type,omitempty" form:"type,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type KeyValuePair1
 */
type KeyValuePair1 struct {
    KeyName         *string         `json:"keyName,omitempty" form:"keyName,omitempty"` //Specifies the name of a key.
    ValueType       *int64          `json:"valueType,omitempty" form:"valueType,omitempty"` //Specifies the type of the value that is associated with the key.
}

/*
 * Structure for the custom type AttributesDescriptor
 */
type AttributesDescriptor struct {
    AttributeVec          []*KeyValuePair1 `json:"attributeVec,omitempty" form:"attributeVec,omitempty"` //Array of Attributes.
    KeyAttributeNameIndex *int64          `json:"keyAttributeNameIndex,omitempty" form:"keyAttributeNameIndex,omitempty"` //Specifies the attribute to use as a unique identifier for the entity.
}

/*
 * Structure for the custom type Entity
 */
type Entity struct {
    AttributeVec    []*KeyValuePair `json:"attributeVec,omitempty" form:"attributeVec,omitempty"` //Array of Attributes.
    EntityId        EntityIdentifier `json:"entityId,omitempty" form:"entityId,omitempty"` //Specifies a unique identifier for the entity.
    LatestMetricVec []*MetricValue  `json:"latestMetricVec,omitempty" form:"latestMetricVec,omitempty"` //Array of Metric Statistics.
}

/*
 * Structure for the custom type EntityPermissionInformation
 */
type EntityPermissionInformation struct {
    EntityId        *int64          `json:"entityId,omitempty" form:"entityId,omitempty"` //Specifies the entity id.
    Groups          []*GroupInformation `json:"groups,omitempty" form:"groups,omitempty"` //Specifies groups that have access to entity in case of restricted user.
    Tenant          TenantInformation `json:"tenant,omitempty" form:"tenant,omitempty"` //Specifies struct with basic tenant details.
    Users           []*UserInformation `json:"users,omitempty" form:"users,omitempty"` //Specifies users that have access to entity in case of restricted user.
}

/*
 * Structure for the custom type EmailDeliveryTarget
 */
type EmailDeliveryTarget struct {
    EmailAddress    *string         `json:"emailAddress,omitempty" form:"emailAddress,omitempty"` //TODO: Write general description for this field
    Locale          *string         `json:"locale,omitempty" form:"locale,omitempty"` //Specifies the language in which the emails sent to the above defined
}

/*
 * Structure for the custom type DownloadFilesAndFoldersParameters
 */
type DownloadFilesAndFoldersParameters struct {
    FilesAndFoldersInfo []*FileAndFoldersInformation `json:"filesAndFoldersInfo,omitempty" form:"filesAndFoldersInfo,omitempty"` //Specifies the absolute paths for list of files and folders to download.
    Name                string          `json:"name" form:"name"` //Specifies the name of the Download Task. This field must be set and must
    SourceObjectInfo    RestoreObject   `json:"sourceObjectInfo,omitempty" form:"sourceObjectInfo,omitempty"` //Specifies an object to recover or clone or an object to restore files
}

/*
 * Structure for the custom type DiskUnit
 */
type DiskUnit struct {
    BusNumber       *int64          `json:"busNumber,omitempty" form:"busNumber,omitempty"` //Specifies the Id of the controller bus that controls the disk.
    ControllerType  *string         `json:"controllerType,omitempty" form:"controllerType,omitempty"` //Specifies the controller type like SCSI, or IDE etc.
    UnitNumber      *int64          `json:"unitNumber,omitempty" form:"unitNumber,omitempty"` //Specifies the disk file name. This is the VMDK name and not the
}

/*
 * Structure for the custom type DiskBlock
 */
type DiskBlock struct {
    LengthBytes     *int64          `json:"lengthBytes,omitempty" form:"lengthBytes,omitempty"` //Specifies the length of the block in bytes.
    OffsetBytes     *int64          `json:"offsetBytes,omitempty" form:"offsetBytes,omitempty"` //Specifies the offset of the block (in bytes) from the beginning
}

/*
 * Structure for the custom type DeviceNode
 */
type DeviceNode struct {
    IntermediateNode DeviceTree      `json:"intermediateNode,omitempty" form:"intermediateNode,omitempty"` //TODO: Write general description for this field
    LeafNode         FilePartitionBlock `json:"leafNode,omitempty" form:"leafNode,omitempty"` //Defines a leaf node of a device tree. This refers to a logical
}

/*
 * Structure for the custom type DeployVMsToGCPParams
 */
type DeployVMsToGCPParams struct {
    ProjectId       Entity          `json:"projectId,omitempty" form:"projectId,omitempty"` //Specifies the attributes and the latest statistics about an entity.
    Region          Entity          `json:"region,omitempty" form:"region,omitempty"` //Specifies the attributes and the latest statistics about an entity.
    Subnet          Entity          `json:"subnet,omitempty" form:"subnet,omitempty"` //Specifies the attributes and the latest statistics about an entity.
    Zone            Entity          `json:"zone,omitempty" form:"zone,omitempty"` //Specifies the attributes and the latest statistics about an entity.
}

/*
 * Structure for the custom type DeployVMsToCloudParams
 */
type DeployVMsToCloudParams struct {
    DeployVmsToAwsParams   DeployVMsToAWSParams `json:"deployVmsToAwsParams,omitempty" form:"deployVmsToAwsParams,omitempty"` //Contains AWS specific information needed to identify various resources
    DeployVmsToAzureParams DeployVMsToAzureParams `json:"deployVmsToAzureParams,omitempty" form:"deployVmsToAzureParams,omitempty"` //Contains Azure specific information needed to identify various resources
    DeployVmsToGcpParams   DeployVMsToGCPParams `json:"deployVmsToGcpParams,omitempty" form:"deployVmsToGcpParams,omitempty"` //Contains GCP specific information needed to identify various resources
}

/*
 * Structure for the custom type DeleteViewUsersQuotaParameters
 */
type DeleteViewUsersQuotaParameters struct {
    DeleteAll       *bool           `json:"deleteAll,omitempty" form:"deleteAll,omitempty"` //Delete all existing user quota override policies.
    UserIds         []*UserID       `json:"userIds,omitempty" form:"userIds,omitempty"` //The user ids whose policy needs to be deleted.
    ViewName        *string         `json:"viewName,omitempty" form:"viewName,omitempty"` //View name of input view.
}

/*
 * Structure for the custom type DeleteRouteParameters
 */
type DeleteRouteParameters struct {
    DestNetwork     *string         `json:"destNetwork,omitempty" form:"destNetwork,omitempty"` //Destination network.
    IfName          *string         `json:"ifName,omitempty" form:"ifName,omitempty"` //Specifies the network interfaces name to use for communicating with the
    IfaceGroupName  *string         `json:"ifaceGroupName,omitempty" form:"ifaceGroupName,omitempty"` //Specifies the network interfaces group or vlan interface group to
}

/*
 * Structure for the custom type DatabaseFileInformation
 */
type DatabaseFileInformation struct {
    FileType        FileTypeEnum    `json:"fileType,omitempty" form:"fileType,omitempty"` //Specifies the format type of the file that SQL database stores the data.
    FullPath        *string         `json:"fullPath,omitempty" form:"fullPath,omitempty"` //Specifies the full path of the database file on the SQL host machine.
    SizeBytes       *int64          `json:"sizeBytes,omitempty" form:"sizeBytes,omitempty"` //Specifies the last known size of the database file.
}

/*
 * Structure for the custom type DatastoreInformation
 */
type DatastoreInformation struct {
    Capacity        *int64          `json:"capacity,omitempty" form:"capacity,omitempty"` //Specifies the capacity of the datastore in bytes.
    FreeSpace       *int64          `json:"freeSpace,omitempty" form:"freeSpace,omitempty"` //Specifies the available space on the datastore in bytes.
}

/*
 * Structure for the custom type DataTransferToVaultPerProtectionJob
 */
type DataTransferToVaultPerProtectionJob struct {
    NumLogicalBytesTransferred  *int64          `json:"numLogicalBytesTransferred,omitempty" form:"numLogicalBytesTransferred,omitempty"` //Specifies the total number of logical bytes that are transferred
    NumPhysicalBytesTransferred *int64          `json:"numPhysicalBytesTransferred,omitempty" form:"numPhysicalBytesTransferred,omitempty"` //Specifies the total number of physical bytes that are transferred
    ProtectionJobName           *string         `json:"protectionJobName,omitempty" form:"protectionJobName,omitempty"` //Specifies the name of the Protection Job.
}

/*
 * Structure for the custom type DashboardResponse
 */
type DashboardResponse struct {
    Dashboard       Dashboard       `json:"dashboard,omitempty" form:"dashboard,omitempty"` //Data shown on Dashboard.
    Dashboards      []*Dashboard    `json:"dashboards,omitempty" form:"dashboards,omitempty"` //Specifies a list of dashboards of all the clusters in the SPOG setup if
}

/*
 * Structure for the custom type CustomUnixIDAttributes
 */
type CustomUnixIDAttributes struct {
    GidAttrName     *string         `json:"gidAttrName,omitempty" form:"gidAttrName,omitempty"` //Specifies the custom field name in Active Directory user properties to get
    UidAttrName     *string         `json:"uidAttrName,omitempty" form:"uidAttrName,omitempty"` //Specifies the custom field name in Active Directory user properties to get
}

/*
 * Structure for the custom type Credentials
 */
type Credentials struct {
    Password        *string         `json:"password,omitempty" form:"password,omitempty"` //Specifies password of the username to access the target source.
    Username        *string         `json:"username,omitempty" form:"username,omitempty"` //Specifies username to access the target source.
}

/*
 * Structure for the custom type BackupPolicyProtoMonthlySchedule
 */
type BackupPolicyProtoMonthlySchedule struct {
    Count           *int64          `json:"count,omitempty" form:"count,omitempty"` //Count of the day on which to perform the backup (look above for a more
    Day             *int64          `json:"day,omitempty" form:"day,omitempty"` //The day of the month the backup is to be performed.
    Time            *time.Time      `json:"time,omitempty" form:"time,omitempty"` //A message to encapusulate time of a day. Users of this proto will have to
}

/*
 * Structure for the custom type BackupPolicyProtoExclusionTimeRange
 */
type BackupPolicyProtoExclusionTimeRange struct {
    Day             *int64          `json:"day,omitempty" form:"day,omitempty"` //If the day is not set, the time range applies to all days.
    EndTime         *time.Time      `json:"endTime,omitempty" form:"endTime,omitempty"` //A message to encapusulate time of a day. Users of this proto will have to
    StartTime       *time.Time      `json:"startTime,omitempty" form:"startTime,omitempty"` //A message to encapusulate time of a day. Users of this proto will have to
}

/*
 * Structure for the custom type BackupPolicyProtoDailySchedule
 */
type BackupPolicyProtoDailySchedule struct {
    Days            *[]int64        `json:"days,omitempty" form:"days,omitempty"` //The days of the week backup must be performed. If no days are specified,
    Time            *time.Time      `json:"time,omitempty" form:"time,omitempty"` //A message to encapusulate time of a day. Users of this proto will have to
}

/*
 * Structure for the custom type BackupPolicyProtoContinuousSchedule
 */
type BackupPolicyProtoContinuousSchedule struct {
    BackupIntervalMins *int64          `json:"backupIntervalMins,omitempty" form:"backupIntervalMins,omitempty"` //If this field is set, backups will be performed periodically every
    ExclusionRanges    []*BackupPolicyProtoExclusionTimeRange `json:"exclusionRanges,omitempty" form:"exclusionRanges,omitempty"` //Do not start backups in these time-ranges. It's possible for a
}

/*
 * Structure for the custom type BackupJobProtoExclusionTimeRange
 */
type BackupJobProtoExclusionTimeRange struct {
    Day             *int64          `json:"day,omitempty" form:"day,omitempty"` //If the day is not set, the time range applies to all days.
    EndTime         *time.Time      `json:"endTime,omitempty" form:"endTime,omitempty"` //A message to encapusulate time of a day. Users of this proto will have to
    StartTime       *time.Time      `json:"startTime,omitempty" form:"startTime,omitempty"` //A message to encapusulate time of a day. Users of this proto will have to
}

/*
 * Structure for the custom type DRToCloudParameters
 */
type DRToCloudParameters struct {
    NeedToFailOver  *bool           `json:"needToFailOver,omitempty" form:"needToFailOver,omitempty"` //Whether the objects in this job will be failed over to cloud.
}

/*
 * Structure for the custom type AzureManagedDiskParams
 */
type AzureManagedDiskParams struct {
    DataDisksSkuType *int64          `json:"dataDisksSkuType,omitempty" form:"dataDisksSkuType,omitempty"` //SKU type for data disks.
    OsDiskSkuType    *int64          `json:"osDiskSkuType,omitempty" form:"osDiskSkuType,omitempty"` //SKU type for OS disk.
}

/*
 * Structure for the custom type AzureCloudCredentials
 */
type AzureCloudCredentials struct {
    StorageAccessKey   *string         `json:"storageAccessKey,omitempty" form:"storageAccessKey,omitempty"` //Specifies the access key to use when accessing a storage tier
    StorageAccountName *string         `json:"storageAccountName,omitempty" form:"storageAccountName,omitempty"` //Specifies the account name to use when accessing a storage tier
    TierType           TierType1Enum   `json:"tierType,omitempty" form:"tierType,omitempty"` //Specifies the storage class of Azure.
}

/*
 * Structure for the custom type AuditLogsTile
 */
type AuditLogsTile struct {
    ClusterAuditLogs []*ClusterAuditLog `json:"clusterAuditLogs,omitempty" form:"clusterAuditLogs,omitempty"` //Array of Cluster Audit Logs.
    TotalCount       *int64          `json:"totalCount,omitempty" form:"totalCount,omitempty"` //Specifies the total number of logs that match the specified
}

/*
 * Structure for the custom type ArchivalTarget1
 */
type ArchivalTarget1 struct {
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //The name of the archival target.
    Type            *int64          `json:"type,omitempty" form:"type,omitempty"` //The type of the archival target.
    VaultId         *int64          `json:"vaultId,omitempty" form:"vaultId,omitempty"` //The id of the archival vault.
}

/*
 * Structure for the custom type ArchivalTarget
 */
type ArchivalTarget struct {
    VaultId         *int64          `json:"vaultId,omitempty" form:"vaultId,omitempty"` //Specifies the id of Archival Vault assigned by the Cohesity Cluster.
    VaultName       *string         `json:"vaultName,omitempty" form:"vaultName,omitempty"` //Name of the Archival Vault.
    VaultType       VaultTypeEnum   `json:"vaultType,omitempty" form:"vaultType,omitempty"` //Specifies the type of the Archival External Target such as 'kCloud',
}

/*
 * Structure for the custom type AthenaAppsConfiguration
 */
type AthenaAppsConfiguration struct {
    AllowExternalTraffic *bool           `json:"allowExternalTraffic,omitempty" form:"allowExternalTraffic,omitempty"` //Whether to allow pod external traffic.
    AppsMode             *int64          `json:"appsMode,omitempty" form:"appsMode,omitempty"` //The apps mode - kDisabled, kBareMetal, kVmOnly.
    AppsSubnet           Subnet          `json:"appsSubnet,omitempty" form:"appsSubnet,omitempty"` //Defines a Subnet (Subnetwork).
}

/*
 * Structure for the custom type ApplicationSpecialJobParameters
 */
type ApplicationSpecialJobParameters struct {
    ApplicationEntityIds *[]int64        `json:"applicationEntityIds,omitempty" form:"applicationEntityIds,omitempty"` //Array of Ids of Application Entities like SQL/Oracle instances, and
}

/*
 * Structure for the custom type ApplicationRunHistory
 */
type ApplicationRunHistory struct {
    AppInfo         MapReduceInformation `json:"appInfo,omitempty" form:"appInfo,omitempty"` //This will be used to encapsulate information about mapper and reducer only.
    MrInstances     []*MapReduceInstanceWrapper `json:"mrInstances,omitempty" form:"mrInstances,omitempty"` //InstancesWrapper is the slice containing the information about the map
}

/*
 * Structure for the custom type AppInstanceSettings
 */
type AppInstanceSettings struct {
    QosTier                 QosTierEnum     `json:"qosTier,omitempty" form:"qosTier,omitempty"` //Specifies QoSTier of the app instance.
    ReadViewPrivileges      ViewPrivileges  `json:"readViewPrivileges,omitempty" form:"readViewPrivileges,omitempty"` //ViewPrivileges specifies which views are allowed to be accessed by an app
    ReadWriteViewPrivileges ViewPrivileges  `json:"readWriteViewPrivileges,omitempty" form:"readWriteViewPrivileges,omitempty"` //ViewPrivileges specifies which views are allowed to be accessed by an app
}

/*
 * Structure for the custom type AlertingPolicyProto
 */
type AlertingPolicyProto struct {
    Emails                       *[]string       `json:"emails,omitempty" form:"emails,omitempty"` //The email addresses to send alerts to.
    Policy                       *int64          `json:"policy,omitempty" form:"policy,omitempty"` //'policy' is declared as int32 because ORing the enums will generate values
    RaiseObjectLevelFailureAlert *bool           `json:"raiseObjectLevelFailureAlert,omitempty" form:"raiseObjectLevelFailureAlert,omitempty"` //Raise per object alert for failures.
}

/*
 * Structure for the custom type AlertingConfig
 */
type AlertingConfig struct {
    EmailAddresses               *[]string       `json:"emailAddresses,omitempty" form:"emailAddresses,omitempty"` //Specifies additional email addresses where alert notifications (configured
    RaiseObjectLevelFailureAlert *bool           `json:"raiseObjectLevelFailureAlert,omitempty" form:"raiseObjectLevelFailureAlert,omitempty"` //Specifies the boolean to raise per object alert for failures.
}

/*
 * Structure for the custom type CreateAlertResolutionRequest
 */
type CreateAlertResolutionRequest struct {
    AlertIdList       *[]string       `json:"alertIdList,omitempty" form:"alertIdList,omitempty"` //Specifies list of alerts resolved by a Resolution, which are specified by
    ResolutionDetails AlertResolutionInfo `json:"resolutionDetails,omitempty" form:"resolutionDetails,omitempty"` //Short description and detailed notes about the Resolution.
}

/*
 * Structure for the custom type AlertResolutionInfo
 */
type AlertResolutionInfo struct {
    ResolutionDetails *string         `json:"resolutionDetails,omitempty" form:"resolutionDetails,omitempty"` //Specifies detailed notes about the Resolution.
    ResolutionSummary *string         `json:"resolutionSummary,omitempty" form:"resolutionSummary,omitempty"` //Specifies short description about the Resolution.
}

/*
 * Structure for the custom type AlertResolution1
 */
type AlertResolution1 struct {
    AlertIdList       *[]string       `json:"alertIdList,omitempty" form:"alertIdList,omitempty"` //Specifies list of Alerts resolved by a Resolution, which are specified by
    ResolutionDetails AlertResolution `json:"resolutionDetails,omitempty" form:"resolutionDetails,omitempty"` //Specifies information about the Alert Resolution such as a summary,
    TenantIds         *[]string       `json:"tenantIds,omitempty" form:"tenantIds,omitempty"` //Specifies unique tenantIds of the alert contained in this resolution.
}

/*
 * Structure for the custom type AlertKeyValuePair
 */
type AlertKeyValuePair struct {
    Key             *string         `json:"key,omitempty" form:"key,omitempty"` //Specifies name of the property.
    Value           *string         `json:"value,omitempty" form:"value,omitempty"` //Specifies value of the property.
}

/*
 * Structure for the custom type AlertCategoryName
 */
type AlertCategoryName struct {
    Category        CategoryEnum    `json:"category,omitempty" form:"category,omitempty"` //Specifies alert category.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies public facing string for alert enums.
}

/*
 * Structure for the custom type CreateAccessTokenCredentialRequest
 */
type CreateAccessTokenCredentialRequest struct {
    Domain          *string         `json:"domain,omitempty" form:"domain,omitempty"` //Specifies the domain the user is logging in to. For a Local user model,
    Password        string          `json:"password" form:"password"` //Specifies the password of the Cohesity user account.
    Username        string          `json:"username" form:"username"` //Specifies the login name of the Cohesity user.
}

/*
 * Structure for the custom type AccessToken
 */
type AccessToken struct {
    AccessToken     *string         `json:"accessToken,omitempty" form:"accessToken,omitempty"` //Generated access token.
    Privileges      *[]string       `json:"privileges,omitempty" form:"privileges,omitempty"` //Privileges for the user.
    TokenType       *string         `json:"tokenType,omitempty" form:"tokenType,omitempty"` //Access token type.
}

/*
 * Structure for the custom type AAGAndDatabases
 */
type AAGAndDatabases struct {
    Aag             ProtectionSource `json:"aag,omitempty" form:"aag,omitempty"` //Specifies a generic structure that represents a node
    Databases       []*ProtectionSource `json:"databases,omitempty" form:"databases,omitempty"` //Specifies databases found that are members of the AAG.
}

/*
 * Structure for the custom type CreateRemoteVaultSearchRequest
 */
type CreateRemoteVaultSearchRequest struct {
    ClusterMatchString *string         `json:"clusterMatchString,omitempty" form:"clusterMatchString,omitempty"` //Filter by specifying a Cluster name prefix string.
    EncryptionKeys     []*VaultEncryptionKey `json:"encryptionKeys,omitempty" form:"encryptionKeys,omitempty"` //Array of Encryption Keys.
    EndTimeUsecs       *int64          `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"` //Filter by a end time specified as a Unix epoch Timestamp
    JobMatchString     *string         `json:"jobMatchString,omitempty" form:"jobMatchString,omitempty"` //Filter by specifying a Protection Job name prefix string.
    SearchJobName      string          `json:"searchJobName" form:"searchJobName"` //Specifies the search Job name.
    StartTimeUsecs     *int64          `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"` //Filter by a start time specified as a Unix epoch Timestamp
    VaultId            int64           `json:"vaultId" form:"vaultId"` //Specifies the id of the Vault to search. This id was assigned by the
}

/*
 * Structure for the custom type CreateIdPConfigurationRequest
 */
type CreateIdPConfigurationRequest struct {
    AllowLocalAuthentication *bool           `json:"allowLocalAuthentication,omitempty" form:"allowLocalAuthentication,omitempty"` //Specifies whether to allow local authentication. When IdP is configured,
    Certificate              *string         `json:"certificate,omitempty" form:"certificate,omitempty"` //Specifies the certificate generated for the app by the IdP service when
    CertificateFilename      *string         `json:"certificateFilename,omitempty" form:"certificateFilename,omitempty"` //Specifies the filename used to upload the certificate.
    Enable                   *bool           `json:"enable,omitempty" form:"enable,omitempty"` //Specifies a flag to enable or disable this IdP service. When it is set
    IssuerId                 *string         `json:"issuerId,omitempty" form:"issuerId,omitempty"` //Specifies the IdP provided Issuer ID for the app.
    Name                     *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the vendor providing IdP service.
    Roles                    *[]string       `json:"roles,omitempty" form:"roles,omitempty"` //Specifies a list roles assigned to an IdP user if samlAttributeName is
    SamlAttributeName        *string         `json:"samlAttributeName,omitempty" form:"samlAttributeName,omitempty"` //Specifies the SAML attribute name that contains a comma separated list
    SsoUrl                   *string         `json:"ssoUrl,omitempty" form:"ssoUrl,omitempty"` //Specifies the SSO URL of the IdP service for the customer. This is the
    TenantId                 *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the Tenant Id if the IdP is configured for a Tenant. If this is
}

/*
 * Structure for the custom type CopyRunTask
 */
type CopyRunTask struct {
    CopySnapshotTasks   []*CopySnapshotTaskStatus `json:"copySnapshotTasks,omitempty" form:"copySnapshotTasks,omitempty"` //Specifies the status information of each task that copies the snapshot
    Error               *string         `json:"error,omitempty" form:"error,omitempty"` //Specifies if an error occurred (if any) while running this task.
    ExpiryTimeUsecs     *int64          `json:"expiryTimeUsecs,omitempty" form:"expiryTimeUsecs,omitempty"` //Specifies expiry time of the copies of the snapshots in this Protection
    HoldForLegalPurpose *bool           `json:"holdForLegalPurpose,omitempty" form:"holdForLegalPurpose,omitempty"` //Specifies whether legal hold is enabled on this run. It is true if the
    LegalHoldings       []*LegalHoldings `json:"legalHoldings,omitempty" form:"legalHoldings,omitempty"` //Specifies the list of Protection Source Ids and the legal hold status.
    RunStartTimeUsecs   *int64          `json:"runStartTimeUsecs,omitempty" form:"runStartTimeUsecs,omitempty"` //Specifies start time of the copy run.
    Stats               CopyRunStats    `json:"stats,omitempty" form:"stats,omitempty"` //Stats for one copy task or aggregated stats of a Copy Run in a
    Status              Status4Enum     `json:"status,omitempty" form:"status,omitempty"` //Specifies the aggregated status of copy tasks such as 'kRunning',
    Target              SnapshotTarget1 `json:"target,omitempty" form:"target,omitempty"` //Specifies settings about a target where a copied Snapshot is stored.
    TaskUid             UniqueGlobalId  `json:"taskUid,omitempty" form:"taskUid,omitempty"` //Specifies a globally unique id of the copy task.
    UserActionMessage   *string         `json:"userActionMessage,omitempty" form:"userActionMessage,omitempty"` //Specifies a message to the user if any manual intervention is needed to
}

/*
 * Structure for the custom type ConnectorParameters
 */
type ConnectorParameters struct {
    Endpoint        *string         `json:"endpoint,omitempty" form:"endpoint,omitempty"` //Specify an IP address or URL of the environment.
    Environment     Environment1Enum `json:"environment,omitempty" form:"environment,omitempty"` //Specifies the environment like VMware, SQL, where the
    Id              *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies a Unique id that is generated when the Source is registered.
    Version         *int64          `json:"version,omitempty" form:"version,omitempty"` //Version is updated each time the connector parameters are updated.
}

/*
 * Structure for the custom type ClusterAuditLog
 */
type ClusterAuditLog struct {
    Action          *string         `json:"action,omitempty" form:"action,omitempty"` //Specifies the action that caused the log to be generated.
    Details         *string         `json:"details,omitempty" form:"details,omitempty"` //Specifies more information about the action.
    Domain          *string         `json:"domain,omitempty" form:"domain,omitempty"` //Specifies the domain of the user who caused the action
    EntityId        *string         `json:"entityId,omitempty" form:"entityId,omitempty"` //Specifies the id of the entity (object) that the action is invoked on.
    EntityName      *string         `json:"entityName,omitempty" form:"entityName,omitempty"` //Specifies the entity (object) name that the action is invoked on.
    EntityType      *string         `json:"entityType,omitempty" form:"entityType,omitempty"` //Specifies the type of the entity (object) that the action is invoked on.
    HumanTimestamp  *string         `json:"humanTimestamp,omitempty" form:"humanTimestamp,omitempty"` //Specifies the time when the log was generated.
    Impersonation   *bool           `json:"impersonation,omitempty" form:"impersonation,omitempty"` //Specifies if the log was generated during impersonation.
    NewRecord       *string         `json:"newRecord,omitempty" form:"newRecord,omitempty"` //Specifies the record after the action is invoked.
    OriginalTenant  TenantDetails   `json:"originalTenant,omitempty" form:"originalTenant,omitempty"` //Specifies details about a tenant.
    PreviousRecord  *string         `json:"previousRecord,omitempty" form:"previousRecord,omitempty"` //Specifies the record before the action is invoked.
    Tenant          TenantDetails   `json:"tenant,omitempty" form:"tenant,omitempty"` //Specifies details about a tenant.
    TimestampUsecs  *int64          `json:"timestampUsecs,omitempty" form:"timestampUsecs,omitempty"` //Specifies the time when the log was generated.
    UserName        *string         `json:"userName,omitempty" form:"userName,omitempty"` //Specifies the user who caused the action that generated the log.
}

/*
 * Structure for the custom type CloneRestoreTaskRequest
 */
type CloneRestoreTaskRequest struct {
    CloneViewParameters  CloneViewRequest `json:"cloneViewParameters,omitempty" form:"cloneViewParameters,omitempty"` //Specifies settings for cloning an existing View.
    ContinueOnError      *bool           `json:"continueOnError,omitempty" form:"continueOnError,omitempty"` //Specifies if the Restore Task should continue when some operations on some
    GlacierRetrievalType GlacierRetrievalTypeEnum `json:"glacierRetrievalType,omitempty" form:"glacierRetrievalType,omitempty"` //Specifies the way data needs to be retrieved from the external target.
    HypervParameters     HypervCloneParameters `json:"hypervParameters,omitempty" form:"hypervParameters,omitempty"` //Specifies information needed when cloning VMs in HyperV enviroment.
    Name                 string          `json:"name" form:"name"` //Specifies the name of the Restore Task. This field must be set and
    NewParentId          *int64          `json:"newParentId,omitempty" form:"newParentId,omitempty"` //Specify a new registered parent Protection Source. If specified
    Objects              []*RestoreObject `json:"objects,omitempty" form:"objects,omitempty"` //Array of Objects.
    TargetViewName       *string         `json:"targetViewName,omitempty" form:"targetViewName,omitempty"` //Specifies the name of the View where the cloned VMs are stored.
    Type                 Type25Enum      `json:"type" form:"type"` //Specifies the type of Restore Task such as 'kCloneVMs' or 'kCloneView'.
    VlanParameters       VLANParameters  `json:"vlanParameters,omitempty" form:"vlanParameters,omitempty"` //Specifies VLAN parameters for the restore operation.
    VmwareParameters     VmwareCloneParameters `json:"vmwareParameters,omitempty" form:"vmwareParameters,omitempty"` //Specifies the information required for recovering or cloning VmWare VMs.
}

/*
 * Structure for the custom type BasicCohesityClusterInformation
 */
type BasicCohesityClusterInformation struct {
    AuthenticationType     *int64          `json:"authenticationType,omitempty" form:"authenticationType,omitempty"` //Specifies the authentication scheme for the cluster.
    ClusterSoftwareVersion *string         `json:"clusterSoftwareVersion,omitempty" form:"clusterSoftwareVersion,omitempty"` //Specifies the current release of the Cohesity software running on
    ClusterType            ClusterTypeEnum `json:"clusterType,omitempty" form:"clusterType,omitempty"` //Specifies the type of Cohesity Cluster.
    Domains                *[]string       `json:"domains,omitempty" form:"domains,omitempty"` //Array of Domains.
    IdpConfigured          *bool           `json:"idpConfigured,omitempty" form:"idpConfigured,omitempty"` //Specifies Idp is configured for the Cluster.
    IdpTenantExists        *bool           `json:"idpTenantExists,omitempty" form:"idpTenantExists,omitempty"` //Specifies Idp is configured for a Tenant.
    LanguageLocale         *string         `json:"languageLocale,omitempty" form:"languageLocale,omitempty"` //Specifies the language and locale for the Cohesity Cluster.
    McmMode                *bool           `json:"mcmMode,omitempty" form:"mcmMode,omitempty"` //Specifies whether server is running in mcm-mode. If set to true,
    MultiTenancyEnabled    *bool           `json:"multiTenancyEnabled,omitempty" form:"multiTenancyEnabled,omitempty"` //Specifies if multi-tenancy is enabled on the cluster.
    Name                   *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the Cohesity Cluster.
}

/*
 * Structure for the custom type BackupRunStatistics
 */
type BackupRunStatistics struct {
    AdmittedTimeUsecs            *int64          `json:"admittedTimeUsecs,omitempty" form:"admittedTimeUsecs,omitempty"` //Specifies the time the task was unqueued from the queue to start running.
    EndTimeUsecs                 *int64          `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"` //Specifies the end time of the Protection Run. The end time
    StartTimeUsecs               *int64          `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"` //Specifies the start time of the Protection Run. The start time
    TimeTakenUsecs               *int64          `json:"timeTakenUsecs,omitempty" form:"timeTakenUsecs,omitempty"` //Specifies the actual execution time for the protection run to complete
    TotalBytesReadFromSource     *int64          `json:"totalBytesReadFromSource,omitempty" form:"totalBytesReadFromSource,omitempty"` //Specifies the total amount of data read from the source (so far).
    TotalBytesToReadFromSource   *int64          `json:"totalBytesToReadFromSource,omitempty" form:"totalBytesToReadFromSource,omitempty"` //Specifies the total amount of data expected to be read from the
    TotalLogicalBackupSizeBytes  *int64          `json:"totalLogicalBackupSizeBytes,omitempty" form:"totalLogicalBackupSizeBytes,omitempty"` //Specifies the size of the source object (such as a VM) protected by
    TotalPhysicalBackupSizeBytes *int64          `json:"totalPhysicalBackupSizeBytes,omitempty" form:"totalPhysicalBackupSizeBytes,omitempty"` //Specifies the total amount of physical space used on the Cohesity
    TotalSourceSizeBytes         *int64          `json:"totalSourceSizeBytes,omitempty" form:"totalSourceSizeBytes,omitempty"` //Specifies the size of the source object (such as a VM) protected by
}

/*
 * Structure for the custom type DEPRECATEDIn501
 */
type DEPRECATEDIn501 struct {
    ContinuousSchedule      BackupPolicyProtoContinuousSchedule `json:"continuousSchedule,omitempty" form:"continuousSchedule,omitempty"` //TODO: Write general description for this field
    DailySchedule           BackupPolicyProtoDailySchedule `json:"dailySchedule,omitempty" form:"dailySchedule,omitempty"` //The daily schedule encompasses weekly schedules as well. This has been
    MonthlySchedule         BackupPolicyProtoMonthlySchedule `json:"monthlySchedule,omitempty" form:"monthlySchedule,omitempty"` //TODO: Write general description for this field
    Name                    *string         `json:"name,omitempty" form:"name,omitempty"` //A backup schedule can have an optional name.
    NumDaysToKeep           *int64          `json:"numDaysToKeep,omitempty" form:"numDaysToKeep,omitempty"` //Specifies how to determine the expiration time for snapshots created by
    NumRetries              *int64          `json:"numRetries,omitempty" form:"numRetries,omitempty"` //The number of retries to perform (for retryable errors) before giving up.
    OneOffSchedule          BackupPolicyProtoOneOffSchedule `json:"oneOffSchedule,omitempty" form:"oneOffSchedule,omitempty"` //TODO: Write general description for this field
    Periodicity             *int64          `json:"periodicity,omitempty" form:"periodicity,omitempty"` //Determines how often the job should be run.
    RetryDelayMins          *int64          `json:"retryDelayMins,omitempty" form:"retryDelayMins,omitempty"` //The number of minutes to wait before retrying a failed job.
    ScheduleEnd             OnlyOneOfTheFollowingFieldsShouldBeSet `json:"scheduleEnd,omitempty" form:"scheduleEnd,omitempty"` //TODO: Write general description for this field
    StartWindowIntervalMins *int64          `json:"startWindowIntervalMins,omitempty" form:"startWindowIntervalMins,omitempty"` //This field determines the amount of time (in minutes) after which a
    TruncateLogs            *bool           `json:"truncateLogs,omitempty" form:"truncateLogs,omitempty"` //Whether to truncate logs after a backup run. This is currently only
}

/*
 * Structure for the custom type AzureParameters
 */
type AzureParameters struct {
    DataDiskType           DataDiskTypeEnum `json:"dataDiskType,omitempty" form:"dataDiskType,omitempty"` //Specifies the disk type used by the data.
    InstanceId             *int64          `json:"instanceId,omitempty" form:"instanceId,omitempty"` //Specifies Type of VM (e.g. small, medium, large) when cloning the VM in
    NetworkResourceGroupId *int64          `json:"networkResourceGroupId,omitempty" form:"networkResourceGroupId,omitempty"` //Specifies id of the resource group for the selected virtual network.
    OsDiskType             OsDiskTypeEnum  `json:"osDiskType,omitempty" form:"osDiskType,omitempty"` //Specifies the disk type used by the OS.
    ResourceGroup          *int64          `json:"resourceGroup,omitempty" form:"resourceGroup,omitempty"` //Specifies id of the Azure resource group. Its value is globally unique
    StorageAccount         *int64          `json:"storageAccount,omitempty" form:"storageAccount,omitempty"` //Specifies id of the storage account that will contain the storage
    StorageContainer       *int64          `json:"storageContainer,omitempty" form:"storageContainer,omitempty"` //Specifies id of the storage container within the above storage account.
    StorageResourceGroupId *int64          `json:"storageResourceGroupId,omitempty" form:"storageResourceGroupId,omitempty"` //Specifies id of the resource group for the selected storage account.
    SubnetId               *int64          `json:"subnetId,omitempty" form:"subnetId,omitempty"` //Specifies Id of the subnet within the above virtual network.
    VirtualNetworkId       *int64          `json:"virtualNetworkId,omitempty" form:"virtualNetworkId,omitempty"` //Specifies Id of the Virtual Network.
}

/*
 * Structure for the custom type CreateApplicationsRestoreTaskRequest
 */
type CreateApplicationsRestoreTaskRequest struct {
    ApplicationEnvironment    ApplicationEnvironmentEnum `json:"applicationEnvironment" form:"applicationEnvironment"` //Specifies the Environment of the Application to restore like 'kSQL', or
    ApplicationRestoreObjects []*ApplicationServerObjectToRestore `json:"applicationRestoreObjects,omitempty" form:"applicationRestoreObjects,omitempty"` //Specifies the Application Server objects whose data should be restored
    HostingProtectionSource   RestoreObject   `json:"hostingProtectionSource" form:"hostingProtectionSource"` //Specifies an object to recover or clone or an object to restore files
    Name                      string          `json:"name" form:"name"` //Specifies a name for the new task to be created. This field has to be
    Password                  *string         `json:"password,omitempty" form:"password,omitempty"` //Specifies password of the username to access the target source.
    Username                  *string         `json:"username,omitempty" form:"username,omitempty"` //Specifies username to access the target source.
    VlanParameters            VLANParameters  `json:"vlanParameters,omitempty" form:"vlanParameters,omitempty"` //Specifies VLAN parameters for the restore operation.
}

/*
 * Structure for the custom type UpgradePhysicalServerAgentsRequest
 */
type UpgradePhysicalServerAgentsRequest struct {
    AgentIds        []int64         `json:"agentIds" form:"agentIds"` //Array of Agent Ids.
}

/*
 * Structure for the custom type StatusOfPhysicalUpgradeRequest
 */
type StatusOfPhysicalUpgradeRequest struct {
    Message         *string         `json:"message,omitempty" form:"message,omitempty"` //Specifies the status message returned after initiating an upgrade request.
}

/*
 * Structure for the custom type UpdateSourcesForPrincipalParameters
 */
type UpdateSourcesForPrincipalParameters struct {
    SourcesForPrincipals []*SourceForPrincipalParameters `json:"sourcesForPrincipals,omitempty" form:"sourcesForPrincipals,omitempty"` //Array of Principals, Sources and Views.
}

/*
 * Structure for the custom type LastProtectionRunSummary
 */
type LastProtectionRunSummary struct {
    NumberOfCancelledProtectionRuns  *int64          `json:"numberOfCancelledProtectionRuns,omitempty" form:"numberOfCancelledProtectionRuns,omitempty"` //Specifies the number of cancelled Protection Runs the given
    NumberOfFailedProtectionRuns     *int64          `json:"numberOfFailedProtectionRuns,omitempty" form:"numberOfFailedProtectionRuns,omitempty"` //Specifies the number of failed Protection Runs the given
    NumberOfProtectedSources         *int64          `json:"numberOfProtectedSources,omitempty" form:"numberOfProtectedSources,omitempty"` //Specifies the number of Protection Sources protected by the given
    NumberOfRunningProtectionRuns    *int64          `json:"numberOfRunningProtectionRuns,omitempty" form:"numberOfRunningProtectionRuns,omitempty"` //Specifies the number of running Protection Runs using the current
    NumberOfSlaViolations            *int64          `json:"numberOfSlaViolations,omitempty" form:"numberOfSlaViolations,omitempty"` //Specifies the number of SLA violations the given
    NumberOfSuccessfulProtectionRuns *int64          `json:"numberOfSuccessfulProtectionRuns,omitempty" form:"numberOfSuccessfulProtectionRuns,omitempty"` //Specifies the number of successful Protection Runs the given
    TotalLogicalBackupSizeInBytes    *int64          `json:"totalLogicalBackupSizeInBytes,omitempty" form:"totalLogicalBackupSizeInBytes,omitempty"` //Specifies the aggregated total logical backup performed in all the
}

/*
 * Structure for the custom type KVMProtectionSource
 */
type KVMProtectionSource struct {
    AgentError      *string         `json:"agentError,omitempty" form:"agentError,omitempty"` //Specifies a message when the agent cannot be reached.
    AgentId         *int64          `json:"agentId,omitempty" form:"agentId,omitempty"` //Specifies the ID of the Agent with which this KVM entity is associated
    ClusterId       *string         `json:"clusterId,omitempty" form:"clusterId,omitempty"` //Specifies the cluster ID for 'kCluster' objects.
    DatacenterId    *string         `json:"datacenterId,omitempty" form:"datacenterId,omitempty"` //Specifies the ID of the 'kDatacenter' objects.
    Description     *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description about the Protection Source.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the KVM entity.
    NetworkId       *string         `json:"networkId,omitempty" form:"networkId,omitempty"` //Specifies the network ID to which Vnic is attached.
    Type            Type9Enum       `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of KVM Protection Source entities such as
    Uuid            *string         `json:"uuid,omitempty" form:"uuid,omitempty"` //Specifies the UUID of the Object. This is unique within the KVM
}

/*
 * Structure for the custom type KMSConfiguration
 */
type KMSConfiguration struct {
    CaCertificate       *string         `json:"caCertificate,omitempty" form:"caCertificate,omitempty"` //Specifies the CA certificate in PEM format.
    ClientCertificate   *string         `json:"clientCertificate,omitempty" form:"clientCertificate,omitempty"` //Specifies the client certificate.
    ClientKey           *string         `json:"clientKey,omitempty" form:"clientKey,omitempty"` //Specifies the client private key.
    KmipProtocolVersion *string         `json:"kmipProtocolVersion,omitempty" form:"kmipProtocolVersion,omitempty"` //Specifies protocol version used to communicate with the KMS.
    ServerIp            *string         `json:"serverIp,omitempty" form:"serverIp,omitempty"` //Specifies the KMS IP address.
    ServerName          *string         `json:"serverName,omitempty" form:"serverName,omitempty"` //Specifies the name given to the KMS Server.
    ServerPort          *int64          `json:"serverPort,omitempty" form:"serverPort,omitempty"` //Specifies port on which the server is listening.
    ServerType          ServerTypeEnum  `json:"serverType,omitempty" form:"serverType,omitempty"` //Specifies the type of key mangement system.
}

/*
 * Structure for the custom type GetKMSConfigurationResponseParameters
 */
type GetKMSConfigurationResponseParameters struct {
    ClientCertificateExpiryDate *int64          `json:"clientCertificateExpiryDate,omitempty" form:"clientCertificateExpiryDate,omitempty"` //Specifies expiry date of client certificate.
    ConnectionStatus            *bool           `json:"connectionStatus,omitempty" form:"connectionStatus,omitempty"` //Specifies if connection to this KMS exists.
    KmipProtocolVersion         *string         `json:"kmipProtocolVersion,omitempty" form:"kmipProtocolVersion,omitempty"` //Specifies protocol version used to communicate with the KMS.
    ServerIp                    *string         `json:"serverIp,omitempty" form:"serverIp,omitempty"` //Specifies the KMS IP address.
    ServerName                  *string         `json:"serverName,omitempty" form:"serverName,omitempty"` //Specifies the name given to the KMS Server.
    ServerPort                  *int64          `json:"serverPort,omitempty" form:"serverPort,omitempty"` //Specifies port on which the server is listening.
}

/*
 * Structure for the custom type RestoreProtectionJobIndexAndSnapshots
 */
type RestoreProtectionJobIndexAndSnapshots struct {
    ArchiveTaskUid         UniqueGlobalId  `json:"archiveTaskUid,omitempty" form:"archiveTaskUid,omitempty"` //Specifies a unique id of the Archive task that originally archived the
    EndTimeUsecs           *int64          `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"` //Specifies the end time as a Unix epoch Timestamp (in microseconds).
    RemoteProtectionJobUid UniqueGlobalId  `json:"remoteProtectionJobUid,omitempty" form:"remoteProtectionJobUid,omitempty"` //Specifies a unique id assigned to the original Protection Job
    StartTimeUsecs         *int64          `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"` //Specifies the start time as a Unix epoch Timestamp (in microseconds).
    ViewBoxId              *int64          `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"` //Specifies the id of the local Storage Domain (View Box) where the index
}

/*
 * Structure for the custom type IsilonProtectionSource
 */
type IsilonProtectionSource struct {
    AccessZone      IsilonAccessZone `json:"accessZone,omitempty" form:"accessZone,omitempty"` //Specifies information about access zone in an Isilon Cluster.
    Cluster         IsilonCluster   `json:"cluster,omitempty" form:"cluster,omitempty"` //Specifies information about an Isilon Cluster.
    MountPoint      IsilonMountPoint `json:"mountPoint,omitempty" form:"mountPoint,omitempty"` //Specifies information about a mount point in an Isilon OneFs Cluster.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies a unique name of the Protection Source.
    Type            Type8Enum       `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of the entity in an Isilon OneFs file system
}

/*
 * Structure for the custom type JobRunsTile
 */
type JobRunsTile struct {
    LastDayNumJobErrors        *int64          `json:"lastDayNumJobErrors,omitempty" form:"lastDayNumJobErrors,omitempty"` //Number of Error runs in the last 24 hours.
    LastDayNumJobRuns          *int64          `json:"lastDayNumJobRuns,omitempty" form:"lastDayNumJobRuns,omitempty"` //Number of Job Runs in the last 24 hours.
    LastDayNumJobSlaViolations *int64          `json:"lastDayNumJobSlaViolations,omitempty" form:"lastDayNumJobSlaViolations,omitempty"` //Number of SLA Violations in the last 24 hours.
    NumJobRunning              *int64          `json:"numJobRunning,omitempty" form:"numJobRunning,omitempty"` //Number of Jobs currently running.
    ObjectsProtectedByPolicy   []*ObjectsProtectedByPolicy `json:"objectsProtectedByPolicy,omitempty" form:"objectsProtectedByPolicy,omitempty"` //Objects Protected By Policy.
}

/*
 * Structure for the custom type IdPUserInformation
 */
type IdPUserInformation struct {
    IdpId           *int64          `json:"idpId,omitempty" form:"idpId,omitempty"` //Specifies the unique Id assigned by the Cluster for the IdP.
    IssuerId        *string         `json:"issuerId,omitempty" form:"issuerId,omitempty"` //Specifies the unique identifier assigned by the vendor for this Cluster.
    UserId          *string         `json:"userId,omitempty" form:"userId,omitempty"` //Specifies the unique identifier assigned by the vendor for the user.
    Vendor          *string         `json:"vendor,omitempty" form:"vendor,omitempty"` //Specifies the vendor providing the IdP service.
}

/*
 * Structure for the custom type HypervVirtualMachineObject
 */
type HypervVirtualMachineObject struct {
    IsHighlyAvailable *bool           `json:"isHighlyAvailable,omitempty" form:"isHighlyAvailable,omitempty"` //Specifies whether the VM is Highly Availabile or not.
    Version           *string         `json:"version,omitempty" form:"version,omitempty"` //Specifies the version of the VM. For example, 8.0, 5.0 etc.
    VmBackupStatus    VmBackupStatusEnum `json:"vmBackupStatus,omitempty" form:"vmBackupStatus,omitempty"` //Specifies the status of the VM for backup purpose.
    VmBackupType      VmBackupTypeEnum `json:"vmBackupType,omitempty" form:"vmBackupType,omitempty"` //Specifies the type of backup supported by the VM.
}

/*
 * Structure for the custom type HypervDatastoreObject
 */
type HypervDatastoreObject struct {
    Capacity        *int64          `json:"capacity,omitempty" form:"capacity,omitempty"` //Specifies the capacity of the datastore in bytes.
    FreeSpace       *int64          `json:"freeSpace,omitempty" form:"freeSpace,omitempty"` //Specifies the available space on the datastore in bytes.
    MountPoints     *[]string       `json:"mountPoints,omitempty" form:"mountPoints,omitempty"` //Specifies the available mount points on the datastore.
    Type            Type6Enum       `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of the datastore object like kFileShare or kVolume.
}

/*
 * Structure for the custom type IOPSTile
 */
type IOPSTile struct {
    MaxReadIops      *int64          `json:"maxReadIops,omitempty" form:"maxReadIops,omitempty"` //Maximum Read IOs per second in last 24 hours.
    MaxWriteIops     *int64          `json:"maxWriteIops,omitempty" form:"maxWriteIops,omitempty"` //Maximum number of Write IOs per second in last 24 hours.
    ReadIopsSamples  []*Sample       `json:"readIopsSamples,omitempty" form:"readIopsSamples,omitempty"` //Read IOs per second samples taken for the past 24 hours at 10 minutes
    WriteIopsSamples []*Sample       `json:"writeIopsSamples,omitempty" form:"writeIopsSamples,omitempty"` //Write IOs per second samples taken for the past 24 hours at 10 minutes
}

/*
 * Structure for the custom type IsilonMountPoint
 */
type IsilonMountPoint struct {
    AccessZoneName  *string         `json:"accessZoneName,omitempty" form:"accessZoneName,omitempty"` //Specifies the name of access zone.
    NfsMountPoint   IsilonNFSMountPoint `json:"nfsMountPoint,omitempty" form:"nfsMountPoint,omitempty"` //Specifies NFS Mount Point exposed by Isilon Protection Source.
    Path            *string         `json:"path,omitempty" form:"path,omitempty"` //Specifies the path of the access zone in ifs. This should include the
    Protocols       *[]Protocol1Enum `json:"protocols,omitempty" form:"protocols,omitempty"` //List of Protocols on Isilon.
    SmbMountPoints  []*IsilonSMBMountPoint `json:"smbMountPoints,omitempty" form:"smbMountPoints,omitempty"` //Specifies information about an SMB share. This field is set if the
}

/*
 * Structure for the custom type HyperFlexStoraeSnapshot
 */
type HyperFlexStoraeSnapshot struct {
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies a unique name of the Protection Source
    ProductVersion  *string         `json:"productVersion,omitempty" form:"productVersion,omitempty"` //Specifies the product version of the protection source.
    Type            Type5Enum       `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of managed Object in a HyperFlex protection source
    Uuid            *string         `json:"uuid,omitempty" form:"uuid,omitempty"` //Specifies the uuid of the protection source.
}

/*
 * Structure for the custom type InputSpecInputVMsSelector
 */
type InputSpecInputVMsSelector struct {
    FileTimeFilter       InputSpecFileTimeFilter `json:"fileTimeFilter,omitempty" form:"fileTimeFilter,omitempty"` //File time filter to filter files based on their last change time. All
    FilenameGlob         *[]string       `json:"filenameGlob,omitempty" form:"filenameGlob,omitempty"` //After VMDKs are selected as above, the files within them can be selected
    JobIds               *[]int64        `json:"jobIds,omitempty" form:"jobIds,omitempty"` //TODO: Write general description for this field
    MaxSnapshotTimestamp *int64          `json:"maxSnapshotTimestamp,omitempty" form:"maxSnapshotTimestamp,omitempty"` //Exclusive end of snapshot_timestamp range. If missing, inf will be used
    MinSnapshotTimestamp *int64          `json:"minSnapshotTimestamp,omitempty" form:"minSnapshotTimestamp,omitempty"` //Inclusive. If missing, 0 will the default lower end of timestamp range
    PartitionIds         *[]int64        `json:"partitionIds,omitempty" form:"partitionIds,omitempty"` //Filters are AND of ORs.
    ProcessLatestOnly    *bool           `json:"processLatestOnly,omitempty" form:"processLatestOnly,omitempty"` //Boolean flag to indicate if only latest snapshot of each object should
    RootDir              *string         `json:"rootDir,omitempty" form:"rootDir,omitempty"` //Within each volume, traversal will be rooted at this directory. A
    SourceEntityIds      *[]int64        `json:"sourceEntityIds,omitempty" form:"sourceEntityIds,omitempty"` //TODO: Write general description for this field
    ViewBoxIds           *[]int64        `json:"viewBoxIds,omitempty" form:"viewBoxIds,omitempty"` //TODO: Write general description for this field
    ViewNames            *[]string       `json:"viewNames,omitempty" form:"viewNames,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type GoogleCloudCredentials
 */
type GoogleCloudCredentials struct {
    ClientEmailAddress *string         `json:"clientEmailAddress,omitempty" form:"clientEmailAddress,omitempty"` //Specifies the client email address used to access Google
    ClientPrivateKey   *string         `json:"clientPrivateKey,omitempty" form:"clientPrivateKey,omitempty"` //Specifies the private key used to access Google Cloud Storage that is
    ProjectId          *string         `json:"projectId,omitempty" form:"projectId,omitempty"` //Specifies the project id of an existing Google Cloud project to store
    TierType           TierType2Enum   `json:"tierType,omitempty" form:"tierType,omitempty"` //Specifies the storage class of GCP.
}

/*
 * Structure for the custom type GDPRCopyTask
 */
type GDPRCopyTask struct {
    JobId           *int64          `json:"JobId,omitempty" form:"JobId,omitempty"` //Specifies the job with which this copy task is tied to.
    CloudTargetType *string         `json:"cloudTargetType,omitempty" form:"cloudTargetType,omitempty"` //Specifies the cloud deploy target type. For example 'kAzure','kAWS',
    ExpiryTimeUsecs *int64          `json:"expiryTimeUsecs,omitempty" form:"expiryTimeUsecs,omitempty"` //Specifies the expiry of the copy task.
    TargetId        *int64          `json:"targetId,omitempty" form:"targetId,omitempty"` //Specifies the id for the target.
    TargetName      *string         `json:"targetName,omitempty" form:"targetName,omitempty"` //Specifies the target of the replication or archival tasks.
    TotalSnapshots  *int64          `json:"totalSnapshots,omitempty" form:"totalSnapshots,omitempty"` //Specifies the total number of snapshots.
    Type            *string         `json:"type,omitempty" form:"type,omitempty"` //Specifies details about the Copy Run of a Job Run.
}

/*
 * Structure for the custom type GCPSourceCredentials
 */
type GCPSourceCredentials struct {
    ClientEmailAddress *string         `json:"clientEmailAddress,omitempty" form:"clientEmailAddress,omitempty"` //Specifies Client email address associated with the service account.
    ClientPrivateKey   *string         `json:"clientPrivateKey,omitempty" form:"clientPrivateKey,omitempty"` //Specifies Client private associated with the service account.
    GcpType            GcpTypeEnum     `json:"gcpType,omitempty" form:"gcpType,omitempty"` //Specifies the entity type such as 'kIAMUser' if the environment is kGCP.
    ProjectId          *string         `json:"projectId,omitempty" form:"projectId,omitempty"` //Specifies Id of the project associated with Google cloud account.
    VpcNetwork         *string         `json:"vpcNetwork,omitempty" form:"vpcNetwork,omitempty"` //Specifies the VPC Network to deploy proxy VMs.
    VpcSubnetwork      *string         `json:"vpcSubnetwork,omitempty" form:"vpcSubnetwork,omitempty"` //Specifies the subnetwork to deploy proxy VMs.
}

/*
 * Structure for the custom type InputSpecInputFilesSelector
 */
type InputSpecInputFilesSelector struct {
    FileTimeFilter       InputSpecFileTimeFilter `json:"fileTimeFilter,omitempty" form:"fileTimeFilter,omitempty"` //File time filter to filter files based on their last change time. All
    FilenameGlob         *[]string       `json:"filenameGlob,omitempty" form:"filenameGlob,omitempty"` //Glob patterns to match on file. e.g. {*.txt, *.cc}
    JobIds               *[]int64        `json:"jobIds,omitempty" form:"jobIds,omitempty"` //TODO: Write general description for this field
    MaxSnapshotTimestamp *int64          `json:"maxSnapshotTimestamp,omitempty" form:"maxSnapshotTimestamp,omitempty"` //Exclusive end of snapshot_timestamp range. If missing, inf will be used
    MinSnapshotTimestamp *int64          `json:"minSnapshotTimestamp,omitempty" form:"minSnapshotTimestamp,omitempty"` //Inclusive. If missing, 0 will the default lower end of timestamp range
    PartitionIds         *[]int64        `json:"partitionIds,omitempty" form:"partitionIds,omitempty"` //TODO: Write general description for this field
    ProcessLatestOnly    *bool           `json:"processLatestOnly,omitempty" form:"processLatestOnly,omitempty"` //Boolean flag to indicate if only latest snapshot of each object should
    RootDir              *string         `json:"rootDir,omitempty" form:"rootDir,omitempty"` //Within each volume, traversal will be rooted at this directory. A
    ViewBoxIds           *[]int64        `json:"viewBoxIds,omitempty" form:"viewBoxIds,omitempty"` //TODO: Write general description for this field
    ViewName             *string         `json:"viewName,omitempty" form:"viewName,omitempty"` //This is the view name user enters manually. If this is set we will
}

/*
 * Structure for the custom type FlashBladeStorageArray
 */
type FlashBladeStorageArray struct {
    CapacityBytes     *int64          `json:"capacityBytes,omitempty" form:"capacityBytes,omitempty"` //Specifies the total capacity in bytes of the Pure Storage FlashBlade
    Id                *string         `json:"id,omitempty" form:"id,omitempty"` //Specifies a unique id of a Pure Storage FlashBlade Array.
    Networks          []*FlashBladeNetworkInterface `json:"networks,omitempty" form:"networks,omitempty"` //Specifies the network interfaces of the Pure Storage FlashBlade Array.
    PhysicalUsedBytes *int64          `json:"physicalUsedBytes,omitempty" form:"physicalUsedBytes,omitempty"` //Specifies the space used for physical data in bytes.
    Revision          *string         `json:"revision,omitempty" form:"revision,omitempty"` //Specifies the revision of the Pure Storage FlashBlade software.
    Version           *string         `json:"version,omitempty" form:"version,omitempty"` //Specifies the software version running on the Pure Storage FlashBlade
}

/*
 * Structure for the custom type PureStorageFlashBladeProtectionSource
 */
type PureStorageFlashBladeProtectionSource struct {
    FileSystem      FlashBladeFileSystem `json:"fileSystem,omitempty" form:"fileSystem,omitempty"` //Specifies information about a Flash Blade File System in a Storage Array.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies a unique name of the Protection Source.
    StorageArray    FlashBladeStorageArray `json:"storageArray,omitempty" form:"storageArray,omitempty"` //Specifies information about a Pure Storage FlashBlade Array.
    Type            Type3Enum       `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of managed object in a Pure Storage FlashBlade
}

/*
 * Structure for the custom type ImportConfigurations
 */
type ImportConfigurations struct {
    ActiveDirectories  *[]string       `json:"activeDirectories,omitempty" form:"activeDirectories,omitempty"` //Selective import of active directories.
    All                *[]string       `json:"all,omitempty" form:"all,omitempty"` //List of which entities to import all.
    Clusters           *[]int64        `json:"clusters,omitempty" form:"clusters,omitempty"` //Selective import certain cluster.
    File               *string         `json:"file,omitempty" form:"file,omitempty"` //File is the config file.
    Groups             *[]string       `json:"groups,omitempty" form:"groups,omitempty"` //Selective import certain groups.
    Partitions         *[]int64        `json:"partitions,omitempty" form:"partitions,omitempty"` //Selective import of Partiton.
    PrincipalSources   *[]string       `json:"principalSources,omitempty" form:"principalSources,omitempty"` //Selective import of principal sources.
    ProtectionJobs     *[]int64        `json:"protectionJobs,omitempty" form:"protectionJobs,omitempty"` //Selective import of protection jobs.
    ProtectionPolicies *[]string       `json:"protectionPolicies,omitempty" form:"protectionPolicies,omitempty"` //Selective import of protection policies.
    ProtectionSources  *[]int64        `json:"protectionSources,omitempty" form:"protectionSources,omitempty"` //Selective import of protection sources.
    RemoteClusters     *[]int64        `json:"remoteClusters,omitempty" form:"remoteClusters,omitempty"` //Selective import certain remote clusters.
    Roles              *[]string       `json:"roles,omitempty" form:"roles,omitempty"` //Selective import certain roles (by username).
    Sql                *[]int64        `json:"sql,omitempty" form:"sql,omitempty"` //Selective import of sql.
    Users              *[]string       `json:"users,omitempty" form:"users,omitempty"` //Selective import certain users.
    Vaults             *[]int64        `json:"vaults,omitempty" form:"vaults,omitempty"` //Selective import certain vaults.
    ViewBoxes          *[]int64        `json:"viewBoxes,omitempty" form:"viewBoxes,omitempty"` //Selective import certain Storage Domains (View Boxes).
    Views              *[]int64        `json:"views,omitempty" form:"views,omitempty"` //Selective import of views.
}

/*
 * Structure for the custom type FileFolderSnapshotInformation
 */
type FileFolderSnapshotInformation struct {
    HasArchivalCopy   *bool           `json:"hasArchivalCopy,omitempty" form:"hasArchivalCopy,omitempty"` //If true, this snapshot is located on an archival target
    HasLocalCopy      *bool           `json:"hasLocalCopy,omitempty" form:"hasLocalCopy,omitempty"` //If true, this snapshot is located on a local Cohesity Cluster.
    HasRemoteCopy     *bool           `json:"hasRemoteCopy,omitempty" form:"hasRemoteCopy,omitempty"` //If true, this snapshot is located on a Remote Cohesity Cluster.
    ModifiedTimeUsecs *int64          `json:"modifiedTimeUsecs,omitempty" form:"modifiedTimeUsecs,omitempty"` //Specifies the time when the file or folder was last modified.
    SizeBytes         *int64          `json:"sizeBytes,omitempty" form:"sizeBytes,omitempty"` //Specifies the size of the file or folder in bytes.
    Snapshot          SnapshotAttempt `json:"snapshot,omitempty" form:"snapshot,omitempty"` //Specifies information about a single snapshot.
}

/*
 * Structure for the custom type IdPServiceConfiguration
 */
type IdPServiceConfiguration struct {
    AllowLocalAuthentication *bool           `json:"allowLocalAuthentication,omitempty" form:"allowLocalAuthentication,omitempty"` //Specifies whether to allow local authentication. When IdP is configured,
    Certificate              *string         `json:"certificate,omitempty" form:"certificate,omitempty"` //Specifies the certificate generated for the app by the IdP service when
    CertificateFilename      *string         `json:"certificateFilename,omitempty" form:"certificateFilename,omitempty"` //Specifies the filename used to upload the certificate.
    Enable                   *bool           `json:"enable,omitempty" form:"enable,omitempty"` //Specifies a flag to enable or disable this IdP service. When it is set
    Id                       *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies the Id assigned by the Cluster for the IdP service.
    IssuerId                 *string         `json:"issuerId,omitempty" form:"issuerId,omitempty"` //Specifies the IdP provided Issuer ID for the app.
    Name                     *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the vendor providing IdP service.
    Roles                    *[]string       `json:"roles,omitempty" form:"roles,omitempty"` //Specifies a list roles assigned to an IdP user if samlAttributeName is
    SamlAttributeName        *string         `json:"samlAttributeName,omitempty" form:"samlAttributeName,omitempty"` //Specifies the SAML attribute name that contains a comma separated list
    SsoUrl                   *string         `json:"ssoUrl,omitempty" form:"ssoUrl,omitempty"` //Specifies the SSO URL of the IdP service for the customer. This is the
    TenantId                 *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the Tenant Id if the IdP is configured for a Tenant. If this is
}

/*
 * Structure for the custom type FilePartitionBlock
 */
type FilePartitionBlock struct {
    DiskFileName    *string         `json:"diskFileName,omitempty" form:"diskFileName,omitempty"` //Specifies the disk file name where the logical partition is.
    LengthBytes     *int64          `json:"lengthBytes,omitempty" form:"lengthBytes,omitempty"` //Specifies the length of the block in bytes.
    Number          *int64          `json:"number,omitempty" form:"number,omitempty"` //Specifies a unique number of the partition within the linear disk file.
    OffsetBytes     *int64          `json:"offsetBytes,omitempty" form:"offsetBytes,omitempty"` //Specifies the offset of the block (in bytes) from the beginning
}

/*
 * Structure for the custom type HypervRestoreParameters
 */
type HypervRestoreParameters struct {
    DatastoreId     *int64          `json:"datastoreId,omitempty" form:"datastoreId,omitempty"` //A datastore entity where the object's files should be restored to. This
    DisableNetwork  *bool           `json:"disableNetwork,omitempty" form:"disableNetwork,omitempty"` //Specifies whether the network should be left in disabled state.
    NetworkId       *int64          `json:"networkId,omitempty" form:"networkId,omitempty"` //Specifies a network configuration to be attached to the cloned or
    PoweredOn       *bool           `json:"poweredOn,omitempty" form:"poweredOn,omitempty"` //Specifies the power state of the cloned or recovered objects.
    Prefix          *string         `json:"prefix,omitempty" form:"prefix,omitempty"` //Specifies a prefix to prepended to the source object name to derive a
    ResourceId      *int64          `json:"resourceId,omitempty" form:"resourceId,omitempty"` //The resource (HyperV host) to which the restored VM will be attached.
    Suffix          *string         `json:"suffix,omitempty" form:"suffix,omitempty"` //Specifies a suffix to appended to the original source object name
}

/*
 * Structure for the custom type FileLockStatus
 */
type FileLockStatus struct {
    ExpiryTimestampMsecs *int64          `json:"expiryTimestampMsecs,omitempty" form:"expiryTimestampMsecs,omitempty"` //Specifies a expiry timestamp in milliseconds until the file is locked.
    HoldTimestampMsecs   *int64          `json:"holdTimestampMsecs,omitempty" form:"holdTimestampMsecs,omitempty"` //Specifies a override timestamp in milliseconds when an expired file is
    LockTimestampMsecs   *int64          `json:"lockTimestampMsecs,omitempty" form:"lockTimestampMsecs,omitempty"` //Specifies the timestamp at which the file was locked.
    State                *int64          `json:"state,omitempty" form:"state,omitempty"` //Specifies the lock state of the file.
}

/*
 * Structure for the custom type ProtectionStatistics
 */
type ProtectionStatistics struct {
    NumFailed       *int64          `json:"numFailed,omitempty" form:"numFailed,omitempty"` //Number of Failed Objects.
    NumObjects      *int64          `json:"numObjects,omitempty" form:"numObjects,omitempty"` //Number of Objects.
    SizeBytes       *int64          `json:"sizeBytes,omitempty" form:"sizeBytes,omitempty"` //Size in Bytes.
}

/*
 * Structure for the custom type ExtractFileRangeResult
 */
type ExtractFileRangeResult struct {
    Data            *[]int64        `json:"data,omitempty" form:"data,omitempty"` //The actual data bytes.
    Eof             *bool           `json:"eof,omitempty" form:"eof,omitempty"` //Will be true if start_offset > file length or EOF is reached. This is an
    Error           ErrorProto      `json:"error,omitempty" form:"error,omitempty"` //TODO: Write general description for this field
    FileLength      *int64          `json:"fileLength,omitempty" form:"fileLength,omitempty"` //The total length of the file. This field would be set provided no error
    StartOffset     *int64          `json:"startOffset,omitempty" form:"startOffset,omitempty"` //The offset from which data was read.
}

/*
 * Structure for the custom type ProtectionSourceAndSnapshots
 */
type ProtectionSourceAndSnapshots struct {
    ProtectionSource ProtectionSource `json:"protectionSource,omitempty" form:"protectionSource,omitempty"` //Specifies the leaf Protection Source Object such as a VM.
    SnapshotsInfo    []*ProtectionSourceSnapshot `json:"snapshotsInfo,omitempty" form:"snapshotsInfo,omitempty"` //Array of Snapshots
}

/*
 * Structure for the custom type ExtendedRetentionPolicy
 */
type ExtendedRetentionPolicy struct {
    BackupRunType   BackupRunTypeEnum `json:"backupRunType,omitempty" form:"backupRunType,omitempty"` //The backup run type to which this extended retention applies to. If this is
    DaysToKeep      *int64          `json:"daysToKeep,omitempty" form:"daysToKeep,omitempty"` //Specifies the number of days to retain copied Snapshots on the target.
    Multiplier      *int64          `json:"multiplier,omitempty" form:"multiplier,omitempty"` //Specifies a factor to multiply the periodicity by, to determine the copy
    Periodicity     PeriodicityEnum `json:"periodicity,omitempty" form:"periodicity,omitempty"` //Specifies the frequency that Snapshots should be copied to the
}

/*
 * Structure for the custom type ProtectionRunResponse
 */
type ProtectionRunResponse struct {
    ArchivalRuns    []*LatestProtectionJobRunInformation `json:"archivalRuns,omitempty" form:"archivalRuns,omitempty"` //Specifies the list of archival job information.
    BackupRuns      []*LatestProtectionJobRunInformation `json:"backupRuns,omitempty" form:"backupRuns,omitempty"` //Specifies the list of local backup job information.
    ReplicationRuns []*LatestProtectionJobRunInformation `json:"replicationRuns,omitempty" form:"replicationRuns,omitempty"` //Specifies the list of replication job information.
}

/*
 * Structure for the custom type ProtectionJobSummaryForPolicies
 */
type ProtectionJobSummaryForPolicies struct {
    BackupRun       BackupRunTask   `json:"backupRun,omitempty" form:"backupRun,omitempty"` //Specifies details about the Backup task for a Job Run.
    CopyRuns        []*CopyRunTask  `json:"copyRuns,omitempty" form:"copyRuns,omitempty"` //Specifies details about the Copy tasks of the Job Run.
    ProtectionJob   ProtectionJob   `json:"protectionJob,omitempty" form:"protectionJob,omitempty"` //Provides details about a Protection Job.
}

/*
 * Structure for the custom type ProtectedVMInformation
 */
type ProtectedVMInformation struct {
    ProtectionJobs     []*ProtectionJob `json:"protectionJobs,omitempty" form:"protectionJobs,omitempty"` //Specifies the list of Protection Jobs that protect the VM.
    ProtectionPolicies []*ProtectionPolicy `json:"protectionPolicies,omitempty" form:"protectionPolicies,omitempty"` //Specifies the list of Policies that are used by the Protection Jobs.
    ProtectionSource   ProtectionSource `json:"protectionSource,omitempty" form:"protectionSource,omitempty"` //Specifies a generic structure that represents a node
    Stats              ProtectionSummary `json:"stats,omitempty" form:"stats,omitempty"` //Specifies the protection stats of VM.
}

/*
 * Structure for the custom type ProductModelAndInterfaceTuple
 */
type ProductModelAndInterfaceTuple struct {
    IfaceName        *string         `json:"ifaceName,omitempty" form:"ifaceName,omitempty"` //Specifies the name of the interface.
    ProductModelName *string         `json:"productModelName,omitempty" form:"productModelName,omitempty"` //Specifies the product model name.
}

/*
 * Structure for the custom type UpdatePreferredDomainControllerRequest
 */
type UpdatePreferredDomainControllerRequest struct {
    DomainControllers *[]string       `json:"domainControllers,omitempty" form:"domainControllers,omitempty"` //List of Domain controllers DCs in FQDN format that are mapped to an Active
    DomainName        *string         `json:"domainName,omitempty" form:"domainName,omitempty"` //Specifies the Domain name or the trusted domain of an Active Directory.
}

/*
 * Structure for the custom type PhysicalSnapshotParams
 */
type PhysicalSnapshotParams struct {
    FetchSnapshotMetadataDisabled *bool           `json:"fetchSnapshotMetadataDisabled,omitempty" form:"fetchSnapshotMetadataDisabled,omitempty"` //Whether fetching and storing of snapshot metadata was disabled.
    NotifyBackupCompleteDisabled  *bool           `json:"notifyBackupCompleteDisabled,omitempty" form:"notifyBackupCompleteDisabled,omitempty"` //Whether notify backup complete step was disabled.
    VssCopyOnlyBackup             *bool           `json:"vssCopyOnlyBackup,omitempty" form:"vssCopyOnlyBackup,omitempty"` //If copy_only_backup option is requrested at the time of the snapshot.
    VssExcludedWriters            *[]string       `json:"vssExcludedWriters,omitempty" form:"vssExcludedWriters,omitempty"` //List of VSS writers that were excluded.
}

/*
 * Structure for the custom type EULAConfiguration
 */
type EULAConfiguration struct {
    LicenseKey      *string         `json:"licenseKey,omitempty" form:"licenseKey,omitempty"` //Specifies the license key.
    SignedByUser    *string         `json:"signedByUser,omitempty" form:"signedByUser,omitempty"` //Specifies the login account name for the Cohesity user who accepted
    SignedTime      *int64          `json:"signedTime,omitempty" form:"signedTime,omitempty"` //Specifies the time that the End User License Agreement was accepted.
    SignedVersion   *int64          `json:"signedVersion,omitempty" form:"signedVersion,omitempty"` //Specifies the version of the End User License Agreement that was accepted.
}

/*
 * Structure for the custom type ErasureCodingInformation
 */
type ErasureCodingInformation struct {
    Algorithm            AlgorithmEnum   `json:"algorithm,omitempty" form:"algorithm,omitempty"` //Algorthm used for erasure coding.
    ErasureCodingEnabled *bool           `json:"erasureCodingEnabled,omitempty" form:"erasureCodingEnabled,omitempty"` //Specifies whether Erasure coding is enabled on the Storage Domain
    InlineErasureCoding  *bool           `json:"inlineErasureCoding,omitempty" form:"inlineErasureCoding,omitempty"` //Specifies if erasure coding should occur inline (as the data is being
    NumCodedStripes      *int64          `json:"numCodedStripes,omitempty" form:"numCodedStripes,omitempty"` //The number of coded stripes.
    NumDataStripes       *int64          `json:"numDataStripes,omitempty" form:"numDataStripes,omitempty"` //The number of stripes containing data.
}

/*
 * Structure for the custom type HypervProtectionSource
 */
type HypervProtectionSource struct {
    Agents          []*AgentInformation `json:"agents,omitempty" form:"agents,omitempty"` //Array of Agents on the Physical Protection Source.
    BackupType      BackupTypeEnum  `json:"backupType,omitempty" form:"backupType,omitempty"` //Specifies the type of backup supported by the VM.
    ClusterName     *string         `json:"clusterName,omitempty" form:"clusterName,omitempty"` //Specifies the cluster name for 'kHostCluster' objects.
    DatastoreInfo   HypervDatastoreObject `json:"datastoreInfo,omitempty" form:"datastoreInfo,omitempty"` //Specifies information about a Datastore Object in HyperV environment.
    Description     *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description about the Protection Source.
    HostType        HostType4Enum   `json:"hostType,omitempty" form:"hostType,omitempty"` //Specifies host OS type for 'kVirtualMachine' objects.
    HypervUuid      *string         `json:"hypervUuid,omitempty" form:"hypervUuid,omitempty"` //Specifies the UUID for 'kVirtualMachine' HyperV objects.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the HyperV Object.
    Type            Type7Enum       `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of an HyperV Protection Source Object such as
    Uuid            *string         `json:"uuid,omitempty" form:"uuid,omitempty"` //Specifies the UUID of the Object. This is unique within the HyperV
    VmInfo          HypervVirtualMachineObject `json:"vmInfo,omitempty" form:"vmInfo,omitempty"` //Specifies information about a VirtualMachine Object in HyperV environment.
}

/*
 * Structure for the custom type EnvironmentSpecificCommonJobParameters
 */
type EnvironmentSpecificCommonJobParameters struct {
    HypervParameters   HypervEnvironmentJobParameters `json:"hypervParameters,omitempty" form:"hypervParameters,omitempty"` //Specifies job parameters applicable for all 'kHyperV' Environment type
    NasParameters      NASEnvironmentJobParameters `json:"nasParameters,omitempty" form:"nasParameters,omitempty"` //Specifies job parameters applicable for all 'kGenericNas' Environment type
    OutlookParameters  OutlookEnvironmentJobParameters `json:"outlookParameters,omitempty" form:"outlookParameters,omitempty"` //Specifies job parameters applicable for all 'kO365Outlook' Environment type
    PhysicalParameters PhysicalEnvironmentJobParameters `json:"physicalParameters,omitempty" form:"physicalParameters,omitempty"` //Protection Job parameters applicable to 'kPhysical' Environment type.
    PureParameters     PureEnvironmentJobParameters `json:"pureParameters,omitempty" form:"pureParameters,omitempty"` //Specifies job parameters applicable for all 'kPure' Environment type
    SqlParameters      SQLEnvironmentJobParameters `json:"sqlParameters,omitempty" form:"sqlParameters,omitempty"` //Specifies job parameters applicable for all 'kSQL' Environment type
    VmwareParameters   VmwareEnvironmentJobParameters `json:"vmwareParameters,omitempty" form:"vmwareParameters,omitempty"` //Specifies job parameters applicable for all 'kVMware' Environment type
}

/*
 * Structure for the custom type Alert
 */
type Alert struct {
    AlertCategory        AlertCategoryEnum `json:"alertCategory,omitempty" form:"alertCategory,omitempty"` //Specifies the category of an Alert.
    AlertCode            *string         `json:"alertCode,omitempty" form:"alertCode,omitempty"` //Specifies a unique code that categorizes the Alert,
    AlertDocument        AlertDocument   `json:"alertDocument,omitempty" form:"alertDocument,omitempty"` //Specifies documentation about the Alert such as name, description, cause
    AlertState           AlertStateEnum  `json:"alertState,omitempty" form:"alertState,omitempty"` //Specifies the current state of the Alert.
    AlertType            *int64          `json:"alertType,omitempty" form:"alertType,omitempty"` //Specifies a 5 digit unique digital id for the Alert Type, such as 00014
    ClusterId            *int64          `json:"clusterId,omitempty" form:"clusterId,omitempty"` //Specifies id of the cluster where the alert was raised.
    ClusterName          *string         `json:"clusterName,omitempty" form:"clusterName,omitempty"` //Specifies name of the cluster where the alert was raised.
    DedupCount           *int64          `json:"dedupCount,omitempty" form:"dedupCount,omitempty"` //Specifies total count of duplicated Alerts even if there are more than
    DedupTimestamps      *[]int64        `json:"dedupTimestamps,omitempty" form:"dedupTimestamps,omitempty"` //Specifies Unix epoch Timestamps (in microseconds) for the last 25
    EventSource          *string         `json:"eventSource,omitempty" form:"eventSource,omitempty"` //Specifies source where the event occurred.
    FirstTimestampUsecs  *int64          `json:"firstTimestampUsecs,omitempty" form:"firstTimestampUsecs,omitempty"` //Specifies Unix epoch Timestamp (in microseconds) of the first
    Id                   *string         `json:"id,omitempty" form:"id,omitempty"` //Specifies unique id of this Alert.
    LatestTimestampUsecs *int64          `json:"latestTimestampUsecs,omitempty" form:"latestTimestampUsecs,omitempty"` //Specifies Unix epoch Timestamp (in microseconds) of the most
    PropertyList         []*AlertKeyValuePair `json:"propertyList,omitempty" form:"propertyList,omitempty"` //Specifies array of key-value pairs associated with the Alert.
    ResolutionDetails    AlertResolution `json:"resolutionDetails,omitempty" form:"resolutionDetails,omitempty"` //Specifies information about the Alert Resolution such as a summary,
    Severity             SeverityEnum    `json:"severity,omitempty" form:"severity,omitempty"` //Specifies the severity level of an Alert.
    SuppressionId        *int64          `json:"suppressionId,omitempty" form:"suppressionId,omitempty"` //Specifies unique id generated when the Alert is suppressed by the admin.
    TenantIds            *[]string       `json:"tenantIds,omitempty" form:"tenantIds,omitempty"` //Specifies the tenants for which this alert has been raised.
}

/*
 * Structure for the custom type AppInstance
 */
type AppInstance struct {
    AppAccessToken   *string         `json:"appAccessToken,omitempty" form:"appAccessToken,omitempty"` //Specifies the token to access with the app.
    AppInstanceId    *int64          `json:"appInstanceId,omitempty" form:"appInstanceId,omitempty"` //Specifies unique id across all instances of all apps.
    AppName          *string         `json:"appName,omitempty" form:"appName,omitempty"` //Specifies name of the app that is launched in this instance.
    AppUid           *int64          `json:"appUid,omitempty" form:"appUid,omitempty"` //Specifies id of the app that is launched in this instance.
    AppVersion       *int64          `json:"appVersion,omitempty" form:"appVersion,omitempty"` //Specifies the version of the app that is launched in this instance.
    CreatedTimeUsecs *int64          `json:"createdTimeUsecs,omitempty" form:"createdTimeUsecs,omitempty"` //Specifies timestamp (in microseconds) when the app instance was first
    CreationUid      *string         `json:"creationUid,omitempty" form:"creationUid,omitempty"` //Specifies an unique identifier generated by the client to let the backend
    Description      *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies user configured description for the app instance.
    DurationUsecs    *int64          `json:"durationUsecs,omitempty" form:"durationUsecs,omitempty"` //Specifies duration (in microseconds) for which the app instance has run.
    HealthDetail     *string         `json:"healthDetail,omitempty" form:"healthDetail,omitempty"` //Specifies the reason the app instance is unhealthy. Only set if app
    HealthStatus     *int64          `json:"healthStatus,omitempty" form:"healthStatus,omitempty"` //Specifies the current health status of the app instance.
    NodeIp           *string         `json:"nodeIp,omitempty" form:"nodeIp,omitempty"` //Specifies the ip of the node which can be used to contact app instance
    NodePort         *int64          `json:"nodePort,omitempty" form:"nodePort,omitempty"` //Specifies the node port on which the app instance services external
    Settings         AppInstanceSettings `json:"settings,omitempty" form:"settings,omitempty"` //AppInstanceSettings provides settings used while launching an app instance.
    State            State1Enum      `json:"state,omitempty" form:"state,omitempty"` //Specifies the current state of the app instance.
    StateDetail      *string         `json:"stateDetail,omitempty" form:"stateDetail,omitempty"` //Specifies the failure reason when the app instance's state is kFailed.
}

/*
 * Structure for the custom type AWSProtectionSource
 */
type AWSProtectionSource struct {
    AccessKey          *string         `json:"accessKey,omitempty" form:"accessKey,omitempty"` //Specifies Access key of the AWS account.
    AmazonResourceName *string         `json:"amazonResourceName,omitempty" form:"amazonResourceName,omitempty"` //Specifies Amazon Resource Name (owner ID) of the IAM user, act as an
    AwsType            AwsTypeEnum     `json:"awsType,omitempty" form:"awsType,omitempty"` //Specifies the entity type such as 'kIAMUser' if the environment is kAWS.
    HostType           HostTypeEnum    `json:"hostType,omitempty" form:"hostType,omitempty"` //Specifies the OS type of the Protection Source of type 'kVirtualMachine'
    IpAddresses        *string         `json:"ipAddresses,omitempty" form:"ipAddresses,omitempty"` //Specifies the IP address of the entity of type 'kVirtualMachine'.
    Name               *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the Object set by the Cloud Provider.
    OwnerId            *string         `json:"ownerId,omitempty" form:"ownerId,omitempty"` //Specifies the owner id of the resource in AWS environment. With type,
    PhysicalSourceId   *int64          `json:"physicalSourceId,omitempty" form:"physicalSourceId,omitempty"` //Specifies the Protection Source id of the registered Physical Host.
    RegionId           *string         `json:"regionId,omitempty" form:"regionId,omitempty"` //Specifies the region Id of the entity if the entity is an EC2 instance.
    ResourceId         *string         `json:"resourceId,omitempty" form:"resourceId,omitempty"` //Specifies the unique Id of the resource given by the cloud provider.
    RestoreTaskId      *int64          `json:"restoreTaskId,omitempty" form:"restoreTaskId,omitempty"` //Specifies the id of the "convert and deploy" restore task that
    SecretAccessKey    *string         `json:"secretAccessKey,omitempty" form:"secretAccessKey,omitempty"` //Specifies Secret Access key of the AWS account.
    TagAttributes      []*VmwareTagAttributes `json:"tagAttributes,omitempty" form:"tagAttributes,omitempty"` //Specifies the list of AWS tag attributes.
    Type               Type1Enum       `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of an AWS Protection Source Object such as
    UserAccountId      *string         `json:"userAccountId,omitempty" form:"userAccountId,omitempty"` //Specifies the account id derived from the ARN of the user.
    UserResourceName   *string         `json:"userResourceName,omitempty" form:"userResourceName,omitempty"` //Specifies the Amazon Resource Name (ARN) of the user.
}

/*
 * Structure for the custom type AzureProtectionSource
 */
type AzureProtectionSource struct {
    ApplicationId    *string         `json:"applicationId,omitempty" form:"applicationId,omitempty"` //Specifies Application Id of the active directory of Azure account.
    ApplicationKey   *string         `json:"applicationKey,omitempty" form:"applicationKey,omitempty"` //Specifies Application key of the active directory of Azure account.
    AzureType        AzureTypeEnum   `json:"azureType,omitempty" form:"azureType,omitempty"` //Specifies the entity type such as 'kSubscription' if the environment is
    HostType         HostTypeEnum    `json:"hostType,omitempty" form:"hostType,omitempty"` //Specifies the OS type of the Protection Source of type 'kVirtualMachine'
    IpAddresses      *[]string       `json:"ipAddresses,omitempty" form:"ipAddresses,omitempty"` //Specifies a list of IP addresses for entities of type 'kVirtualMachine'.
    Location         *string         `json:"location,omitempty" form:"location,omitempty"` //Specifies the physical location of the resource group.
    MemoryMbytes     *int64          `json:"memoryMbytes,omitempty" form:"memoryMbytes,omitempty"` //Specifies the amount of memory in MegaBytes of the Azure resource of
    Name             *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the Object set by the Cloud Provider.
    NumCores         *int64          `json:"numCores,omitempty" form:"numCores,omitempty"` //Specifies the number of CPU cores of the Azure resource of
    PhysicalSourceId *int64          `json:"physicalSourceId,omitempty" form:"physicalSourceId,omitempty"` //Specifies the Protection Source id of the registered Physical Host.
    ResourceId       *string         `json:"resourceId,omitempty" form:"resourceId,omitempty"` //Specifies the unique Id of the resource given by the cloud provider.
    RestoreTaskId    *int64          `json:"restoreTaskId,omitempty" form:"restoreTaskId,omitempty"` //Specifies the id of the "convert and deploy" restore task that
    SubscriptionId   *string         `json:"subscriptionId,omitempty" form:"subscriptionId,omitempty"` //Specifies Subscription id inside a customer's Azure account. It represents
    TenantId         *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies Tenant Id of the active directory of Azure account.
    Type             Type2Enum       `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of an Azure Protection Source Object such as
}

/*
 * Structure for the custom type BackupJobProto
 */
type BackupJobProto struct {
    AbortInExclusionWindow               *bool           `json:"abortInExclusionWindow,omitempty" form:"abortInExclusionWindow,omitempty"` //This field determines whether a backup run should be aborted when it hits
    AlertingPolicy                       AlertingPolicyProto `json:"alertingPolicy,omitempty" form:"alertingPolicy,omitempty"` //TODO: Write general description for this field
    BackupQosPrincipal                   *int64          `json:"backupQosPrincipal,omitempty" form:"backupQosPrincipal,omitempty"` //The backup QoS principal to use for the backup job.
    BackupSourceParams                   []*BackupSourceParameters `json:"backupSourceParams,omitempty" form:"backupSourceParams,omitempty"` //This contains additional backup params that are applicable to sources
    ContinueOnQuiesceFailure             *bool           `json:"continueOnQuiesceFailure,omitempty" form:"continueOnQuiesceFailure,omitempty"` //Whether to continue backing up on quiesce failure.
    DedupDisabledSourceIdVec             *[]int64        `json:"dedupDisabledSourceIdVec,omitempty" form:"dedupDisabledSourceIdVec,omitempty"` //List of source ids for which source side dedup is disabled from the backup
    DeletionStatus                       *int64          `json:"deletionStatus,omitempty" form:"deletionStatus,omitempty"` //Determines if the job (and associated backups) should be deleted. Once a
    Description                          *string         `json:"description,omitempty" form:"description,omitempty"` //Job description (as entered by the user).
    DrToCloudParams                      DRToCloudParameters `json:"drToCloudParams,omitempty" form:"drToCloudParams,omitempty"` //A Proto needed in case objects backed up by this job need to DR to cloud.
    EhParentSource                       Entity          `json:"ehParentSource,omitempty" form:"ehParentSource,omitempty"` //Specifies the attributes and the latest statistics about an entity.
    EndTimeUsecs                         *int64          `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"` //The time (in usecs) after which no backup for the job should be scheduled.
    EnvBackupParams                      EnvBackupParams `json:"envBackupParams,omitempty" form:"envBackupParams,omitempty"` //Message to capture any additional environment specific backup params at the
    ExcludeSources                       []*BackupJobProtoExcludeSource `json:"excludeSources,omitempty" form:"excludeSources,omitempty"` //The list of sources to exclude from backups. These can have non-leaf-level
    ExcludeSourcesDEPRECATED             []*Entity       `json:"excludeSources_DEPRECATED,omitempty" form:"excludeSources_DEPRECATED,omitempty"` //The list of sources to exclude from backups. These can have non-leaf-level
    ExclusionRanges                      []*BackupJobProtoExclusionTimeRange `json:"exclusionRanges,omitempty" form:"exclusionRanges,omitempty"` //Do not run backups in these time-ranges.
    FullBackupJobPolicy                  DEPRECATEDIn50  `json:"fullBackupJobPolicy,omitempty" form:"fullBackupJobPolicy,omitempty"` //A message that specifies the policies to use for a job.
    FullBackupSlaTimeMins                *int64          `json:"fullBackupSlaTimeMins,omitempty" form:"fullBackupSlaTimeMins,omitempty"` //Same as 'sla_time_mins' above, but applies to full backups.
    IndexingPolicy                       IndexingPolicyProto `json:"indexingPolicy,omitempty" form:"indexingPolicy,omitempty"` //Proto to encapsulate file level indexing policy for VMs in a backup job.
    IsActive                             *bool           `json:"isActive,omitempty" form:"isActive,omitempty"` //Whether the backup job is active or not. Details about what an active job
    IsDeleted                            *bool           `json:"isDeleted,omitempty" form:"isDeleted,omitempty"` //Tracks whether the backup job has actually been deleted.
    IsPaused                             *bool           `json:"isPaused,omitempty" form:"isPaused,omitempty"` //Whether the backup job is paused. New backup runs are not scheduled for
    IsRpoJob                             *bool           `json:"isRpoJob,omitempty" form:"isRpoJob,omitempty"` //Whether the backup job is an RPO policy job. These jobs are hidden from
    JobCreationTimeUsecs                 *int64          `json:"jobCreationTimeUsecs,omitempty" form:"jobCreationTimeUsecs,omitempty"` //Time when this job was first created.
    JobId                                *int64          `json:"jobId,omitempty" form:"jobId,omitempty"` //A unique id for locally created jobs. This should only be used to identify
    JobPolicy                            DEPRECATEDIn50  `json:"jobPolicy,omitempty" form:"jobPolicy,omitempty"` //A message that specifies the policies to use for a job.
    JobUid                               UniversalIdProto `json:"jobUid,omitempty" form:"jobUid,omitempty"` //TODO: Write general description for this field
    LastModificationTimeUsecs            *int64          `json:"lastModificationTimeUsecs,omitempty" form:"lastModificationTimeUsecs,omitempty"` //Time when this job description was last updated.
    LastPauseModificationTimeUsecs       *int64          `json:"lastPauseModificationTimeUsecs,omitempty" form:"lastPauseModificationTimeUsecs,omitempty"` //Time when the job was last paused or unpaused.
    LastPauseReason                      *int64          `json:"lastPauseReason,omitempty" form:"lastPauseReason,omitempty"` //Last reason for pausing the backup job. Capturing the reason will help in
    LastUpdatedUsername                  *string         `json:"lastUpdatedUsername,omitempty" form:"lastUpdatedUsername,omitempty"` //The user who modified the job most recently.
    LeverageStorageSnapshots             *bool           `json:"leverageStorageSnapshots,omitempty" form:"leverageStorageSnapshots,omitempty"` //Whether to leverage the storage array based snapshots for this backup
    LeverageStorageSnapshotsForHyperflex *bool           `json:"leverageStorageSnapshotsForHyperflex,omitempty" form:"leverageStorageSnapshotsForHyperflex,omitempty"` //This is set to true by the user if hyperflex snapshots are requested
    LogBackupJobPolicy                   DEPRECATEDIn50  `json:"logBackupJobPolicy,omitempty" form:"logBackupJobPolicy,omitempty"` //A message that specifies the policies to use for a job.
    Name                                 *string         `json:"name,omitempty" form:"name,omitempty"` //The name of this backup job. This must be unique across all jobs.
    NumSnapshotsToKeepOnPrimary          *int64          `json:"numSnapshotsToKeepOnPrimary,omitempty" form:"numSnapshotsToKeepOnPrimary,omitempty"` //Specifies how many recent snapshots of each backed up entity to retain on
    ParentSource                         Entity          `json:"parentSource,omitempty" form:"parentSource,omitempty"` //Specifies the attributes and the latest statistics about an entity.
    PerformSourceSideDedup               *bool           `json:"performSourceSideDedup,omitempty" form:"performSourceSideDedup,omitempty"` //Whether or not to perform source side dedup.
    PolicyAppliedTimeMsecs               *int64          `json:"policyAppliedTimeMsecs,omitempty" form:"policyAppliedTimeMsecs,omitempty"` //Epoch time in milliseconds when the policy was last applied to this job.
    PolicyId                             *string         `json:"policyId,omitempty" form:"policyId,omitempty"` //Id of the policy being applied to the backup job. It is expected to be of
    PolicyName                           *string         `json:"policyName,omitempty" form:"policyName,omitempty"` //The name of the policy referred to by policy_uid. This field can be stale
    PostBackupScript                     BackupJobPreOrPostScript `json:"postBackupScript,omitempty" form:"postBackupScript,omitempty"` //A message to encapsulate the pre and post scripts associated with a backup
    PreScript                            BackupJobPreOrPostScript `json:"preScript,omitempty" form:"preScript,omitempty"` //A message to encapsulate the pre and post scripts associated with a backup
    PrimaryJobUid                        UniversalIdProto `json:"primaryJobUid,omitempty" form:"primaryJobUid,omitempty"` //TODO: Write general description for this field
    Priority                             *int64          `json:"priority,omitempty" form:"priority,omitempty"` //The priority for the job. This is used at admission time - all admitted
    Quiesce                              *bool           `json:"quiesce,omitempty" form:"quiesce,omitempty"` //Whether to take app-consistent snapshots by quiescing apps and the
    RemoteJobUids                        []*UniversalIdProto `json:"remoteJobUids,omitempty" form:"remoteJobUids,omitempty"` //The globally unique ids of all remote jobs that are linked to this job
    RequiredFeatureVec                   *[]string       `json:"requiredFeatureVec,omitempty" form:"requiredFeatureVec,omitempty"` //The features that are strictly required to be supported by the cluster
    SlaTimeMins                          *int64          `json:"slaTimeMins,omitempty" form:"slaTimeMins,omitempty"` //If specified, this variable determines the amount of time (after backup
    Sources                              []*BackupJobProtoBackupSource `json:"sources,omitempty" form:"sources,omitempty"` //The list of sources that should be backed up. A source in this list could
    StartTime                            *time.Time      `json:"startTime,omitempty" form:"startTime,omitempty"` //A message to encapusulate time of a day. Users of this proto will have to
    TagVec                               *[]string       `json:"tagVec,omitempty" form:"tagVec,omitempty"` //Tags associated with the job. User can specify tags/keywords that can
    Timezone                             *string         `json:"timezone,omitempty" form:"timezone,omitempty"` //Timezone of the backup job. All time fields (i.e., TimeOfDay) in this
    TruncateLogs                         *bool           `json:"truncateLogs,omitempty" form:"truncateLogs,omitempty"` //Whether to truncate logs after a backup run. This is currently only
    Type                                 *int64          `json:"type,omitempty" form:"type,omitempty"` //The type of environment this backup job corresponds to.
    UserInfo                             UserInformation1 `json:"userInfo,omitempty" form:"userInfo,omitempty"` //A message to encapsulate information about the user who made the request.
    ViewBoxId                            *int64          `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"` //The view box to which data will be written.
}

/*
 * Structure for the custom type BackupRunTask
 */
type BackupRunTask struct {
    Environment               Environment6Enum `json:"environment,omitempty" form:"environment,omitempty"` //Specifies the environment type that the task is protecting.
    Error                     *string         `json:"error,omitempty" form:"error,omitempty"` //Specifies if an error occurred (if any) while running this task.
    JobRunId                  *int64          `json:"jobRunId,omitempty" form:"jobRunId,omitempty"` //Specifies the id of the Job Run that ran the backup task and
    Message                   *string         `json:"message,omitempty" form:"message,omitempty"` //Specifies a message after finishing the task successfully. This field
    MetadataDeleted           *bool           `json:"metadataDeleted,omitempty" form:"metadataDeleted,omitempty"` //Specifies if the metadata and snapshots associated with this Job Run
    Quiesced                  *bool           `json:"quiesced,omitempty" form:"quiesced,omitempty"` //Specifies if app-consistent snapshot was captured. This field is set to
    RunType                   RunTypeEnum     `json:"runType,omitempty" form:"runType,omitempty"` //Specifies the type of backup such as 'kRegular', 'kFull', 'kLog' or
    SlaViolated               *bool           `json:"slaViolated,omitempty" form:"slaViolated,omitempty"` //Specifies if the SLA was violated for the Job Run. This field is set
    SnapshotsDeleted          *bool           `json:"snapshotsDeleted,omitempty" form:"snapshotsDeleted,omitempty"` //Specifies if backup snapshots associated
    SnapshotsDeletedTimeUsecs *int64          `json:"snapshotsDeletedTimeUsecs,omitempty" form:"snapshotsDeletedTimeUsecs,omitempty"` //Specifies if backup snapshots associated
    SourceBackupStatus        []*SourceObjectBackupStatus `json:"sourceBackupStatus,omitempty" form:"sourceBackupStatus,omitempty"` //Array of Source Object Backup Status.
    Stats                     ProtectionJobRunStatistics `json:"stats,omitempty" form:"stats,omitempty"` //Specifies statistics about a Protection Job Run.
    Status                    Status2Enum     `json:"status,omitempty" form:"status,omitempty"` //Specifies the status of Backup task such as 'kRunning', 'kSuccess',
    Warnings                  *[]string       `json:"warnings,omitempty" form:"warnings,omitempty"` //Array of Warnings.
    WormRetentionType         WormRetentionTypeEnum `json:"wormRetentionType,omitempty" form:"wormRetentionType,omitempty"` //Specifies WORM retention type for the snapshot as given by the policy.
}

/*
 * Structure for the custom type CloneViewRequest
 */
type CloneViewRequest struct {
    AccessSids                      *[]string       `json:"accessSids,omitempty" form:"accessSids,omitempty"` //Array of Security Identifiers (SIDs)
    CloneViewName                   *string         `json:"cloneViewName,omitempty" form:"cloneViewName,omitempty"` //Specifies the name of the new View that is cloned from the source View.
    DataLockExpiryUsecs             *int64          `json:"dataLockExpiryUsecs,omitempty" form:"dataLockExpiryUsecs,omitempty"` //DataLock (Write Once Read Many) lock expiry epoch time in microseconds. If
    Description                     *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies an optional text description about the View.
    EnableFilerAuditLogging         *bool           `json:"enableFilerAuditLogging,omitempty" form:"enableFilerAuditLogging,omitempty"` //Specifies if Filer Audit Logging is enabled for this view.
    EnableMixedModePermissions      *bool           `json:"enableMixedModePermissions,omitempty" form:"enableMixedModePermissions,omitempty"` //If set, mixed mode (NFS and SMB) access is enabled for this view.
    EnableNfsViewDiscovery          *bool           `json:"enableNfsViewDiscovery,omitempty" form:"enableNfsViewDiscovery,omitempty"` //If set, it enables discovery of view for NFS.
    EnableSmbAccessBasedEnumeration *bool           `json:"enableSmbAccessBasedEnumeration,omitempty" form:"enableSmbAccessBasedEnumeration,omitempty"` //Specifies if access-based enumeration should be enabled.
    EnableSmbEncryption             *bool           `json:"enableSmbEncryption,omitempty" form:"enableSmbEncryption,omitempty"` //Specifies the SMB encryption for the View. If set, it enables the SMB
    EnableSmbViewDiscovery          *bool           `json:"enableSmbViewDiscovery,omitempty" form:"enableSmbViewDiscovery,omitempty"` //If set, it enables discovery of view for SMB.
    EnforceSmbEncryption            *bool           `json:"enforceSmbEncryption,omitempty" form:"enforceSmbEncryption,omitempty"` //Specifies the SMB encryption for all the sessions for the View.
    FileExtensionFilter             FileExtensionFilter `json:"fileExtensionFilter,omitempty" form:"fileExtensionFilter,omitempty"` //TODO: Write general description for this field
    FileLockConfig                  FileLevelDataLockConfigurations `json:"fileLockConfig,omitempty" form:"fileLockConfig,omitempty"` //Specifies a config to lock files in a view - to protect from malicious or
    LogicalQuota                    QuotaPolicy     `json:"logicalQuota,omitempty" form:"logicalQuota,omitempty"` //Specifies an optional logical quota limit (in bytes) for the usage allowed
    ProtocolAccess                  ProtocolAccessEnum `json:"protocolAccess,omitempty" form:"protocolAccess,omitempty"` //Specifies the supported Protocols for the View.
    Qos                             QoS             `json:"qos,omitempty" form:"qos,omitempty"` //Specifies the Quality of Service (QoS) Policy for the View.
    SecurityMode                    SecurityModeEnum `json:"securityMode,omitempty" form:"securityMode,omitempty"` //Specifies the security mode used for this view.
    SmbPermissionsInfo              SMBPermissionsInformation `json:"smbPermissionsInfo,omitempty" form:"smbPermissionsInfo,omitempty"` //Specifies information about SMB permissions.
    SourceViewName                  *string         `json:"sourceViewName,omitempty" form:"sourceViewName,omitempty"` //Specifies the name of the source View that will be cloned.
    StoragePolicyOverride           StoragePolicyOverride `json:"storagePolicyOverride,omitempty" form:"storagePolicyOverride,omitempty"` //Specifies if inline deduplication and compression settings inherited from
    SubnetWhitelist                 []*Subnet       `json:"subnetWhitelist,omitempty" form:"subnetWhitelist,omitempty"` //Array of Subnets.
    TenantId                        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Optional tenant id who has access to this View.
}

/*
 * Structure for the custom type CohesityCluster
 */
type CohesityCluster struct {
    AppsSettings                    AthenaAppsConfiguration `json:"appsSettings,omitempty" form:"appsSettings,omitempty"` //TODO: Write general description for this field
    AvailableMetadataSpace          *int64          `json:"availableMetadataSpace,omitempty" form:"availableMetadataSpace,omitempty"` //Information about storage available for metadata
    BondingMode                     BondingModeEnum `json:"bondingMode,omitempty" form:"bondingMode,omitempty"` //Specifies the bonding mode to use when bonding NICs to this Cluster.
    ClusterAuditLogConfig           ClusterAuditLogConfiguration `json:"clusterAuditLogConfig,omitempty" form:"clusterAuditLogConfig,omitempty"` //Specifies the settings of the Cluster audit log configuration.
    ClusterSoftwareVersion          *string         `json:"clusterSoftwareVersion,omitempty" form:"clusterSoftwareVersion,omitempty"` //Specifies the current release of the Cohesity software running on
    ClusterType                     ClusterType1Enum `json:"clusterType,omitempty" form:"clusterType,omitempty"` //Specifies the type of Cluster such as kPhysical.
    CreatedTimeMsecs                *int64          `json:"createdTimeMsecs,omitempty" form:"createdTimeMsecs,omitempty"` //Specifies the time when the Cohesity Cluster was created.
    CurrentOpScheduledTimeSecs      *int64          `json:"currentOpScheduledTimeSecs,omitempty" form:"currentOpScheduledTimeSecs,omitempty"` //Specifies the time scheduled by the Cohesity Cluster to
    CurrentOperation                CurrentOperationEnum `json:"currentOperation,omitempty" form:"currentOperation,omitempty"` //Specifies the current Cluster-level operation in progress.
    CurrentTimeMsecs                *int64          `json:"currentTimeMsecs,omitempty" form:"currentTimeMsecs,omitempty"` //Specifies the current system time on the Cohesity Cluster.
    DnsServerIps                    *[]string       `json:"dnsServerIps,omitempty" form:"dnsServerIps,omitempty"` //Array of IP Addresses of DNS Servers.
    DomainNames                     *[]string       `json:"domainNames,omitempty" form:"domainNames,omitempty"` //Array of Domain Names.
    EnableActiveMonitoring          *bool           `json:"enableActiveMonitoring,omitempty" form:"enableActiveMonitoring,omitempty"` //Specifies if Cohesity can receive monitoring information from the
    EnableUpgradePkgPolling         *bool           `json:"enableUpgradePkgPolling,omitempty" form:"enableUpgradePkgPolling,omitempty"` //If 'true', Cohesity's upgrade server is polled for new releases.
    EncryptionEnabled               *bool           `json:"encryptionEnabled,omitempty" form:"encryptionEnabled,omitempty"` //If 'true', the entire Cohesity Cluster is encrypted including all View
    EncryptionKeyRotationPeriodSecs *int64          `json:"encryptionKeyRotationPeriodSecs,omitempty" form:"encryptionKeyRotationPeriodSecs,omitempty"` //Specifies the period of time (in seconds) when encryption keys are rotated.
    EulaConfig                      EULAConfiguration `json:"eulaConfig,omitempty" form:"eulaConfig,omitempty"` //Specifies the End User License Agreement (EULA) acceptance information.
    FilerAuditLogConfig             FilerAuditLogConfiguration `json:"filerAuditLogConfig,omitempty" form:"filerAuditLogConfig,omitempty"` //Specifies the settings of the filer audit log configuration.
    FipsModeEnabled                 *bool           `json:"fipsModeEnabled,omitempty" form:"fipsModeEnabled,omitempty"` //Specifies if the Cohesity Cluster should operate in the FIPS mode,
    Gateway                         *string         `json:"gateway,omitempty" form:"gateway,omitempty"` //Specifies the gateway IP address.
    GoogleAnalyticsEnabled          *bool           `json:"googleAnalyticsEnabled,omitempty" form:"googleAnalyticsEnabled,omitempty"` //Specified whether Google Analytics is enabled.
    HardwareInfo                    CohesityClusterHardware `json:"hardwareInfo,omitempty" form:"hardwareInfo,omitempty"` //Specifies a hardware type for motherboard of the Nodes
    Id                              *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies the unique id of Cohesity Cluster.
    IncarnationId                   *int64          `json:"incarnationId,omitempty" form:"incarnationId,omitempty"` //Specifies the unique incarnation id of the Cohesity Cluster.
    IsDocumentationLocal            *bool           `json:"isDocumentationLocal,omitempty" form:"isDocumentationLocal,omitempty"` //Specifies what version of the documentation is used.
    LanguageLocale                  *string         `json:"languageLocale,omitempty" form:"languageLocale,omitempty"` //Specifies the language and locale for this Cohesity Cluster.
    LicenseServerClaimed            *bool           `json:"licenseServerClaimed,omitempty" form:"licenseServerClaimed,omitempty"` //Speifies if cluster is claimed by Helios or not.
    MetadataFaultToleranceFactor    *int64          `json:"metadataFaultToleranceFactor,omitempty" form:"metadataFaultToleranceFactor,omitempty"` //Specifies metadata fault tolerance setting for the cluster. This denotes
    Mtu                             *int64          `json:"mtu,omitempty" form:"mtu,omitempty"` //Specifies the Maxium Transmission Unit (MTU) in bytes of
    MultiTenancyEnabled             *bool           `json:"multiTenancyEnabled,omitempty" form:"multiTenancyEnabled,omitempty"` //Specifies if multi tenancy is enabled in the cluster. Authentication &
    Name                            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the Cohesity Cluster.
    NodeCount                       *int64          `json:"nodeCount,omitempty" form:"nodeCount,omitempty"` //Specifies the number of Nodes in the Cohesity Cluster.
    NtpSettings                     NtpSettingsConfig `json:"ntpSettings,omitempty" form:"ntpSettings,omitempty"` //TODO: Write general description for this field
    ReverseTunnelEnabled            *bool           `json:"reverseTunnelEnabled,omitempty" form:"reverseTunnelEnabled,omitempty"` //If 'true', Cohesity's Remote Tunnel is enabled.
    ReverseTunnelEndTimeMsecs       *int64          `json:"reverseTunnelEndTimeMsecs,omitempty" form:"reverseTunnelEndTimeMsecs,omitempty"` //ReverseTunnelEndTimeMsecs specifies the end time in milliseconds since
    SmbAdDisabled                   *bool           `json:"smbAdDisabled,omitempty" form:"smbAdDisabled,omitempty"` //Specifies if Active Directory should be disabled for authentication of SMB
    Stats                           CohesityClusterStatistics `json:"stats,omitempty" form:"stats,omitempty"` //Specifies statistics about this Cohesity Cluster.
    StigMode                        *bool           `json:"stigMode,omitempty" form:"stigMode,omitempty"` //Specifies if STIG mode is enabled or not.
    SupportedConfig                 SupportedErasureCodingAndNodeConfigurations `json:"supportedConfig,omitempty" form:"supportedConfig,omitempty"` //Lists the supported Erasure Coding options for the number of
    SyslogServers                   []*SyslogServerConfiguration `json:"syslogServers,omitempty" form:"syslogServers,omitempty"` //Array of Syslog Servers.
    TargetSoftwareVersion           *string         `json:"targetSoftwareVersion,omitempty" form:"targetSoftwareVersion,omitempty"` //Specifies the Cohesity release that this Cluster is being upgraded to
    TenantViewboxSharingEnabled     *bool           `json:"tenantViewboxSharingEnabled,omitempty" form:"tenantViewboxSharingEnabled,omitempty"` //In case multi tenancy is enabled, this flag controls whether multiple
    Timezone                        *string         `json:"timezone,omitempty" form:"timezone,omitempty"` //Specifies the timezone to use for showing time in emails, reports,
    TurboMode                       *bool           `json:"turboMode,omitempty" form:"turboMode,omitempty"` //Specifies if the cluster is in Turbo mode.
    UsedMetadataSpacePct            *float64        `json:"usedMetadataSpacePct,omitempty" form:"usedMetadataSpacePct,omitempty"` //UsedMetadataSpacePct measures the percentage about storage used for
}

/*
 * Structure for the custom type StorageDomainViewBoxRequest
 */
type StorageDomainViewBoxRequest struct {
    AdDomainName                    *string         `json:"adDomainName,omitempty" form:"adDomainName,omitempty"` //Specifies an active directory domain that this view box is mapped to.
    ClientSubnetWhiteList           []*Subnet       `json:"clientSubnetWhiteList,omitempty" form:"clientSubnetWhiteList,omitempty"` //Array of Subnets.
    CloudDownWaterfallThresholdPct  *int64          `json:"cloudDownWaterfallThresholdPct,omitempty" form:"cloudDownWaterfallThresholdPct,omitempty"` //Specifies the cloud down water-fall threshold percentage. This indicates
    CloudDownWaterfallThresholdSecs *int64          `json:"cloudDownWaterfallThresholdSecs,omitempty" form:"cloudDownWaterfallThresholdSecs,omitempty"` //Specifies the cloud down water-fall threshold seconds. This indicates
    ClusterPartitionId              int64           `json:"clusterPartitionId" form:"clusterPartitionId"` //Specifies the Cluster Partition id where the Storage Domain (View Box) is
    DefaultUserQuotaPolicy          QuotaPolicy     `json:"defaultUserQuotaPolicy,omitempty" form:"defaultUserQuotaPolicy,omitempty"` //Specifies an optional quota policy/limits that are inherited by all users
    DefaultViewQuotaPolicy          QuotaPolicy     `json:"defaultViewQuotaPolicy,omitempty" form:"defaultViewQuotaPolicy,omitempty"` //Specifies an optional default logical quota limit (in bytes)
    LdapProviderId                  *int64          `json:"ldapProviderId,omitempty" form:"ldapProviderId,omitempty"` //When set, the following provides the LDAP provider the view box is
    Name                            string          `json:"name" form:"name"` //Specifies the name of the Storage Domain (View Box).
    PhysicalQuota                   QuotaPolicy     `json:"physicalQuota,omitempty" form:"physicalQuota,omitempty"` //Specifies an optional quota limit (in bytes) for the physical
    S3BucketsAllowed                *bool           `json:"s3BucketsAllowed,omitempty" form:"s3BucketsAllowed,omitempty"` //Specifies whether creation of a S3 bucket is allowed in this
    StoragePolicy                   StoragePolicy   `json:"storagePolicy,omitempty" form:"storagePolicy,omitempty"` //Specifies the storage options applied to a Storage Domain (View Box).
    TenantIdVec                     *[]string       `json:"tenantIdVec,omitempty" form:"tenantIdVec,omitempty"` //Optional ids for the tenants that this view box belongs. This must be
}

/*
 * Structure for the custom type CreateViewRequest
 */
type CreateViewRequest struct {
    AccessSids                      *[]string       `json:"accessSids,omitempty" form:"accessSids,omitempty"` //Array of Security Identifiers (SIDs)
    CaseInsensitiveNamesEnabled     *bool           `json:"caseInsensitiveNamesEnabled,omitempty" form:"caseInsensitiveNamesEnabled,omitempty"` //Specifies whether to support case insensitive file/folder names. This
    Description                     *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies an optional text description about the View.
    EnableFilerAuditLogging         *bool           `json:"enableFilerAuditLogging,omitempty" form:"enableFilerAuditLogging,omitempty"` //Specifies if Filer Audit Logging is enabled for this view.
    EnableMixedModePermissions      *bool           `json:"enableMixedModePermissions,omitempty" form:"enableMixedModePermissions,omitempty"` //If set, mixed mode (NFS and SMB) access is enabled for this view.
    EnableNfsViewDiscovery          *bool           `json:"enableNfsViewDiscovery,omitempty" form:"enableNfsViewDiscovery,omitempty"` //If set, it enables discovery of view for NFS.
    EnableSmbAccessBasedEnumeration *bool           `json:"enableSmbAccessBasedEnumeration,omitempty" form:"enableSmbAccessBasedEnumeration,omitempty"` //Specifies if access-based enumeration should be enabled.
    EnableSmbEncryption             *bool           `json:"enableSmbEncryption,omitempty" form:"enableSmbEncryption,omitempty"` //Specifies the SMB encryption for the View. If set, it enables the SMB
    EnableSmbViewDiscovery          *bool           `json:"enableSmbViewDiscovery,omitempty" form:"enableSmbViewDiscovery,omitempty"` //If set, it enables discovery of view for SMB.
    EnforceSmbEncryption            *bool           `json:"enforceSmbEncryption,omitempty" form:"enforceSmbEncryption,omitempty"` //Specifies the SMB encryption for all the sessions for the View.
    FileExtensionFilter             FileExtensionFilter `json:"fileExtensionFilter,omitempty" form:"fileExtensionFilter,omitempty"` //TODO: Write general description for this field
    FileLockConfig                  FileLevelDataLockConfigurations `json:"fileLockConfig,omitempty" form:"fileLockConfig,omitempty"` //Specifies a config to lock files in a view - to protect from malicious or
    LogicalQuota                    QuotaPolicy     `json:"logicalQuota,omitempty" form:"logicalQuota,omitempty"` //Specifies an optional logical quota limit (in bytes) for the usage allowed
    Name                            string          `json:"name" form:"name"` //Specifies the name of the new View to create.
    ProtocolAccess                  ProtocolAccessEnum `json:"protocolAccess,omitempty" form:"protocolAccess,omitempty"` //Specifies the supported Protocols for the View.
    Qos                             QoS             `json:"qos,omitempty" form:"qos,omitempty"` //Specifies the Quality of Service (QoS) Policy for the View.
    SecurityMode                    SecurityModeEnum `json:"securityMode,omitempty" form:"securityMode,omitempty"` //Specifies the security mode used for this view.
    SmbPermissionsInfo              SMBPermissionsInformation `json:"smbPermissionsInfo,omitempty" form:"smbPermissionsInfo,omitempty"` //Specifies information about SMB permissions.
    StoragePolicyOverride           StoragePolicyOverride `json:"storagePolicyOverride,omitempty" form:"storagePolicyOverride,omitempty"` //Specifies if inline deduplication and compression settings inherited from
    SubnetWhitelist                 []*Subnet       `json:"subnetWhitelist,omitempty" form:"subnetWhitelist,omitempty"` //Array of Subnets.
    TenantId                        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Optional tenant id who has access to this View.
    ViewBoxId                       int64           `json:"viewBoxId" form:"viewBoxId"` //Specifies the id of the Storage Domain (View Box) where the View will be
}

/*
 * Structure for the custom type GCPProtectionSource
 */
type GCPProtectionSource struct {
    ClientEmailAddress *string         `json:"clientEmailAddress,omitempty" form:"clientEmailAddress,omitempty"` //Specifies Client email address associated with the service account.
    ClientPrivateKey   *string         `json:"clientPrivateKey,omitempty" form:"clientPrivateKey,omitempty"` //Specifies Client private associated with the service account.
    GcpType            GcpTypeEnum     `json:"gcpType,omitempty" form:"gcpType,omitempty"` //Specifies the entity type such as 'kIAMUser' if the environment is kGCP.
    HostType           HostTypeEnum    `json:"hostType,omitempty" form:"hostType,omitempty"` //Specifies the OS type of the Protection Source of type 'kVirtualMachine'
    IpAddressesVM      *string         `json:"ipAddressesVM,omitempty" form:"ipAddressesVM,omitempty"` //Specifies the IP address of the entity of type 'kVirtualMachine'.
    Name               *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the Object set by the Cloud Provider.
    OwnerId            *string         `json:"ownerId,omitempty" form:"ownerId,omitempty"` //Specifies the owner id of the resource in GCP environment. With type,
    PhysicalSourceId   *int64          `json:"physicalSourceId,omitempty" form:"physicalSourceId,omitempty"` //Specifies the Protection Source id of the registered Physical Host.
    ProjectId          *string         `json:"projectId,omitempty" form:"projectId,omitempty"` //Specifies the project Id.
    RegionId           *string         `json:"regionId,omitempty" form:"regionId,omitempty"` //Specifies the region Id.
    ResourceId         *string         `json:"resourceId,omitempty" form:"resourceId,omitempty"` //Specifies the unique Id of the resource given by the cloud provider.
    RestoreTaskId      *int64          `json:"restoreTaskId,omitempty" form:"restoreTaskId,omitempty"` //Specifies the id of the "convert and deploy" restore task that
    Type               Type4Enum       `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of an GCP Protection Source Object such as
    VpcNetwork         *string         `json:"vpcNetwork,omitempty" form:"vpcNetwork,omitempty"` //Specifies the VPC Network to deploy proxy VMs.
    VpcSubnetwork      *string         `json:"vpcSubnetwork,omitempty" form:"vpcSubnetwork,omitempty"` //Specifies the subnetwork to deploy proxy VMs.
}

/*
 * Structure for the custom type LatestProtectionRun
 */
type LatestProtectionRun struct {
    BackupRun           SourceObjectBackupStatus `json:"backupRun,omitempty" form:"backupRun,omitempty"` //Specifies the source object to protect and the current backup status.
    ChangeEventId       *int64          `json:"changeEventId,omitempty" form:"changeEventId,omitempty"` //Specifies the event id which caused last update on this object.
    CopyRun             CopyRunTask     `json:"copyRun,omitempty" form:"copyRun,omitempty"` //Specifies details about the Copy Run for a backup run of a Job Run.
    JobRunId            *int64          `json:"jobRunId,omitempty" form:"jobRunId,omitempty"` //Specifies job run id of the latest successful Protection Job Run.
    ProtectionJobRunUid RunUID          `json:"protectionJobRunUid,omitempty" form:"protectionJobRunUid,omitempty"` //Specifies the universal id of the latest successful Protection Job Run.
    SnapshotTarget      *string         `json:"snapshotTarget,omitempty" form:"snapshotTarget,omitempty"` //Specifies the cluster id in case of local or replication snapshots and
    SnapshotTargetType  *int64          `json:"snapshotTargetType,omitempty" form:"snapshotTargetType,omitempty"` //Specifies the snapshot target type of the latest snapshot.
    TaskStatus          *int64          `json:"taskStatus,omitempty" form:"taskStatus,omitempty"` //Specifies the task status of the Protecion Job Run in the final attempt.
    Uuid                *string         `json:"uuid,omitempty" form:"uuid,omitempty"` //Specifies the unique id of the Protection Source for which a snapshot is
}

/*
 * Structure for the custom type LaunchAppInstance
 */
type LaunchAppInstance struct {
    AppUid          *int64          `json:"appUid,omitempty" form:"appUid,omitempty"` //Specifies the app Id.
    AppVersion      *int64          `json:"appVersion,omitempty" form:"appVersion,omitempty"` //Specifies the app version.
    CreationUid     *string         `json:"creationUid,omitempty" form:"creationUid,omitempty"` //Specifies unique identifier generated by the client to let the backend distinguish
    Description     *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies user configured description for the app instance.
    Settings        AppInstanceSettings `json:"settings,omitempty" form:"settings,omitempty"` //AppInstanceSettings provides settings used while launching an app instance.
}

/*
 * Structure for the custom type LDAP
 */
type LDAP struct {
    AdDomainName            *string         `json:"adDomainName,omitempty" form:"adDomainName,omitempty"` //Specifies the domain name of an Active Directory which is mapped to this
    AuthType                AuthTypeEnum    `json:"authType,omitempty" form:"authType,omitempty"` //Specifies the authentication type used while connecting to LDAP servers.
    BaseDistinguishedName   *string         `json:"baseDistinguishedName,omitempty" form:"baseDistinguishedName,omitempty"` //Specifies the base distinguished name used as the base for LDAP
    DomainName              *string         `json:"domainName,omitempty" form:"domainName,omitempty"` //Specifies the name of the domain name to be used for querying LDAP servers
    Name                    *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the LDAP provider.
    Port                    *int64          `json:"port,omitempty" form:"port,omitempty"` //Specifies LDAP server port.
    PreferredLdapServerList *[]string       `json:"preferredLdapServerList,omitempty" form:"preferredLdapServerList,omitempty"` //Specifies the preferred LDAP servers. Server names should be either in
    TenantId                *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
    UseSsl                  *bool           `json:"useSsl,omitempty" form:"useSsl,omitempty"` //Specifies whether to use SSL for LDAP connections.
    UserDistinguishedName   *string         `json:"userDistinguishedName,omitempty" form:"userDistinguishedName,omitempty"` //Specifies the user distinguished name that is used for LDAP authentication.
    UserPassword            *string         `json:"userPassword,omitempty" form:"userPassword,omitempty"` //Specifies the user password that is used for LDAP authentication.
}

/*
 * Structure for the custom type LDAPProviderResponse
 */
type LDAPProviderResponse struct {
    AdDomainName            *string         `json:"adDomainName,omitempty" form:"adDomainName,omitempty"` //Specifies the domain name of an Active Directory which is mapped to this
    AuthType                AuthTypeEnum    `json:"authType,omitempty" form:"authType,omitempty"` //Specifies the authentication type used while connecting to LDAP servers.
    BaseDistinguishedName   *string         `json:"baseDistinguishedName,omitempty" form:"baseDistinguishedName,omitempty"` //Specifies the base distinguished name used as the base for LDAP
    DomainName              *string         `json:"domainName,omitempty" form:"domainName,omitempty"` //Specifies the name of the domain name to be used for querying LDAP servers
    Id                      *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies the ID of the LDAP provider.
    Name                    *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the LDAP provider.
    Port                    *int64          `json:"port,omitempty" form:"port,omitempty"` //Specifies LDAP server port.
    PreferredLdapServerList *[]string       `json:"preferredLdapServerList,omitempty" form:"preferredLdapServerList,omitempty"` //Specifies the preferred LDAP servers. Server names should be either in
    TenantId                *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
    UseSsl                  *bool           `json:"useSsl,omitempty" form:"useSsl,omitempty"` //Specifies whether to use SSL for LDAP connections.
    UserDistinguishedName   *string         `json:"userDistinguishedName,omitempty" form:"userDistinguishedName,omitempty"` //Specifies the user distinguished name that is used for LDAP authentication.
    UserPassword            *string         `json:"userPassword,omitempty" form:"userPassword,omitempty"` //Specifies the user password that is used for LDAP authentication.
}

/*
 * Structure for the custom type ListCentrifyZone
 */
type ListCentrifyZone struct {
    CentrifySchema    CentrifySchemaEnum `json:"centrifySchema,omitempty" form:"centrifySchema,omitempty"` //Specifies the schema of this Centrify zone.
    Description       *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description of the Centrify zone.
    DistinguishedName *string         `json:"distinguishedName,omitempty" form:"distinguishedName,omitempty"` //Specifies the distinguished name of the Centrify zone.
    ZoneName          *string         `json:"zoneName,omitempty" form:"zoneName,omitempty"` //Specifies the zone name.
}

/*
 * Structure for the custom type LogicalVolume
 */
type LogicalVolume struct {
    DeviceRootNode  DeviceTree      `json:"deviceRootNode,omitempty" form:"deviceRootNode,omitempty"` //Specifies a logical volume stored as a tree where the leaves are
    GroupName       *string         `json:"groupName,omitempty" form:"groupName,omitempty"` //Specifies the group name of the logical volume.
    GroupUuid       *string         `json:"groupUuid,omitempty" form:"groupUuid,omitempty"` //Specifies the group uuid of the logical volume.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the logical volume.
    Uuid            *string         `json:"uuid,omitempty" form:"uuid,omitempty"` //Specifies the uuid of the logical volume.
}

/*
 * Structure for the custom type MapReduceInformation
 */
type MapReduceInformation struct {
    AppProperty           MapReduceInfoAppProperty `json:"appProperty,omitempty" form:"appProperty,omitempty"` //AppProperty message encapsulates properties that are same across all run
    AuxData               MapReduceAuxData `json:"auxData,omitempty" form:"auxData,omitempty"` //This message encapsulates auxillary data for a MapReduce. One example of
    Description           *string         `json:"description,omitempty" form:"description,omitempty"` //Map reduce job description.
    ExcludedDataSourceVec *[]int64        `json:"excludedDataSourceVec,omitempty" form:"excludedDataSourceVec,omitempty"` //List of all excluded data sources for this app.
    Id                    *int64          `json:"id,omitempty" form:"id,omitempty"` //ID of map reduce job.
    IsSystemDefined       *bool           `json:"isSystemDefined,omitempty" form:"isSystemDefined,omitempty"` //Flag to denote if this is system pre-defined app or user has written this
    MapperId              *int64          `json:"mapperId,omitempty" form:"mapperId,omitempty"` //ID of the mapper to process the input.
    Name                  *string         `json:"name,omitempty" form:"name,omitempty"` //Map reduce job name.
    ReducerId             *int64          `json:"reducerId,omitempty" form:"reducerId,omitempty"` //ID of the reducer.
    RequiredPropertyVec   []*MapReduceInfoRequiredProperty `json:"requiredPropertyVec,omitempty" form:"requiredPropertyVec,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type MapReduceInstance
 */
type MapReduceInstance struct {
    Id              *int64          `json:"id,omitempty" form:"id,omitempty"` //System generated ID of map reduce instance.
    InputParams     []*MapReduceInstanceInputParam `json:"inputParams,omitempty" form:"inputParams,omitempty"` //TODO: Write general description for this field
    InputSpec       InputSelectorSelectsTheFilesToMapOver `json:"inputSpec,omitempty" form:"inputSpec,omitempty"` //TODO: Write general description for this field
    MapReduceInfoId *int64          `json:"mapReduceInfoId,omitempty" form:"mapReduceInfoId,omitempty"` //ID of Map reduce info.
    OutputSpec      OutputSpecificationForTheMapreduce `json:"outputSpec,omitempty" form:"outputSpec,omitempty"` //TODO: Write general description for this field
    RunInfo         StoresTheProgressOfRunOfThisInstance `json:"runInfo,omitempty" form:"runInfo,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type StoresTheProgressOfRunOfThisInstance
 */
type StoresTheProgressOfRunOfThisInstance struct {
    EndTime                     *int64          `json:"endTime,omitempty" form:"endTime,omitempty"` //Time when map redcue job completed.
    ErrorMessage                *string         `json:"errorMessage,omitempty" form:"errorMessage,omitempty"` //If this run failed, then error message for failure.
    ExecutionStartTimeUsecs     *int64          `json:"executionStartTimeUsecs,omitempty" form:"executionStartTimeUsecs,omitempty"` //Time (in usecs) when job was picked up for execution.
    FilesProcessed              *int64          `json:"filesProcessed,omitempty" form:"filesProcessed,omitempty"` //Number of files processed in this run.
    MapDoneTimeUsecs            *int64          `json:"mapDoneTimeUsecs,omitempty" form:"mapDoneTimeUsecs,omitempty"` //Time (in usecs) when map tasks were done.
    MapInputBytes               *int64          `json:"mapInputBytes,omitempty" form:"mapInputBytes,omitempty"` //Total size of data processed by this run in bytes.
    MappersSpawned              *int64          `json:"mappersSpawned,omitempty" form:"mappersSpawned,omitempty"` //Number of mappers spawned till now.
    NumMapOutputs               *int64          `json:"numMapOutputs,omitempty" form:"numMapOutputs,omitempty"` //Number of outputs from mappers.
    NumReduceOutputs            *int64          `json:"numReduceOutputs,omitempty" form:"numReduceOutputs,omitempty"` //Number of outputs from reducers.
    PercentageCompletion        *float64        `json:"percentageCompletion,omitempty" form:"percentageCompletion,omitempty"` //Percentage completion of this run so far.
    PercentageMapperCompletion  *float64        `json:"percentageMapperCompletion,omitempty" form:"percentageMapperCompletion,omitempty"` //Percentage of mapper phase completed.
    PercentageReducerCompletion *float64        `json:"percentageReducerCompletion,omitempty" form:"percentageReducerCompletion,omitempty"` //Percentage of reducer phase completed.
    ReducersSpawned             *int64          `json:"reducersSpawned,omitempty" form:"reducersSpawned,omitempty"` //Number of reducers spawned till now.
    RemainingTimeMins           *int64          `json:"remainingTimeMins,omitempty" form:"remainingTimeMins,omitempty"` //Expected remaining time in minutes for completion of this run.
    StartTime                   *int64          `json:"startTime,omitempty" form:"startTime,omitempty"` //Time when map reduce job was started by user.
    Status                      *int64          `json:"status,omitempty" form:"status,omitempty"` //Status of this run.
    TotalNumMappers             *int64          `json:"totalNumMappers,omitempty" form:"totalNumMappers,omitempty"` //Total number of mappers to be spawned.
    TotalNumReducers            *int64          `json:"totalNumReducers,omitempty" form:"totalNumReducers,omitempty"` //Total number of reducers to be spawned.
}

/*
 * Structure for the custom type InformationAboutAMapper
 */
type InformationAboutAMapper struct {
    Code            *string         `json:"code,omitempty" form:"code,omitempty"` //The code of the mapper in the specified language. Should be UTF-8.
    Id              *int64          `json:"id,omitempty" form:"id,omitempty"` //Mapper ID generated by system. Absent when user is creating a new mapper.
    IsSystemDefined *bool           `json:"isSystemDefined,omitempty" form:"isSystemDefined,omitempty"` //Whether the mapper is system defined.
    JarName         *string         `json:"jarName,omitempty" form:"jarName,omitempty"` //User can write their own mapper/reducer or upload jar files containing
    JarPath         *string         `json:"jarPath,omitempty" form:"jarPath,omitempty"` //path of JAR in which this mapper was found. This is applicable only when
    Language        *int64          `json:"language,omitempty" form:"language,omitempty"` //Language of the mapper.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Name of the mapper.
}

/*
 * Structure for the custom type MonthlySchedule
 */
type MonthlySchedule struct {
    Day             Day3Enum        `json:"day,omitempty" form:"day,omitempty"` //Specifies the day of the week (such as 'kMonday') to start the Job Run.
    DayCount        DayCountEnum    `json:"dayCount,omitempty" form:"dayCount,omitempty"` //Specifies the day count in the month (such as 'kThird') to start
}

/*
 * Structure for the custom type MountingVolumes
 */
type MountingVolumes struct {
    BringDisksOnline *bool           `json:"bringDisksOnline,omitempty" form:"bringDisksOnline,omitempty"` //Optional setting that determines if the volumes are brought
    Password         *string         `json:"password,omitempty" form:"password,omitempty"` //Specifies password of the username to access the target source.
    TargetSourceId   *int64          `json:"targetSourceId,omitempty" form:"targetSourceId,omitempty"` //Specifies the target Protection Source id where the volumes will be
    Username         *string         `json:"username,omitempty" form:"username,omitempty"` //Specifies username to access the target source.
    VolumeNames      *[]string       `json:"volumeNames,omitempty" form:"volumeNames,omitempty"` //Array of Volume Names.
}

/*
 * Structure for the custom type MountVolumesStates
 */
type MountVolumesStates struct {
    BringDisksOnline   *bool           `json:"bringDisksOnline,omitempty" form:"bringDisksOnline,omitempty"` //Optional setting that determines if the volumes are brought
    MountVolumeResults []*MountVolumeResultDetails `json:"mountVolumeResults,omitempty" form:"mountVolumeResults,omitempty"` //Array of Mount Volume Results.
    OtherError         Error           `json:"otherError,omitempty" form:"otherError,omitempty"` //Specifies an error that did not occur during the mount operation.
    TargetSourceId     *int64          `json:"targetSourceId,omitempty" form:"targetSourceId,omitempty"` //Specifies the target Protection Source Id where the volumes will be
    Username           *string         `json:"username,omitempty" form:"username,omitempty"` //Specifies the username to access the mount target.
}

/*
 * Structure for the custom type NASBackupParameters
 */
type NASBackupParameters struct {
    ContinueOnError       *bool           `json:"continueOnError,omitempty" form:"continueOnError,omitempty"` //Whether the backup job should continue on errors for snapshot based
    FilteringPolicy       FilteringPolicyProto `json:"filteringPolicy,omitempty" form:"filteringPolicy,omitempty"` //Proto to encapsulate the filtering policy for backup objects like files or
    MixedModePreference   *int64          `json:"mixedModePreference,omitempty" form:"mixedModePreference,omitempty"` //If the target entity is a mixed mode volume, which NAS protocol type the
    SnapshotChangeEnabled *bool           `json:"snapshotChangeEnabled,omitempty" form:"snapshotChangeEnabled,omitempty"` //Whether this backup job should utilize changelist like API when available
}

/*
 * Structure for the custom type NASServerCredentials
 */
type NASServerCredentials struct {
    Host            *string         `json:"host,omitempty" form:"host,omitempty"` //Specifies the hostname or IP address of the NAS server.
    MountPath       *string         `json:"mountPath,omitempty" form:"mountPath,omitempty"` //Specifies the mount path to the NAS server.
    Password        *string         `json:"password,omitempty" form:"password,omitempty"` //If using the CIFS protocol and a Username was specified, specify
    ShareType       ShareTypeEnum   `json:"shareType,omitempty" form:"shareType,omitempty"` //Specifies the sharing protocol type used to mount the file system.
    Username        *string         `json:"username,omitempty" form:"username,omitempty"` //If using the CIFS protocol, you can optional specify a username
}

/*
 * Structure for the custom type NASEnvironmentJobParameters
 */
type NASEnvironmentJobParameters struct {
    ContinueOnError                *bool           `json:"continueOnError,omitempty" form:"continueOnError,omitempty"` //Specifies if the backup should continue on with other Protection Sources
    EnableFasterIncrementalBackups *bool           `json:"enableFasterIncrementalBackups,omitempty" form:"enableFasterIncrementalBackups,omitempty"` //Specifies whether this job will enable faster incremental backups using
    FilePathFilters                FileOrDirectoryPathFilter `json:"filePathFilters,omitempty" form:"filePathFilters,omitempty"` //Specifies filters to match files and directories on a Server.
    NasProtocol                    NasProtocol1Enum `json:"nasProtocol,omitempty" form:"nasProtocol,omitempty"` //Specifies the preferred protocol to use for backup. This does not
}

/*
 * Structure for the custom type NASMountCredentials
 */
type NASMountCredentials struct {
    Domain          *string         `json:"domain,omitempty" form:"domain,omitempty"` //Specifies the domain in which this credential is valid.
    NasProtocol     NasProtocolEnum `json:"nasProtocol,omitempty" form:"nasProtocol,omitempty"` //Specifies the protocol used by the NAS server.
    NasType         NasTypeEnum     `json:"nasType,omitempty" form:"nasType,omitempty"` //Specifies the type of a NAS Object such as 'kGroup', or 'kHost'.
    Password        *string         `json:"password,omitempty" form:"password,omitempty"` //Specifies the password for the username to use for mounting the NAS.
    SkipValidation  *bool           `json:"skipValidation,omitempty" form:"skipValidation,omitempty"` //Specifies the flag to disable mount point validation during registration
    Username        *string         `json:"username,omitempty" form:"username,omitempty"` //Specifies a username to use for mounting the NAS.
}

/*
 * Structure for the custom type GenericNASProtectionSource
 */
type GenericNASProtectionSource struct {
    Description     *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description about the Protection Source.
    MountPath       *string         `json:"mountPath,omitempty" form:"mountPath,omitempty"` //Specifies the mount path of this NAS. For example, for a NFS mount point,
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the NetApp Object.
    Protocol        Protocol2Enum   `json:"protocol,omitempty" form:"protocol,omitempty"` //Specifies the protocol used by the NAS server.
    SkipValidation  *bool           `json:"skipValidation,omitempty" form:"skipValidation,omitempty"` //Specifies whether to skip validation of the given mount point.
    Type            Type10Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of a Protection Source Object in a generic NAS Source
}

/*
 * Structure for the custom type NetAppProtectionSource
 */
type NetAppProtectionSource struct {
    ClusterInfo     NetAppClusterInformation `json:"clusterInfo,omitempty" form:"clusterInfo,omitempty"` //Specifies information about a NetApp Cluster Protection Source.
    IsTopLevel      *bool           `json:"isTopLevel,omitempty" form:"isTopLevel,omitempty"` //Specifies if this Object is a top level Object.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the NetApp Object.
    Type            Type11Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of managed NetApp Object in a NetApp Protection Source
    Uuid            *string         `json:"uuid,omitempty" form:"uuid,omitempty"` //Specifies the globally unique ID of this Object assigned by the
    VolumeInfo      NetAppVolumeInformation `json:"volumeInfo,omitempty" form:"volumeInfo,omitempty"` //Specifies information about a volume in a NetApp cluster.
    VserverInfo     NetAppVserverInformation `json:"vserverInfo,omitempty" form:"vserverInfo,omitempty"` //Specifies information about a NetApp Vserver in a NetApp Protection Source.
}

/*
 * Structure for the custom type NetAppVolumeInformation
 */
type NetAppVolumeInformation struct {
    AggregateName     *string         `json:"aggregateName,omitempty" form:"aggregateName,omitempty"` //Specifies the containing aggregate name of this volume.
    CapacityBytes     *int64          `json:"capacityBytes,omitempty" form:"capacityBytes,omitempty"` //Specifies the total capacity in bytes of this volume.
    CifsShares        []*CIFSShareInformation `json:"cifsShares,omitempty" form:"cifsShares,omitempty"` //Array of CIFS Shares.
    CreationTimeUsecs *int64          `json:"creationTimeUsecs,omitempty" form:"creationTimeUsecs,omitempty"` //Specifies the creation time of the volume specified in Unix epoch time
    DataProtocols     *[]DataProtocolEnum `json:"dataProtocols,omitempty" form:"dataProtocols,omitempty"` //Array of Data Protocols.
    ExportPolicyName  *string         `json:"exportPolicyName,omitempty" form:"exportPolicyName,omitempty"` //Specifies the name of the export policy (which defines the access
    JunctionPath      *string         `json:"junctionPath,omitempty" form:"junctionPath,omitempty"` //Specifies the junction path of this volume.
    Name              *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the NetApp Vserver that this volume belongs to.
    SecurityInfo      VolumeSecurityInformation `json:"securityInfo,omitempty" form:"securityInfo,omitempty"` //Specifies information about NetApp volume security settings.
    State             StateEnum       `json:"state,omitempty" form:"state,omitempty"` //Specifies the state of this volume.
    Type              Type12Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the NetApp type of this volume.
    UsedBytes         *int64          `json:"usedBytes,omitempty" form:"usedBytes,omitempty"` //Specifies the total space (in bytes) used in this volume.
}

/*
 * Structure for the custom type NetAppVserverInformation
 */
type NetAppVserverInformation struct {
    DataProtocols   *[]DataProtocolEnum `json:"dataProtocols,omitempty" form:"dataProtocols,omitempty"` //Array of Data Protocols.
    Interfaces      []*VserverNetworkInterface `json:"interfaces,omitempty" form:"interfaces,omitempty"` //Array of Interfaces.
    RootCifsShare   CIFSShareInformation `json:"rootCifsShare,omitempty" form:"rootCifsShare,omitempty"` //Specifies information about a CIFS share of a NetApp volume.
    Type            Type13Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of this Vserver.
}

/*
 * Structure for the custom type Node
 */
type Node struct {
    CapacityByTier           []*CapacityByTier `json:"capacityByTier,omitempty" form:"capacityByTier,omitempty"` //CapacityByTier describes the capacity of each storage tier.
    ChassisInfo              ChassisInformation `json:"chassisInfo,omitempty" form:"chassisInfo,omitempty"` //ChassisInfo is the struct for the Chassis.
    ClusterPartitionId       *int64          `json:"clusterPartitionId,omitempty" form:"clusterPartitionId,omitempty"` //ClusterPartitionId is the Id of the cluster partition to which
    ClusterPartitionName     *string         `json:"clusterPartitionName,omitempty" form:"clusterPartitionName,omitempty"` //ClusterPartitionName is the name of the cluster to which the Node
    DiskCount                *int64          `json:"diskCount,omitempty" form:"diskCount,omitempty"` //DiskCount is the number of disks in a node.
    Id                       *int64          `json:"id,omitempty" form:"id,omitempty"` //Id is the Id of the Node.
    Ip                       *string         `json:"ip,omitempty" form:"ip,omitempty"` //Ip is the IP address of the Node.
    IsMarkedForRemoval       *bool           `json:"isMarkedForRemoval,omitempty" form:"isMarkedForRemoval,omitempty"` //IsMarkedForRemoval specifies whether the node has been marked for
    MaxPhysicalCapacityBytes *int64          `json:"maxPhysicalCapacityBytes,omitempty" form:"maxPhysicalCapacityBytes,omitempty"` //MaxPhysicalCapacityBytes specifies the maximum physical capacity of the
    NodeHardwareInfo         NodeHardwareInformation `json:"nodeHardwareInfo,omitempty" form:"nodeHardwareInfo,omitempty"` //NodeHardwareInfo provides the information regarding the hardware.
    NodeIncarnationId        *int64          `json:"nodeIncarnationId,omitempty" form:"nodeIncarnationId,omitempty"` //NodeIncarnationId is the incarnation id  of this node. The incarnation
    NodeSoftwareVersion      *string         `json:"nodeSoftwareVersion,omitempty" form:"nodeSoftwareVersion,omitempty"` //NodeSoftwareVersion is the current version of Cohesity software installed
    OfflineMountPathsOfDisks *[]string       `json:"offlineMountPathsOfDisks,omitempty" form:"offlineMountPathsOfDisks,omitempty"` //OfflineMountPathsOfDisks provides the corresponding mount paths for
    RemovalReason            *[]RemovalReasonEnum `json:"removalReason,omitempty" form:"removalReason,omitempty"` //RemovalReason specifies the removal reason of the node.
    RemovalState             RemovalStateEnum `json:"removalState,omitempty" form:"removalState,omitempty"` //RemovalState specifies the removal state of the node.
    Stats                    NodeStatistics  `json:"stats,omitempty" form:"stats,omitempty"` //NodeStats provides various statistics for the node.
    SystemDisks              []*NodeSystemDiskInfo `json:"systemDisks,omitempty" form:"systemDisks,omitempty"` //SystemDisk describes the node system disks.
}

/*
 * Structure for the custom type NotificationRule
 */
type NotificationRule struct {
    AlertTypeList          *[]int64        `json:"alertTypeList,omitempty" form:"alertTypeList,omitempty"` //Specifies alert types this rule is applicable to.
    Categories             *[]Category2Enum `json:"categories,omitempty" form:"categories,omitempty"` //Specifies alert categories this rule is applicable to.
    EmailDeliveryTargets   []*EmailDeliveryTarget `json:"emailDeliveryTargets,omitempty" form:"emailDeliveryTargets,omitempty"` //Specifies email addresses to be notified when an alert matching this
    RuleId                 *int64          `json:"ruleId,omitempty" form:"ruleId,omitempty"` //Specifies id of the alert delivery rule.
    RuleName               *string         `json:"ruleName,omitempty" form:"ruleName,omitempty"` //Specifies name of the alert delivery rule.
    Severities             *[]Severity1Enum `json:"severities,omitempty" form:"severities,omitempty"` //Specifies alert severity types this rule is applicable to.
    TenantId               *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies tenant id this rule is applicable to.
    WebHookDeliveryTargets []*WebHookDeliveryTarget `json:"webHookDeliveryTargets,omitempty" form:"webHookDeliveryTargets,omitempty"` //Specifies external api urls to be invoked when an alert matching this
}

/*
 * Structure for the custom type ObjectInformation
 */
type ObjectInformation struct {
    AccessibleUsers *[]string       `json:"accessibleUsers,omitempty" form:"accessibleUsers,omitempty"` //Species the list of user who have access to this object.
    AuditLogs       []*ClusterAuditLog `json:"auditLogs,omitempty" form:"auditLogs,omitempty"` //Specifies the audit log information.
    CopyTaskInfo    []*GDPRCopyTask `json:"copyTaskInfo,omitempty" form:"copyTaskInfo,omitempty"` //Specifies the copy task information.
    IsProtected     *bool           `json:"isProtected,omitempty" form:"isProtected,omitempty"` //Specifies the protection status of the object.
    Location        *string         `json:"location,omitempty" form:"location,omitempty"` //Specifies the location of the parent source.
    ProtectionInfo  []*ProtectionInformation `json:"protectionInfo,omitempty" form:"protectionInfo,omitempty"` //Specifies the data locations for the protected objects.
    RootNodeId      *int64          `json:"rootNodeId,omitempty" form:"rootNodeId,omitempty"` //Specifies the id of the root node.
    SourceId        *int64          `json:"sourceId,omitempty" form:"sourceId,omitempty"` //Specifies the id of the Protection Source.
    SourceName      *string         `json:"sourceName,omitempty" form:"sourceName,omitempty"` //Specifies the name of the object.
}

/*
 * Structure for the custom type ObjectSnapshot
 */
type ObjectSnapshot struct {
    ClusterPartitionId *int64          `json:"clusterPartitionId,omitempty" form:"clusterPartitionId,omitempty"` //Specifies the Cohesity Cluster partition id where this object is stored.
    JobId              *int64          `json:"jobId,omitempty" form:"jobId,omitempty"` //Specifies the id for the Protection Job that is currently
    JobName            *string         `json:"jobName,omitempty" form:"jobName,omitempty"` //Specifies the name of the Protection Job that captured the backup.
    JobUid             UniqueGlobalId  `json:"jobUid,omitempty" form:"jobUid,omitempty"` //Specifies the globally unique id of the Protection Job that backed up
    ObjectName         *string         `json:"objectName,omitempty" form:"objectName,omitempty"` //Specifies the primary name of the object.
    OsType             *string         `json:"osType,omitempty" form:"osType,omitempty"` //Specifies the inferred OS type.
    RegisteredSource   ProtectionSource `json:"registeredSource,omitempty" form:"registeredSource,omitempty"` //Specifies a generic structure that represents a node
    SnapshottedSource  ProtectionSource `json:"snapshottedSource,omitempty" form:"snapshottedSource,omitempty"` //Specifies a generic structure that represents a node
    Versions           []*SnapshotVersion `json:"versions,omitempty" form:"versions,omitempty"` //Array of Snapshots.
    ViewBoxId          *int64          `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"` //Specifies the id of the Domain (View Box) where this
    ViewName           *string         `json:"viewName,omitempty" form:"viewName,omitempty"` //Specifies the View name where this object is stored.
}

/*
 * Structure for the custom type Office365ProtectionSource
 */
type Office365ProtectionSource struct {
    Description        *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies the description of the Office 365 entity.
    Name               *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the office 365 entity.
    PrimarySMTPAddress *string         `json:"primarySMTPAddress,omitempty" form:"primarySMTPAddress,omitempty"` //Specifies the SMTP address for the Outlook source.
    Type               *int64          `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of the Office 365 entity.
    Uuid               *string         `json:"uuid,omitempty" form:"uuid,omitempty"` //Specifies the UUID of the Office 365 entity.
}

/*
 * Structure for the custom type OracleCloudCredentials
 */
type OracleCloudCredentials struct {
    AccessKeyId     *string         `json:"accessKeyId,omitempty" form:"accessKeyId,omitempty"` //Specifies access key to connect to Oracle S3 Compatible vault account.
    Region          *string         `json:"region,omitempty" form:"region,omitempty"` //Specifies the region for Oracle S3 Compatible vault account.
    SecretAccessKey *string         `json:"secretAccessKey,omitempty" form:"secretAccessKey,omitempty"` //Specifies the secret access key for Oracle S3 Compatible vault account.
    Tenant          *string         `json:"tenant,omitempty" form:"tenant,omitempty"` //Specifies the tenant which is part of the REST endpoints for Oracle S3
    TierType        TierType3Enum   `json:"tierType,omitempty" form:"tierType,omitempty"` //Specifies the storage class of Oracle vault.
}

/*
 * Structure for the custom type OracleProtectionSource
 */
type OracleProtectionSource struct {
    ArchiveLogEnabled *bool           `json:"archiveLogEnabled,omitempty" form:"archiveLogEnabled,omitempty"` //Specifies whether the database is running in ARCHIVELOG mode. It enables
    BctEnabled        *bool           `json:"bctEnabled,omitempty" form:"bctEnabled,omitempty"` //Specifies whether the Block Change Tracking is enabled. BCT improves the
    DbType            *int64          `json:"dbType,omitempty" form:"dbType,omitempty"` //Specifies the type of the database in Oracle Protection Source.
    Hosts             []*OracleHost   `json:"hosts,omitempty" form:"hosts,omitempty"` //Specifies the list of hosts for the current DB entity.
    Name              *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the instance name of the Oracle entity.
    OwnerId           *int64          `json:"ownerId,omitempty" form:"ownerId,omitempty"` //Specifies the entity id of the owner entity (such as a VM). This is only
    Size              *int64          `json:"size,omitempty" form:"size,omitempty"` //Specifies database size.
    Type              Type14Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of the managed Object in Oracle Protection Source.
    Uuid              *string         `json:"uuid,omitempty" form:"uuid,omitempty"` //Specifies the UUID for the Oracle entity.
    Version           *string         `json:"version,omitempty" form:"version,omitempty"` //Specifies the Oracle database instance version.
}

/*
 * Structure for the custom type OutlookFolder
 */
type OutlookFolder struct {
    FolderId            *string         `json:"folderId,omitempty" form:"folderId,omitempty"` //Specifies the unique ID of the folder.
    FolderKey           *int64          `json:"folderKey,omitempty" form:"folderKey,omitempty"` //Specifies the key unique within the mailbox of the folder.
    OutlookItemIdList   *[]string       `json:"outlookItemIdList,omitempty" form:"outlookItemIdList,omitempty"` //Specifies the outlook items within the folder to be restored incase the
    RestoreEntireFolder *bool           `json:"restoreEntireFolder,omitempty" form:"restoreEntireFolder,omitempty"` //Specifies whether the entire folder is to be restored.
}

/*
 * Structure for the custom type OutputSpecificationForTheMapreduce
 */
type OutputSpecificationForTheMapreduce struct {
    NumReduceShards    *int64          `json:"numReduceShards,omitempty" form:"numReduceShards,omitempty"` //Number of reduce shards.
    OutputDir          *string         `json:"outputDir,omitempty" form:"outputDir,omitempty"` //Name of output directory.
    PartitionId        *int64          `json:"partitionId,omitempty" form:"partitionId,omitempty"` //Partition id where output will go.
    ReduceOutputPrefix *string         `json:"reduceOutputPrefix,omitempty" form:"reduceOutputPrefix,omitempty"` //Prefix of the reduce output files. File names will be:
    ViewBoxId          *int64          `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"` //Viewbox id where the output will go.
    ViewName           *string         `json:"viewName,omitempty" form:"viewName,omitempty"` //Name of the view where output will go. This will be filled up by yoda.
}

/*
 * Structure for the custom type PhysicalBackupSourceParameters
 */
type PhysicalBackupSourceParameters struct {
    EnableSystemBackup *bool           `json:"enableSystemBackup,omitempty" form:"enableSystemBackup,omitempty"` //Allows Magneto to drive a "system" backup using a 3rd-party tool installed
    FileBackupParams   PhysicalFileBackupParameters `json:"fileBackupParams,omitempty" form:"fileBackupParams,omitempty"` //Message to capture params when backing up files on a Physical source.
    SnapshotParams     PhysicalSnapshotParams `json:"snapshotParams,omitempty" form:"snapshotParams,omitempty"` //This message contains params that controls the snapshot process for a
    SourceAppParams    SourceAppParams `json:"sourceAppParams,omitempty" form:"sourceAppParams,omitempty"` //This message contains params specific to application running on the source
    VolumeGuidVec      *[]string       `json:"volumeGuidVec,omitempty" form:"volumeGuidVec,omitempty"` //If this list is non-empty, then only volumes in this will be
}

/*
 * Structure for the custom type BackupPathInformation
 */
type BackupPathInformation struct {
    ExcludePaths      *[]string       `json:"excludePaths,omitempty" form:"excludePaths,omitempty"` //A list of absolute paths on the Physical source that should not be
    IncludePath       *string         `json:"includePath,omitempty" form:"includePath,omitempty"` //An absolute path on the Physical source that should be backed up. Any
    SkipNestedVolumes *bool           `json:"skipNestedVolumes,omitempty" form:"skipNestedVolumes,omitempty"` //Whether to skip any nested volumes (both local and network) that are
}

/*
 * Structure for the custom type PhysicalProtectionSource
 */
type PhysicalProtectionSource struct {
    Agents          []*AgentInformation `json:"agents,omitempty" form:"agents,omitempty"` //Array of Agents on the Physical Protection Source.
    HostType        HostType5Enum   `json:"hostType,omitempty" form:"hostType,omitempty"` //Specifies the environment type for the host.
    Id              UniqueGlobalId  `json:"id,omitempty" form:"id,omitempty"` //Specifies a unique id of a Physical Protection Source.
    MemorySizeBytes *int64          `json:"memorySizeBytes,omitempty" form:"memorySizeBytes,omitempty"` //Specifies the total memory ont the host in bytes.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies a human readable name of the Protection Source.
    NetworkingInfo  NetworkingInformation `json:"networkingInfo,omitempty" form:"networkingInfo,omitempty"` //Specifies the struct containing information about network addresses
    NumProcessors   *int64          `json:"numProcessors,omitempty" form:"numProcessors,omitempty"` //Specifies the number of processors on the host.
    OsName          *string         `json:"osName,omitempty" form:"osName,omitempty"` //Specifies a human readable name of the OS of the Protection Source.
    Type            Type15Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of managed Object in a Physical Protection Source.
    Volumes         []*PhysicalVolume `json:"volumes,omitempty" form:"volumes,omitempty"` //Array of Physical Volumes.
}

/*
 * Structure for the custom type PhysicalSourceSpecialJobParameters
 */
type PhysicalSourceSpecialJobParameters struct {
    ApplicationParameters ApplicationSpecificParameters `json:"applicationParameters,omitempty" form:"applicationParameters,omitempty"` //TODO: Write general description for this field
    EnableSystemBackup    *bool           `json:"enableSystemBackup,omitempty" form:"enableSystemBackup,omitempty"` //Specifies whether to allow system backup using 3rd party tools installed
    FilePaths             []*FileOrDirectoryToProtect `json:"filePaths,omitempty" form:"filePaths,omitempty"` //Array of File Paths to Back Up.
    VolumeGuid            *[]string       `json:"volumeGuid,omitempty" form:"volumeGuid,omitempty"` //Array of Mounted Volumes to Back Up.
    WindowsParameters     WindowsHostSnapshotParameters `json:"windowsParameters,omitempty" form:"windowsParameters,omitempty"` //Specifies settings that are meaningful only on Windows hosts.
}

/*
 * Structure for the custom type PhysicalVolume
 */
type PhysicalVolume struct {
    DevicePath                    *string         `json:"devicePath,omitempty" form:"devicePath,omitempty"` //Specifies the path to the device that hosts the volume locally.
    Guid                          *string         `json:"guid,omitempty" form:"guid,omitempty"` //Specifies an id for the Physical Volume.
    IsExtendedAttributesSupported *bool           `json:"isExtendedAttributesSupported,omitempty" form:"isExtendedAttributesSupported,omitempty"` //Specifies whether this volume supports extended attributes (like ACLs)
    IsProtected                   *bool           `json:"isProtected,omitempty" form:"isProtected,omitempty"` //Specifies if a volume is protected by a Job.
    Label                         *string         `json:"label,omitempty" form:"label,omitempty"` //Specifies a volume label that can be used for displaying additional
    LogicalSizeBytes              *int64          `json:"logicalSizeBytes,omitempty" form:"logicalSizeBytes,omitempty"` //Specifies the logical size of the volume in bytes that is
    MountPoints                   *[]string       `json:"mountPoints,omitempty" form:"mountPoints,omitempty"` //Array of Mount Points.
    NetworkPath                   *string         `json:"networkPath,omitempty" form:"networkPath,omitempty"` //Specifies the full path to connect to the network attached volume.
    UsedSizeBytes                 *int64          `json:"usedSizeBytes,omitempty" form:"usedSizeBytes,omitempty"` //Specifies the size used by the volume in bytes.
}

/*
 * Structure for the custom type Principal
 */
type Principal struct {
    Domain          *string         `json:"domain,omitempty" form:"domain,omitempty"` //Specifies the domain name of the where the principal' account is
    FullName        *string         `json:"fullName,omitempty" form:"fullName,omitempty"` //Specifies the full name (first and last names) of the principal.
    ObjectClass     ObjectClassEnum `json:"objectClass,omitempty" form:"objectClass,omitempty"` //Specifies the object class of the principal (either 'kGroup' or 'kUser').
    PrincipalName   *string         `json:"principalName,omitempty" form:"principalName,omitempty"` //Specifies the name of the principal.
    Sid             *string         `json:"sid,omitempty" form:"sid,omitempty"` //Specifies the unique Security id (SID) of the principal.
}

/*
 * Structure for the custom type PrivilegeInformation
 */
type PrivilegeInformation struct {
    Category            *string         `json:"category,omitempty" form:"category,omitempty"` //Specifies a category for the privilege such as 'Access Management'.
    Description         *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description defining what the privilege provides.
    IsCustomRoleDefault *bool           `json:"isCustomRoleDefault,omitempty" form:"isCustomRoleDefault,omitempty"` //Specifies if this privilege is automatically assigned to custom roles.
    IsSpecial           *bool           `json:"isSpecial,omitempty" form:"isSpecial,omitempty"` //Specifies if this privilege is automatically assigned to the default
    IsViewOnly          *bool           `json:"isViewOnly,omitempty" form:"isViewOnly,omitempty"` //Specifies if privilege is view-only privilege that cannot make changes.
    Label               *string         `json:"label,omitempty" form:"label,omitempty"` //Specifies the label for the privilege as displayed on the Cohesity
    Name                *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the Cluster name for the privilege such as PRINCIPAL_VIEW.
}

/*
 * Structure for the custom type ProtectedObject
 */
type ProtectedObject struct {
    Cancelled        *int64          `json:"cancelled,omitempty" form:"cancelled,omitempty"` //Specifies number of cancelled runs across trends.
    Environment      Environment9Enum `json:"environment,omitempty" form:"environment,omitempty"` //Specifies environment.
    Failed           *int64          `json:"failed,omitempty" form:"failed,omitempty"` //Specifies number of failed runs across trends.
    Id               *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies protected object's Id.
    Name             *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies protected object's name.
    ParentSourceId   *int64          `json:"parentSourceId,omitempty" form:"parentSourceId,omitempty"` //Specifies protected object's parent id.
    ParentSourceName *string         `json:"parentSourceName,omitempty" form:"parentSourceName,omitempty"` //Specifies protected object's parent name.
    Running          *int64          `json:"running,omitempty" form:"running,omitempty"` //Specifies number of in-progress runs across trends.
    Successful       *int64          `json:"successful,omitempty" form:"successful,omitempty"` //Specifies number of successful runs across trends.
    Total            *int64          `json:"total,omitempty" form:"total,omitempty"` //Specifies total number of runs across trends.
    Trends           []*TrendingData `json:"trends,omitempty" form:"trends,omitempty"` //Aggregated protection runs information by days/weeks.
}

/*
 * Structure for the custom type ProtectedObjectsByEnviournment
 */
type ProtectedObjectsByEnviournment struct {
    EnvType              *string         `json:"envType,omitempty" form:"envType,omitempty"` //Environment Type.
    ProtectedCount       *int64          `json:"protectedCount,omitempty" form:"protectedCount,omitempty"` //Number of Protected Objects.
    ProtectedSizeBytes   *int64          `json:"protectedSizeBytes,omitempty" form:"protectedSizeBytes,omitempty"` //Size of Protected Objects.
    UnprotectedCount     *int64          `json:"unprotectedCount,omitempty" form:"unprotectedCount,omitempty"` //Number of Unprotected Objects.
    UnprotectedSizeBytes *int64          `json:"unprotectedSizeBytes,omitempty" form:"unprotectedSizeBytes,omitempty"` //Size of Unprotected Objects.
}

/*
 * Structure for the custom type ProtectedObjectsTile
 */
type ProtectedObjectsTile struct {
    ObjectsProtected     []*ProtectedObjectsByEnviournment `json:"objectsProtected,omitempty" form:"objectsProtected,omitempty"` //Protected Objects breakdown by object type.
    ProtectedCount       *int64          `json:"protectedCount,omitempty" form:"protectedCount,omitempty"` //Number of Protected Objects.
    ProtectedSizeBytes   *int64          `json:"protectedSizeBytes,omitempty" form:"protectedSizeBytes,omitempty"` //Size of Protected Objects.
    UnprotectedCount     *int64          `json:"unprotectedCount,omitempty" form:"unprotectedCount,omitempty"` //Number of Unprotected Objects.
    UnprotectedSizeBytes *int64          `json:"unprotectedSizeBytes,omitempty" form:"unprotectedSizeBytes,omitempty"` //Size of Unprotected Objects.
}

/*
 * Structure for the custom type ProtectedSourceSummary
 */
type ProtectedSourceSummary struct {
    BackupRun                  BackupRunTask   `json:"backupRun,omitempty" form:"backupRun,omitempty"` //Specifies details about the Backup task for a Job Run.
    CopyRuns                   []*CopyRunTask  `json:"copyRuns,omitempty" form:"copyRuns,omitempty"` //Specifies details about the Copy tasks of the Job Run.
    JobUid                     UniqueGlobalId  `json:"jobUid,omitempty" form:"jobUid,omitempty"` //Specifies an id for an object that is unique across Cohesity Clusters.
    NextProtectionRunTimeUsecs *int64          `json:"nextProtectionRunTimeUsecs,omitempty" form:"nextProtectionRunTimeUsecs,omitempty"` //Specifies the time at which the next Protection Run is scheduled for the
    ProtectionSource           ProtectionSource `json:"protectionSource,omitempty" form:"protectionSource,omitempty"` //Specifies a generic structure that represents a node
}

/*
 * Structure for the custom type ProtectionInformation
 */
type ProtectionInformation struct {
    EndTimeUsecs      *int64          `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"` //Specifies the end time for object retention.
    Location          *string         `json:"location,omitempty" form:"location,omitempty"` //Specifies the location of the object.
    PolicyId          *string         `json:"policyId,omitempty" form:"policyId,omitempty"` //Specifies the id of the policy.
    ProtectionJobId   *int64          `json:"protectionJobId,omitempty" form:"protectionJobId,omitempty"` //Specifies the id of the protection job.
    ProtectionJobName *string         `json:"protectionJobName,omitempty" form:"protectionJobName,omitempty"` //Specifies the protection job name which protects this object.
    RetentionPeriod   *int64          `json:"retentionPeriod,omitempty" form:"retentionPeriod,omitempty"` //Specifies the retention period.
    StartTimeUsecs    *int64          `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"` //Specifies the start time for object retention.
    StorageDomain     *string         `json:"storageDomain,omitempty" form:"storageDomain,omitempty"` //Specifies the storage domain name.
    TotalSnapshots    *int64          `json:"totalSnapshots,omitempty" form:"totalSnapshots,omitempty"` //Specifies the total number of snapshots.
}

/*
 * Structure for the custom type ProtectionJob
 */
type ProtectionJob struct {
    AbortInBlackoutPeriod                *bool           `json:"abortInBlackoutPeriod,omitempty" form:"abortInBlackoutPeriod,omitempty"` //If true, the Cohesity Cluster aborts any currently executing Job Runs
    AlertingConfig                       AlertingConfig  `json:"alertingConfig,omitempty" form:"alertingConfig,omitempty"` //Specifies optional settings for alerting.
    AlertingPolicy                       *[]AlertingPolicyEnum `json:"alertingPolicy,omitempty" form:"alertingPolicy,omitempty"` //Array of Job Events.
    CloudParameters                      CloudParameters `json:"cloudParameters,omitempty" form:"cloudParameters,omitempty"` //Specifies Cloud parameters that are applicable to all Protection
    ContinueOnQuiesceFailure             *bool           `json:"continueOnQuiesceFailure,omitempty" form:"continueOnQuiesceFailure,omitempty"` //Whether to continue backing up on quiesce failure.
    CreationTimeUsecs                    *int64          `json:"creationTimeUsecs,omitempty" form:"creationTimeUsecs,omitempty"` //Specifies the time when the Protection Job was created.
    DedupDisabledSourceIds               *[]int64        `json:"dedupDisabledSourceIds,omitempty" form:"dedupDisabledSourceIds,omitempty"` //List of source ids for which source side dedup is disabled from the backup
    Description                          *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a text description about the Protection Job.
    EndTimeUsecs                         *int64          `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"` //Specifies the epoch time (in microseconds) after which the Protection Job
    Environment                          Environment10Enum `json:"environment,omitempty" form:"environment,omitempty"` //Specifies the environment type (such as kVMware or kSQL)
    EnvironmentParameters                EnvironmentSpecificCommonJobParameters `json:"environmentParameters,omitempty" form:"environmentParameters,omitempty"` //Specifies additional parameters that are common to all Protection
    ExcludeSourceIds                     *[]int64        `json:"excludeSourceIds,omitempty" form:"excludeSourceIds,omitempty"` //Array of Excluded Source Objects.
    ExcludeVmTagIds                      *[]int64        `json:"excludeVmTagIds,omitempty" form:"excludeVmTagIds,omitempty"` //Array of Arrays of VM Tag Ids that Specify VMs to Exclude.
    FullProtectionSlaTimeMins            *int64          `json:"fullProtectionSlaTimeMins,omitempty" form:"fullProtectionSlaTimeMins,omitempty"` //If specified, this setting is number of minutes that a Job Run
    FullProtectionStartTime              TimeOfDay       `json:"fullProtectionStartTime,omitempty" form:"fullProtectionStartTime,omitempty"` //Specifies the time of day to start the Full Protection Schedule.
    Id                                   *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies an id for the Protection Job.
    IncrementalProtectionSlaTimeMins     *int64          `json:"incrementalProtectionSlaTimeMins,omitempty" form:"incrementalProtectionSlaTimeMins,omitempty"` //If specified, this setting is number of minutes that a Job Run
    IncrementalProtectionStartTime       TimeOfDay       `json:"incrementalProtectionStartTime,omitempty" form:"incrementalProtectionStartTime,omitempty"` //Specifies the time of day to start the CBT-based Protection Schedule.
    IndexingPolicy                       IndexingPolicy  `json:"indexingPolicy,omitempty" form:"indexingPolicy,omitempty"` //Specifies settings for indexing files found in an Object
    IsActive                             *bool           `json:"isActive,omitempty" form:"isActive,omitempty"` //Indicates if the current state of the Protection Job is Active
    IsDeleted                            *bool           `json:"isDeleted,omitempty" form:"isDeleted,omitempty"` //Equals 'true' if the Protection Job was deleted but some Snapshots
    IsPaused                             *bool           `json:"isPaused,omitempty" form:"isPaused,omitempty"` //Indicates if the Protection Job is paused, which means that no new
    LastRun                              ProtectionJobRunInstance `json:"lastRun,omitempty" form:"lastRun,omitempty"` //Specifies the status of one Job Run.
    LeverageStorageSnapshots             *bool           `json:"leverageStorageSnapshots,omitempty" form:"leverageStorageSnapshots,omitempty"` //Specifies whether to leverage the storage array based snapshots for this
    LeverageStorageSnapshotsForHyperflex *bool           `json:"leverageStorageSnapshotsForHyperflex,omitempty" form:"leverageStorageSnapshotsForHyperflex,omitempty"` //Specifies whether to leverage Hyperflex as the storage snapshot array
    ModificationTimeUsecs                *int64          `json:"modificationTimeUsecs,omitempty" form:"modificationTimeUsecs,omitempty"` //Specifies the last time this Job was updated.
    ModifiedByUser                       *string         `json:"modifiedByUser,omitempty" form:"modifiedByUser,omitempty"` //Specifies the last Cohesity user who updated this Job.
    Name                                 string          `json:"name" form:"name"` //Specifies the name of the Protection Job.
    ParentSourceId                       *int64          `json:"parentSourceId,omitempty" form:"parentSourceId,omitempty"` //Specifies the id of the registered Protection Source that is the
    PerformSourceSideDedup               *bool           `json:"performSourceSideDedup,omitempty" form:"performSourceSideDedup,omitempty"` //Specifies whether source side dedupe should be performed or not.
    PolicyAppliedTimeMsecs               *int64          `json:"policyAppliedTimeMsecs,omitempty" form:"policyAppliedTimeMsecs,omitempty"` //Specifies the epoch time (in milliseconds) when the
    PolicyId                             string          `json:"policyId" form:"policyId"` //Specifies the unique id of the Protection Policy associated with
    PostBackupScript                     BackupScript    `json:"postBackupScript,omitempty" form:"postBackupScript,omitempty"` //Specifies the script associated with the backup job. This field must be
    PreBackupScript                      BackupScript    `json:"preBackupScript,omitempty" form:"preBackupScript,omitempty"` //Specifies the script associated with the backup job. This field must be
    Priority                             PriorityEnum    `json:"priority,omitempty" form:"priority,omitempty"` //Specifies the priority of execution for a Protection Job.
    QosType                              QosTypeEnum     `json:"qosType,omitempty" form:"qosType,omitempty"` //Specifies the QoS policy type to use for this Protection Job.
    Quiesce                              *bool           `json:"quiesce,omitempty" form:"quiesce,omitempty"` //Indicates if the App-Consistent option is enabled for this Job.
    RemoteScript                         RemoteAdapter   `json:"remoteScript,omitempty" form:"remoteScript,omitempty"` //For a Remote Adapter 'kPuppeteer' Job, this field specifies the
    SourceIds                            *[]int64        `json:"sourceIds,omitempty" form:"sourceIds,omitempty"` //Array of Protected Source Objects.
    SourceSpecialParameters              []*SourceSpecialParameters `json:"sourceSpecialParameters,omitempty" form:"sourceSpecialParameters,omitempty"` //Array of Special Source Parameters.
    StartTime                            TimeOfDay       `json:"startTime,omitempty" form:"startTime,omitempty"` //Specifies the time of day to start the Protection Schedule.
    SummaryStats                         ProtectionJobSummaryStatistics `json:"summaryStats,omitempty" form:"summaryStats,omitempty"` //Specifies statistics about a Protection Job.
    TenantId                             *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
    Timezone                             *string         `json:"timezone,omitempty" form:"timezone,omitempty"` //Specifies the timezone to use when calculating time for this
    Uid                                  UniqueGlobalId  `json:"uid,omitempty" form:"uid,omitempty"` //Specifies a global Protection Job id that is unique across Cohesity
    ViewBoxId                            int64           `json:"viewBoxId" form:"viewBoxId"` //Specifies the Storage Domain (View Box) id where this Job writes data.
    ViewName                             *string         `json:"viewName,omitempty" form:"viewName,omitempty"` //For a Remote Adapter 'kPuppeteer' Job or a 'kView' Job, this field
    VmTagIds                             *[]int64        `json:"vmTagIds,omitempty" form:"vmTagIds,omitempty"` //Array of Arrays of VMs Tags Ids that Specify VMs to Protect.
}

/*
 * Structure for the custom type ProtectionJobInfo
 */
type ProtectionJobInfo struct {
    JobId           *int64          `json:"jobId,omitempty" form:"jobId,omitempty"` //Specifies the id of the Protection Job.
    JobName         *string         `json:"jobName,omitempty" form:"jobName,omitempty"` //Specifies the name of the Protection Job.
    Type            Type23Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of the Protection Job such as kView or kPuppeteer.
}

/*
 * Structure for the custom type ProtectionJobRequest
 */
type ProtectionJobRequest struct {
    AbortInBlackoutPeriod                *bool           `json:"abortInBlackoutPeriod,omitempty" form:"abortInBlackoutPeriod,omitempty"` //If true, the Cohesity Cluster aborts any currently executing Job Runs
    AlertingConfig                       AlertingConfig  `json:"alertingConfig,omitempty" form:"alertingConfig,omitempty"` //Specifies optional settings for alerting.
    AlertingPolicy                       *[]AlertingPolicyEnum `json:"alertingPolicy,omitempty" form:"alertingPolicy,omitempty"` //Array of Job Events.
    CloudParameters                      CloudParameters `json:"cloudParameters,omitempty" form:"cloudParameters,omitempty"` //Specifies Cloud parameters that are applicable to all Protection
    ContinueOnQuiesceFailure             *bool           `json:"continueOnQuiesceFailure,omitempty" form:"continueOnQuiesceFailure,omitempty"` //Whether to continue backing up on quiesce failure.
    DedupDisabledSourceIds               *[]int64        `json:"dedupDisabledSourceIds,omitempty" form:"dedupDisabledSourceIds,omitempty"` //List of source ids for which source side dedup is disabled from the backup
    Description                          *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a text description about the Protection Job.
    EndTimeUsecs                         *int64          `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"` //Specifies the epoch time (in microseconds) after which the Protection Job
    Environment                          Environment10Enum `json:"environment,omitempty" form:"environment,omitempty"` //Specifies the environment type (such as kVMware or kSQL)
    EnvironmentParameters                EnvironmentSpecificCommonJobParameters `json:"environmentParameters,omitempty" form:"environmentParameters,omitempty"` //Specifies additional parameters that are common to all Protection
    ExcludeSourceIds                     *[]int64        `json:"excludeSourceIds,omitempty" form:"excludeSourceIds,omitempty"` //Array of Excluded Source Objects.
    ExcludeVmTagIds                      *[]int64        `json:"excludeVmTagIds,omitempty" form:"excludeVmTagIds,omitempty"` //Array of Arrays of VM Tag Ids that Specify VMs to Exclude.
    FullProtectionSlaTimeMins            *int64          `json:"fullProtectionSlaTimeMins,omitempty" form:"fullProtectionSlaTimeMins,omitempty"` //If specified, this setting is number of minutes that a Job Run
    FullProtectionStartTime              TimeOfDay       `json:"fullProtectionStartTime,omitempty" form:"fullProtectionStartTime,omitempty"` //Specifies the time of day to start the Full Protection Schedule.
    IncrementalProtectionSlaTimeMins     *int64          `json:"incrementalProtectionSlaTimeMins,omitempty" form:"incrementalProtectionSlaTimeMins,omitempty"` //If specified, this setting is number of minutes that a Job Run
    IncrementalProtectionStartTime       TimeOfDay       `json:"incrementalProtectionStartTime,omitempty" form:"incrementalProtectionStartTime,omitempty"` //Specifies the time of day to start the CBT-based Protection Schedule.
    IndexingPolicy                       IndexingPolicy  `json:"indexingPolicy,omitempty" form:"indexingPolicy,omitempty"` //Specifies settings for indexing files found in an Object
    LeverageStorageSnapshots             *bool           `json:"leverageStorageSnapshots,omitempty" form:"leverageStorageSnapshots,omitempty"` //Specifies whether to leverage the storage array based snapshots for this
    LeverageStorageSnapshotsForHyperflex *bool           `json:"leverageStorageSnapshotsForHyperflex,omitempty" form:"leverageStorageSnapshotsForHyperflex,omitempty"` //Specifies whether to leverage Hyperflex as the storage snapshot array
    Name                                 string          `json:"name" form:"name"` //Specifies the name of the Protection Job.
    ParentSourceId                       *int64          `json:"parentSourceId,omitempty" form:"parentSourceId,omitempty"` //Specifies the id of the registered Protection Source that is the
    PerformSourceSideDedup               *bool           `json:"performSourceSideDedup,omitempty" form:"performSourceSideDedup,omitempty"` //Specifies whether source side dedupe should be performed or not.
    PolicyId                             string          `json:"policyId" form:"policyId"` //Specifies the unique id of the Protection Policy associated with
    PostBackupScript                     BackupScript    `json:"postBackupScript,omitempty" form:"postBackupScript,omitempty"` //Specifies the script associated with the backup job. This field must be
    PreBackupScript                      BackupScript    `json:"preBackupScript,omitempty" form:"preBackupScript,omitempty"` //Specifies the script associated with the backup job. This field must be
    Priority                             PriorityEnum    `json:"priority,omitempty" form:"priority,omitempty"` //Specifies the priority of execution for a Protection Job.
    QosType                              QosTypeEnum     `json:"qosType,omitempty" form:"qosType,omitempty"` //Specifies the QoS policy type to use for this Protection Job.
    Quiesce                              *bool           `json:"quiesce,omitempty" form:"quiesce,omitempty"` //Indicates if the App-Consistent option is enabled for this Job.
    RemoteScript                         RemoteAdapter   `json:"remoteScript,omitempty" form:"remoteScript,omitempty"` //For a Remote Adapter 'kPuppeteer' Job, this field specifies the
    SourceIds                            *[]int64        `json:"sourceIds,omitempty" form:"sourceIds,omitempty"` //Array of Protected Source Objects.
    SourceSpecialParameters              []*SourceSpecialParameters `json:"sourceSpecialParameters,omitempty" form:"sourceSpecialParameters,omitempty"` //Array of Special Source Parameters.
    StartTime                            TimeOfDay       `json:"startTime,omitempty" form:"startTime,omitempty"` //Specifies the time of day to start the Protection Schedule.
    Timezone                             *string         `json:"timezone,omitempty" form:"timezone,omitempty"` //Specifies the timezone to use when calculating time for this
    ViewBoxId                            int64           `json:"viewBoxId" form:"viewBoxId"` //Specifies the Storage Domain (View Box) id where this Job writes data.
    ViewName                             *string         `json:"viewName,omitempty" form:"viewName,omitempty"` //For a Remote Adapter 'kPuppeteer' Job or a 'kView' Job, this field
    VmTagIds                             *[]int64        `json:"vmTagIds,omitempty" form:"vmTagIds,omitempty"` //Array of Arrays of VMs Tags Ids that Specify VMs to Protect.
}

/*
 * Structure for the custom type ProtectionJobRunStatistics
 */
type ProtectionJobRunStatistics struct {
    AdmittedTimeUsecs            *int64          `json:"admittedTimeUsecs,omitempty" form:"admittedTimeUsecs,omitempty"` //Specifies the time the task was unqueued from the queue to start running.
    EndTimeUsecs                 *int64          `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"` //Specifies the end time of the Protection Run. The end time
    NumAppInstances              *int64          `json:"numAppInstances,omitempty" form:"numAppInstances,omitempty"` //Specifies the number of application instances backed up by this Run.
    NumCanceledTasks             *int64          `json:"numCanceledTasks,omitempty" form:"numCanceledTasks,omitempty"` //Specifies the number of backup tasks that were canceled.
    NumFailedAppObjects          *int64          `json:"numFailedAppObjects,omitempty" form:"numFailedAppObjects,omitempty"` //Specifies the number of application objects that were cancelled in this
    NumFailedTasks               *int64          `json:"numFailedTasks,omitempty" form:"numFailedTasks,omitempty"` //Specifies the number of backup tasks that failed.
    NumSuccessfulAppObjects      *int64          `json:"numSuccessfulAppObjects,omitempty" form:"numSuccessfulAppObjects,omitempty"` //Specifies the number of application objects successfully backed up by this
    NumSuccessfulTasks           *int64          `json:"numSuccessfulTasks,omitempty" form:"numSuccessfulTasks,omitempty"` //Specifies the number of backup tasks that completed successfully.
    StartTimeUsecs               *int64          `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"` //Specifies the start time of the Protection Run. The start time
    TimeTakenUsecs               *int64          `json:"timeTakenUsecs,omitempty" form:"timeTakenUsecs,omitempty"` //Specifies the actual execution time for the protection run to complete
    TotalBytesReadFromSource     *int64          `json:"totalBytesReadFromSource,omitempty" form:"totalBytesReadFromSource,omitempty"` //Specifies the total amount of data read from the source (so far).
    TotalBytesToReadFromSource   *int64          `json:"totalBytesToReadFromSource,omitempty" form:"totalBytesToReadFromSource,omitempty"` //Specifies the total amount of data expected to be read from the
    TotalLogicalBackupSizeBytes  *int64          `json:"totalLogicalBackupSizeBytes,omitempty" form:"totalLogicalBackupSizeBytes,omitempty"` //Specifies the size of the source object (such as a VM) protected by
    TotalPhysicalBackupSizeBytes *int64          `json:"totalPhysicalBackupSizeBytes,omitempty" form:"totalPhysicalBackupSizeBytes,omitempty"` //Specifies the total amount of physical space used on the Cohesity
    TotalSourceSizeBytes         *int64          `json:"totalSourceSizeBytes,omitempty" form:"totalSourceSizeBytes,omitempty"` //Specifies the size of the source object (such as a VM) protected by
}

/*
 * Structure for the custom type SpecifiesProtectionJobSummaryOfAnObject
 */
type SpecifiesProtectionJobSummaryOfAnObject struct {
    ClusterId                  *int64          `json:"clusterId,omitempty" form:"clusterId,omitempty"` //Specifies the id of the cluster on which object is protected.
    ClusterIncarnationId       *int64          `json:"clusterIncarnationId,omitempty" form:"clusterIncarnationId,omitempty"` //Specifies the incarnation id of the cluster on which object is protected.
    JobId                      *int64          `json:"jobId,omitempty" form:"jobId,omitempty"` //Specifies the id of the Protection Job.
    JobName                    *string         `json:"jobName,omitempty" form:"jobName,omitempty"` //Specifies the name of the Protection Job.
    LastProtectionJobRunStatus *int64          `json:"lastProtectionJobRunStatus,omitempty" form:"lastProtectionJobRunStatus,omitempty"` //Specifies the last job run status.
    PolicyId                   *string         `json:"policyId,omitempty" form:"policyId,omitempty"` //Specifies the id of the policy that is used by a Protection Job.
    PolicyName                 *string         `json:"policyName,omitempty" form:"policyName,omitempty"` //Specifies the name of the policy that is used by a Protection Job.
}

/*
 * Structure for the custom type ProtectionJobSummaryStatistics
 */
type ProtectionJobSummaryStatistics struct {
    AverageRunTimeUsecs          *int64          `json:"averageRunTimeUsecs,omitempty" form:"averageRunTimeUsecs,omitempty"` //Specifies the average run time of all successful Protection Runs.
    FastestRunTimeUsecs          *int64          `json:"fastestRunTimeUsecs,omitempty" form:"fastestRunTimeUsecs,omitempty"` //Specifies the time taken for a fastest successful Protection Run so far.
    NumCanceledRuns              *int64          `json:"numCanceledRuns,omitempty" form:"numCanceledRuns,omitempty"` //Specifies the number of runs that were canceled.
    NumFailedRuns                *int64          `json:"numFailedRuns,omitempty" form:"numFailedRuns,omitempty"` //Specifies the number of runs that failed to finish.
    NumSlaViolations             *int64          `json:"numSlaViolations,omitempty" form:"numSlaViolations,omitempty"` //Specifies the number of runs having SLA violations.
    NumSuccessfulRuns            *int64          `json:"numSuccessfulRuns,omitempty" form:"numSuccessfulRuns,omitempty"` //Specifies the number of runs that finished successfully.
    SlowestRunTimeUsecs          *int64          `json:"slowestRunTimeUsecs,omitempty" form:"slowestRunTimeUsecs,omitempty"` //Specifies the time taken for a slowest successful Protection Run so far.
    TotalBytesReadFromSource     *int64          `json:"totalBytesReadFromSource,omitempty" form:"totalBytesReadFromSource,omitempty"` //Specifies the total amount of data read from the source (so far).
    TotalLogicalBackupSizeBytes  *int64          `json:"totalLogicalBackupSizeBytes,omitempty" form:"totalLogicalBackupSizeBytes,omitempty"` //Specifies the size of the source object (such as a VM) protected by
    TotalPhysicalBackupSizeBytes *int64          `json:"totalPhysicalBackupSizeBytes,omitempty" form:"totalPhysicalBackupSizeBytes,omitempty"` //Specifies the total amount of physical space used on the Cohesity
}

/*
 * Structure for the custom type ProtectionPolicy
 */
type ProtectionPolicy struct {
    BlackoutPeriods                 []*BlackoutPeriod `json:"blackoutPeriods,omitempty" form:"blackoutPeriods,omitempty"` //Array of Blackout Periods.
    CloudDeployPolicies             []*CloudDeployPolicy `json:"cloudDeployPolicies,omitempty" form:"cloudDeployPolicies,omitempty"` //Array of Cloud Deploy Policies.
    DaysToKeep                      *int64          `json:"daysToKeep,omitempty" form:"daysToKeep,omitempty"` //Specifies how many days to retain Snapshots on the Cohesity Cluster.
    DaysToKeepLog                   *int64          `json:"daysToKeepLog,omitempty" form:"daysToKeepLog,omitempty"` //Specifies the number of days to retain log run if Log Schedule exists.
    DaysToKeepSystem                *int64          `json:"daysToKeepSystem,omitempty" form:"daysToKeepSystem,omitempty"` //Specifies the number of days to retain system backups made for bare metal
    Description                     *string         `json:"description,omitempty" form:"description,omitempty"` //Description of the Protection Policy.
    ExtendedRetentionPolicies       []*ExtendedRetentionPolicy `json:"extendedRetentionPolicies,omitempty" form:"extendedRetentionPolicies,omitempty"` //Specifies additional retention policies that should be applied to the
    FullSchedulingPolicy            SchedulingPolicy `json:"fullSchedulingPolicy,omitempty" form:"fullSchedulingPolicy,omitempty"` //Specifies the Full (no CBT) backup schedule of a Protection Job and
    Id                              *string         `json:"id,omitempty" form:"id,omitempty"` //Specifies a unique Policy id assigned by the Cohesity Cluster.
    IncrementalSchedulingPolicy     SchedulingPolicy `json:"incrementalSchedulingPolicy,omitempty" form:"incrementalSchedulingPolicy,omitempty"` //Specifies the CBT-based backup schedule of a Protection Job and
    LastModifiedTimeMsecs           *int64          `json:"lastModifiedTimeMsecs,omitempty" form:"lastModifiedTimeMsecs,omitempty"` //Specifies the epoch time (in milliseconds) when the Protection Policy
    LogSchedulingPolicy             SchedulingPolicy `json:"logSchedulingPolicy,omitempty" form:"logSchedulingPolicy,omitempty"` //Specifies settings that define a backup schedule for a Protection Job.
    Name                            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the Protection Policy.
    Retries                         *int64          `json:"retries,omitempty" form:"retries,omitempty"` //Specifies the number of times to retry capturing Snapshots before
    RetryIntervalMins               *int64          `json:"retryIntervalMins,omitempty" form:"retryIntervalMins,omitempty"` //Specifies the number of minutes before retrying a failed Protection Job.
    RpoPolicySettings               RPOPolicy       `json:"rpoPolicySettings,omitempty" form:"rpoPolicySettings,omitempty"` //Specifies all the additional settings that are applicable only
    SkipIntervalMins                *int64          `json:"skipIntervalMins,omitempty" form:"skipIntervalMins,omitempty"` //Specifies the period of time before skipping the execution of new
    SnapshotArchivalCopyPolicies    []*SnapshotCopyArchivalPolicy `json:"snapshotArchivalCopyPolicies,omitempty" form:"snapshotArchivalCopyPolicies,omitempty"` //Array of External Targets.
    SnapshotReplicationCopyPolicies []*SnapshotCopyReplicationPolicy `json:"snapshotReplicationCopyPolicies,omitempty" form:"snapshotReplicationCopyPolicies,omitempty"` //Array of Remote Clusters.
    SystemSchedulingPolicy          SchedulingPolicy `json:"systemSchedulingPolicy,omitempty" form:"systemSchedulingPolicy,omitempty"` //Specifies settings that define a backup schedule for a Protection Job.
    TenantIds                       *[]string       `json:"tenantIds,omitempty" form:"tenantIds,omitempty"` //Specifies which organizations have been assigned this policy.
    Type                            Type29Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of the protection policy.
    WormRetentionType               WormRetentionType1Enum `json:"wormRetentionType,omitempty" form:"wormRetentionType,omitempty"` //Specifies WORM retention type for the snapshots. When a WORM retention
}

/*
 * Structure for the custom type ProtectionPolicyRequest
 */
type ProtectionPolicyRequest struct {
    BlackoutPeriods                 []*BlackoutPeriod `json:"blackoutPeriods,omitempty" form:"blackoutPeriods,omitempty"` //Array of Blackout Periods.
    CloudDeployPolicies             []*CloudDeployPolicy `json:"cloudDeployPolicies,omitempty" form:"cloudDeployPolicies,omitempty"` //Array of Cloud Deploy Policies.
    DaysToKeep                      *int64          `json:"daysToKeep,omitempty" form:"daysToKeep,omitempty"` //Specifies how many days to retain Snapshots on the Cohesity Cluster.
    DaysToKeepLog                   *int64          `json:"daysToKeepLog,omitempty" form:"daysToKeepLog,omitempty"` //Specifies the number of days to retain log run if Log Schedule exists.
    DaysToKeepSystem                *int64          `json:"daysToKeepSystem,omitempty" form:"daysToKeepSystem,omitempty"` //Specifies the number of days to retain system backups made for bare metal
    Description                     *string         `json:"description,omitempty" form:"description,omitempty"` //Description of the Protection Policy.
    ExtendedRetentionPolicies       []*ExtendedRetentionPolicy `json:"extendedRetentionPolicies,omitempty" form:"extendedRetentionPolicies,omitempty"` //Specifies additional retention policies that should be applied to the
    FullSchedulingPolicy            SchedulingPolicy `json:"fullSchedulingPolicy,omitempty" form:"fullSchedulingPolicy,omitempty"` //Specifies the Full (no CBT) backup schedule of a Protection Job and
    IncrementalSchedulingPolicy     SchedulingPolicy `json:"incrementalSchedulingPolicy,omitempty" form:"incrementalSchedulingPolicy,omitempty"` //Specifies the CBT-based backup schedule of a Protection Job and
    LogSchedulingPolicy             SchedulingPolicy `json:"logSchedulingPolicy,omitempty" form:"logSchedulingPolicy,omitempty"` //Specifies settings that define a backup schedule for a Protection Job.
    Name                            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the Protection Policy.
    Retries                         *int64          `json:"retries,omitempty" form:"retries,omitempty"` //Specifies the number of times to retry capturing Snapshots before
    RetryIntervalMins               *int64          `json:"retryIntervalMins,omitempty" form:"retryIntervalMins,omitempty"` //Specifies the number of minutes before retrying a failed Protection Job.
    SkipIntervalMins                *int64          `json:"skipIntervalMins,omitempty" form:"skipIntervalMins,omitempty"` //Specifies the period of time before skipping the execution of new
    SnapshotArchivalCopyPolicies    []*SnapshotCopyArchivalPolicy `json:"snapshotArchivalCopyPolicies,omitempty" form:"snapshotArchivalCopyPolicies,omitempty"` //Array of External Targets.
    SnapshotReplicationCopyPolicies []*SnapshotCopyReplicationPolicy `json:"snapshotReplicationCopyPolicies,omitempty" form:"snapshotReplicationCopyPolicies,omitempty"` //Array of Remote Clusters.
    SystemSchedulingPolicy          SchedulingPolicy `json:"systemSchedulingPolicy,omitempty" form:"systemSchedulingPolicy,omitempty"` //Specifies settings that define a backup schedule for a Protection Job.
    WormRetentionType               WormRetentionType1Enum `json:"wormRetentionType,omitempty" form:"wormRetentionType,omitempty"` //Specifies WORM retention type for the snapshots. When a WORM retention
}

/*
 * Structure for the custom type ProtectionPolicySummary
 */
type ProtectionPolicySummary struct {
    LastProtectionRunSummary LastProtectionRunSummary `json:"lastProtectionRunSummary,omitempty" form:"lastProtectionRunSummary,omitempty"` //LastProtectionRunsSummary is the summary of the last Protection Run for the
    ProtectedSourcesSummary  []*ProtectedSourceSummary `json:"protectedSourcesSummary,omitempty" form:"protectedSourcesSummary,omitempty"` //Specifies the list of Protection Sources which are protected under the
    ProtectionJobsSummary    []*ProtectionJobSummaryForPolicies `json:"protectionJobsSummary,omitempty" form:"protectionJobsSummary,omitempty"` //Specifies the list of Protection Jobs associated with the given
    ProtectionPolicy         ProtectionPolicy `json:"protectionPolicy,omitempty" form:"protectionPolicy,omitempty"` //TODO: Write general description for this field
    ProtectionRunsSummary    ProtectionRunsSummary `json:"protectionRunsSummary,omitempty" form:"protectionRunsSummary,omitempty"` //ProtectionRunsSummary is the summary of the all the Protection Runs for the
}

/*
 * Structure for the custom type ProtectionJobRunInstance
 */
type ProtectionJobRunInstance struct {
    BackupRun       BackupRunTask   `json:"backupRun,omitempty" form:"backupRun,omitempty"` //Specifies details about the Backup task for a Job Run.
    CopyRun         []*CopyRunTask  `json:"copyRun,omitempty" form:"copyRun,omitempty"` //Array of Copy Run Tasks.
    JobId           *int64          `json:"jobId,omitempty" form:"jobId,omitempty"` //Specifies the id of the Protection Job that was run.
    JobName         *string         `json:"jobName,omitempty" form:"jobName,omitempty"` //Specifies the name of the Protection Job name that was run.
    JobUid          UniqueGlobalId  `json:"jobUid,omitempty" form:"jobUid,omitempty"` //Specifies the globally unique id of the Protection Job that was run.
    ViewBoxId       *int64          `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"` //Specifies the Storage Domain (View Box) to store the backed up data.
}

/*
 * Structure for the custom type ProtectionRunsSummary
 */
type ProtectionRunsSummary struct {
    NumberOfArchivalRuns              *int64          `json:"numberOfArchivalRuns,omitempty" form:"numberOfArchivalRuns,omitempty"` //Specifies the total number of Archival Runs using the current
    NumberOfProtectionRuns            *int64          `json:"numberOfProtectionRuns,omitempty" form:"numberOfProtectionRuns,omitempty"` //Specifies the total number of Protection Runs by the given Protection
    NumberOfReplicationRuns           *int64          `json:"numberOfReplicationRuns,omitempty" form:"numberOfReplicationRuns,omitempty"` //Specifies the total number of Replication Runs using the current
    NumberOfSuccessfulArchivalRuns    *int64          `json:"numberOfSuccessfulArchivalRuns,omitempty" form:"numberOfSuccessfulArchivalRuns,omitempty"` //Specifies the number of total successful Archival Runs using the
    NumberOfSuccessfulProtectionRuns  *int64          `json:"numberOfSuccessfulProtectionRuns,omitempty" form:"numberOfSuccessfulProtectionRuns,omitempty"` //Specifies the number of successful Protection Runs using the current
    NumberOfSuccessfulReplicationRuns *int64          `json:"numberOfSuccessfulReplicationRuns,omitempty" form:"numberOfSuccessfulReplicationRuns,omitempty"` //Specifies the number of total successful Replication Runs using the
}

/*
 * Structure for the custom type ProtectionSource
 */
type ProtectionSource struct {
    AcropolisProtectionSource  AcropolisProtectionSource `json:"acropolisProtectionSource,omitempty" form:"acropolisProtectionSource,omitempty"` //Specifies details about an Acropolis Protection Source
    AwsProtectionSource        AWSProtectionSource `json:"awsProtectionSource,omitempty" form:"awsProtectionSource,omitempty"` //Specifies details about an AWS Protection Source
    AzureProtectionSource      AzureProtectionSource `json:"azureProtectionSource,omitempty" form:"azureProtectionSource,omitempty"` //Specifies details about an Azure Protection Source
    Environment                EnvironmentEnum `json:"environment,omitempty" form:"environment,omitempty"` //Specifies the environment (such as 'kVMware' or 'kSQL') where the
    FlashBladeProtectionSource PureStorageFlashBladeProtectionSource `json:"flashBladeProtectionSource,omitempty" form:"flashBladeProtectionSource,omitempty"` //Specifies details about a Pure Storage FlashBlade Protection Source
    GcpProtectionSource        GCPProtectionSource `json:"gcpProtectionSource,omitempty" form:"gcpProtectionSource,omitempty"` //Specifies details about an GCP Protection Source
    HyperFlexProtectionSource  HyperFlexStoraeSnapshot `json:"hyperFlexProtectionSource,omitempty" form:"hyperFlexProtectionSource,omitempty"` //Specifies details about a HyperFlex Storage Snapshot source
    HypervProtectionSource     HypervProtectionSource `json:"hypervProtectionSource,omitempty" form:"hypervProtectionSource,omitempty"` //Specifies details about a HyperV Protection Source
    Id                         *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies an id of the Protection Source.
    IsilonProtectionSource     IsilonProtectionSource `json:"isilonProtectionSource,omitempty" form:"isilonProtectionSource,omitempty"` //Specifies details about an Isilon OneFs Protection Source
    KvmProtectionSource        KVMProtectionSource `json:"kvmProtectionSource,omitempty" form:"kvmProtectionSource,omitempty"` //Specifies details about a KVM Protection Source
    Name                       *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies a name of the Protection Source.
    NasProtectionSource        GenericNASProtectionSource `json:"nasProtectionSource,omitempty" form:"nasProtectionSource,omitempty"` //Specifies details about a Generic NAS Protection Source
    NetappProtectionSource     NetAppProtectionSource `json:"netappProtectionSource,omitempty" form:"netappProtectionSource,omitempty"` //Specifies details about a NetApp Protection Source
    Office365ProtectionSource  Office365ProtectionSource `json:"office365ProtectionSource,omitempty" form:"office365ProtectionSource,omitempty"` //Specifies details about an Office 365 Protection Source
    OracleProtectionSource     OracleProtectionSource `json:"oracleProtectionSource,omitempty" form:"oracleProtectionSource,omitempty"` //Specifies details about an Oracle Protection Source
    ParentId                   *int64          `json:"parentId,omitempty" form:"parentId,omitempty"` //Specifies an id of the parent of the Protection Source.
    PhysicalProtectionSource   PhysicalProtectionSource `json:"physicalProtectionSource,omitempty" form:"physicalProtectionSource,omitempty"` //Specifies details about a Physical Protection Source
    PureProtectionSource       PureProtectionSource `json:"pureProtectionSource,omitempty" form:"pureProtectionSource,omitempty"` //Specifies details about a Pure Protection Source
    SqlProtectionSource        SQLProtectionSource `json:"sqlProtectionSource,omitempty" form:"sqlProtectionSource,omitempty"` //Specifies details about a SQL Protection Source
    ViewProtectionSource       ViewProtectionSource `json:"viewProtectionSource,omitempty" form:"viewProtectionSource,omitempty"` //Specifies details about a View Protection Source
    VmwareProtectionSource     VmwareProtectionSource `json:"vmWareProtectionSource,omitempty" form:"vmWareProtectionSource,omitempty"` //Specifies details about a VMware Protection Source
}

/*
 * Structure for the custom type NodeInAProtectionSourcesTree
 */
type NodeInAProtectionSourcesTree struct {
    ApplicationNodes          *[]interface{}  `json:"applicationNodes,omitempty" form:"applicationNodes,omitempty"` //Array of Child Subtrees.
    EntityPermissionInfo      EntityPermissionInformation `json:"entityPermissionInfo,omitempty" form:"entityPermissionInfo,omitempty"` //Specifies the permission information of entities.
    LogicalSize               *int64          `json:"logicalSize,omitempty" form:"logicalSize,omitempty"` //Specifies the logical size of the data in bytes for the Object
    Nodes                     *[]interface{}  `json:"nodes,omitempty" form:"nodes,omitempty"` //Array of Child Nodes.
    ProtectedSourcesSummary   []*AggregatedSubtreeInfo `json:"protectedSourcesSummary,omitempty" form:"protectedSourcesSummary,omitempty"` //Array of Protected Objects.
    ProtectionSource          ProtectionSource `json:"protectionSource,omitempty" form:"protectionSource,omitempty"` //Specifies the Protection Source for the current node.
    RegistrationInfo          RegisteredSourceInfo `json:"registrationInfo,omitempty" form:"registrationInfo,omitempty"` //Specifies registration information for a root node in a Protection
    UnprotectedSourcesSummary []*AggregatedSubtreeInfo `json:"unprotectedSourcesSummary,omitempty" form:"unprotectedSourcesSummary,omitempty"` //Array of Unprotected Sources.
}

/*
 * Structure for the custom type ProtectionSourceResponse
 */
type ProtectionSourceResponse struct {
    Jobs                    []*SpecifiesProtectionJobSummaryOfAnObject `json:"jobs,omitempty" form:"jobs,omitempty"` //Specifies the list of Protection Jobs that protect the object.
    LogicalSizeInBytes      *int64          `json:"logicalSizeInBytes,omitempty" form:"logicalSizeInBytes,omitempty"` //Specifies the logical size of Protection Source in bytes.
    ParentSource            ProtectionSource `json:"parentSource,omitempty" form:"parentSource,omitempty"` //Specifies a generic structure that represents a node
    ProtectionSourceUidList []*PreotectionSourceUID `json:"protectionSourceUidList,omitempty" form:"protectionSourceUidList,omitempty"` //Specifies the list of universal ids of the Protection Source.
    Source                  ProtectionSource `json:"source,omitempty" form:"source,omitempty"` //Specifies a generic structure that represents a node
    Uuid                    *string         `json:"uuid,omitempty" form:"uuid,omitempty"` //Specifies the unique id of the Protection Source.
}

/*
 * Structure for the custom type ProtectionSourceSnapshot
 */
type ProtectionSourceSnapshot struct {
    CopyTasks                []*SnapshotCopyTask `json:"copyTasks,omitempty" form:"copyTasks,omitempty"` //Array of Snapshot Copy Tasks.
    JobId                    *int64          `json:"jobId,omitempty" form:"jobId,omitempty"` //Specifies the id of the Protection Job.
    JobName                  *string         `json:"jobName,omitempty" form:"jobName,omitempty"` //Specifies the name of the Protection Job.
    JobRunId                 *int64          `json:"jobRunId,omitempty" form:"jobRunId,omitempty"` //Specifies the id of the Job Run.
    JobRunStartTimeUsecs     *int64          `json:"jobRunStartTimeUsecs,omitempty" form:"jobRunStartTimeUsecs,omitempty"` //Specifies the start time of the Job which this object is part of.
    LastRunEndTimeUsecs      *int64          `json:"lastRunEndTimeUsecs,omitempty" form:"lastRunEndTimeUsecs,omitempty"` //Specifies the end time of the last Run of this object's snapshot.
    LastRunStartTimeUsecs    *int64          `json:"lastRunStartTimeUsecs,omitempty" form:"lastRunStartTimeUsecs,omitempty"` //Specifies the start time of the last Run of this object's snapshot.
    Message                  *string         `json:"message,omitempty" form:"message,omitempty"` //Specifies warning or error information when the Job Run is not
    NumBytesRead             *int64          `json:"numBytesRead,omitempty" form:"numBytesRead,omitempty"` //Specifies the total number of bytes read.
    NumLogicalBytesProtected *int64          `json:"numLogicalBytesProtected,omitempty" form:"numLogicalBytesProtected,omitempty"` //Specifies the total number of logical bytes that are protected. The
    PaginationCookie         *int64          `json:"paginationCookie,omitempty" form:"paginationCookie,omitempty"` //Specifies an opaque string to pass into the next request to get
    RunStatus                RunStatusEnum   `json:"runStatus,omitempty" form:"runStatus,omitempty"` //Specifies the type of the Job Run.
    RunType                  RunType1Enum    `json:"runType,omitempty" form:"runType,omitempty"` //Specifies the status of the Job Run.
}

/*
 * Structure for the custom type RegistrationAndProtectionInformation
 */
type RegistrationAndProtectionInformation struct {
    Applications         []*ApplicationInformation `json:"applications,omitempty" form:"applications,omitempty"` //Array of applications hierarchy registered on this node.
    EntityPermissionInfo EntityPermissionInformation `json:"entityPermissionInfo,omitempty" form:"entityPermissionInfo,omitempty"` //Specifies the permission information of entities.
    LogicalSizeBytes     *int64          `json:"logicalSizeBytes,omitempty" form:"logicalSizeBytes,omitempty"` //Specifies the logical size of the Protection Source in bytes.
    RegistrationInfo     RegisteredSourceInfo `json:"registrationInfo,omitempty" form:"registrationInfo,omitempty"` //Specifies registration information for a root node in a Protection
    RootNode             ProtectionSource `json:"rootNode,omitempty" form:"rootNode,omitempty"` //Specifies the Protection Source for the root node of the Protection
    Stats                ProtectionSummary `json:"stats,omitempty" form:"stats,omitempty"` //Specifies the stats of protection for a Protection Source Tree.
    StatsByEnv           []*ProtectionSummaryByEnvironment `json:"statsByEnv,omitempty" form:"statsByEnv,omitempty"` //Specifies the breakdown of the stats of protection by environment.
}

/*
 * Structure for the custom type PreotectionSourceUID
 */
type PreotectionSourceUID struct {
    ClusterId            *int64          `json:"clusterId,omitempty" form:"clusterId,omitempty"` //Specifies the id of the cluster on which object is present.
    ClusterIncarnationId *int64          `json:"clusterIncarnationId,omitempty" form:"clusterIncarnationId,omitempty"` //Specifies the incarnation id of the cluster on which object is present.
    ParentSourceId       *int64          `json:"parentSourceId,omitempty" form:"parentSourceId,omitempty"` //Specifies parent source id of an object.
    SourceId             *int64          `json:"sourceId,omitempty" form:"sourceId,omitempty"` //Specifies source id of an object.
}

/*
 * Structure for the custom type ProtectionSourcesSummaryStatistics
 */
type ProtectionSourcesSummaryStatistics struct {
    FirstFailedRunTimeUsecs     *int64          `json:"firstFailedRunTimeUsecs,omitempty" form:"firstFailedRunTimeUsecs,omitempty"` //Specifies the start time of the first failed Job Run protecting this
    FirstSuccessfulRunTimeUsecs *int64          `json:"firstSuccessfulRunTimeUsecs,omitempty" form:"firstSuccessfulRunTimeUsecs,omitempty"` //Specifies the start time of the first successful Job Run protecting this
    LastFailedRunTimeUsecs      *int64          `json:"lastFailedRunTimeUsecs,omitempty" form:"lastFailedRunTimeUsecs,omitempty"` //Specifies the start time of the last failed Job Run protecting this
    LastRunEndTimeUsecs         *int64          `json:"lastRunEndTimeUsecs,omitempty" form:"lastRunEndTimeUsecs,omitempty"` //Specifies the end time of the last Job Run protecting this
    LastRunErrorMsg             *string         `json:"lastRunErrorMsg,omitempty" form:"lastRunErrorMsg,omitempty"` //Specifies the error message associated with last run, if the last run
    LastRunStartTimeUsecs       *int64          `json:"lastRunStartTimeUsecs,omitempty" form:"lastRunStartTimeUsecs,omitempty"` //Specifies the start time of the last Job Run protecting this
    LastRunStatus               LastRunStatusEnum `json:"lastRunStatus,omitempty" form:"lastRunStatus,omitempty"` //Specifies the Job Run status of the last Job Run protecting
    LastRunType                 LastRunTypeEnum `json:"lastRunType,omitempty" form:"lastRunType,omitempty"` //Specifies the Job Run type of the last Job Run protecting
    LastSuccessfulRunTimeUsecs  *int64          `json:"lastSuccessfulRunTimeUsecs,omitempty" form:"lastSuccessfulRunTimeUsecs,omitempty"` //Specifies the start time of the last successful Job Run protecting this
    NumDataReadBytes            *int64          `json:"numDataReadBytes,omitempty" form:"numDataReadBytes,omitempty"` //Specifies the total number of bytes read while protecting this
    NumErrors                   *int64          `json:"numErrors,omitempty" form:"numErrors,omitempty"` //Specifies the total number of errors reported during Job Runs
    NumLogicalBytesProtected    *int64          `json:"numLogicalBytesProtected,omitempty" form:"numLogicalBytesProtected,omitempty"` //Specifies the total logical bytes protected for this
    NumSnapshots                *int64          `json:"numSnapshots,omitempty" form:"numSnapshots,omitempty"` //Specifies the total number of Snapshots that are backing up this
    NumSuccessRuns              *int64          `json:"numSuccessRuns,omitempty" form:"numSuccessRuns,omitempty"` //Specifies the total number of successful Job Runs protecting this
    NumWarnings                 *int64          `json:"numWarnings,omitempty" form:"numWarnings,omitempty"` //Specifies the total number of warnings reported during Job Runs
    ProtectionSource            ProtectionSource `json:"protectionSource,omitempty" form:"protectionSource,omitempty"` //Specifies the leaf Protection Source Object (such as VM).
    RegisteredSource            *string         `json:"registeredSource,omitempty" form:"registeredSource,omitempty"` //Specifies the name of the Registered Source that is the top level
    Tenants                     []*TenantInformation `json:"tenants,omitempty" form:"tenants,omitempty"` //Specifies basic information about tenants having access to the protection
}

/*
 * Structure for the custom type ProtectionSummary
 */
type ProtectionSummary struct {
    ProtectedCount   *int64          `json:"protectedCount,omitempty" form:"protectedCount,omitempty"` //Specifies the number of objects that are protected under the given
    ProtectedSize    *int64          `json:"protectedSize,omitempty" form:"protectedSize,omitempty"` //Specifies the total size of the protected objects under the given entity.
    UnprotectedCount *int64          `json:"unprotectedCount,omitempty" form:"unprotectedCount,omitempty"` //Specifies the number of objects that are not protected under the given
    UnprotectedSize  *int64          `json:"unprotectedSize,omitempty" form:"unprotectedSize,omitempty"` //Specifies the total size of the unprotected objects under the given
}

/*
 * Structure for the custom type ProtectionSummaryByEnvironment
 */
type ProtectionSummaryByEnvironment struct {
    Environment      Environment8Enum `json:"environment,omitempty" form:"environment,omitempty"` //Specifies the type of environment of the source object like kSQL etc.
    ProtectedCount   *int64          `json:"protectedCount,omitempty" form:"protectedCount,omitempty"` //Specifies the number of objects that are protected under the given
    ProtectedSize    *int64          `json:"protectedSize,omitempty" form:"protectedSize,omitempty"` //Specifies the total size of the protected objects under the given entity.
    UnprotectedCount *int64          `json:"unprotectedCount,omitempty" form:"unprotectedCount,omitempty"` //Specifies the number of objects that are not protected under the given
    UnprotectedSize  *int64          `json:"unprotectedSize,omitempty" form:"unprotectedSize,omitempty"` //Specifies the total size of the unprotected objects under the given
}

/*
 * Structure for the custom type ProtectionTile
 */
type ProtectionTile struct {
    LastDayArchival       ProtectionStatistics `json:"lastDayArchival,omitempty" form:"lastDayArchival,omitempty"` //Protection Statistics.
    LastDayBackup         ProtectionStatistics `json:"lastDayBackup,omitempty" form:"lastDayBackup,omitempty"` //Protection Statistics.
    LastDayReplicationIn  ProtectionStatistics `json:"lastDayReplicationIn,omitempty" form:"lastDayReplicationIn,omitempty"` //Protection Statistics.
    LastDayReplicationOut ProtectionStatistics `json:"lastDayReplicationOut,omitempty" form:"lastDayReplicationOut,omitempty"` //Protection Statistics.
}

/*
 * Structure for the custom type PureProtectionSource
 */
type PureProtectionSource struct {
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies a unique name of the Protection Source
    StorageArray    PureStorageArray `json:"storageArray,omitempty" form:"storageArray,omitempty"` //Specifies a Pure Storage Array.
    Type            Type16Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of managed Object in a pure Protection Source like
    Volume          PureVolume      `json:"volume,omitempty" form:"volume,omitempty"` //Specifies a Pure Volume in a Pure Storage Array.
}

/*
 * Structure for the custom type PureStorageArray
 */
type PureStorageArray struct {
    Id              *string         `json:"id,omitempty" form:"id,omitempty"` //Specifies a unique id of a Pure Storage Array.
    Ports           []*ISCSISANPort `json:"ports,omitempty" form:"ports,omitempty"` //Specifies the SAN ports of the Pure Storage Array.
    Revision        *string         `json:"revision,omitempty" form:"revision,omitempty"` //Specifies the revision of the Pure Storage Array.
    Version         *string         `json:"version,omitempty" form:"version,omitempty"` //Specifies the version of the Pure Storage Array.
}

/*
 * Structure for the custom type PureVolume
 */
type PureVolume struct {
    CreatedTime     *string         `json:"createdTime,omitempty" form:"createdTime,omitempty"` //Specifies the created time (e.g., "2015-07-21T17:59:41Z") of the volume.
    ParentVolume    *string         `json:"parentVolume,omitempty" form:"parentVolume,omitempty"` //Specifies the name of the source volume, if this volume was
    SerialNumber    *string         `json:"serialNumber,omitempty" form:"serialNumber,omitempty"` //Specifies the serial number of the volume.
    SizeBytes       *int64          `json:"sizeBytes,omitempty" form:"sizeBytes,omitempty"` //Specifies the provisioned size in bytes of the volume.
    UsedBytes       *int64          `json:"usedBytes,omitempty" form:"usedBytes,omitempty"` //Specifies the total space actually used by the volume.
}

/*
 * Structure for the custom type QStarServerCredentials
 */
type QStarServerCredentials struct {
    Host                *string         `json:"host,omitempty" form:"host,omitempty"` //Specifies the IP address or DNS name of the server where QStar
    IntegralVolumeNames *[]string       `json:"integralVolumeNames,omitempty" form:"integralVolumeNames,omitempty"` //Array of Integral Volume Names.
    Password            *string         `json:"password,omitempty" form:"password,omitempty"` //Specifies the password used to access the QStar host.
    Port                *int64          `json:"port,omitempty" form:"port,omitempty"` //Specifies the listening port where QStar WEB API service is running.
    ShareType           *string         `json:"shareType,omitempty" form:"shareType,omitempty"` //Specifies the sharing protocol type used by QStar to mount
    UseHttps            *bool           `json:"useHttps,omitempty" form:"useHttps,omitempty"` //Specifies whether to use http or https to connect to the service.
    Username            *string         `json:"username,omitempty" form:"username,omitempty"` //Specifies the account name used to access the QStar host.
}

/*
 * Structure for the custom type QuotaAndUsageInView
 */
type QuotaAndUsageInView struct {
    Quota           QuotaPolicy     `json:"quota,omitempty" form:"quota,omitempty"` //Specifies a quota limit that can be optionally applied to Views and
    UsageBytes      *int64          `json:"usageBytes,omitempty" form:"usageBytes,omitempty"` //Usage in bytes of this user in this view.
    ViewId          *int64          `json:"viewId,omitempty" form:"viewId,omitempty"` //The usage and quota policy information of this user for this view.
    ViewName        *string         `json:"viewName,omitempty" form:"viewName,omitempty"` //View name.
}

/*
 * Structure for the custom type QuotaPolicy
 */
type QuotaPolicy struct {
    AlertLimit               *string         `json:"AlertLimit,omitempty" form:"AlertLimit,omitempty"` //AlertLimitBytes converted to GiB format for report purposes.
    HardLimit                *string         `json:"HardLimit,omitempty" form:"HardLimit,omitempty"` //HardLimitBytes converted to GiB format for report purposes.
    AlertLimitBytes          *int64          `json:"alertLimitBytes,omitempty" form:"alertLimitBytes,omitempty"` //Specifies if an alert should be triggered when the usage of this
    AlertThresholdPercentage *int64          `json:"alertThresholdPercentage,omitempty" form:"alertThresholdPercentage,omitempty"` //Supported only for user quota policy.
    HardLimitBytes           *int64          `json:"hardLimitBytes,omitempty" form:"hardLimitBytes,omitempty"` //Specifies an optional quota limit on the usage allowed for this
}

/*
 * Structure for the custom type CreateRestoreTaskRequest
 */
type CreateRestoreTaskRequest struct {
    AcropolisParameters          AcropolisRestoreParameters `json:"acropolisParameters,omitempty" form:"acropolisParameters,omitempty"` //This field defines the Acropolis specific params for restore tasks of type
    ContinueOnError              *bool           `json:"continueOnError,omitempty" form:"continueOnError,omitempty"` //Specifies if the Restore Task should continue when some operations on some
    GlacierRetrievalType         GlacierRetrievalTypeEnum `json:"glacierRetrievalType,omitempty" form:"glacierRetrievalType,omitempty"` //Specifies the way data needs to be retrieved from the external target.
    HypervParameters             HypervRestoreParameters `json:"hypervParameters,omitempty" form:"hypervParameters,omitempty"` //Specifies information needed when restoring VMs in HyperV enviroment.
    MountParameters              MountingVolumes `json:"mountParameters,omitempty" form:"mountParameters,omitempty"` //Specifies the information required for mounting volumes.
    Name                         string          `json:"name" form:"name"` //Specifies the name of the Restore Task. This field must be set and
    NewParentId                  *int64          `json:"newParentId,omitempty" form:"newParentId,omitempty"` //Specify a new registered parent Protection Source. If specified
    Objects                      []*RestoreObject `json:"objects,omitempty" form:"objects,omitempty"` //Array of Objects.
    OutlookParameters            OutlookRestoreParameters `json:"outlookParameters,omitempty" form:"outlookParameters,omitempty"` //Specifies information needed for recovering Mailboxes in O365Outlook
    RestoreViewParameters        View1           `json:"restoreViewParameters,omitempty" form:"restoreViewParameters,omitempty"` //Specifies the settings that define a View.
    Type                         Type30Enum      `json:"type" form:"type"` //Specifies the type of Restore Task such as 'kRecoverVMs' or
    ViewName                     *string         `json:"viewName,omitempty" form:"viewName,omitempty"` //Specifie target view into which the objects are to be cloned when doing
    VirtualDiskRestoreParameters VirtualDiskRestoreParameters `json:"virtualDiskRestoreParameters,omitempty" form:"virtualDiskRestoreParameters,omitempty"` //Specifies the parameters to recover virtual disks of a vm.
    VlanParameters               VLANParameters  `json:"vlanParameters,omitempty" form:"vlanParameters,omitempty"` //Specifies VLAN parameters for the restore operation.
    VmwareParameters             VmwareRestoreParameters `json:"vmwareParameters,omitempty" form:"vmwareParameters,omitempty"` //Specifies the information required for recovering or cloning VmWare VMs.
}

/*
 * Structure for the custom type RecoveriesTile
 */
type RecoveriesTile struct {
    LastMonthNumRecoveries     *int64          `json:"lastMonthNumRecoveries,omitempty" form:"lastMonthNumRecoveries,omitempty"` //Number of Recoveries in the last 30 days.
    LastMonthRecoveriesByType  []*RestoreCountByObjectType `json:"lastMonthRecoveriesByType,omitempty" form:"lastMonthRecoveriesByType,omitempty"` //Recoveries by Type in the last month.
    LastMonthRecoverySizeBytes *int64          `json:"lastMonthRecoverySizeBytes,omitempty" form:"lastMonthRecoverySizeBytes,omitempty"` //Bytes recovered in the last 30 days.
    RecoveryNumRunning         *int64          `json:"recoveryNumRunning,omitempty" form:"recoveryNumRunning,omitempty"` //Number of recoveries that are currently running.
}

/*
 * Structure for the custom type RecoveryTaskInformation
 */
type RecoveryTaskInformation struct {
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Name of the recovery task.
    TaskId          *string         `json:"taskId,omitempty" form:"taskId,omitempty"` //Id of the recovery task.
    Type            Type28Enum      `json:"type,omitempty" form:"type,omitempty"` //Denotes if the recovery task has an archival target.
}

/*
 * Structure for the custom type InformationAboutAReducer
 */
type InformationAboutAReducer struct {
    Code            *string         `json:"code,omitempty" form:"code,omitempty"` //The code of the reducer in the specified language.
    Id              *int64          `json:"id,omitempty" form:"id,omitempty"` //Reduced ID generated by system. Absent when user is creating a new
    IsSystemDefined *bool           `json:"isSystemDefined,omitempty" form:"isSystemDefined,omitempty"` //Whether the mapper is system defined.
    JarName         *string         `json:"jarName,omitempty" form:"jarName,omitempty"` //User can write their own mapper/reducer or upload jar files containing
    JarPath         *string         `json:"jarPath,omitempty" form:"jarPath,omitempty"` //path of JAR in which this reducer was found. This is applicable only when
    Language        *int64          `json:"language,omitempty" form:"language,omitempty"` //Programming language used by the reducer.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Name of the reducer.
}

/*
 * Structure for the custom type RegisterApplicationServersParameters
 */
type RegisterApplicationServersParameters struct {
    Applications       *[]ApplicationEnum `json:"applications,omitempty" form:"applications,omitempty"` //Specifies the types of applications such as 'kSQL', 'kExchange' running
    HasPersistentAgent *bool           `json:"hasPersistentAgent,omitempty" form:"hasPersistentAgent,omitempty"` //Set this to true if a persistent agent is running on the host. If this is
    Password           *string         `json:"password,omitempty" form:"password,omitempty"` //Specifies password of the username to access the target source.
    ProtectionSourceId *int64          `json:"protectionSourceId,omitempty" form:"protectionSourceId,omitempty"` //Specifies the Id of the Protection Source that contains one or more
    Username           *string         `json:"username,omitempty" form:"username,omitempty"` //Specifies username to access the target source.
}

/*
 * Structure for the custom type RegisterProtectionSourceParameters
 */
type RegisterProtectionSourceParameters struct {
    AwsCredentials            AWSSourceCredentials `json:"awsCredentials,omitempty" form:"awsCredentials,omitempty"` //Specifies the credentials to authenticate with AWS Cloud Platform.
    AzureCredentials          AzureSourceCredentials `json:"azureCredentials,omitempty" form:"azureCredentials,omitempty"` //Specifies the credentials to authenticate with Azure Cloud Platform.
    Endpoint                  *string         `json:"endpoint,omitempty" form:"endpoint,omitempty"` //Specifies the network endpoint of the Protection Source where it is
    Environment               Environment12Enum `json:"environment,omitempty" form:"environment,omitempty"` //Specifies the environment such as 'kPhysical' or 'kVMware' of the
    ForceRegister             *bool           `json:"forceRegister,omitempty" form:"forceRegister,omitempty"` //TODO: Write general description for this field
    GcpCredentials            GCPSourceCredentials `json:"gcpCredentials,omitempty" form:"gcpCredentials,omitempty"` //Specifies the credentials to authenticate with Google Cloud Platform.
    HostType                  HostType7Enum   `json:"hostType,omitempty" form:"hostType,omitempty"` //Specifies the optional OS type of the Protection Source (such as kWindows
    NasMountCredentials       NASMountCredentials `json:"nasMountCredentials,omitempty" form:"nasMountCredentials,omitempty"` //Specifies the server credentials to connect to a NetApp server.
    NetappType                NetappTypeEnum  `json:"netappType,omitempty" form:"netappType,omitempty"` //Specifies the entity type such as 'kCluster,' if the environment is
    Office365Type             *int64          `json:"office365Type,omitempty" form:"office365Type,omitempty"` //Specifies the entity type such as 'kDomain', 'kOutlook', 'kMailbox', if the
    Password                  *string         `json:"password,omitempty" form:"password,omitempty"` //Specifies password of the username to access the target source.
    PhysicalType              PhysicalTypeEnum `json:"physicalType,omitempty" form:"physicalType,omitempty"` //Specifies the entity type such as 'kPhysicalHost' if the environment is
    PureType                  PureTypeEnum    `json:"pureType,omitempty" form:"pureType,omitempty"` //Specifies the entity type such as 'kStorageArray' if the environment is
    SourceSideDedupEnabled    *bool           `json:"sourceSideDedupEnabled,omitempty" form:"sourceSideDedupEnabled,omitempty"` //This controls whether to use source side dedup on the source or not.
    SslVerification           SSLCertificationVerificationSettings `json:"sslVerification,omitempty" form:"sslVerification,omitempty"` //Specifies information about SSL verification when registering certain
    ThrottlingPolicy          ThrottlingPolicy `json:"throttlingPolicy,omitempty" form:"throttlingPolicy,omitempty"` //Specifies the throttling policy that should be applied to this Source.
    ThrottlingPolicyOverrides []*ThrottlingPolicyOverride `json:"throttlingPolicyOverrides,omitempty" form:"throttlingPolicyOverrides,omitempty"` //Array of Throttling Policy Overrides for Datastores.
    Username                  *string         `json:"username,omitempty" form:"username,omitempty"` //Specifies username to access the target source.
    VmwareType                VmwareTypeEnum  `json:"vmwareType,omitempty" form:"vmwareType,omitempty"` //Specifies the entity type such as 'kVCenter' if the environment is
}

/*
 * Structure for the custom type RegisterRemoteCluster
 */
type RegisterRemoteCluster struct {
    AllEndpointsReachable   *bool           `json:"allEndpointsReachable,omitempty" form:"allEndpointsReachable,omitempty"` //Specifies whether any endpoint (such as a Node) on the remote Cluster
    BandwidthLimit          BandwidthLimit  `json:"bandwidthLimit,omitempty" form:"bandwidthLimit,omitempty"` //Specifies settings for limiting the data transfer rate between
    ClearInterfaces         *bool           `json:"clearInterfaces,omitempty" form:"clearInterfaces,omitempty"` //TODO: Write general description for this field
    ClearVlanId             *bool           `json:"clearVlanId,omitempty" form:"clearVlanId,omitempty"` //Specifies whether to clear the vlanId field, and thus stop
    ClusterId               *int64          `json:"clusterId,omitempty" form:"clusterId,omitempty"` //Specifies the unique id of the remote Cluster.
    CompressionEnabled      *bool           `json:"compressionEnabled,omitempty" form:"compressionEnabled,omitempty"` //Specifies whether to compress the outbound data when
    EncryptionKey           *string         `json:"encryptionKey,omitempty" form:"encryptionKey,omitempty"` //Specifies the encryption key used for encrypting the replication data
    IfaceName               *string         `json:"ifaceName,omitempty" form:"ifaceName,omitempty"` //Specifies the interface name of the VLAN to use for communicating with
    NetworkInterfaceGroup   *string         `json:"networkInterfaceGroup,omitempty" form:"networkInterfaceGroup,omitempty"` //Specifies the group name of the network interfaces to
    NetworkInterfaceIds     *[]int64        `json:"networkInterfaceIds,omitempty" form:"networkInterfaceIds,omitempty"` //Array of Network Interface Ids.
    Password                *string         `json:"password,omitempty" form:"password,omitempty"` //Specifies the password for Cohesity user to use when
    PurposeRemoteAccess     *bool           `json:"purposeRemoteAccess,omitempty" form:"purposeRemoteAccess,omitempty"` //Whether the remote cluster will be used for remote access for SPOG.
    PurposeReplication      *bool           `json:"purposeReplication,omitempty" form:"purposeReplication,omitempty"` //Whether the remote cluster will be used for replication.
    RemoteAccessCredentials CreateAccessTokenCredentialRequest `json:"remoteAccessCredentials,omitempty" form:"remoteAccessCredentials,omitempty"` //Specifies the Cohesity credentials required for generating an access token.
    RemoteIps               *[]string       `json:"remoteIps,omitempty" form:"remoteIps,omitempty"` //Array of Remote Node IP Addresses.
    RemoteIrisPorts         *[]int64        `json:"remoteIrisPorts,omitempty" form:"remoteIrisPorts,omitempty"` //Array of Ports.
    UserName                *string         `json:"userName,omitempty" form:"userName,omitempty"` //Specifies the Cohesity user name used to connect to the
    ValidateOnly            *bool           `json:"validateOnly,omitempty" form:"validateOnly,omitempty"` //Whether to only validate the credentials without saving the information.
    ViewBoxPairInfo         []*StorageDomainViewBoxPairing `json:"viewBoxPairInfo,omitempty" form:"viewBoxPairInfo,omitempty"` //Array of Storage Domain (View Box) Pairs.
    VlanId                  *int64          `json:"vlanId,omitempty" form:"vlanId,omitempty"` //Specifies the Id of the VLAN to use for communicating with the remote
}

/*
 * Structure for the custom type RegisteredApplicationServer
 */
type RegisteredApplicationServer struct {
    ApplicationServer          NodeInAProtectionSourcesTree `json:"applicationServer,omitempty" form:"applicationServer,omitempty"` //Specifies the child subtree used to store additional application-level
    RegisteredProtectionSource ProtectionSource `json:"registeredProtectionSource,omitempty" form:"registeredProtectionSource,omitempty"` //Specifies the Protection Source like a VM or Physical Server that
}

/*
 * Structure for the custom type RegisteredSourceInfo
 */
type RegisteredSourceInfo struct {
    AccessInfo                 ConnectorParameters `json:"accessInfo,omitempty" form:"accessInfo,omitempty"` //Specifies the parameters required to establish a connection with
    AuthenticationErrorMessage *string         `json:"authenticationErrorMessage,omitempty" form:"authenticationErrorMessage,omitempty"` //Specifies an authentication error message. This indicates the given
    AuthenticationStatus       AuthenticationStatusEnum `json:"authenticationStatus,omitempty" form:"authenticationStatus,omitempty"` //Specifies the status of the authenticating to the Protection Source
    Environments               *[]Environment2Enum `json:"environments,omitempty" form:"environments,omitempty"` //Specifies a list of applications environment that are registered
    MinimumFreeSpaceGB         *int64          `json:"minimumFreeSpaceGB,omitempty" form:"minimumFreeSpaceGB,omitempty"` //Specifies the minimum free space in GiB of the space expected to be
    NasMountCredentials        NASMountCredentials `json:"nasMountCredentials,omitempty" form:"nasMountCredentials,omitempty"` //Specifies the credentials required to mount directories on the NetApp
    Password                   *string         `json:"password,omitempty" form:"password,omitempty"` //Specifies password of the username to access the target source.
    RefreshErrorMessage        *string         `json:"refreshErrorMessage,omitempty" form:"refreshErrorMessage,omitempty"` //Specifies a message if there was any error encountered during the last
    RefreshTimeUsecs           *int64          `json:"refreshTimeUsecs,omitempty" form:"refreshTimeUsecs,omitempty"` //Specifies the Unix epoch time (in microseconds) when the Protection
    RegistrationTimeUsecs      *int64          `json:"registrationTimeUsecs,omitempty" form:"registrationTimeUsecs,omitempty"` //Specifies the Unix epoch time (in microseconds) when the Protection
    ThrottlingPolicy           ThrottlingPolicy `json:"throttlingPolicy,omitempty" form:"throttlingPolicy,omitempty"` //Specifies the throttling policy for a registered Protection Source.
    ThrottlingPolicyOverrides  []*ThrottlingPolicyOverride `json:"throttlingPolicyOverrides,omitempty" form:"throttlingPolicyOverrides,omitempty"` //Array of Throttling Policy Overrides for Datastores.
    UseVmBiosUuid              *bool           `json:"useVmBiosUuid,omitempty" form:"useVmBiosUuid,omitempty"` //Specifies if registered vCenter is using BIOS UUID to track virtual
    Username                   *string         `json:"username,omitempty" form:"username,omitempty"` //Specifies username to access the target source.
    WarningMessages            *[]string       `json:"warningMessages,omitempty" form:"warningMessages,omitempty"` //Specifies a list of warnings encountered during registration.
}

/*
 * Structure for the custom type RemoteCluster
 */
type RemoteCluster struct {
    AllEndpointsReachable   *bool           `json:"allEndpointsReachable,omitempty" form:"allEndpointsReachable,omitempty"` //Specifies whether any endpoint (such as a Node) on the remote Cluster
    BandwidthLimit          BandwidthLimit  `json:"bandwidthLimit,omitempty" form:"bandwidthLimit,omitempty"` //Specifies settings for limiting the data transfer rate between
    ClearInterfaces         *bool           `json:"clearInterfaces,omitempty" form:"clearInterfaces,omitempty"` //TODO: Write general description for this field
    ClearVlanId             *bool           `json:"clearVlanId,omitempty" form:"clearVlanId,omitempty"` //Specifies whether to clear the vlanId field, and thus stop
    ClusterId               *int64          `json:"clusterId,omitempty" form:"clusterId,omitempty"` //Specifies the unique id of the remote Cluster.
    ClusterIncarnationId    *int64          `json:"clusterIncarnationId,omitempty" form:"clusterIncarnationId,omitempty"` //Specifies the unique incarnation id of the remote Cluster. This
    CompressionEnabled      *bool           `json:"compressionEnabled,omitempty" form:"compressionEnabled,omitempty"` //Specifies whether to compress the outbound data when
    EncryptionKey           *string         `json:"encryptionKey,omitempty" form:"encryptionKey,omitempty"` //Specifies the encryption key used for encrypting the replication data
    IfaceName               *string         `json:"ifaceName,omitempty" form:"ifaceName,omitempty"` //Specifies the interface name of the VLAN to use for communicating with
    LocalIps                *[]string       `json:"localIps,omitempty" form:"localIps,omitempty"` //Array of Local IP Addresses.
    Name                    *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the remote Cluster.
    NetworkInterfaceGroup   *string         `json:"networkInterfaceGroup,omitempty" form:"networkInterfaceGroup,omitempty"` //Specifies the group name of the network interfaces to
    NetworkInterfaceIds     *[]int64        `json:"networkInterfaceIds,omitempty" form:"networkInterfaceIds,omitempty"` //Array of Network Interface Ids.
    PurposeRemoteAccess     *bool           `json:"purposeRemoteAccess,omitempty" form:"purposeRemoteAccess,omitempty"` //Whether the remote cluster will be used for remote access for SPOG.
    PurposeReplication      *bool           `json:"purposeReplication,omitempty" form:"purposeReplication,omitempty"` //Whether the remote cluster will be used for replication.
    RemoteAccessCredentials CreateAccessTokenCredentialRequest `json:"remoteAccessCredentials,omitempty" form:"remoteAccessCredentials,omitempty"` //Specifies the Cohesity credentials required for generating an access token.
    RemoteIps               *[]string       `json:"remoteIps,omitempty" form:"remoteIps,omitempty"` //Array of Remote Node IP Addresses.
    TenantId                *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the tenant Id of the organization that created this remote
    UserName                *string         `json:"userName,omitempty" form:"userName,omitempty"` //Specifies the Cohesity user name used to connect to the
    ViewBoxPairInfo         []*StorageDomainViewBoxPairing `json:"viewBoxPairInfo,omitempty" form:"viewBoxPairInfo,omitempty"` //Array of Storage Domain (View Box) Pairs.
    VlanId                  *int64          `json:"vlanId,omitempty" form:"vlanId,omitempty"` //Specifies the id of the VLAN to use when communicating with the remote
}

/*
 * Structure for the custom type RemoteHost
 */
type RemoteHost struct {
    Address         *string         `json:"address,omitempty" form:"address,omitempty"` //Specifies the address (IP, hostname or FQDN) of the remote host
    Type            Type24Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the OS type of the remote host that will run the script.
}

/*
 * Structure for the custom type ContainsParametersToConnectToARemoteHost
 */
type ContainsParametersToConnectToARemoteHost struct {
    Credentials     Credentials     `json:"credentials,omitempty" form:"credentials,omitempty"` //Specifies credentials to access a target source.
    HostAddress     *string         `json:"hostAddress,omitempty" form:"hostAddress,omitempty"` //Address (i.e., IP, hostname or FQDN) of the host to connect to. Magneto
    HostType        *int64          `json:"hostType,omitempty" form:"hostType,omitempty"` //Type of host to connect to.
}

/*
 * Structure for the custom type RemoteAdapter
 */
type RemoteAdapter struct {
    FullBackupScript        RemoteScript    `json:"fullBackupScript,omitempty" form:"fullBackupScript,omitempty"` //Specifies the script that should run for the Full (no CBT) backup schedule
    IncrementalBackupScript RemoteScript    `json:"incrementalBackupScript,omitempty" form:"incrementalBackupScript,omitempty"` //Specifies the script that should run for the CBT-based backup
    LogBackupScript         RemoteScript    `json:"logBackupScript,omitempty" form:"logBackupScript,omitempty"` //Specifies the script that should run for the Log backup schedule
    RemoteHost              RemoteHost      `json:"remoteHost,omitempty" form:"remoteHost,omitempty"` //Specifies the remote host where the remote scripts are executed.
    Username                *string         `json:"username,omitempty" form:"username,omitempty"` //Specifies the username that will be used to login to the remote host.
}

/*
 * Structure for the custom type RemoteProtectionJobInformation
 */
type RemoteProtectionJobInformation struct {
    ClusterName     *string         `json:"clusterName,omitempty" form:"clusterName,omitempty"` //Specifies the name of the original Cluster that archived the data to the
    Environment     Environment13Enum `json:"environment,omitempty" form:"environment,omitempty"` //Specifies the environment type (such as kVMware or kSQL)
    JobName         *string         `json:"jobName,omitempty" form:"jobName,omitempty"` //Specifies the name of the Protection Job on the original Cluster.
    JobUid          UniqueGlobalId  `json:"jobUid,omitempty" form:"jobUid,omitempty"` //Specifies the globally unique id of the original Protection Job
}

/*
 * Structure for the custom type RemoteProtectionJobRunInformation
 */
type RemoteProtectionJobRunInformation struct {
    ClusterName       *string         `json:"clusterName,omitempty" form:"clusterName,omitempty"` //Specifies the name of the original Cluster that archived the data to the
    Environment       Environment13Enum `json:"environment,omitempty" form:"environment,omitempty"` //Specifies the environment type (such as kVMware or kSQL)
    JobName           *string         `json:"jobName,omitempty" form:"jobName,omitempty"` //Specifies the name of the Protection Job on the original Cluster.
    JobUid            UniqueGlobalId  `json:"jobUid,omitempty" form:"jobUid,omitempty"` //Specifies the globally unique id of the original Protection Job
    ProtectionJobRuns []*RemoteProtectionJobRunInstance `json:"protectionJobRuns,omitempty" form:"protectionJobRuns,omitempty"` //Array of Protection Job Run Details.
}

/*
 * Structure for the custom type RemoteProtectionJobRunInstance
 */
type RemoteProtectionJobRunInstance struct {
    ArchiveTaskUid    UniqueGlobalId  `json:"archiveTaskUid,omitempty" form:"archiveTaskUid,omitempty"` //Specifies the globally unique id of the archival task that archived
    ArchiveVersion    *int64          `json:"archiveVersion,omitempty" form:"archiveVersion,omitempty"` //Specifies the version of the archive.
    ExpiryTimeUsecs   *int64          `json:"expiryTimeUsecs,omitempty" form:"expiryTimeUsecs,omitempty"` //Specifies the time when the archive expires.
    IndexSizeBytes    *int64          `json:"indexSizeBytes,omitempty" form:"indexSizeBytes,omitempty"` //Specifies the size of the index for the archive.
    JobRunId          *int64          `json:"jobRunId,omitempty" form:"jobRunId,omitempty"` //Specifies the instance id of the Job Run task capturing the Snapshot.
    MetadataComplete  *bool           `json:"metadataComplete,omitempty" form:"metadataComplete,omitempty"` //Specifies whether a full set of metadata is available now.
    SnapshotTimeUsecs *int64          `json:"snapshotTimeUsecs,omitempty" form:"snapshotTimeUsecs,omitempty"` //Specify the time the Snapshot was captured as a Unix epoch Timestamp (in
}

/*
 * Structure for the custom type StatusOfIndexingTask
 */
type StatusOfIndexingTask struct {
    EndTimeUsecs               *int64          `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"` //Specifies the end time of the time range that is being indexed.
    Error                      *string         `json:"error,omitempty" form:"error,omitempty"` //Specifies the error message if the indexing Job/task fails.
    IndexingTaskEndTimeUsecs   *int64          `json:"indexingTaskEndTimeUsecs,omitempty" form:"indexingTaskEndTimeUsecs,omitempty"` //Specifies when the indexing task completed. This time is recorded
    IndexingTaskStartTimeUsecs *int64          `json:"indexingTaskStartTimeUsecs,omitempty" form:"indexingTaskStartTimeUsecs,omitempty"` //Specifies when the indexing task started. This time is recorded
    IndexingTaskStatus         IndexingTaskStatusEnum `json:"indexingTaskStatus,omitempty" form:"indexingTaskStatus,omitempty"` //Specifies the status of the indexing Job/task.
    IndexingTaskUid            UniqueGlobalId  `json:"indexingTaskUid,omitempty" form:"indexingTaskUid,omitempty"` //Specifies the unique id of the indexing task assigned by this Cluster.
    LatestExpiryTimeUsecs      *int64          `json:"latestExpiryTimeUsecs,omitempty" form:"latestExpiryTimeUsecs,omitempty"` //For all the Snapshots retrieved by this Job, specifies the latest time
    ProgressMonitorTask        *string         `json:"progressMonitorTask,omitempty" form:"progressMonitorTask,omitempty"` //Specifies the path to progress monitor task to track the progress
    StartTimeUsecs             *int64          `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"` //Specifies the start time of the time range that is being indexed.
}

/*
 * Structure for the custom type StatusOfRestoreSnapshotTask
 */
type StatusOfRestoreSnapshotTask struct {
    ArchiveTaskUid      UniqueGlobalId  `json:"archiveTaskUid,omitempty" form:"archiveTaskUid,omitempty"` //Specifies the globally unique id of the archival task that archived
    Error               *string         `json:"error,omitempty" form:"error,omitempty"` //Specifies the error message if the indexing task fails.
    ExpiryTimeUsecs     *int64          `json:"expiryTimeUsecs,omitempty" form:"expiryTimeUsecs,omitempty"` //Specifies the time when the Snapshot expires on the remote Vault.
    JobRunId            *int64          `json:"jobRunId,omitempty" form:"jobRunId,omitempty"` //Specifies the id of the Job Run that originally captured the Snapshot.
    ProgressMonitorTask *string         `json:"progressMonitorTask,omitempty" form:"progressMonitorTask,omitempty"` //Specifies the path to the progress monitor task that tracks the progress
    SnapshotTaskStatus  SnapshotTaskStatusEnum `json:"snapshotTaskStatus,omitempty" form:"snapshotTaskStatus,omitempty"` //Specifies the status of the indexing task.
    SnapshotTaskUid     UniqueGlobalId  `json:"snapshotTaskUid,omitempty" form:"snapshotTaskUid,omitempty"` //Specifies the globally unique id of the task capturing the Snapshot.
    SnapshotTimeUsecs   *int64          `json:"snapshotTimeUsecs,omitempty" form:"snapshotTimeUsecs,omitempty"` //Specify the time the Snapshot was captured.
}

/*
 * Structure for the custom type RemoteScript
 */
type RemoteScript struct {
    ContinueOnError *bool           `json:"continueOnError,omitempty" form:"continueOnError,omitempty"` //Specifies if the script needs to continue even if there is an occurence of
    IsActive        *bool           `json:"isActive,omitempty" form:"isActive,omitempty"` //Specifies if the script is active. If set to false, this script will not
    ScriptParams    *string         `json:"scriptParams,omitempty" form:"scriptParams,omitempty"` //Specifies the parameters and values to pass into the remote script.
    ScriptPath      *string         `json:"scriptPath,omitempty" form:"scriptPath,omitempty"` //Specifies the path to the script on the remote host.
    TimeoutSecs     *int64          `json:"timeoutSecs,omitempty" form:"timeoutSecs,omitempty"` //Specifies the timeout of the script in seconds. The script will be killed
}

/*
 * Structure for the custom type RemoteVaultRestoreTaskStatus
 */
type RemoteVaultRestoreTaskStatus struct {
    CurrentIndexingStatus          StatusOfIndexingTask `json:"currentIndexingStatus,omitempty" form:"currentIndexingStatus,omitempty"` //Specifies the status of an indexing task that builds an index from
    CurrentSnapshotStatus          StatusOfRestoreSnapshotTask `json:"currentSnapshotStatus,omitempty" form:"currentSnapshotStatus,omitempty"` //Specifies the status of the Snapshot restore task.
    LocalProtectionJobUid          UniqueGlobalId  `json:"localProtectionJobUid,omitempty" form:"localProtectionJobUid,omitempty"` //Specifies the globally unique id of the new inactive Protection Job
    ParentJobUid                   UniqueGlobalId  `json:"parentJobUid,omitempty" form:"parentJobUid,omitempty"` //Specifies the unique id of the parent Job/task that spawned the
    RemoteProtectionJobInformation RemoteProtectionJobInformation `json:"remoteProtectionJobInformation,omitempty" form:"remoteProtectionJobInformation,omitempty"` //Specifies details about the original Protection Job and its
    SearchJobUid                   UniqueGlobalId  `json:"searchJobUid,omitempty" form:"searchJobUid,omitempty"` //Specifies the unique id of the search Job that searched the remote Vault.
}

/*
 * Structure for the custom type RemoteVaultSearchInformation
 */
type RemoteVaultSearchInformation struct {
    ClusterCount    *int64          `json:"clusterCount,omitempty" form:"clusterCount,omitempty"` //Specifies number of Clusters that have archived to the remote Vault
    EndTimeUsecs    *int64          `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"` //Specifies the end time of the search as a Unix epoch
    Error           *string         `json:"error,omitempty" form:"error,omitempty"` //Specifies the error message reported when a search fails.
    JobCount        *int64          `json:"jobCount,omitempty" form:"jobCount,omitempty"` //Specifies number of Protection Jobs that have archived to the remote Vault
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the search Job.
    SearchJobStatus SearchJobStatusEnum `json:"searchJobStatus,omitempty" form:"searchJobStatus,omitempty"` //Specifies the status of the search.
    SearchJobUid    UniqueGlobalId  `json:"searchJobUid,omitempty" form:"searchJobUid,omitempty"` //Specifies the unique id assigned to the search Job by the Cluster.
    StartTimeUsecs  *int64          `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"` //Specifies the start time of the search as a Unix epoch
    VaultId         *int64          `json:"vaultId,omitempty" form:"vaultId,omitempty"` //Specifies the id of the remote Vault (External Target) that was searched.
    VaultName       *string         `json:"vaultName,omitempty" form:"vaultName,omitempty"` //Specifies the name of the remote Vault (External Target) that was searched.
}

/*
 * Structure for the custom type RemoteVaultSearchJobResult
 */
type RemoteVaultSearchJobResult struct {
    ClusterCount       *int64          `json:"clusterCount,omitempty" form:"clusterCount,omitempty"` //Specifies number of Clusters that have archived to the remote Vault
    ClusterMatchString *string         `json:"clusterMatchString,omitempty" form:"clusterMatchString,omitempty"` //Specifies the value of the clusterMatchSting if it was set in the
    Cookie             *string         `json:"cookie,omitempty" form:"cookie,omitempty"` //Specifies an opaque string to pass to the next request to get the
    EndTimeUsecs       *int64          `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"` //Specifies the value of endTimeUsecs if it was set in the original
    Error              *string         `json:"error,omitempty" form:"error,omitempty"` //Specifies the error message if the search fails.
    JobCount           *int64          `json:"jobCount,omitempty" form:"jobCount,omitempty"` //Specifies number of Protection Jobs that have archived to the remote Vault
    JobMatchString     *string         `json:"jobMatchString,omitempty" form:"jobMatchString,omitempty"` //Specifies the value of the jobMatchSting if it was set in the
    ProtectionJobs     []*RemoteProtectionJobRunInformation `json:"protectionJobs,omitempty" form:"protectionJobs,omitempty"` //Array of Protection Jobs.
    SearchJobStatus    SearchJobStatus1Enum `json:"searchJobStatus,omitempty" form:"searchJobStatus,omitempty"` //Specifies the status of the search Job.
    SearchJobUid       UniqueGlobalId  `json:"searchJobUid,omitempty" form:"searchJobUid,omitempty"` //Specifies the unique id of the search Job assigned by the Cluster.
    StartTimeUsecs     *int64          `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"` //Specifies the value of startTimeUsecs if it was set in the original
    VaultId            *int64          `json:"vaultId,omitempty" form:"vaultId,omitempty"` //Specifies the id of the remote Vault that was searched.
    VaultName          *string         `json:"vaultName,omitempty" form:"vaultName,omitempty"` //Specifies the name of the remote Vault that was searched.
}

/*
 * Structure for the custom type ReplicationTarget
 */
type ReplicationTarget struct {
    ClusterId       *int64          `json:"clusterId,omitempty" form:"clusterId,omitempty"` //The id of the remote cluster.
    ClusterName     *string         `json:"clusterName,omitempty" form:"clusterName,omitempty"` //The name of the remote cluster.
}

/*
 * Structure for the custom type ReplicationTarget1
 */
type ReplicationTarget1 struct {
    ClusterId       *int64          `json:"clusterId,omitempty" form:"clusterId,omitempty"` //Specifies the id of the Remote Cluster.
    ClusterName     *string         `json:"clusterName,omitempty" form:"clusterName,omitempty"` //Specifies the name of the Remote Cluster.
}

/*
 * Structure for the custom type Error
 */
type Error struct {
    ErrorCode       *int64          `json:"errorCode,omitempty" form:"errorCode,omitempty"` //Operation response error code.
    Message         *string         `json:"message,omitempty" form:"message,omitempty"` //Description of the error.
}

/*
 * Structure for the custom type ResetS3SecretAccessKey
 */
type ResetS3SecretAccessKey struct {
    Domain          *string         `json:"domain,omitempty" form:"domain,omitempty"` //Specifies the fully qualified domain name (FQDN) of an Active Directory
    TenantId        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the tenant for which the the users are to be deleted.
    Username        *string         `json:"username,omitempty" form:"username,omitempty"` //Specifies the Cohesity user.
}

/*
 * Structure for the custom type RestoreCountByObjectType
 */
type RestoreCountByObjectType struct {
    ObjectCount     *int64          `json:"objectCount,omitempty" form:"objectCount,omitempty"` //Specifies the number of restores of the object type.
    ObjectType      *string         `json:"objectType,omitempty" form:"objectType,omitempty"` //Specifies the type of the restored object.
}

/*
 * Structure for the custom type RestoreTask
 */
type RestoreTask struct {
    ContinueOnError          *bool           `json:"continueOnError,omitempty" form:"continueOnError,omitempty"` //Specifies if the Restore Task should continue even if the copy operation
    Filenames                *[]string       `json:"filenames,omitempty" form:"filenames,omitempty"` //Array of Files or Folders.
    IsFileBasedVolumeRestore *bool           `json:"isFileBasedVolumeRestore,omitempty" form:"isFileBasedVolumeRestore,omitempty"` //Specifies whether this is a file based volume restore.
    Name                     *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the Restore Task. This field must be set and
    NewBaseDirectory         *string         `json:"newBaseDirectory,omitempty" form:"newBaseDirectory,omitempty"` //Specifies an optional root folder where to recover the selected
    Overwrite                *bool           `json:"overwrite,omitempty" form:"overwrite,omitempty"` //If true, any existing files and folders on the operating system
    Password                 *string         `json:"password,omitempty" form:"password,omitempty"` //Specifies password of the username to access the target source.
    PreserveAttributes       *bool           `json:"preserveAttributes,omitempty" form:"preserveAttributes,omitempty"` //If true, the Restore Tasks preserves the original file and
    SourceObjectInfo         RestoreObject   `json:"sourceObjectInfo,omitempty" form:"sourceObjectInfo,omitempty"` //Specifies information about the source object (such as a VM)
    TargetHostType           TargetHostTypeEnum `json:"targetHostType,omitempty" form:"targetHostType,omitempty"` //Specifies the target host types to be restored.
    TargetParentSourceId     *int64          `json:"targetParentSourceId,omitempty" form:"targetParentSourceId,omitempty"` //Specifies the registered source (such as a vCenter Server)
    TargetSourceId           *int64          `json:"targetSourceId,omitempty" form:"targetSourceId,omitempty"` //Specifies the id of the target protection source (such as a VM)
    Username                 *string         `json:"username,omitempty" form:"username,omitempty"` //Specifies username to access the target source.
}

/*
 * Structure for the custom type FullSnapshotInformation
 */
type FullSnapshotInformation struct {
    ArchivalTarget          ArchivalTarget  `json:"archivalTarget,omitempty" form:"archivalTarget,omitempty"` //Specifies settings about the Archival External Target (such as Tape or AWS).
    AttemptNumber           *int64          `json:"attemptNumber,omitempty" form:"attemptNumber,omitempty"` //Specifies the attempt number.
    CloudDeployTarget       CloudDeployTargetDetails `json:"cloudDeployTarget,omitempty" form:"cloudDeployTarget,omitempty"` //Message that specifies the details about CloudDeploy target where backup
    JobRunId                *int64          `json:"jobRunId,omitempty" form:"jobRunId,omitempty"` //Specifies the id of the job run.
    JobUid                  UniqueGlobalId  `json:"jobUid,omitempty" form:"jobUid,omitempty"` //Specifies an id for an object that is unique across Cohesity Clusters.
    ParentSource            ProtectionSource `json:"parentSource,omitempty" form:"parentSource,omitempty"` //Specifies a generic structure that represents a node
    SnapshotRelativeDirPath *string         `json:"snapshotRelativeDirPath,omitempty" form:"snapshotRelativeDirPath,omitempty"` //Specifies the relative path of the snapshot directory.
    Source                  ProtectionSource `json:"source,omitempty" form:"source,omitempty"` //Specifies a generic structure that represents a node
    StartTimeUsecs          *int64          `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"` //Specifies the start time specified as a Unix epoch Timestamp
    ViewName                *string         `json:"viewName,omitempty" form:"viewName,omitempty"` //Specifies the name of the view.
    VmHadIndependentDisks   *bool           `json:"vmHadIndependentDisks,omitempty" form:"vmHadIndependentDisks,omitempty"` //Specifies if the VM had independent disks.
}

/*
 * Structure for the custom type RestoreObject
 */
type RestoreObject struct {
    ArchivalTarget     ArchivalTarget  `json:"archivalTarget,omitempty" form:"archivalTarget,omitempty"` //Specifies settings about the Archival Target (such as Tape or AWS).
    CloudDeployTarget  CloudDeployTargetDetails `json:"cloudDeployTarget,omitempty" form:"cloudDeployTarget,omitempty"` //Specifies settings about the Cloud Deploy target.
    Environment        Environment5Enum `json:"environment,omitempty" form:"environment,omitempty"` //Specifies the type of the Protection Source.
    JobId              *int64          `json:"jobId,omitempty" form:"jobId,omitempty"` //Protection Job Id.
    JobRunId           *int64          `json:"jobRunId,omitempty" form:"jobRunId,omitempty"` //Specifies the id of the Job Run that captured the snapshot.
    JobUid             UniqueGlobalId  `json:"jobUid,omitempty" form:"jobUid,omitempty"` //Specifies the universal id of the Protection Job that backed up
    ProtectionSourceId *int64          `json:"protectionSourceId,omitempty" form:"protectionSourceId,omitempty"` //Specifies the id of the leaf object to recover, clone or recover
    SourceName         *string         `json:"sourceName,omitempty" form:"sourceName,omitempty"` //Specifies the name of the Protection Source.
    StartedTimeUsecs   *int64          `json:"startedTimeUsecs,omitempty" form:"startedTimeUsecs,omitempty"` //Specifies the time when the Job Run starts capturing a snapshot.
}

/*
 * Structure for the custom type RestoreObjectState
 */
type RestoreObjectState struct {
    Error            Error           `json:"error,omitempty" form:"error,omitempty"` //Details about the Error.
    ObjectStatus     ObjectStatusEnum `json:"objectStatus,omitempty" form:"objectStatus,omitempty"` //Specifies the status of an object during a Restore Task.
    ResourcePoolId   *int64          `json:"resourcePoolId,omitempty" form:"resourcePoolId,omitempty"` //Specifies the id of the Resource Pool that the restored
    RestoredObjectId *int64          `json:"restoredObjectId,omitempty" form:"restoredObjectId,omitempty"` //Specifies the Id of the recovered or cloned object.
    SourceObjectId   *int64          `json:"sourceObjectId,omitempty" form:"sourceObjectId,omitempty"` //Specifies the Protection Source id of the object to be recovered or
}

/*
 * Structure for the custom type RestorePointsForTimeRange
 */
type RestorePointsForTimeRange struct {
    FullSnapshotInfo []*FullSnapshotInfo `json:"fullSnapshotInfo,omitempty" form:"fullSnapshotInfo,omitempty"` //Specifies the info related to the recovery object.
    TimeRanges       []*TimeRange    `json:"timeRanges,omitempty" form:"timeRanges,omitempty"` //Specifies the time ranges of the restore object between full snapshots.
}

/*
 * Structure for the custom type RestorePointsForTimeRangeParameters
 */
type RestorePointsForTimeRangeParameters struct {
    EndTimeUsecs       *int64          `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"` //Specifies the end time specified as a Unix epoch Timestamp
    Environment        Environment15Enum `json:"environment,omitempty" form:"environment,omitempty"` //Specifies the protection source environment type.
    JobUids            []*UniqueGlobalId `json:"jobUids" form:"jobUids"` //Specifies the jobs for which to get the full snapshot information.
    ProtectionSourceId *int64          `json:"protectionSourceId,omitempty" form:"protectionSourceId,omitempty"` //Specifies the id of the Protection Source which is to be restored.
    StartTimeUsecs     *int64          `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"` //Specifies the start time specified as a Unix epoch Timestamp
}

/*
 * Structure for the custom type RestoreSourceSummaryByObjectTypeElement
 */
type RestoreSourceSummaryByObjectTypeElement struct {
    DatastoreId          *int64          `json:"datastoreId,omitempty" form:"datastoreId,omitempty"` //Specifies the datastore where the object's files are recovered to.
    FileRestoreInfo      []*FileRestoreInformation `json:"fileRestoreInfo,omitempty" form:"fileRestoreInfo,omitempty"` //Specifies a list of restore information of files.
    Name                 string          `json:"name" form:"name"` //Specifies the name of the Restore Task. This field must be set and
    Objects              []*RestoreObject `json:"objects,omitempty" form:"objects,omitempty"` //Array of Objects.
    ProtectionSourceName *string         `json:"protectionSourceName,omitempty" form:"protectionSourceName,omitempty"` //The protection source name.
    StartTimeUsecs       *int64          `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"` //Specifies the start time of the Restore Task as a Unix epoch
    Type                 Type31Enum      `json:"type,omitempty" form:"type,omitempty"` //Specify the object type to filter with.
    Username             *string         `json:"username,omitempty" form:"username,omitempty"` //Specifies the Cohesity user who requested this Restore Task.
}

/*
 * Structure for the custom type RestoreTask1
 */
type RestoreTask1 struct {
    AcropolisParameters     AcropolisRestoreParameters `json:"acropolisParameters,omitempty" form:"acropolisParameters,omitempty"` //This field defines the Acropolis specific params for restore tasks of type
    ArchiveTaskUid          UniqueGlobalId  `json:"archiveTaskUid,omitempty" form:"archiveTaskUid,omitempty"` //Specifies the uid of the Restore Task that retrieves objects from
    CloneViewParameters     View1           `json:"cloneViewParameters,omitempty" form:"cloneViewParameters,omitempty"` //Specifies the View settings used when cloning a View.
    ContinueOnError         *bool           `json:"continueOnError,omitempty" form:"continueOnError,omitempty"` //Specifies if the Restore Task should continue when some operations on some
    DatastoreId             *int64          `json:"datastoreId,omitempty" form:"datastoreId,omitempty"` //Specifies the datastore where the object's files are recovered to.
    EndTimeUsecs            *int64          `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"` //Specifies the end time of the Restore Task as a Unix epoch
    Error                   Error           `json:"error,omitempty" form:"error,omitempty"` //Specifies the error reported by the Restore Task (if any) after the
    FullViewName            *string         `json:"fullViewName,omitempty" form:"fullViewName,omitempty"` //Specifies the full name of a View.
    HypervParameters        HypervRestoreParameters `json:"hypervParameters,omitempty" form:"hypervParameters,omitempty"` //Specifies information needed when restoring VMs in HyperV enviroment.
    Id                      *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies the id of the Restore Task assigned by
    MountVolumesState       MountVolumesStates `json:"mountVolumesState,omitempty" form:"mountVolumesState,omitempty"` //Specifies the states of mounting all the volumes onto a mount target
    Name                    string          `json:"name" form:"name"` //Specifies the name of the Restore Task. This field must be set and
    NewParentId             *int64          `json:"newParentId,omitempty" form:"newParentId,omitempty"` //Specify a new registered parent Protection Source. If specified
    Objects                 []*RestoreObject `json:"objects,omitempty" form:"objects,omitempty"` //Array of Objects.
    OutlookParameters       OutlookRestoreParameters `json:"outlookParameters,omitempty" form:"outlookParameters,omitempty"` //Specifies information needed for recovering Mailboxes in O365Outlook
    RestoreObjectState      []*RestoreObjectState `json:"restoreObjectState,omitempty" form:"restoreObjectState,omitempty"` //Array of Object States.
    StartTimeUsecs          *int64          `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"` //Specifies the start time for the Restore Task as a Unix epoch
    Status                  Status6Enum     `json:"status,omitempty" form:"status,omitempty"` //Specifies the overall status of the Restore Task.
    TargetViewCreated       *bool           `json:"targetViewCreated,omitempty" form:"targetViewCreated,omitempty"` //Is true if a new View was created by a 'kCloneVMs' Restore Task.
    Type                    Type32Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of Restore Task.
    Username                *string         `json:"username,omitempty" form:"username,omitempty"` //Specifies the Cohesity user who requested this Restore Task.
    ViewBoxId               *int64          `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"` //Specifies the id of the Domain (View Box) where the View is stored.
    VirtualDiskRestoreState VirtualDiskRecoverState `json:"virtualDiskRestoreState,omitempty" form:"virtualDiskRestoreState,omitempty"` //Specifies the complete information about a recover virtual disk task state.
    VlanParameters          VLANParameters  `json:"vlanParameters,omitempty" form:"vlanParameters,omitempty"` //Specifies VLAN parameters for the restore operation.
    VmwareParameters        VmwareRestoreParameters `json:"vmwareParameters,omitempty" form:"vmwareParameters,omitempty"` //Specifies the information required for recovering or cloning VmWare VMs.
}

/*
 * Structure for the custom type RetentionPolicyProto
 */
type RetentionPolicyProto struct {
    NumDaysToKeep   *int64          `json:"numDaysToKeep,omitempty" form:"numDaysToKeep,omitempty"` //The number of days to keep the snapshots for a backup run.
    WormRetention   WormRetentionProto `json:"wormRetention,omitempty" form:"wormRetention,omitempty"` //Message that specifies the WORM attributes. WORM attributes can be
}

/*
 * Structure for the custom type RoleInformation
 */
type RoleInformation struct {
    CreatedTimeMsecs     *int64          `json:"createdTimeMsecs,omitempty" form:"createdTimeMsecs,omitempty"` //Specifies the epoch time in milliseconds when the role was created.
    Description          *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description about the role.
    IsCustomRole         *bool           `json:"isCustomRole,omitempty" form:"isCustomRole,omitempty"` //Specifies if the role is a user-defined custom role.
    Label                *string         `json:"label,omitempty" form:"label,omitempty"` //Specifies the label for the role as displayed on the Cohesity
    LastUpdatedTimeMsecs *int64          `json:"lastUpdatedTimeMsecs,omitempty" form:"lastUpdatedTimeMsecs,omitempty"` //Specifies the epoch time in milliseconds when the role was last modified.
    Name                 *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the internal Cluster name for the role such as COHESITY_VIEWER.
    Privileges           *[]string       `json:"privileges,omitempty" form:"privileges,omitempty"` //Array of Privileges.
    TenantId             *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies unique id of the tenant owning the role.
    TenantIds            *[]string       `json:"tenantIds,omitempty" form:"tenantIds,omitempty"` //Specifies id of tenants using this role.
}

/*
 * Structure for the custom type RoleCreate
 */
type RoleCreate struct {
    Description     *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description about the role.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the custom role.
    Privileges      *[]string       `json:"privileges,omitempty" form:"privileges,omitempty"` //Array of Privileges.
}

/*
 * Structure for the custom type RoleUpdate
 */
type RoleUpdate struct {
    Description     *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description about the role.
    Privileges      *[]string       `json:"privileges,omitempty" form:"privileges,omitempty"` //Array of Privileges.
}

/*
 * Structure for the custom type Route
 */
type Route struct {
    Description     *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description of the Static Route.
    DestNetwork     *string         `json:"destNetwork,omitempty" form:"destNetwork,omitempty"` //Destination network.
    IfName          *string         `json:"ifName,omitempty" form:"ifName,omitempty"` //Specifies the network interfaces name to use for communicating with the
    IfaceGroupName  *string         `json:"ifaceGroupName,omitempty" form:"ifaceGroupName,omitempty"` //Specifies the network interfaces group or interface group with vlan id to
    NextHop         *string         `json:"nextHop,omitempty" form:"nextHop,omitempty"` //Specifies the next hop to the destination network.
}

/*
 * Structure for the custom type RPOPolicy
 */
type RPOPolicy struct {
    AlertingConfig           AlertingConfig  `json:"alertingConfig,omitempty" form:"alertingConfig,omitempty"` //Specifies optional settings for alerting.
    AlertingPolicy           *[]AlertingPolicyEnum `json:"alertingPolicy,omitempty" form:"alertingPolicy,omitempty"` //Array of Job Events.
    EnvironmentTypeJobParams EnvironmentSpecificCommonJobParameters `json:"environmentTypeJobParams,omitempty" form:"environmentTypeJobParams,omitempty"` //Specifies additional parameters that are common to all Protection
    IndexingPolicy           IndexingPolicy  `json:"indexingPolicy,omitempty" form:"indexingPolicy,omitempty"` //Specifies settings for indexing files found in an Object
    QosType                  QosType1Enum    `json:"qosType,omitempty" form:"qosType,omitempty"` //Specifies the QoS policy type to use.
    StorageDomainId          *int64          `json:"storageDomainId,omitempty" form:"storageDomainId,omitempty"` //Specifies the Storage Domain to which data will be written.
}

/*
 * Structure for the custom type CopyTaskJobRunParameters
 */
type CopyTaskJobRunParameters struct {
    ArchivalTarget      ArchivalTarget  `json:"archivalTarget,omitempty" form:"archivalTarget,omitempty"` //Specifies settings about the Archival External Target (such as Tape or AWS).
    DaysToKeep          *int64          `json:"daysToKeep,omitempty" form:"daysToKeep,omitempty"` //Specifies the number of days to retain copied Snapshots on the target.
    HoldForLegalPurpose *bool           `json:"holdForLegalPurpose,omitempty" form:"holdForLegalPurpose,omitempty"` //Specifies optionally whether to retain the snapshot for legal purpose.
    ReplicationTarget   ReplicationTarget1 `json:"replicationTarget,omitempty" form:"replicationTarget,omitempty"` //Specifies settings about the Remote Cohesity Cluster where Snapshots
    Type                Type26Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of a Snapshot target such as 'kLocal', 'kRemote' or
}

/*
 * Structure for the custom type ResultOfRunMapReduceInstanceCall
 */
type ResultOfRunMapReduceInstanceCall struct {
    Error               ErrorProto      `json:"error,omitempty" form:"error,omitempty"` //TODO: Write general description for this field
    MapReduceInstanceId *int64          `json:"mapReduceInstanceId,omitempty" form:"mapReduceInstanceId,omitempty"` //Return the ID of instance.
}

/*
 * Structure for the custom type GetRunMapReduceAppParameters
 */
type GetRunMapReduceAppParameters struct {
    AppId           *int64          `json:"appId,omitempty" form:"appId,omitempty"` //ApplicationId is the Id of the map reduce application to run.
    InputParams     []*MapReduceInstanceInputParam `json:"inputParams,omitempty" form:"inputParams,omitempty"` //InputParams specifies optional list of key=value input params specified
    MrInput         InputSelectorSelectsTheFilesToMapOver `json:"mrInput,omitempty" form:"mrInput,omitempty"` //TODO: Write general description for this field
    MrOutput        OutputSpecificationForTheMapreduce `json:"mrOutput,omitempty" form:"mrOutput,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type RunNowParameters
 */
type RunNowParameters struct {
    DatabaseIds     *[]int64        `json:"databaseIds,omitempty" form:"databaseIds,omitempty"` //Specifies the ids of the DB's to perform run now on.
    SourceId        *int64          `json:"sourceId,omitempty" form:"sourceId,omitempty"` //Specifies the source id of the Databases to perform the Run Now
}

/*
 * Structure for the custom type ProtectionRunParameters
 */
type ProtectionRunParameters struct {
    CopyRunTargets   []*CopyTaskJobRunParameters `json:"copyRunTargets,omitempty" form:"copyRunTargets,omitempty"` //Optional parameter to be set if you want specific replication or archival
    RunNowParameters []*RunNowParameters `json:"runNowParameters,omitempty" form:"runNowParameters,omitempty"` //Optional parameters of a Run Now operation.
    RunType          RunType2Enum    `json:"runType,omitempty" form:"runType,omitempty"` //Specifies the type of backup. If not specified, 'kRegular' is assumed.
    SourceIds        *[]int64        `json:"sourceIds,omitempty" form:"sourceIds,omitempty"` //Optional parameter if you want to back up only a subset of sources that
}

/*
 * Structure for the custom type RunUID
 */
type RunUID struct {
    JobUid          UniqueGlobalId  `json:"jobUid,omitempty" form:"jobUid,omitempty"` //Specifies an id for an object that is unique across Cohesity Clusters.
    StartTimeUsecs  *int64          `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"` //Specifies the start time of the Protection Job Run.
}

/*
 * Structure for the custom type SQLServerInstanceVersion
 */
type SQLServerInstanceVersion struct {
    Build           *int64          `json:"build,omitempty" form:"build,omitempty"` //Specfies the build.
    MajorVersion    *int64          `json:"majorVersion,omitempty" form:"majorVersion,omitempty"` //Specfies the major version.
    MinorVersion    *int64          `json:"minorVersion,omitempty" form:"minorVersion,omitempty"` //Specfies the minor version.
    Revision        *int64          `json:"revision,omitempty" form:"revision,omitempty"` //Specfies the revision.
    VersionString   *string         `json:"versionString,omitempty" form:"versionString,omitempty"` //Specfies the version string.
}

/*
 * Structure for the custom type SalesforceAccountInformation
 */
type SalesforceAccountInformation struct {
    AccountId       *string         `json:"accountId,omitempty" form:"accountId,omitempty"` //Specifies the Account Id assigned by Salesforce.
    UserId          *string         `json:"userId,omitempty" form:"userId,omitempty"` //Specifies the User Id assigned by Salesforce.
}

/*
 * Structure for the custom type Sample
 */
type Sample struct {
    FloatValue      *float64        `json:"floatValue,omitempty" form:"floatValue,omitempty"` //Specifies the value of the data sample if the type is float64.
    TimestampMsecs  *int64          `json:"timestampMsecs,omitempty" form:"timestampMsecs,omitempty"` //Specifies the timestamp when the data sample occured.
    Value           *int64          `json:"value,omitempty" form:"value,omitempty"` //Specifies the value of the data sample if the type is int64.
}

/*
 * Structure for the custom type SchedulingPolicy
 */
type SchedulingPolicy struct {
    ContinuousSchedule ContinuousSchedule `json:"continuousSchedule,omitempty" form:"continuousSchedule,omitempty"` //Specifies the time interval between two Job Runs of a continuous backup
    DailySchedule      DailyWeeklySchedule `json:"dailySchedule,omitempty" form:"dailySchedule,omitempty"` //Specifies a daily or weekly backup schedule.
    MonthlySchedule    MonthlySchedule `json:"monthlySchedule,omitempty" form:"monthlySchedule,omitempty"` //Specifies a monthly backup schedule.
    Periodicity        Periodicity2Enum `json:"periodicity,omitempty" form:"periodicity,omitempty"` //Specifies how often to start new Job Runs of a Protection Job.
    RpoSchedule        RPOSchedule     `json:"rpoSchedule,omitempty" form:"rpoSchedule,omitempty"` //Specifies an RPO policy schedule.
}

/*
 * Structure for the custom type ScriptPathAndParams
 */
type ScriptPathAndParams struct {
    ContinueOnError *bool           `json:"continueOnError,omitempty" form:"continueOnError,omitempty"` //Applicable only for pre backup scripts. If this flag is set to true, then
    IsActive        *bool           `json:"isActive,omitempty" form:"isActive,omitempty"` //Indicates if the script is active. If 'is_active' is set to false, this
    ScriptParams    *string         `json:"scriptParams,omitempty" form:"scriptParams,omitempty"` //Custom parameters that users want to pass to the script. For example,
    ScriptPath      *string         `json:"scriptPath,omitempty" form:"scriptPath,omitempty"` //For backup jobs of type 'kPuppeteer', 'script_path' is full path of
    TimeoutSecs     *int64          `json:"timeoutSecs,omitempty" form:"timeoutSecs,omitempty"` //Timeout of the script. The script will be killed if it exceeds this value.
}

/*
 * Structure for the custom type Share
 */
type Share struct {
    AllSmbMountPaths *[]string       `json:"allSmbMountPaths,omitempty" form:"allSmbMountPaths,omitempty"` //Array of SMB Paths.
    NfsMountPath     *string         `json:"nfsMountPath,omitempty" form:"nfsMountPath,omitempty"` //Specifies the path for mounting this Share as an NFS share.
    Path             *string         `json:"path,omitempty" form:"path,omitempty"` //Specifies the path information for this share.
    S3AccessPath     *string         `json:"s3AccessPath,omitempty" form:"s3AccessPath,omitempty"` //Specifies the path to access this View as an S3 share.
    ShareName        *string         `json:"shareName,omitempty" form:"shareName,omitempty"` //The name of the share.
    SmbMountPath     *string         `json:"smbMountPath,omitempty" form:"smbMountPath,omitempty"` //Specifies the main path for mounting this Share as an SMB share.
    TenantId         *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
    ViewName         *string         `json:"viewName,omitempty" form:"viewName,omitempty"` //Specifies the view name this share belongs to.
}

/*
 * Structure for the custom type SMBActiveFileOpenResponse
 */
type SMBActiveFileOpenResponse struct {
    ActiveFilePaths []*ActiveSMBFilePath `json:"activeFilePaths,omitempty" form:"activeFilePaths,omitempty"` //Specifies the active opens for an SMB file in a view.
    Cookie          *string         `json:"cookie,omitempty" form:"cookie,omitempty"` //Specifies an opaque string to pass to get the next set of active opens.
}

/*
 * Structure for the custom type ActiveSMBFilePath
 */
type ActiveSMBFilePath struct {
    ActiveSessions  []*SMBActiveSession `json:"activeSessions,omitempty" form:"activeSessions,omitempty"` //Specifies the sessions where the file is open.
    FilePath        *string         `json:"filePath,omitempty" form:"filePath,omitempty"` //Specifies the filepath in the view.
    ViewId          *int64          `json:"viewId,omitempty" form:"viewId,omitempty"` //Specifies the id of the View assigned by the Cohesity Cluster.
    ViewName        *string         `json:"viewName,omitempty" form:"viewName,omitempty"` //Specifies the name of the View.
}

/*
 * Structure for the custom type SMBActiveOpenFile
 */
type SMBActiveOpenFile struct {
    AccessInfoList  *[]AccessInfoListEnum `json:"accessInfoList,omitempty" form:"accessInfoList,omitempty"` //Specifies the access information.
    OpenId          *int64          `json:"openId,omitempty" form:"openId,omitempty"` //Specifies the id of the active open.
    OthersCanDelete *bool           `json:"othersCanDelete,omitempty" form:"othersCanDelete,omitempty"` //Specifies whether others are allowed to delete.
    OthersCanRead   *bool           `json:"othersCanRead,omitempty" form:"othersCanRead,omitempty"` //Specifies whether others are allowed to read.
    OthersCanWrite  *bool           `json:"othersCanWrite,omitempty" form:"othersCanWrite,omitempty"` //Specifies whether others are allowed to write.
}

/*
 * Structure for the custom type SMBActiveSession
 */
type SMBActiveSession struct {
    ActiveOpens     []*SMBActiveOpenFile `json:"activeOpens,omitempty" form:"activeOpens,omitempty"` //Specifies the list of active opens of the file in this session.
    ClientIp        *string         `json:"clientIp,omitempty" form:"clientIp,omitempty"` //Specifies the IP address from which the file is still open.
    Domain          *string         `json:"domain,omitempty" form:"domain,omitempty"` //Specifies the domain of the user.
    ServerIp        *string         `json:"serverIp,omitempty" form:"serverIp,omitempty"` //Specifies the IP address of the server where the file exists.
    SessionId       *int64          `json:"sessionId,omitempty" form:"sessionId,omitempty"` //Specifies the id of the session.
    Username        *string         `json:"username,omitempty" form:"username,omitempty"` //Specifies the username who keeps the file open.
}

/*
 * Structure for the custom type SMBPermission
 */
type SMBPermission struct {
    Access            AccessEnum      `json:"access,omitempty" form:"access,omitempty"` //Specifies the read/write access to the SMB share.
    Mode              Mode1Enum       `json:"mode,omitempty" form:"mode,omitempty"` //Specifies how the permission should be applied to folders and/or files.
    Sid               *string         `json:"sid,omitempty" form:"sid,omitempty"` //Specifies the security identifier (SID) of the principal.
    SpecialAccessMask *int64          `json:"specialAccessMask,omitempty" form:"specialAccessMask,omitempty"` //Specifies custom access permissions.
    SpecialType       *int64          `json:"specialType,omitempty" form:"specialType,omitempty"` //Specifies a custom type.
    Type              Type22Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of permission.
}

/*
 * Structure for the custom type SMBPermissionsInformation
 */
type SMBPermissionsInformation struct {
    OwnerSid        *string         `json:"ownerSid,omitempty" form:"ownerSid,omitempty"` //Specifies the security identifier (SID) of the owner of the SMB share.
    Permissions     []*SMBPermission `json:"permissions,omitempty" form:"permissions,omitempty"` //Array of SMB Permissions.
}

/*
 * Structure for the custom type SpecifiesStructWithSMBPrincipalDetails
 */
type SpecifiesStructWithSMBPrincipalDetails struct {
    Domain          *string         `json:"domain,omitempty" form:"domain,omitempty"` //Specifies domain name of the principal.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies name of the SMB principal which may be a group or user.
    Sid             *string         `json:"sid,omitempty" form:"sid,omitempty"` //Specifies unique Security ID (SID) of the principal that look similar to
    Type            *string         `json:"type,omitempty" form:"type,omitempty"` //Specifies the type. This can be a user or a group.
}

/*
 * Structure for the custom type SnapshotCopyArchivalPolicy
 */
type SnapshotCopyArchivalPolicy struct {
    CopyPartial     *bool           `json:"copyPartial,omitempty" form:"copyPartial,omitempty"` //Specifies if Snapshots are copied from the first completely successful
    DaysToKeep      *int64          `json:"daysToKeep,omitempty" form:"daysToKeep,omitempty"` //Specifies the number of days to retain copied Snapshots on the target.
    Multiplier      *int64          `json:"multiplier,omitempty" form:"multiplier,omitempty"` //Specifies a factor to multiply the periodicity by, to determine the copy
    Periodicity     PeriodicityEnum `json:"periodicity,omitempty" form:"periodicity,omitempty"` //Specifies the frequency that Snapshots should be copied to the
    Target          ArchivalTarget  `json:"target,omitempty" form:"target,omitempty"` //Specifies the archival target to copy the Snapshots to.
}

/*
 * Structure for the custom type SnapshotAttempt
 */
type SnapshotAttempt struct {
    AttemptNumber    *int64          `json:"attemptNumber,omitempty" form:"attemptNumber,omitempty"` //Specifies the number of the attempts made by the Job Run
    JobRunId         *int64          `json:"jobRunId,omitempty" form:"jobRunId,omitempty"` //Specifies the id of the Job Run that captured the snapshot.
    StartedTimeUsecs *int64          `json:"startedTimeUsecs,omitempty" form:"startedTimeUsecs,omitempty"` //Specifies the time when the Job Run starts capturing a snapshot.
}

/*
 * Structure for the custom type CloudDeployPolicy
 */
type CloudDeployPolicy struct {
    CopyPartial     *bool           `json:"copyPartial,omitempty" form:"copyPartial,omitempty"` //Specifies if Snapshots are copied from the first completely successful
    DaysToKeep      *int64          `json:"daysToKeep,omitempty" form:"daysToKeep,omitempty"` //Specifies the number of days to retain copied Snapshots on the target.
    Multiplier      *int64          `json:"multiplier,omitempty" form:"multiplier,omitempty"` //Specifies a factor to multiply the periodicity by, to determine the copy
    Periodicity     PeriodicityEnum `json:"periodicity,omitempty" form:"periodicity,omitempty"` //Specifies the frequency that Snapshots should be copied to the
    Target          CloudDeployTargetDetails `json:"target,omitempty" form:"target,omitempty"` //Message that specifies the details about CloudDeploy target where backup
}

/*
 * Structure for the custom type SnapshotCopyTask
 */
type SnapshotCopyTask struct {
    CopyStatus      *string         `json:"copyStatus,omitempty" form:"copyStatus,omitempty"` //Specifies the status of the copy task.
    ExpiryTimeUsecs *int64          `json:"expiryTimeUsecs,omitempty" form:"expiryTimeUsecs,omitempty"` //Specifies when the Snapshot expires on the target.
    Message         *string         `json:"message,omitempty" form:"message,omitempty"` //Specifies warning or error information when the copy task is not
    SnapshotTarget  SnapshotTarget1 `json:"snapshotTarget,omitempty" form:"snapshotTarget,omitempty"` //Specifies settings about a target where a copied Snapshot is stored.
}

/*
 * Structure for the custom type SnapshotInformation
 */
type SnapshotInformation struct {
    Environment               Environment7Enum `json:"environment,omitempty" form:"environment,omitempty"` //Specifies the environment type (such as kVMware or kSQL) that
    RelativeSnapshotDirectory *string         `json:"relativeSnapshotDirectory,omitempty" form:"relativeSnapshotDirectory,omitempty"` //Specifies the relative directory path from root path where the snapshot
    RootPath                  *string         `json:"rootPath,omitempty" form:"rootPath,omitempty"` //Specifies the root path where the snapshot is stored, using the following
    ViewName                  *string         `json:"viewName,omitempty" form:"viewName,omitempty"` //Specifies the name of the View that is cloned.
}

/*
 * Structure for the custom type SnapshotCopyReplicationPolicy
 */
type SnapshotCopyReplicationPolicy struct {
    CopyPartial     *bool           `json:"copyPartial,omitempty" form:"copyPartial,omitempty"` //Specifies if Snapshots are copied from the first completely successful
    DaysToKeep      *int64          `json:"daysToKeep,omitempty" form:"daysToKeep,omitempty"` //Specifies the number of days to retain copied Snapshots on the target.
    Multiplier      *int64          `json:"multiplier,omitempty" form:"multiplier,omitempty"` //Specifies a factor to multiply the periodicity by, to determine the copy
    Periodicity     PeriodicityEnum `json:"periodicity,omitempty" form:"periodicity,omitempty"` //Specifies the frequency that Snapshots should be copied to the
    Target          ReplicationTarget1 `json:"target,omitempty" form:"target,omitempty"` //Specifies the replication target to copy the Snapshots to.
}

/*
 * Structure for the custom type SnapshotTarget
 */
type SnapshotTarget struct {
    ArchivalTarget    ArchivalTarget1 `json:"archivalTarget,omitempty" form:"archivalTarget,omitempty"` //Message that specifies the details about an archival target (such as cloud
    CloudDeployTarget CloudDeployTarget `json:"cloudDeployTarget,omitempty" form:"cloudDeployTarget,omitempty"` //Message that specifies the details about CloudDeploy target where backup
    ReplicationTarget ReplicationTarget `json:"replicationTarget,omitempty" form:"replicationTarget,omitempty"` //Message that specifies the details about a remote cluster where backup
    Type              *int64          `json:"type,omitempty" form:"type,omitempty"` //The type of snapshot target this proto represents.
}

/*
 * Structure for the custom type SnapshotTargetPolicyProto
 */
type SnapshotTargetPolicyProto struct {
    CopyPartiallySuccessfulRun *bool           `json:"copyPartiallySuccessfulRun,omitempty" form:"copyPartiallySuccessfulRun,omitempty"` //If this is false, then only snapshots from the first completely successful
    GranularityBucket          GranularityBucket `json:"granularityBucket,omitempty" form:"granularityBucket,omitempty"` //Message that specifies the frequency granularity at which to copy the
    NumDaysToKeep              *int64          `json:"numDaysToKeep,omitempty" form:"numDaysToKeep,omitempty"` //Specifies how to determine the expiration time for snapshots copied due to
    RetentionPolicy            RetentionPolicyProto `json:"retentionPolicy,omitempty" form:"retentionPolicy,omitempty"` //Message that specifies the retention policy for backup snapshots.
    SnapshotTarget             SnapshotTarget  `json:"snapshotTarget,omitempty" form:"snapshotTarget,omitempty"` //Message that specifies details about a target (such as a replication or
}

/*
 * Structure for the custom type SnapshotTarget1
 */
type SnapshotTarget1 struct {
    ArchivalTarget    ArchivalTarget  `json:"archivalTarget,omitempty" form:"archivalTarget,omitempty"` //Specifies settings about the Archival External Target (such as Tape or AWS).
    ReplicationTarget ReplicationTarget1 `json:"replicationTarget,omitempty" form:"replicationTarget,omitempty"` //Specifies settings about the Remote Cohesity Cluster where Snapshots
    Type              Type26Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of a Snapshot target such as 'kLocal', 'kRemote' or
}

/*
 * Structure for the custom type SnapshotVersion
 */
type SnapshotVersion struct {
    AttemptNumber            *int64          `json:"attemptNumber,omitempty" form:"attemptNumber,omitempty"` //Specifies the number of the attempts made by the Job Run
    DeltaSizeBytes           *int64          `json:"deltaSizeBytes,omitempty" form:"deltaSizeBytes,omitempty"` //Specifies the size of the data captured from the source object.
    IsAppConsistent          *bool           `json:"isAppConsistent,omitempty" form:"isAppConsistent,omitempty"` //Specifies if an app-consistent snapshot was captured. For example,
    IsFullBackup             *bool           `json:"isFullBackup,omitempty" form:"isFullBackup,omitempty"` //Specifies if the snapshot is a full backup. For example, all blocks
    JobRunId                 *int64          `json:"jobRunId,omitempty" form:"jobRunId,omitempty"` //Specifies the id of the Job Run that captured the snapshot.
    LocalMountPath           *string         `json:"localMountPath,omitempty" form:"localMountPath,omitempty"` //Specifies the local path relative to the View, without the
    LogicalSizeBytes         *int64          `json:"logicalSizeBytes,omitempty" form:"logicalSizeBytes,omitempty"` //Specifies the size of the snapshot if the data
    PhysicalSizeBytes        *int64          `json:"physicalSizeBytes,omitempty" form:"physicalSizeBytes,omitempty"` //Specifies the amount of data actually used on the disk to store this
    PrimaryPhysicalSizeBytes *int64          `json:"primaryPhysicalSizeBytes,omitempty" form:"primaryPhysicalSizeBytes,omitempty"` //Specifies the total amount of disk space used to store this
    StartedTimeUsecs         *int64          `json:"startedTimeUsecs,omitempty" form:"startedTimeUsecs,omitempty"` //Specifies the time when the Job Run starts capturing a snapshot.
}

/*
 * Structure for the custom type SourceAppParams
 */
type SourceAppParams struct {
    IsVssCopyOnly    *bool           `json:"isVssCopyOnly,omitempty" form:"isVssCopyOnly,omitempty"` //If the backup is a VSS full backup with the copy-only option specified.
    MsExchangeParams MSExchangeParameters `json:"msExchangeParams,omitempty" form:"msExchangeParams,omitempty"` //All the params specific to MS exchange application.
}

/*
 * Structure for the custom type SourceObjectBackupStatus
 */
type SourceObjectBackupStatus struct {
    CurrentSnapshotInfo SnapshotInformation `json:"currentSnapshotInfo,omitempty" form:"currentSnapshotInfo,omitempty"` //Specifies details about the snapshot task created to backup or copy one
    Error               *string         `json:"error,omitempty" form:"error,omitempty"` //Specifies if an error occurred (if any) while running this task.
    IsFullBackup        *bool           `json:"isFullBackup,omitempty" form:"isFullBackup,omitempty"` //Specifies whether this is a 'kFull' or 'kRegular' backup of the Run.
    NumRestarts         *int64          `json:"numRestarts,omitempty" form:"numRestarts,omitempty"` //Specifies the number of times the the task was restarted because of the
    ParentSourceId      *int64          `json:"parentSourceId,omitempty" form:"parentSourceId,omitempty"` //Specifies the id of the registered Protection Source that is the
    Quiesced            *bool           `json:"quiesced,omitempty" form:"quiesced,omitempty"` //Specifies if app-consistent snapshot was captured. This field is set to
    SlaViolated         *bool           `json:"slaViolated,omitempty" form:"slaViolated,omitempty"` //Specifies if the SLA was violated for the Job Run. This field is set
    Source              ProtectionSource `json:"source,omitempty" form:"source,omitempty"` //Specifies a generic structure that represents a node
    Stats               BackupRunStatistics `json:"stats,omitempty" form:"stats,omitempty"` //Specifies statistics about a Backup task in a Protection Job Run.
    Status              Status1Enum     `json:"status,omitempty" form:"status,omitempty"` //Specifies the status of the source object being protected.
    Warnings            *[]string       `json:"warnings,omitempty" form:"warnings,omitempty"` //Array of Warnings.
}

/*
 * Structure for the custom type SourceForPrincipalParameters
 */
type SourceForPrincipalParameters struct {
    ProtectionSourceIds *[]int64        `json:"protectionSourceIds,omitempty" form:"protectionSourceIds,omitempty"` //Array of Protection Source Ids.
    Sid                 *string         `json:"sid,omitempty" form:"sid,omitempty"` //Specifies the SID of the principal to grant access permissions to.
    ViewNames           *[]string       `json:"viewNames,omitempty" form:"viewNames,omitempty"` //Array of View names.
}

/*
 * Structure for the custom type SourceSpecialParameters
 */
type SourceSpecialParameters struct {
    OracleSpecialParameters   ApplicationSpecialJobParameters `json:"oracleSpecialParameters,omitempty" form:"oracleSpecialParameters,omitempty"` //Specifies additional special settings applicable for a Protection Source
    PhysicalSpecialParameters PhysicalSourceSpecialJobParameters `json:"physicalSpecialParameters,omitempty" form:"physicalSpecialParameters,omitempty"` //Specifies additional special settings applicable for a Protection Source
    SkipIndexing              *bool           `json:"skipIndexing,omitempty" form:"skipIndexing,omitempty"` //Specifies not to index the objects in the Protection Source when
    SourceId                  *int64          `json:"sourceId,omitempty" form:"sourceId,omitempty"` //Specifies the object id of the Protection Source that these
    SqlSpecialParameters      ApplicationSpecialJobParameters `json:"sqlSpecialParameters,omitempty" form:"sqlSpecialParameters,omitempty"` //Specifies additional special settings applicable for a Protection Source
    TruncateExchangeLog       *bool           `json:"truncateExchangeLog,omitempty" form:"truncateExchangeLog,omitempty"` //If true, after the Cohesity Cluster successfully captures a Snapshot
    VmCredentials             Credentials     `json:"vmCredentials,omitempty" form:"vmCredentials,omitempty"` //Specifies the administrator credentials to log in to the
    VmwareSpecialParameters   VmwareSourceSpecialJobParameters `json:"vmwareSpecialParameters,omitempty" form:"vmwareSpecialParameters,omitempty"` //Specifies additional special settings applicable for a Protection Source
}

/*
 * Structure for the custom type SourcesForSid
 */
type SourcesForSid struct {
    ProtectionSources []*ProtectionSource `json:"protectionSources,omitempty" form:"protectionSources,omitempty"` //Array of Protection Sources.
    Sid               *string         `json:"sid,omitempty" form:"sid,omitempty"` //Specifies the security identifier (SID) of the principal.
    Views             []*View         `json:"views,omitempty" form:"views,omitempty"` //Array of View Names.
}

/*
 * Structure for the custom type SQLAAGHostAndDatabases
 */
type SQLAAGHostAndDatabases struct {
    AagDatabases    []*AAGAndDatabases `json:"aagDatabases,omitempty" form:"aagDatabases,omitempty"` //Specifies a list of AAGs and database members in each AAG.
    ApplicationNode NodeInAProtectionSourcesTree `json:"applicationNode,omitempty" form:"applicationNode,omitempty"` //Many different node types are supported such as
    Databases       []*ProtectionSource `json:"databases,omitempty" form:"databases,omitempty"` //Specifies all database entities found on the server. Database may or
    ErrorMessage    *string         `json:"errorMessage,omitempty" form:"errorMessage,omitempty"` //Specifies an error message when the host is not registered as an SQL
    UnknownHostName *string         `json:"unknownHostName,omitempty" form:"unknownHostName,omitempty"` //Specifies the name of the host that is not registered as an SQL server
}

/*
 * Structure for the custom type SQLBackupJobParameters
 */
type SQLBackupJobParameters struct {
    AagBackupPreferenceType        *int64          `json:"aagBackupPreferenceType,omitempty" form:"aagBackupPreferenceType,omitempty"` //Preference type for backing up databases that are part of an AAG.
    BackupDatabaseVolumesOnly      *bool           `json:"backupDatabaseVolumesOnly,omitempty" form:"backupDatabaseVolumesOnly,omitempty"` //If set to true, only the volumes associated with databases should be
    BackupSystemDbs                *bool           `json:"backupSystemDbs,omitempty" form:"backupSystemDbs,omitempty"` //Set to true if system databases should be backed up.
    FullBackupType                 *int64          `json:"fullBackupType,omitempty" form:"fullBackupType,omitempty"` //The type of SQL full backup to be used for this job.
    IsCopyOnlyFull                 *bool           `json:"isCopyOnlyFull,omitempty" form:"isCopyOnlyFull,omitempty"` //Whether full backups should be copy-only.
    UseAagPreferencesFromSqlServer *bool           `json:"useAagPreferencesFromSqlServer,omitempty" form:"useAagPreferencesFromSqlServer,omitempty"` //Set to true if we should use AAG preferences specified at the SQL server
    UserDbPreferenceType           *int64          `json:"userDbPreferenceType,omitempty" form:"userDbPreferenceType,omitempty"` //Preference type for backing up user databases on the host.
}

/*
 * Structure for the custom type SQLEnvironmentJobParameters
 */
type SQLEnvironmentJobParameters struct {
    AagPreference              AagPreferenceEnum `json:"aagPreference,omitempty" form:"aagPreference,omitempty"` //Specifies the preference for backing up databases that are part of an AAG.
    AagPreferenceFromSqlServer *bool           `json:"aagPreferenceFromSqlServer,omitempty" form:"aagPreferenceFromSqlServer,omitempty"` //If true, AAG preferences are taken from the SQL server host. If this is
    BackupSystemDatabases      *bool           `json:"backupSystemDatabases,omitempty" form:"backupSystemDatabases,omitempty"` //If true, system databases are backed up. If this is set to false,
    BackupType                 BackupType1Enum `json:"backupType,omitempty" form:"backupType,omitempty"` //Specifies the type of the 'kFull' backup job. Specifies whether it is
    BackupVolumesOnly          *bool           `json:"backupVolumesOnly,omitempty" form:"backupVolumesOnly,omitempty"` //If set to true, only the volumes associated with databases should be
    IsCopyOnlyFull             *bool           `json:"isCopyOnlyFull,omitempty" form:"isCopyOnlyFull,omitempty"` //If true, the backup is a full backup with the copy-only option specified.
    UserDatabasePreference     UserDatabasePreferenceEnum `json:"userDatabasePreference,omitempty" form:"userDatabasePreference,omitempty"` //Specifies the preference for backing up user databases on the host.
}

/*
 * Structure for the custom type SQLProtectionSource
 */
type SQLProtectionSource struct {
    IsAvailableForVssBackup  *bool           `json:"IsAvailableForVssBackup,omitempty" form:"IsAvailableForVssBackup,omitempty"` //Specifies whether the database is marked as available for backup according
    CreatedTimestamp         *string         `json:"createdTimestamp,omitempty" form:"createdTimestamp,omitempty"` //Specifies the time when the database was created. It is displayed in the
    DatabaseName             *string         `json:"databaseName,omitempty" form:"databaseName,omitempty"` //Specifies the database name of the SQL Protection Source, if the type
    DbAagEntityId            *int64          `json:"dbAagEntityId,omitempty" form:"dbAagEntityId,omitempty"` //Specifies the AAG entity id if the database is part of an AAG.
    DbAagName                *string         `json:"dbAagName,omitempty" form:"dbAagName,omitempty"` //Specifies the name of the AAG if the database is part of an AAG.
    DbCompatibilityLevel     *int64          `json:"dbCompatibilityLevel,omitempty" form:"dbCompatibilityLevel,omitempty"` //Specifies the versions of SQL server that the database is compatible
    DbFileGroups             *[]string       `json:"dbFileGroups,omitempty" form:"dbFileGroups,omitempty"` //Specifies the information about the set of file groups for this db on
    DbFiles                  []*DatabaseFileInformation `json:"dbFiles,omitempty" form:"dbFiles,omitempty"` //Specifies the last known information about the set of database files
    DbOwnerUsername          *string         `json:"dbOwnerUsername,omitempty" form:"dbOwnerUsername,omitempty"` //Specifies the name of the database owner.
    Id                       SQLSourceId     `json:"id,omitempty" form:"id,omitempty"` //Specifies a unique id for a SQL Protection Source.
    Name                     *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the instance name of the SQL Protection Source
    OwnerId                  *int64          `json:"ownerId,omitempty" form:"ownerId,omitempty"` //Specifies the id of the container VM for the SQL Protection Source.
    RecoveryModel            RecoveryModelEnum `json:"recoveryModel,omitempty" form:"recoveryModel,omitempty"` //Specifies the Recovery Model for the database in SQL environment.
    SqlServerDbState         SqlServerDbStateEnum `json:"sqlServerDbState,omitempty" form:"sqlServerDbState,omitempty"` //The state of the database as returned by SQL Server.
    SqlServerInstanceVersion SQLServerInstanceVersion `json:"sqlServerInstanceVersion,omitempty" form:"sqlServerInstanceVersion,omitempty"` //Specifies the Server Instance Version.
    Type                     Type17Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of the managed Object in a SQL Protection Source.
}

/*
 * Structure for the custom type SQLApplicationServerRestoreParameters
 */
type SQLApplicationServerRestoreParameters struct {
    CaptureTailLogs                       *bool           `json:"captureTailLogs,omitempty" form:"captureTailLogs,omitempty"` //Set this to true if tail logs are to be captured before the restore
    KeepOffline                           *bool           `json:"keepOffline,omitempty" form:"keepOffline,omitempty"` //Set this to true if we want to restore the database and do not want to
    NewDatabaseName                       *string         `json:"newDatabaseName,omitempty" form:"newDatabaseName,omitempty"` //Specifies optionally a new name for the restored database.
    NewInstanceName                       *string         `json:"newInstanceName,omitempty" form:"newInstanceName,omitempty"` //Specifies an instance name of the SQL Server that should be restored.
    RestoreTimeSecs                       *int64          `json:"restoreTimeSecs,omitempty" form:"restoreTimeSecs,omitempty"` //Specifies the time in the past to which the SQL database needs to be
    TargetDataFilesDirectory              *string         `json:"targetDataFilesDirectory,omitempty" form:"targetDataFilesDirectory,omitempty"` //Specifies the directory where to put the database data files.
    TargetLogFilesDirectory               *string         `json:"targetLogFilesDirectory,omitempty" form:"targetLogFilesDirectory,omitempty"` //Specifies the directory where to put the database log files. Missing
    TargetSecondaryDataFilesDirectoryList []*FilenamePatternToDirectory `json:"targetSecondaryDataFilesDirectoryList,omitempty" form:"targetSecondaryDataFilesDirectoryList,omitempty"` //Specifies the secondary data filename pattern and corresponding
}

/*
 * Structure for the custom type SQLSourceId
 */
type SQLSourceId struct {
    CreatedDateMsecs *int64          `json:"createdDateMsecs,omitempty" form:"createdDateMsecs,omitempty"` //Specifies a unique identifier generated from the date the database is
    DatabaseId       *int64          `json:"databaseId,omitempty" form:"databaseId,omitempty"` //Specifies a unique id of the database but only for the life of the
    InstanceId       *[]int64        `json:"instanceId,omitempty" form:"instanceId,omitempty"` //Array of bytes that stores the SQL Server Instance id.
}

/*
 * Structure for the custom type SSLCertificateConfiguration
 */
type SSLCertificateConfiguration struct {
    Certificate         *string         `json:"certificate,omitempty" form:"certificate,omitempty"` //Certificate is a SSL certificate used by Iris HTTPS webserver.
    LastUpdateTimeMsecs *int64          `json:"lastUpdateTimeMsecs,omitempty" form:"lastUpdateTimeMsecs,omitempty"` //LastUpdateTimeMsecs is a time in milliseconds at which certificate was
    PrivateKey          *string         `json:"privateKey,omitempty" form:"privateKey,omitempty"` //PrivateKey is a matching private key of the above certificate.
}

/*
 * Structure for the custom type SSLCertificationVerificationSettings
 */
type SSLCertificationVerificationSettings struct {
    CaCertificate   *string         `json:"caCertificate,omitempty" form:"caCertificate,omitempty"` //Contains the contents of CA cert/cert chain.
    IsEnabled       *bool           `json:"isEnabled,omitempty" form:"isEnabled,omitempty"` //Whether SSL verification should be performed.
}

/*
 * Structure for the custom type StaticRoute
 */
type StaticRoute struct {
    Description           *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description of the Static Route.
    IsUpdate              *bool           `json:"isUpdate,omitempty" form:"isUpdate,omitempty"` //Specifies if the route is currently being updated on the Cohesity Cluster.
    NetworkInterfaceGroup *string         `json:"networkInterfaceGroup,omitempty" form:"networkInterfaceGroup,omitempty"` //Specifies the group name of the network interfaces to
    NetworkInterfaceIds   *[]int64        `json:"networkInterfaceIds,omitempty" form:"networkInterfaceIds,omitempty"` //Array of Network Interface Ids.
    Subnet                Subnet          `json:"subnet,omitempty" form:"subnet,omitempty"` //Specifies the destination subnet of the Static Route.
    VlanId                *int64          `json:"vlanId,omitempty" form:"vlanId,omitempty"` //Specifies the ID of the VLAN to use for communication with the
}

/*
 * Structure for the custom type StorageEfficiencyTile
 */
type StorageEfficiencyTile struct {
    DataInBytes               *int64          `json:"dataInBytes,omitempty" form:"dataInBytes,omitempty"` //Specifies the size of data brought into the cluster. This is the usage
    DataInBytesSamples        []*Sample       `json:"dataInBytesSamples,omitempty" form:"dataInBytesSamples,omitempty"` //Specifies the samples taken for Data brought into the cluster in bytes
    DataInDedupedBytes        *int64          `json:"dataInDedupedBytes,omitempty" form:"dataInDedupedBytes,omitempty"` //Specifies the size of data after compression and or dedupe operations
    DataInDedupedBytesSamples []*Sample       `json:"dataInDedupedBytesSamples,omitempty" form:"dataInDedupedBytesSamples,omitempty"` //Specifies the samples taken for morphed data in bytes in ascending order
    DedupeRatio               *float64        `json:"dedupeRatio,omitempty" form:"dedupeRatio,omitempty"` //Specifies the current dedupe ratio on the cluster. It is the ratio of
    DedupeRatioSamples        []*Sample       `json:"dedupeRatioSamples,omitempty" form:"dedupeRatioSamples,omitempty"` //Specifies the samples for data reduction ratio in ascending order of time.
    DurationDays              *int64          `json:"durationDays,omitempty" form:"durationDays,omitempty"` //Specifies the duration in days in which the samples were taken.
    IntervalSeconds           *int64          `json:"intervalSeconds,omitempty" form:"intervalSeconds,omitempty"` //Specifies the interval between the samples in seconds.
    LogicalUsedBytes          *int64          `json:"logicalUsedBytes,omitempty" form:"logicalUsedBytes,omitempty"` //Specifies the size of logical data currently represented on the cluster.
    LogicalUsedBytesSamples   []*Sample       `json:"logicalUsedBytesSamples,omitempty" form:"logicalUsedBytesSamples,omitempty"` //Specifies the samples taken for logical data represented in bytes in
    PhysicalUsedBytes         *int64          `json:"physicalUsedBytes,omitempty" form:"physicalUsedBytes,omitempty"` //Specifies the size of physical data currently consumed on the cluster.
    PhysicalUsedBytesSamples  []*Sample       `json:"physicalUsedBytesSamples,omitempty" form:"physicalUsedBytesSamples,omitempty"` //Specifies the samples taken for physical data consumed in bytes in
    StorageReductionRatio     *float64        `json:"storageReductionRatio,omitempty" form:"storageReductionRatio,omitempty"` //Specifies the current storage reduction ratio on the cluster.
    StorageReductionSamples   []*Sample       `json:"storageReductionSamples,omitempty" form:"storageReductionSamples,omitempty"` //Specifies the samples for storage reduction ratio in ascending order of
}

/*
 * Structure for the custom type StoragePolicy
 */
type StoragePolicy struct {
    CloudSpillVaultId            *int64          `json:"cloudSpillVaultId,omitempty" form:"cloudSpillVaultId,omitempty"` //Specifies the vault id assigned for an external Storage
    CompressionPolicy            CompressionPolicyEnum `json:"compressionPolicy,omitempty" form:"compressionPolicy,omitempty"` //Specifies the compression setting to be applied to a Storage Domain
    DeduplicateCompressDelaySecs *int64          `json:"deduplicateCompressDelaySecs,omitempty" form:"deduplicateCompressDelaySecs,omitempty"` //Specifies the time in seconds when deduplication and compression
    DeduplicationEnabled         *bool           `json:"deduplicationEnabled,omitempty" form:"deduplicationEnabled,omitempty"` //Specifies if deduplication is enabled for the Storage Domain (View Box).
    EncryptionPolicy             EncryptionPolicyEnum `json:"encryptionPolicy,omitempty" form:"encryptionPolicy,omitempty"` //Specifies the encryption setting for the Storage Domain (View Box).
    ErasureCodingInfo            ErasureCodingInformation `json:"erasureCodingInfo,omitempty" form:"erasureCodingInfo,omitempty"` //Specifies information for erasure coding.
    InlineCompress               *bool           `json:"inlineCompress,omitempty" form:"inlineCompress,omitempty"` //Specifies if compression should occur inline (as the data is being
    InlineDeduplicate            *bool           `json:"inlineDeduplicate,omitempty" form:"inlineDeduplicate,omitempty"` //Specifies if deduplication should occur inline (as the data is being
    NumFailuresTolerated         *int64          `json:"numFailuresTolerated,omitempty" form:"numFailuresTolerated,omitempty"` //Number of disk failures to tolerate. This is an optional field. Default value
    NumNodeFailuresTolerated     *int64          `json:"numNodeFailuresTolerated,omitempty" form:"numNodeFailuresTolerated,omitempty"` //Number of node failures to tolerate. If NumNodeFailuresTolerated is set to
}

/*
 * Structure for the custom type Subnet
 */
type Subnet struct {
    Component       *string         `json:"component,omitempty" form:"component,omitempty"` //Component that has reserved the subnet.
    Description     *string         `json:"description,omitempty" form:"description,omitempty"` //Description of the subnet.
    Id              *int64          `json:"id,omitempty" form:"id,omitempty"` //ID of the subnet.
    Ip              *string         `json:"ip,omitempty" form:"ip,omitempty"` //Specifies either an IPv6 address or an IPv4 address.
    NetmaskBits     *int64          `json:"netmaskBits,omitempty" form:"netmaskBits,omitempty"` //Specifies the netmask using bits.
    NetmaskIp4      *string         `json:"netmaskIp4,omitempty" form:"netmaskIp4,omitempty"` //Specifies the netmask using an IP4 address.
    NfsAccess       NfsAccessEnum   `json:"nfsAccess,omitempty" form:"nfsAccess,omitempty"` //Specifies whether clients from this subnet can mount using NFS protocol.
    NfsRootSquash   *bool           `json:"nfsRootSquash,omitempty" form:"nfsRootSquash,omitempty"` //Specifies whether clients from this subnet can mount as root on NFS.
    SmbAccess       SmbAccessEnum   `json:"smbAccess,omitempty" form:"smbAccess,omitempty"` //Specifies whether clients from this subnet can mount using SMB protocol.
}

/*
 * Structure for the custom type SupportedErasureCodingAndNodeConfigurations
 */
type SupportedErasureCodingAndNodeConfigurations struct {
    MinNodesAllowed        *int64          `json:"minNodesAllowed,omitempty" form:"minNodesAllowed,omitempty"` //Specifies the minimum number of Nodes supported for this Cluster type.
    SupportedErasureCoding *[]string       `json:"supportedErasureCoding,omitempty" form:"supportedErasureCoding,omitempty"` //Array of Supported Erasure Coding Options.
}

/*
 * Structure for the custom type SupportedPattern
 */
type SupportedPattern struct {
    IsSystemDefined *bool           `json:"isSystemDefined,omitempty" form:"isSystemDefined,omitempty"` //Specifies whether the pattern has been defined by the system or the user.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the Pattern.
    Pattern         *string         `json:"pattern,omitempty" form:"pattern,omitempty"` //Specifies the value of the pattern(Regex).
    PatternType     PatternTypeEnum `json:"patternType,omitempty" form:"patternType,omitempty"` //Specifies the Pattern type.
}

/*
 * Structure for the custom type SyslogServerConfiguration
 */
type SyslogServerConfiguration struct {
    Address                  string          `json:"address" form:"address"` //Specifies the IP address or hostname of the syslog server.
    IsClusterAuditingEnabled *bool           `json:"isClusterAuditingEnabled,omitempty" form:"isClusterAuditingEnabled,omitempty"` //Specifies if Cluster audit logs should be sent to this syslog server.
    IsFilerAuditingEnabled   *bool           `json:"isFilerAuditingEnabled,omitempty" form:"isFilerAuditingEnabled,omitempty"` //Specifies if filer audit logs should be sent to this syslog server.
    Name                     *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies a unique name for the syslog server on the Cluster.
    Port                     int64           `json:"port" form:"port"` //Specifies the port where the syslog server listens.
    Protocol                 Protocol3Enum   `json:"protocol" form:"protocol"` //Specifies the protocol used to send the logs.
}

/*
 * Structure for the custom type VmwareTagAttributes
 */
type VmwareTagAttributes struct {
    Id              *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies the Coheisty id of the VM tag.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the VMware name of the VM tag.
    Uuid            *string         `json:"uuid,omitempty" form:"uuid,omitempty"` //Specifies the VMware Universally Unique Identifier (UUID) of the
}

/*
 * Structure for the custom type TapeMediaInformation
 */
type TapeMediaInformation struct {
    Barcode         *string         `json:"barcode,omitempty" form:"barcode,omitempty"` //Specifies a unique identifier for the media.
    Location        *string         `json:"location,omitempty" form:"location,omitempty"` //Specifies the location of the offline media as recorded by the
    Online          *bool           `json:"online,omitempty" form:"online,omitempty"` //Specifies a flag that indicates if the media is online or offline.
}

/*
 * Structure for the custom type TaskNotification
 */
type TaskNotification struct {
    BackupTask        ParametersForABackupOp `json:"backupTask,omitempty" form:"backupTask,omitempty"` //TODO: Write general description for this field
    CloneTask         CloneTaskInformation `json:"cloneTask,omitempty" form:"cloneTask,omitempty"` //Parameters for a clone op.
    CreatedTimeSecs   *int64          `json:"createdTimeSecs,omitempty" form:"createdTimeSecs,omitempty"` //Timestamp at which the notification was created.
    Description       *string         `json:"description,omitempty" form:"description,omitempty"` //Description holds the actual notification text generated for
    Dismissed         *bool           `json:"dismissed,omitempty" form:"dismissed,omitempty"` //Dismissed keeps track of whether a notification has been seen
    DismissedTimeSecs *int64          `json:"dismissedTimeSecs,omitempty" form:"dismissedTimeSecs,omitempty"` //Timestamp at which user dismissed this notification event.
    FieldMessageTask  BasicTaskInfoHandlesTheBasicElementsOfTheNotificationTask `json:"fieldMessageTask,omitempty" form:"fieldMessageTask,omitempty"` //TODO: Write general description for this field
    Id                *string         `json:"id,omitempty" form:"id,omitempty"` //id identifies a user notification event uniquely.
    RecoveryTask      RecoveryTaskInformation `json:"recoveryTask,omitempty" form:"recoveryTask,omitempty"` //Parameters for a recovery op.
    Status            Status5Enum     `json:"status,omitempty" form:"status,omitempty"` //Status of the task.
    TaskType          TaskTypeEnum    `json:"taskType,omitempty" form:"taskType,omitempty"` //Task type denotes which type of task this notification is for.
    Visited           *bool           `json:"visited,omitempty" form:"visited,omitempty"` //Visited keeps track of whether a notification has been seen or not.
    VisitedTimeSecs   *int64          `json:"visitedTimeSecs,omitempty" form:"visitedTimeSecs,omitempty"` //Timestamp at which user visited this notification event.
}

/*
 * Structure for the custom type TenantDetails
 */
type TenantDetails struct {
    ActiveDirectories      []*ActiveDirectory `json:"activeDirectories,omitempty" form:"activeDirectories,omitempty"` //Specifies the active directories this tenant is associated to.
    BifrostEnabled         *bool           `json:"bifrostEnabled,omitempty" form:"bifrostEnabled,omitempty"` //Specifies whether bifrost (Ambassador proxy) is enabled for tenant.
    CreatedTimeMsecs       *int64          `json:"createdTimeMsecs,omitempty" form:"createdTimeMsecs,omitempty"` //Specifies the epoch time in milliseconds when the tenant account
    Deleted                *bool           `json:"deleted,omitempty" form:"deleted,omitempty"` //Specifies if the Tenant is deleted.
    DeletedTimeMsecs       *int64          `json:"deletedTimeMsecs,omitempty" form:"deletedTimeMsecs,omitempty"` //Specifies the timestamp at which the tenant was deleted.
    DeletionFinished       *bool           `json:"deletionFinished,omitempty" form:"deletionFinished,omitempty"` //Specifies if the object collection is complete for the tenant.
    DeletionInfoVec        []*TenantDeletionInfo `json:"deletionInfoVec,omitempty" form:"deletionInfoVec,omitempty"` //Specifies the current deletion state of object categories.
    Description            *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies the description of this tenant.
    EntityIds              *[]int64        `json:"entityIds,omitempty" form:"entityIds,omitempty"` //Specifies the EntityIds this tenant is associated to.
    LastUpdatedTimeMsecs   *int64          `json:"lastUpdatedTimeMsecs,omitempty" form:"lastUpdatedTimeMsecs,omitempty"` //Specifies the epoch time in milliseconds when the tenant account was last
    LdapProviders          []*LDAPProviderResponse `json:"ldapProviders,omitempty" form:"ldapProviders,omitempty"` //Specifies the ldap providers this tenant is associated to.
    Name                   *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the tenant.
    OrgSuffix              *string         `json:"orgSuffix,omitempty" form:"orgSuffix,omitempty"` //Specifies the organization suffix needed to construct tenant id. Tenant id
    ParentTenantId         *string         `json:"parentTenantId,omitempty" form:"parentTenantId,omitempty"` //Specifies the parent tenant of this tenant if available.
    PolicyIds              *[]string       `json:"policyIds,omitempty" form:"policyIds,omitempty"` //Specifies the PolicyIds this tenant is associated to.
    ProtectionJobs         []*BackupJobProto `json:"protectionJobs,omitempty" form:"protectionJobs,omitempty"` //Specifies the ProtectionJobs this tenant is associated to.
    SubscribeToAlertEmails *bool           `json:"subscribeToAlertEmails,omitempty" form:"subscribeToAlertEmails,omitempty"` //Service provider can optionally unsubscribe from the Tenant Alert Emails.
    TenantId               *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
    ViewBoxIds             *[]int64        `json:"viewBoxIds,omitempty" form:"viewBoxIds,omitempty"` //Specifies the ViewBoxIds this tenant is associated to.
    Views                  []*View         `json:"views,omitempty" form:"views,omitempty"` //Specifies the Views this tenant is associated to.
    VlanIfaceNames         *[]string       `json:"vlanIfaceNames,omitempty" form:"vlanIfaceNames,omitempty"` //Specifies the VlanIfaceNames this tenant is associated to,
}

/*
 * Structure for the custom type TenantActiveDirectoryUpdate
 */
type TenantActiveDirectoryUpdate struct {
    ActiveDirectoryDomains *[]string       `json:"activeDirectoryDomains,omitempty" form:"activeDirectoryDomains,omitempty"` //Specifies the ActiveDirectoryDomain vec for respective tenant.
    TenantId               *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantActiveDirectoryUpdateDetails
 */
type TenantActiveDirectoryUpdateDetails struct {
    ActiveDirectoryDomains *[]string       `json:"activeDirectoryDomains,omitempty" form:"activeDirectoryDomains,omitempty"` //Specifies the ActiveDirectoryDomain vec for respective tenant.
    TenantId               *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantCreateRequest
 */
type TenantCreateRequest struct {
    BifrostEnabled         *bool           `json:"bifrostEnabled,omitempty" form:"bifrostEnabled,omitempty"` //Specifies whether bifrost (Ambassador proxy) is enabled for tenant.
    Description            *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies the description of this tenant.
    Name                   *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the tenant.
    OrgSuffix              *string         `json:"orgSuffix,omitempty" form:"orgSuffix,omitempty"` //Specifies the organization suffix needed to construct tenant id. Tenant id
    ParentTenantId         *string         `json:"parentTenantId,omitempty" form:"parentTenantId,omitempty"` //Specifies the parent tenant of this tenant if available.
    SubscribeToAlertEmails *bool           `json:"subscribeToAlertEmails,omitempty" form:"subscribeToAlertEmails,omitempty"` //Service provider can optionally unsubscribe from the Tenant Alert Emails.
}

/*
 * Structure for the custom type TenantDeletionInfo
 */
type TenantDeletionInfo struct {
    Category            *int64          `json:"category,omitempty" form:"category,omitempty"` //Specifies the category of objects whose deletion state is being captured.
    FinishedAtTimeMsecs *int64          `json:"finishedAtTimeMsecs,omitempty" form:"finishedAtTimeMsecs,omitempty"` //Specifies the time when the process finished.
    ProcessedAtNode     *string         `json:"processedAtNode,omitempty" form:"processedAtNode,omitempty"` //Specifies the node ip where the process ran. Typically this would be
    RetryCount          *int64          `json:"retryCount,omitempty" form:"retryCount,omitempty"` //Specifies the number of times this task has been retried.
    StartedAtTimeMsecs  *int64          `json:"startedAtTimeMsecs,omitempty" form:"startedAtTimeMsecs,omitempty"` //Specifies the time when the process started.
    State               *int64          `json:"state,omitempty" form:"state,omitempty"` //Specifies the deletion completion state of the object category.
}

/*
 * Structure for the custom type TenantEntityUpdate
 */
type TenantEntityUpdate struct {
    EntityIds       *[]int64        `json:"entityIds,omitempty" form:"entityIds,omitempty"` //Specifies the EntityIds for respective tenant.
    TenantId        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantEntityUpdateDetails
 */
type TenantEntityUpdateDetails struct {
    EntityIds       *[]int64        `json:"entityIds,omitempty" form:"entityIds,omitempty"` //Specifies the EntityIds for respective tenant.
    TenantId        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantInformation
 */
type TenantInformation struct {
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies name of the tenant.
    TenantId        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantLdapProviderUpdate
 */
type TenantLdapProviderUpdate struct {
    LdapProviderIds *[]int64        `json:"ldapProviderIds,omitempty" form:"ldapProviderIds,omitempty"` //Specifies the ids of ldap providers for respective tenant.
    TenantId        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantLdapProviderUpdateDetails
 */
type TenantLdapProviderUpdateDetails struct {
    LdapProviderIds *[]int64        `json:"ldapProviderIds,omitempty" form:"ldapProviderIds,omitempty"` //Specifies the ids of ldap providers for respective tenant.
    TenantId        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantProtectionJobUpdate
 */
type TenantProtectionJobUpdate struct {
    ProtectionJobIds *[]int64        `json:"protectionJobIds,omitempty" form:"protectionJobIds,omitempty"` //Specifies the ProtectionJobIds vec for respective tenant.
    TenantId         *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantProtectionJobUpdateDetails
 */
type TenantProtectionJobUpdateDetails struct {
    ProtectionJobIds *[]int64        `json:"protectionJobIds,omitempty" form:"protectionJobIds,omitempty"` //Specifies the ProtectionJobIds vec for respective tenant.
    TenantId         *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantProtectionPolicyUpdate
 */
type TenantProtectionPolicyUpdate struct {
    PolicyIds       *[]string       `json:"policyIds,omitempty" form:"policyIds,omitempty"` //Specifies the PolicyIds for respective tenant.
    TenantId        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantProtectionPolicyUpdateDetails
 */
type TenantProtectionPolicyUpdateDetails struct {
    PolicyIds       *[]string       `json:"policyIds,omitempty" form:"policyIds,omitempty"` //Specifies the PolicyIds for respective tenant.
    TenantId        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantProxy
 */
type TenantProxy struct {
    IpAddress       *string         `json:"ipAddress,omitempty" form:"ipAddress,omitempty"` //Specifies the ip address of the proxy.
    TenantId        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantUpdateDetails
 */
type TenantUpdateDetails struct {
    BifrostEnabled         *bool           `json:"bifrostEnabled,omitempty" form:"bifrostEnabled,omitempty"` //Specifies whether bifrost (Ambassador proxy) is enabled for tenant.
    Description            *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies the description of this tenant.
    Name                   *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the tenant.
    SubscribeToAlertEmails *bool           `json:"subscribeToAlertEmails,omitempty" form:"subscribeToAlertEmails,omitempty"` //Service provider can optionally unsubscribe from the Tenant Alert Emails.
    TenantId               *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantUserUpdateDetails
 */
type TenantUserUpdateDetails struct {
    Sids            *[]string       `json:"sids,omitempty" form:"sids,omitempty"` //Specifies the array of Sid of the users.
    TenantId        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
}

/*
 * Structure for the custom type TenantViewBoxUpdate
 */
type TenantViewBoxUpdate struct {
    TenantId        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
    ViewBoxIds      *[]int64        `json:"viewBoxIds,omitempty" form:"viewBoxIds,omitempty"` //Specifies the ViewBoxIds for respective tenant.
}

/*
 * Structure for the custom type TenantViewBoxUpdateDetails
 */
type TenantViewBoxUpdateDetails struct {
    TenantId        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
    ViewBoxIds      *[]int64        `json:"viewBoxIds,omitempty" form:"viewBoxIds,omitempty"` //Specifies the ViewBoxIds for respective tenant.
}

/*
 * Structure for the custom type TenantViewUpdate
 */
type TenantViewUpdate struct {
    TenantId        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
    ViewNames       *[]string       `json:"viewNames,omitempty" form:"viewNames,omitempty"` //Specifies the PolicyIds for respective tenant.
}

/*
 * Structure for the custom type TenantViewUpdateDetails
 */
type TenantViewUpdateDetails struct {
    TenantId        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
    ViewNames       *[]string       `json:"viewNames,omitempty" form:"viewNames,omitempty"` //Specifies the PolicyIds for respective tenant.
}

/*
 * Structure for the custom type TenantVlanUpdate
 */
type TenantVlanUpdate struct {
    TenantId        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
    VlanIfaceNames  *[]string       `json:"vlanIfaceNames,omitempty" form:"vlanIfaceNames,omitempty"` //Specifies the VlanIfaceNames for respective tenant,
}

/*
 * Structure for the custom type TenantVlanUpdateDetails
 */
type TenantVlanUpdateDetails struct {
    TenantId        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
    VlanIfaceNames  *[]string       `json:"vlanIfaceNames,omitempty" form:"vlanIfaceNames,omitempty"` //Specifies the VlanIfaceNames for respective tenant,
}

/*
 * Structure for the custom type TestIdPReachable
 */
type TestIdPReachable struct {
    IssuerId        *string         `json:"issuerId,omitempty" form:"issuerId,omitempty"` //Specifies the IdP provided Issuer ID for the app.
    SsoUrl          *string         `json:"ssoUrl,omitempty" form:"ssoUrl,omitempty"` //Specifies the SSO URL of the IdP service for the customer. This is the
}

/*
 * Structure for the custom type ThrottlingPolicyOverride
 */
type ThrottlingPolicyOverride struct {
    DatastoreId      *int64          `json:"datastoreId,omitempty" form:"datastoreId,omitempty"` //Specifies the Protection Source id of the Datastore.
    DatastoreName    *string         `json:"datastoreName,omitempty" form:"datastoreName,omitempty"` //Specifies the display name of the Datastore.
    ThrottlingPolicy ThrottlingPolicy `json:"throttlingPolicy,omitempty" form:"throttlingPolicy,omitempty"` //Specifies the throttling policy for a registered Protection Source.
}

/*
 * Structure for the custom type ThrottlingPolicy
 */
type ThrottlingPolicy struct {
    EnforceMaxStreams    *bool           `json:"enforceMaxStreams,omitempty" form:"enforceMaxStreams,omitempty"` //Specifies whether datastore streams are configured for all datastores
    IsEnabled            *bool           `json:"isEnabled,omitempty" form:"isEnabled,omitempty"` //Indicates whether read operations to the datastores, which are
    LatencyThresholds    LatencyThresholds `json:"latencyThresholds,omitempty" form:"latencyThresholds,omitempty"` //Specifies latency thresholds that trigger throttling for all datastores
    MaxConcurrentStreams *int64          `json:"maxConcurrentStreams,omitempty" form:"maxConcurrentStreams,omitempty"` //Specifies the limit on the number of streams Cohesity cluster will make
}

/*
 * Structure for the custom type ThroughputTile
 */
type ThroughputTile struct {
    MaxReadThroughput      *int64          `json:"maxReadThroughput,omitempty" form:"maxReadThroughput,omitempty"` //Maxium Read throughput in last 24 hours.
    MaxWriteThroughput     *int64          `json:"maxWriteThroughput,omitempty" form:"maxWriteThroughput,omitempty"` //Maximum Write throughput in last 24 hours.
    ReadThroughputSamples  []*Sample       `json:"readThroughputSamples,omitempty" form:"readThroughputSamples,omitempty"` //Read throughput samples taken for the past 24 hours at 10 minutes
    WriteThroughputSamples []*Sample       `json:"writeThroughputSamples,omitempty" form:"writeThroughputSamples,omitempty"` //Write throughput samples taken for the past 24 hours at 10 minutes
}

/*
 * Structure for the custom type Time
 */
type Time struct {
    Hour            *int64          `json:"hour,omitempty" form:"hour,omitempty"` //The hour when backup should be performed (0 - 23).
    Minute          *int64          `json:"minute,omitempty" form:"minute,omitempty"` //The minute when backup should be performed (0 - 59).
}

/*
 * Structure for the custom type TimePeriod
 */
type TimePeriod struct {
    Days            *[]DayEnum      `json:"days,omitempty" form:"days,omitempty"` //Array of Week Days.
    EndTime         TimeOfDay       `json:"endTime,omitempty" form:"endTime,omitempty"` //Specifies the end time for the daily time period.
    StartTime       TimeOfDay       `json:"startTime,omitempty" form:"startTime,omitempty"` //Specifies the start time for the daily time period.
}

/*
 * Structure for the custom type TimeOfDay
 */
type TimeOfDay struct {
    Hour            *int64          `json:"hour,omitempty" form:"hour,omitempty"` //Specifies an (0-23) hour in a day.
    Minute          *int64          `json:"minute,omitempty" form:"minute,omitempty"` //Specifies a (0-59) minute in an hour.
}

/*
 * Structure for the custom type TimeRange
 */
type TimeRange struct {
    EndTimeUsecs    *int64          `json:"endTimeUsecs,omitempty" form:"endTimeUsecs,omitempty"` //Specifies the end time specified as a Unix epoch Timestamp
    JobUid          UniqueGlobalId  `json:"jobUid,omitempty" form:"jobUid,omitempty"` //Specifies an id for an object that is unique across Cohesity Clusters.
    StartTimeUsecs  *int64          `json:"startTimeUsecs,omitempty" form:"startTimeUsecs,omitempty"` //Specifies the start time specified as a Unix epoch Timestamp
}

/*
 * Structure for the custom type TrendingData
 */
type TrendingData struct {
    Cancelled           *int64          `json:"cancelled,omitempty" form:"cancelled,omitempty"` //Specifies number of cancelled runs.
    Failed              *int64          `json:"failed,omitempty" form:"failed,omitempty"` //Specifies number of failed runs.
    Running             *int64          `json:"running,omitempty" form:"running,omitempty"` //Specifies number of in-progress runs.
    Successful          *int64          `json:"successful,omitempty" form:"successful,omitempty"` //Specifies number of successful runs.
    Total               *int64          `json:"total,omitempty" form:"total,omitempty"` //Specifies total number of runs.
    TrendName           *string         `json:"trendName,omitempty" form:"trendName,omitempty"` //Specifies trend name. This is start of the day/week/month.
    TrendStartTimeUsecs *int64          `json:"trendStartTimeUsecs,omitempty" form:"trendStartTimeUsecs,omitempty"` //Specifies start of the day/week/month in micro seconds
}

/*
 * Structure for the custom type UniqueGlobalId
 */
type UniqueGlobalId struct {
    ClusterId            *int64          `json:"clusterId,omitempty" form:"clusterId,omitempty"` //Specifies the Cohesity Cluster id where the object was created.
    ClusterIncarnationId *int64          `json:"clusterIncarnationId,omitempty" form:"clusterIncarnationId,omitempty"` //Specifies an id for the Cohesity Cluster that is generated when
    Id                   *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies a unique id assigned to an object (such as a Job)
}

/*
 * Structure for the custom type UniversalIdProto
 */
type UniversalIdProto struct {
    ClusterId            *int64          `json:"clusterId,omitempty" form:"clusterId,omitempty"` //The id of the cluster at which the object was created.
    ClusterIncarnationId *int64          `json:"clusterIncarnationId,omitempty" form:"clusterIncarnationId,omitempty"` //The incarnation id of the above cluster.
    ObjectId             *int64          `json:"objectId,omitempty" form:"objectId,omitempty"` //The object id - this is unique within the above cluster.
}

/*
 * Structure for the custom type UpdateAppInstanceStateParameters
 */
type UpdateAppInstanceStateParameters struct {
    State           State2Enum      `json:"state,omitempty" form:"state,omitempty"` //Specifies the desired app instance state type.
}

/*
 * Structure for the custom type UpdateApplicationServerParameters
 */
type UpdateApplicationServerParameters struct {
    Applications       *[]ApplicationEnum `json:"applications,omitempty" form:"applications,omitempty"` //Specifies the types of applications such as 'kSQL', 'kExchange' running
    HasPersistentAgent *bool           `json:"hasPersistentAgent,omitempty" form:"hasPersistentAgent,omitempty"` //Set this to true if a persistent agent is running on the host. If this is
    Password           *string         `json:"password,omitempty" form:"password,omitempty"` //Specifies password of the username to access the target source.
    ProtectionSourceId *int64          `json:"protectionSourceId,omitempty" form:"protectionSourceId,omitempty"` //Specifies the Id of the Protection Source that contains one or more
    Username           *string         `json:"username,omitempty" form:"username,omitempty"` //Specifies username to access the target source.
}

/*
 * Structure for the custom type UpdateCluster
 */
type UpdateCluster struct {
    AppsSettings                    AthenaAppsConfiguration `json:"appsSettings,omitempty" form:"appsSettings,omitempty"` //TODO: Write general description for this field
    BondingMode                     BondingModeEnum `json:"bondingMode,omitempty" form:"bondingMode,omitempty"` //Specifies the bonding mode to use when bonding NICs to this Cluster.
    ClusterAuditLogConfig           ClusterAuditLogConfiguration `json:"clusterAuditLogConfig,omitempty" form:"clusterAuditLogConfig,omitempty"` //Specifies the settings of the Cluster audit log configuration.
    DnsServerIps                    *[]string       `json:"dnsServerIps,omitempty" form:"dnsServerIps,omitempty"` //Array of IP Addresses of DNS Servers.
    DomainNames                     *[]string       `json:"domainNames,omitempty" form:"domainNames,omitempty"` //Array of Domain Names.
    EnableActiveMonitoring          *bool           `json:"enableActiveMonitoring,omitempty" form:"enableActiveMonitoring,omitempty"` //Specifies if Cohesity can receive monitoring information from the
    EnableUpgradePkgPolling         *bool           `json:"enableUpgradePkgPolling,omitempty" form:"enableUpgradePkgPolling,omitempty"` //If 'true', Cohesity's upgrade server is polled for new releases.
    EncryptionKeyRotationPeriodSecs *int64          `json:"encryptionKeyRotationPeriodSecs,omitempty" form:"encryptionKeyRotationPeriodSecs,omitempty"` //Specifies the period of time (in seconds) when encryption keys are rotated.
    FilerAuditLogConfig             FilerAuditLogConfiguration `json:"filerAuditLogConfig,omitempty" form:"filerAuditLogConfig,omitempty"` //Specifies the settings of the filer audit log configuration.
    Gateway                         *string         `json:"gateway,omitempty" form:"gateway,omitempty"` //Specifies the gateway IP address.
    GoogleAnalyticsEnabled          *bool           `json:"googleAnalyticsEnabled,omitempty" form:"googleAnalyticsEnabled,omitempty"` //Specified whether Google Analytics is enabled.
    IsDocumentationLocal            *bool           `json:"isDocumentationLocal,omitempty" form:"isDocumentationLocal,omitempty"` //Specifies what version of the documentation is used.
    LanguageLocale                  *string         `json:"languageLocale,omitempty" form:"languageLocale,omitempty"` //Specifies the language and locale for this Cohesity Cluster.
    LicenseServerClaimed            *bool           `json:"licenseServerClaimed,omitempty" form:"licenseServerClaimed,omitempty"` //Speifies if cluster is claimed by Helios or not.
    MetadataFaultToleranceFactor    *int64          `json:"metadataFaultToleranceFactor,omitempty" form:"metadataFaultToleranceFactor,omitempty"` //Specifies metadata fault tolerance setting for the cluster. This denotes
    Mtu                             *int64          `json:"mtu,omitempty" form:"mtu,omitempty"` //Specifies the Maxium Transmission Unit (MTU) in bytes of
    MultiTenancyEnabled             *bool           `json:"multiTenancyEnabled,omitempty" form:"multiTenancyEnabled,omitempty"` //Specifies if multi tenancy is enabled in the cluster. Authentication &
    Name                            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the Cohesity Cluster.
    NtpSettings                     NtpSettingsConfig `json:"ntpSettings,omitempty" form:"ntpSettings,omitempty"` //TODO: Write general description for this field
    ReverseTunnelEnabled            *bool           `json:"reverseTunnelEnabled,omitempty" form:"reverseTunnelEnabled,omitempty"` //If 'true', Cohesity's Remote Tunnel is enabled.
    ReverseTunnelEndTimeMsecs       *int64          `json:"reverseTunnelEndTimeMsecs,omitempty" form:"reverseTunnelEndTimeMsecs,omitempty"` //ReverseTunnelEndTimeMsecs specifies the end time in milliseconds since
    SmbAdDisabled                   *bool           `json:"smbAdDisabled,omitempty" form:"smbAdDisabled,omitempty"` //Specifies if Active Directory should be disabled for authentication of SMB
    StigMode                        *bool           `json:"stigMode,omitempty" form:"stigMode,omitempty"` //Specifies if STIG mode is enabled or not.
    SyslogServers                   []*SyslogServerConfiguration `json:"syslogServers,omitempty" form:"syslogServers,omitempty"` //Array of Syslog Servers.
    TenantViewboxSharingEnabled     *bool           `json:"tenantViewboxSharingEnabled,omitempty" form:"tenantViewboxSharingEnabled,omitempty"` //In case multi tenancy is enabled, this flag controls whether multiple
    Timezone                        *string         `json:"timezone,omitempty" form:"timezone,omitempty"` //Specifies the timezone to use for showing time in emails, reports,
    TurboMode                       *bool           `json:"turboMode,omitempty" form:"turboMode,omitempty"` //Specifies if the cluster is in Turbo mode.
}

/*
 * Structure for the custom type UpdateIdPConfigurationRequest
 */
type UpdateIdPConfigurationRequest struct {
    AllowLocalAuthentication *bool           `json:"allowLocalAuthentication,omitempty" form:"allowLocalAuthentication,omitempty"` //Specifies whether to allow local authentication. When IdP is configured,
    Certificate              *string         `json:"certificate,omitempty" form:"certificate,omitempty"` //Specifies the certificate generated for the app by the IdP service when
    CertificateFilename      *string         `json:"certificateFilename,omitempty" form:"certificateFilename,omitempty"` //Specifies the filename used to upload the certificate.
    Enable                   *bool           `json:"enable,omitempty" form:"enable,omitempty"` //Specifies a flag to enable or disable this IdP service. When it is set
    IssuerId                 *string         `json:"issuerId,omitempty" form:"issuerId,omitempty"` //Specifies the IdP provided Issuer ID for the app.
    Roles                    *[]string       `json:"roles,omitempty" form:"roles,omitempty"` //Specifies a list roles assigned to an IdP user if samlAttributeName is
    SamlAttributeName        *string         `json:"samlAttributeName,omitempty" form:"samlAttributeName,omitempty"` //Specifies the SAML attribute name that contains a comma separated list
    SsoUrl                   *string         `json:"ssoUrl,omitempty" form:"ssoUrl,omitempty"` //Specifies the SSO URL of the IdP service for the customer. This is the
}

/*
 * Structure for the custom type UpdateLDAPProviderParameters
 */
type UpdateLDAPProviderParameters struct {
    AdDomainName            *string         `json:"adDomainName,omitempty" form:"adDomainName,omitempty"` //Specifies the domain name of an Active Directory which is mapped to this
    AuthType                AuthTypeEnum    `json:"authType,omitempty" form:"authType,omitempty"` //Specifies the authentication type used while connecting to LDAP servers.
    BaseDistinguishedName   *string         `json:"baseDistinguishedName,omitempty" form:"baseDistinguishedName,omitempty"` //Specifies the base distinguished name used as the base for LDAP
    DomainName              *string         `json:"domainName,omitempty" form:"domainName,omitempty"` //Specifies the name of the domain name to be used for querying LDAP servers
    Id                      *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies the ID of the LDAP provider.
    Name                    *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the LDAP provider.
    Port                    *int64          `json:"port,omitempty" form:"port,omitempty"` //Specifies LDAP server port.
    PreferredLdapServerList *[]string       `json:"preferredLdapServerList,omitempty" form:"preferredLdapServerList,omitempty"` //Specifies the preferred LDAP servers. Server names should be either in
    TenantId                *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the unique id of the tenant.
    UseSsl                  *bool           `json:"useSsl,omitempty" form:"useSsl,omitempty"` //Specifies whether to use SSL for LDAP connections.
    UserDistinguishedName   *string         `json:"userDistinguishedName,omitempty" form:"userDistinguishedName,omitempty"` //Specifies the user distinguished name that is used for LDAP authentication.
    UserPassword            *string         `json:"userPassword,omitempty" form:"userPassword,omitempty"` //Specifies the user password that is used for LDAP authentication.
}

/*
 * Structure for the custom type UpdateMachineAccountsRequest
 */
type UpdateMachineAccountsRequest struct {
    MachineAccounts           *[]string       `json:"machineAccounts,omitempty" form:"machineAccounts,omitempty"` //Array of Machine Accounts.
    OverwriteExistingAccounts *bool           `json:"overwriteExistingAccounts,omitempty" form:"overwriteExistingAccounts,omitempty"` //Specifies whether the specified machine accounts should overwrite the
    Password                  *string         `json:"password,omitempty" form:"password,omitempty"` //Specifies the password for the specified userName.
    UserName                  *string         `json:"userName,omitempty" form:"userName,omitempty"` //Specifies a userName that has administrative privileges in the domain.
}

/*
 * Structure for the custom type UpdateProtectionJobRun
 */
type UpdateProtectionJobRun struct {
    CopyRunTargets    []*CopyTaskJobRunParameters `json:"copyRunTargets,omitempty" form:"copyRunTargets,omitempty"` //Specifies the retention for archival, replication or extended local
    ExpiryTimeUsecs   *int64          `json:"expiryTimeUsecs,omitempty" form:"expiryTimeUsecs,omitempty"` //Specifies a new expiration time as a Unix epoch Timestamp
    JobUid            UniqueGlobalId  `json:"jobUid,omitempty" form:"jobUid,omitempty"` //Specifies a unique universal id for the Job.
    RunStartTimeUsecs *int64          `json:"runStartTimeUsecs,omitempty" form:"runStartTimeUsecs,omitempty"` //Specifies the start time of the Job Run to update. The start time
    SourceIds         *[]int64        `json:"sourceIds,omitempty" form:"sourceIds,omitempty"` //Ids of the Protection Sources. If this is specified, retention time will
}

/*
 * Structure for the custom type SpecifiesTheResponseOfUpdationOfStateOfMultipleProtectionJobs
 */
type SpecifiesTheResponseOfUpdationOfStateOfMultipleProtectionJobs struct {
    FailedJobIds     *[]int64        `json:"failedJobIds,omitempty" form:"failedJobIds,omitempty"` //Specifies a list of Protection Job ids for which updation of state failed.
    SuccessfulJobIds *[]int64        `json:"successfulJobIds,omitempty" form:"successfulJobIds,omitempty"` //Specifies a list of Protection Job ids for which updation of state is
}

/*
 * Structure for the custom type UpdateStateParametersOfProtectionJobs
 */
type UpdateStateParametersOfProtectionJobs struct {
    Action          ActionEnum      `json:"action,omitempty" form:"action,omitempty"` //Specifies the action to be performed on all the specfied Protection Jobs.
    JobIds          *[]int64        `json:"jobIds,omitempty" form:"jobIds,omitempty"` //Specifies a list of Protection Job ids for which the state should change.
}

/*
 * Structure for the custom type UpdateRestoreTaskParams
 */
type UpdateRestoreTaskParams struct {
    RestoreTaskId   *int64          `json:"restoreTaskId,omitempty" form:"restoreTaskId,omitempty"` //Specifies the ID of the existing Restore Task to update.
    SqlOptions      SqlOptionsEnum  `json:"sqlOptions,omitempty" form:"sqlOptions,omitempty"` //Specifies the sql options to update the Restore Task with.
}

/*
 * Structure for the custom type UpdateUserQuotaSettingsForView
 */
type UpdateUserQuotaSettingsForView struct {
    DefaultUserQuotaPolicy          QuotaPolicy     `json:"defaultUserQuotaPolicy,omitempty" form:"defaultUserQuotaPolicy,omitempty"` //Specifies a quota limit that can be optionally applied to Views and
    EnableUserQuota                 *bool           `json:"enableUserQuota,omitempty" form:"enableUserQuota,omitempty"` //If set, it enables/disables the user quota overrides for a view.
    InheritDefaultPolicyFromViewbox *bool           `json:"inheritDefaultPolicyFromViewbox,omitempty" form:"inheritDefaultPolicyFromViewbox,omitempty"` //If set to true, the default_policy in view metadata will be cleared and
    ViewName                        *string         `json:"viewName,omitempty" form:"viewName,omitempty"` //View name of input view.
}

/*
 * Structure for the custom type View1
 */
type View1 struct {
    AccessSids                      *[]string       `json:"accessSids,omitempty" form:"accessSids,omitempty"` //Array of Security Identifiers (SIDs)
    Description                     *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies an optional text description about the View.
    EnableFilerAuditLogging         *bool           `json:"enableFilerAuditLogging,omitempty" form:"enableFilerAuditLogging,omitempty"` //Specifies if Filer Audit Logging is enabled for this view.
    EnableMixedModePermissions      *bool           `json:"enableMixedModePermissions,omitempty" form:"enableMixedModePermissions,omitempty"` //If set, mixed mode (NFS and SMB) access is enabled for this view.
    EnableNfsViewDiscovery          *bool           `json:"enableNfsViewDiscovery,omitempty" form:"enableNfsViewDiscovery,omitempty"` //If set, it enables discovery of view for NFS.
    EnableSmbAccessBasedEnumeration *bool           `json:"enableSmbAccessBasedEnumeration,omitempty" form:"enableSmbAccessBasedEnumeration,omitempty"` //Specifies if access-based enumeration should be enabled.
    EnableSmbEncryption             *bool           `json:"enableSmbEncryption,omitempty" form:"enableSmbEncryption,omitempty"` //Specifies the SMB encryption for the View. If set, it enables the SMB
    EnableSmbViewDiscovery          *bool           `json:"enableSmbViewDiscovery,omitempty" form:"enableSmbViewDiscovery,omitempty"` //If set, it enables discovery of view for SMB.
    EnforceSmbEncryption            *bool           `json:"enforceSmbEncryption,omitempty" form:"enforceSmbEncryption,omitempty"` //Specifies the SMB encryption for all the sessions for the View.
    FileExtensionFilter             FileExtensionFilter `json:"fileExtensionFilter,omitempty" form:"fileExtensionFilter,omitempty"` //TODO: Write general description for this field
    FileLockConfig                  FileLevelDataLockConfigurations `json:"fileLockConfig,omitempty" form:"fileLockConfig,omitempty"` //Specifies a config to lock files in a view - to protect from malicious or
    LogicalQuota                    QuotaPolicy     `json:"logicalQuota,omitempty" form:"logicalQuota,omitempty"` //Specifies an optional logical quota limit (in bytes) for the usage allowed
    ProtocolAccess                  ProtocolAccessEnum `json:"protocolAccess,omitempty" form:"protocolAccess,omitempty"` //Specifies the supported Protocols for the View.
    Qos                             QoS             `json:"qos,omitempty" form:"qos,omitempty"` //Specifies the Quality of Service (QoS) Policy for the View.
    SecurityMode                    SecurityModeEnum `json:"securityMode,omitempty" form:"securityMode,omitempty"` //Specifies the security mode used for this view.
    SmbPermissionsInfo              SMBPermissionsInformation `json:"smbPermissionsInfo,omitempty" form:"smbPermissionsInfo,omitempty"` //Specifies information about SMB permissions.
    StoragePolicyOverride           StoragePolicyOverride `json:"storagePolicyOverride,omitempty" form:"storagePolicyOverride,omitempty"` //Specifies if inline deduplication and compression settings inherited from
    SubnetWhitelist                 []*Subnet       `json:"subnetWhitelist,omitempty" form:"subnetWhitelist,omitempty"` //Array of Subnets.
    TenantId                        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Optional tenant id who has access to this View.
}

/*
 * Structure for the custom type UploadJARRequest
 */
type UploadJARRequest struct {
    JarName         *string         `json:"jarName,omitempty" form:"jarName,omitempty"` //JarName is the name of the uploaded jar.
    JarPath         *string         `json:"jarPath,omitempty" form:"jarPath,omitempty"` //JarPath is the path for the directory where uploaded jar is stored.
}

/*
 * Structure for the custom type UsageAndPerformanceStatistics
 */
type UsageAndPerformanceStatistics struct {
    DataInBytes                    *int64          `json:"dataInBytes,omitempty" form:"dataInBytes,omitempty"` //Data brought into the cluster. This is the usage before data reduction if
    DataInBytesAfterReduction      *int64          `json:"dataInBytesAfterReduction,omitempty" form:"dataInBytesAfterReduction,omitempty"` //Morphed Usage before data is replicated to other nodes as per RF or
    MinUsablePhysicalCapacityBytes *int64          `json:"minUsablePhysicalCapacityBytes,omitempty" form:"minUsablePhysicalCapacityBytes,omitempty"` //Specifies the minimum usable capacity available
    NumBytesRead                   *int64          `json:"numBytesRead,omitempty" form:"numBytesRead,omitempty"` //Provides the total number of bytes read in the last 30 seconds.
    NumBytesWritten                *int64          `json:"numBytesWritten,omitempty" form:"numBytesWritten,omitempty"` //Provides the total number of bytes written in the last 30 second.
    PhysicalCapacityBytes          *int64          `json:"physicalCapacityBytes,omitempty" form:"physicalCapacityBytes,omitempty"` //Provides the total physical capacity in bytes as computed
    ReadIos                        *int64          `json:"readIos,omitempty" form:"readIos,omitempty"` //Provides the number of Read IOs that occurred in the last 30 seconds.
    ReadLatencyMsecs               *float64        `json:"readLatencyMsecs,omitempty" form:"readLatencyMsecs,omitempty"` //Provides the Read latency in milliseconds for the Read IOs that occurred
    SystemCapacityBytes            *int64          `json:"systemCapacityBytes,omitempty" form:"systemCapacityBytes,omitempty"` //Provides the total available capacity as computed by
    SystemUsageBytes               *int64          `json:"systemUsageBytes,omitempty" form:"systemUsageBytes,omitempty"` //Provides the usage of bytes, as computed by the Linux 'statfs' command,
    TotalPhysicalRawUsageBytes     *int64          `json:"totalPhysicalRawUsageBytes,omitempty" form:"totalPhysicalRawUsageBytes,omitempty"` //Provides the usage of bytes, as computed by the Cohesity Cluster,
    TotalPhysicalUsageBytes        *int64          `json:"totalPhysicalUsageBytes,omitempty" form:"totalPhysicalUsageBytes,omitempty"` //Provides the total capacity, as computed by the Cohesity Cluster,
    WriteIos                       *int64          `json:"writeIos,omitempty" form:"writeIos,omitempty"` //Provides the number of Write IOs that occurred in the last 30 seconds.
    WriteLatencyMsecs              *float64        `json:"writeLatencyMsecs,omitempty" form:"writeLatencyMsecs,omitempty"` //Provides the Write latency in milliseconds for the Write IOs that occurred
}

/*
 * Structure for the custom type UserDetails
 */
type UserDetails struct {
    AdditionalGroupNames *[]string       `json:"additionalGroupNames,omitempty" form:"additionalGroupNames,omitempty"` //Array of Additional Groups.
    CreatedTimeMsecs     *int64          `json:"createdTimeMsecs,omitempty" form:"createdTimeMsecs,omitempty"` //Specifies the epoch time in milliseconds when the user account
    Description          *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description about the user.
    Domain               *string         `json:"domain,omitempty" form:"domain,omitempty"` //Specifies the fully qualified domain name (FQDN) of an Active Directory
    EffectiveTimeMsecs   *int64          `json:"effectiveTimeMsecs,omitempty" form:"effectiveTimeMsecs,omitempty"` //Specifies the epoch time in milliseconds when the user becomes
    EmailAddress         *string         `json:"emailAddress,omitempty" form:"emailAddress,omitempty"` //Specifies the email address of the user.
    ExpiredTimeMsecs     *int64          `json:"expiredTimeMsecs,omitempty" form:"expiredTimeMsecs,omitempty"` //Specifies the epoch time in milliseconds when the user becomes
    GoogleAccount        GoogleAccountInformation `json:"googleAccount,omitempty" form:"googleAccount,omitempty"` //Google Account Information of a Helios BaaS user.
    IdpUserInfo          IdPUserInformation `json:"idpUserInfo,omitempty" form:"idpUserInfo,omitempty"` //Specifies an IdP User's information logged in using an IdP.
    LastUpdatedTimeMsecs *int64          `json:"lastUpdatedTimeMsecs,omitempty" form:"lastUpdatedTimeMsecs,omitempty"` //Specifies the epoch time in milliseconds when the user account was last
    Password             *string         `json:"password,omitempty" form:"password,omitempty"` //Specifies the password of this user.
    PrimaryGroupName     *string         `json:"primaryGroupName,omitempty" form:"primaryGroupName,omitempty"` //Specifies the name of the primary group of this User.
    Restricted           *bool           `json:"restricted,omitempty" form:"restricted,omitempty"` //Whether the user is a restricted user. A restricted user can only view
    Roles                *[]string       `json:"roles,omitempty" form:"roles,omitempty"` //Array of Roles.
    S3AccessKeyId        *string         `json:"s3AccessKeyId,omitempty" form:"s3AccessKeyId,omitempty"` //Specifies the S3 Account Access Key ID.
    S3AccountId          *string         `json:"s3AccountId,omitempty" form:"s3AccountId,omitempty"` //Specifies the S3 Account Canonical User ID.
    S3SecretKey          *string         `json:"s3SecretKey,omitempty" form:"s3SecretKey,omitempty"` //Specifies the S3 Account Secret Key.
    SalesforceAccount    SalesforceAccountInformation `json:"salesforceAccount,omitempty" form:"salesforceAccount,omitempty"` //Salesforce Account Information of a Helios user.
    Sid                  *string         `json:"sid,omitempty" form:"sid,omitempty"` //Specifies the unique Security ID (SID) of the user.
    TenantId             *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the effective Tenant ID of the user.
    Username             *string         `json:"username,omitempty" form:"username,omitempty"` //Specifies the login name of the user.
}

/*
 * Structure for the custom type DeleteUsersRequest
 */
type DeleteUsersRequest struct {
    Domain          *string         `json:"domain,omitempty" form:"domain,omitempty"` //Specifies the domain associated with the users to delete.
    TenantId        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Specifies the tenant for which the the users are to be deleted.
    Users           *[]string       `json:"users,omitempty" form:"users,omitempty"` //Array of Users.
}

/*
 * Structure for the custom type UserID
 */
type UserID struct {
    Sid             *string         `json:"sid,omitempty" form:"sid,omitempty"` //If interested in a user via smb_client, include SID.
    UnixUid         *int64          `json:"unixUid,omitempty" form:"unixUid,omitempty"` //If interested in a user via unix-identifier, include UnixUid.
}

/*
 * Structure for the custom type UserIDMapping
 */
type UserIDMapping struct {
    CentrifyZoneMapping     CentrifyZone    `json:"centrifyZoneMapping,omitempty" form:"centrifyZoneMapping,omitempty"` //Specifies the properties associated to a Centrify zone of an Active
    CustomAttributesMapping CustomUnixIDAttributes `json:"customAttributesMapping,omitempty" form:"customAttributesMapping,omitempty"` //Specifies the custom attributes when mapping type is set to
    FixedMapping            FixedUnixIDMapping `json:"fixedMapping,omitempty" form:"fixedMapping,omitempty"` //Specifies the fields when mapping type is set to 'kFixed'. It maps all
    Type                    Type20Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the mapping type used.
}

/*
 * Structure for the custom type UserInformation
 */
type UserInformation struct {
    Domain          *string         `json:"domain,omitempty" form:"domain,omitempty"` //Specifies domain name of the user.
    Sid             *string         `json:"sid,omitempty" form:"sid,omitempty"` //Specifies unique Security ID (SID) of the user.
    UserName        *string         `json:"userName,omitempty" form:"userName,omitempty"` //Specifies user name of the user.
}

/*
 * Structure for the custom type UserInformation1
 */
type UserInformation1 struct {
    IncludeSubtenantObjects *bool           `json:"includeSubtenantObjects,omitempty" form:"includeSubtenantObjects,omitempty"` //Whether objects owned by subtenants should be returned. This would
    PulseAttributeVec       []*KeyValuePair `json:"pulseAttributeVec,omitempty" form:"pulseAttributeVec,omitempty"` //Specifies the KeyValuePair that client (eg. Iris) wants to persist along
    SidVec                  []*ClusterConfigProtoSID `json:"sidVec,omitempty" form:"sidVec,omitempty"` //If specified, only the objects associated with these SIDs should be
    TenantIdVec             *[]string       `json:"tenantIdVec,omitempty" form:"tenantIdVec,omitempty"` //If specified, only the objects associated with this tenant should be
}

/*
 * Structure for the custom type UserRequest
 */
type UserRequest struct {
    AdditionalGroupNames *[]string       `json:"additionalGroupNames,omitempty" form:"additionalGroupNames,omitempty"` //Array of Additional Groups.
    Description          *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description about the user.
    Domain               *string         `json:"domain,omitempty" form:"domain,omitempty"` //Specifies the fully qualified domain name (FQDN) of an Active Directory
    EffectiveTimeMsecs   *int64          `json:"effectiveTimeMsecs,omitempty" form:"effectiveTimeMsecs,omitempty"` //Specifies the epoch time in milliseconds when the user becomes
    EmailAddress         *string         `json:"emailAddress,omitempty" form:"emailAddress,omitempty"` //Specifies the email address of the user.
    ExpiredTimeMsecs     *int64          `json:"expiredTimeMsecs,omitempty" form:"expiredTimeMsecs,omitempty"` //Specifies the epoch time in milliseconds when the user becomes
    Password             *string         `json:"password,omitempty" form:"password,omitempty"` //Specifies the password of this user.
    PrimaryGroupName     *string         `json:"primaryGroupName,omitempty" form:"primaryGroupName,omitempty"` //Specifies the name of the primary group of this User.
    Restricted           *bool           `json:"restricted,omitempty" form:"restricted,omitempty"` //Whether the user is a restricted user. A restricted user can only view
    Roles                *[]string       `json:"roles,omitempty" form:"roles,omitempty"` //Array of Roles.
    Username             *string         `json:"username,omitempty" form:"username,omitempty"` //Specifies the login name of the user.
}

/*
 * Structure for the custom type UserPreferencesProtoUserPreferencesPreference
 */
type UserPreferencesProtoUserPreferencesPreference struct {
    Key             *string         `json:"key,omitempty" form:"key,omitempty"` //Property name.
    Value           *string         `json:"value,omitempty" form:"value,omitempty"` //Property value.
}

/*
 * Structure for the custom type UserQuotaPolicy
 */
type UserQuotaPolicy struct {
    QuotaPolicy     QuotaPolicy     `json:"quotaPolicy,omitempty" form:"quotaPolicy,omitempty"` //Specifies a quota limit that can be optionally applied to Views and
    Sid             *string         `json:"sid,omitempty" form:"sid,omitempty"` //If interested in a user via smb_client, include SID.
    UnixUid         *int64          `json:"unixUid,omitempty" form:"unixUid,omitempty"` //If interested in a user via unix-identifier, include UnixUid.
}

/*
 * Structure for the custom type UserQuotaAndUsage
 */
type UserQuotaAndUsage struct {
    QuotaPolicy     QuotaPolicy     `json:"quotaPolicy,omitempty" form:"quotaPolicy,omitempty"` //Specifies a quota limit that can be optionally applied to Views and
    Sid             *string         `json:"sid,omitempty" form:"sid,omitempty"` //If interested in a user via smb_client, include SID.
    UnixUid         *int64          `json:"unixUid,omitempty" form:"unixUid,omitempty"` //If interested in a user via unix-identifier, include UnixUid.
    UsageBytes      *int64          `json:"usageBytes,omitempty" form:"usageBytes,omitempty"` //Current logical usage of user id in the input view.
}

/*
 * Structure for the custom type UserQuotaSettings
 */
type UserQuotaSettings struct {
    DefaultUserQuotaPolicy QuotaPolicy     `json:"defaultUserQuotaPolicy,omitempty" form:"defaultUserQuotaPolicy,omitempty"` //Specifies a quota limit that can be optionally applied to Views and
    EnableUserQuota        *bool           `json:"enableUserQuota,omitempty" form:"enableUserQuota,omitempty"` //If set, it enables/disables the user quota overrides for a view.
}

/*
 * Structure for the custom type UserQuotaSummaryForUser
 */
type UserQuotaSummaryForUser struct {
    NumViewsAboveAlertThreshold *int64          `json:"numViewsAboveAlertThreshold,omitempty" form:"numViewsAboveAlertThreshold,omitempty"` //Number of views in which user has exceeded alert threshold limit.
    NumViewsAboveHardLimit      *int64          `json:"numViewsAboveHardLimit,omitempty" form:"numViewsAboveHardLimit,omitempty"` //Number of views in which the user has exceeded hard limit.
    TotalNumViews               *int64          `json:"totalNumViews,omitempty" form:"totalNumViews,omitempty"` //Total number of views in which the user has a quota policy specified
}

/*
 * Structure for the custom type UserQuotaSummaryForView
 */
type UserQuotaSummaryForView struct {
    DefaultUserQuotaPolicy      QuotaPolicy     `json:"defaultUserQuotaPolicy,omitempty" form:"defaultUserQuotaPolicy,omitempty"` //Specifies a quota limit that can be optionally applied to Views and
    NumUsersAboveAlertThreshold *int64          `json:"numUsersAboveAlertThreshold,omitempty" form:"numUsersAboveAlertThreshold,omitempty"` //Number of users who has exceeded their specified alert limit.
    NumUsersAboveHardLimit      *int64          `json:"numUsersAboveHardLimit,omitempty" form:"numUsersAboveHardLimit,omitempty"` //Number of users who has exceeded their specified quota hard limit.
    TotalNumUsers               *int64          `json:"totalNumUsers,omitempty" form:"totalNumUsers,omitempty"` //Total number of users who has either a user quota policy override
}

/*
 * Structure for the custom type VcloudDirectorInfoInformation
 */
type VcloudDirectorInfoInformation struct {
    Endpoint        *string         `json:"endpoint,omitempty" form:"endpoint,omitempty"` //vCenter endpoint.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //vCenter name.
}

/*
 * Structure for the custom type VmwareBackupEnvironmentParameters
 */
type VmwareBackupEnvironmentParameters struct {
    AllowCrashConsistentSnapshot *bool           `json:"allowCrashConsistentSnapshot,omitempty" form:"allowCrashConsistentSnapshot,omitempty"` //Whether to fallback to take a crash-consistent snapshot incase taking
    AllowVmsWithPhysicalRdmDisks *bool           `json:"allowVmsWithPhysicalRdmDisks,omitempty" form:"allowVmsWithPhysicalRdmDisks,omitempty"` //Physical RDM disks cannot be backed up using VADP. By default the backups
    VmwareDiskExclusionInfo      []*VmwareDiskExclusionProto `json:"vmwareDiskExclusionInfo,omitempty" form:"vmwareDiskExclusionInfo,omitempty"` //List of Virtual Disk(s) to be excluded from the backup job. These disks
}

/*
 * Structure for the custom type VmwareBackupSourceParameters
 */
type VmwareBackupSourceParameters struct {
    SourceAppParams         SourceAppParams `json:"sourceAppParams,omitempty" form:"sourceAppParams,omitempty"` //This message contains params specific to application running on the source
    VmCredentials           Credentials     `json:"vmCredentials,omitempty" form:"vmCredentials,omitempty"` //Specifies credentials to access a target source.
    VmwareDiskExclusionInfo []*VmwareDiskExclusionProto `json:"vmwareDiskExclusionInfo,omitempty" form:"vmwareDiskExclusionInfo,omitempty"` //List of Virtual Disk(s) to be excluded from the backup job for the source.
}

/*
 * Structure for the custom type VmwareDiskExclusionProto
 */
type VmwareDiskExclusionProto struct {
    ControllerBusNumber *int64          `json:"controllerBusNumber,omitempty" form:"controllerBusNumber,omitempty"` //Controller's bus-id controlling the virtual disk in question.
    ControllerType      *string         `json:"controllerType,omitempty" form:"controllerType,omitempty"` //Controller's type (SCSI, IDE etc).
    UnitNumber          *int64          `json:"unitNumber,omitempty" form:"unitNumber,omitempty"` //Disk unit number to identify the virtual disk within a controller.
}

/*
 * Structure for the custom type VmwareObjectId
 */
type VmwareObjectId struct {
    MorItem         *string         `json:"morItem,omitempty" form:"morItem,omitempty"` //Specifies the Managed Object Reference Item.
    MorType         *string         `json:"morType,omitempty" form:"morType,omitempty"` //Specifies the Managed Object Reference Type.
    Uuid            *string         `json:"uuid,omitempty" form:"uuid,omitempty"` //Specifies a Universally Unique Identifier (UUID) of a VMware Object.
}

/*
 * Structure for the custom type VmwareProtectionSource
 */
type VmwareProtectionSource struct {
    AgentId            *int64          `json:"agentId,omitempty" form:"agentId,omitempty"` //Specifies the id of the persistent agent.
    Agents             []*AgentInformation `json:"agents,omitempty" form:"agents,omitempty"` //Specifies the list of agent information on the Virtual Machine.
    ConnectionState    ConnectionStateEnum `json:"connectionState,omitempty" form:"connectionState,omitempty"` //Specifies the connection state of the Object and are only valid for
    DatastoreInfo      DatastoreInformation `json:"datastoreInfo,omitempty" form:"datastoreInfo,omitempty"` //TODO: Write general description for this field
    FolderType         FolderTypeEnum  `json:"folderType,omitempty" form:"folderType,omitempty"` //Specifies the folder type for the 'kFolder' Object.
    HasPersistentAgent *bool           `json:"hasPersistentAgent,omitempty" form:"hasPersistentAgent,omitempty"` //Set to true if a persistent agent is running on the Virtual Machine.
    HostType           HostType6Enum   `json:"hostType,omitempty" form:"hostType,omitempty"` //Specifies the host type for the 'kVirtualMachine' Object.
    Id                 VmwareObjectId  `json:"id,omitempty" form:"id,omitempty"` //Specifies a unique Protection Source id across Cohesity Clusters.
    IsVmTemplate       *bool           `json:"isVmTemplate,omitempty" form:"isVmTemplate,omitempty"` //IsTemplate specifies if the VM is a template or not.
    Name               *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies a human readable name of the Protection Source.
    TagAttributes      []*VmwareTagAttributes `json:"tagAttributes,omitempty" form:"tagAttributes,omitempty"` //Specifies the optional list of VM Tag attributes associated with this
    ToolsRunningStatus ToolsRunningStatusEnum `json:"toolsRunningStatus,omitempty" form:"toolsRunningStatus,omitempty"` //Specifies the status of VMware Tools for the guest OS on the VM.
    Type               Type19Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of managed Object in a VMware Protection Source.
    VcloudDirectorInfo []*VcloudDirectorInfoInformation `json:"vCloudDirectorInfo,omitempty" form:"vCloudDirectorInfo,omitempty"` //Specifies an array of vCenters to be registered
    VirtualDisks       []*VirtualDiskInformation `json:"virtualDisks,omitempty" form:"virtualDisks,omitempty"` //Specifies an array of virtual disks that are part of the Virtual Machine.
}

/*
 * Structure for the custom type Value
 */
type Value struct {
    Data            Data            `json:"data,omitempty" form:"data,omitempty"` //Specifies the fields to store data of a given type.
    Type            *int64          `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of value.
}

/*
 * Structure for the custom type Data
 */
type Data struct {
    BytesValue      *[]int64        `json:"bytesValue,omitempty" form:"bytesValue,omitempty"` //Specifies the field to store an array of bytes if the current
    DoubleValue     *float64        `json:"doubleValue,omitempty" form:"doubleValue,omitempty"` //Specifies the field to store data if the current data type is a
    Int64Value      *int64          `json:"int64Value,omitempty" form:"int64Value,omitempty"` //Specifies the field to store data if the current data type is a
    StringValue     *string         `json:"stringValue,omitempty" form:"stringValue,omitempty"` //Specifies the field to store data if the current data type is a
}

/*
 * Structure for the custom type Vault
 */
type Vault struct {
    CaTrustedCertificate           *string         `json:"caTrustedCertificate,omitempty" form:"caTrustedCertificate,omitempty"` //Specifies the CA (certificate authority) trusted certificate.
    ClientCertificate              *string         `json:"clientCertificate,omitempty" form:"clientCertificate,omitempty"` //Specifies the client CA  certificate. This certificate is in pem format.
    ClientPrivateKey               *string         `json:"clientPrivateKey,omitempty" form:"clientPrivateKey,omitempty"` //Specifies the client private key. This certificate is in pem format.
    CompressionPolicy              CompressionPolicy1Enum `json:"compressionPolicy,omitempty" form:"compressionPolicy,omitempty"` //Specifies whether to send data to the Vault in a
    Config                         VaultConfiguration `json:"config,omitempty" form:"config,omitempty"` //Specifies the settings required to connect to a specific Vault type.
    CustomerManagingEncryptionKeys *bool           `json:"customerManagingEncryptionKeys,omitempty" form:"customerManagingEncryptionKeys,omitempty"` //Specifies whether to manage the encryption key manually or let the
    DedupEnabled                   *bool           `json:"dedupEnabled,omitempty" form:"dedupEnabled,omitempty"` //Specifies whether to deduplicate data before sending it to the Vault.
    Description                    *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description about the Vault.
    DesiredWalLocation             DesiredWalLocationEnum `json:"desiredWalLocation,omitempty" form:"desiredWalLocation,omitempty"` //Desired location for write ahead logs(wal).
    EncryptionKeyFileDownloaded    *bool           `json:"encryptionKeyFileDownloaded,omitempty" form:"encryptionKeyFileDownloaded,omitempty"` //Specifies if the encryption key file has been downloaded using the
    EncryptionPolicy               EncryptionPolicy1Enum `json:"encryptionPolicy,omitempty" form:"encryptionPolicy,omitempty"` //Specifies whether to send and store data in an encrypted format.
    ExternalTargetType             ExternalTargetTypeEnum `json:"externalTargetType,omitempty" form:"externalTargetType,omitempty"` //Specifies the type of Vault.
    FullArchiveIntervalDays        *int64          `json:"fullArchiveIntervalDays,omitempty" form:"fullArchiveIntervalDays,omitempty"` //Specifies the number days between full archives to the Vault.
    Id                             *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies an id that identifies the Vault.
    IncrementalArchivesEnabled     *bool           `json:"incrementalArchivesEnabled,omitempty" form:"incrementalArchivesEnabled,omitempty"` //Specifies whether to perform incremental archival when sending data
    KeyFileDownloadTimeUsecs       *int64          `json:"keyFileDownloadTimeUsecs,omitempty" form:"keyFileDownloadTimeUsecs,omitempty"` //Specifies the time (in microseconds) when the encryption key file was
    KeyFileDownloadUser            *string         `json:"keyFileDownloadUser,omitempty" form:"keyFileDownloadUser,omitempty"` //Specifies the user who downloaded the encryption key from the
    Name                           *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the Vault.
    Type                           Type34Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of Vault.
    UsageType                      UsageTypeEnum   `json:"usageType,omitempty" form:"usageType,omitempty"` //Specifies the usage type of the Vault.
    VaultBandwidthLimits           VaultBandwidthLimits `json:"vaultBandwidthLimits,omitempty" form:"vaultBandwidthLimits,omitempty"` //VaultBandwidthLimits represents the network bandwidth limits
}

/*
 * Structure for the custom type VaultBandwidthLimits
 */
type VaultBandwidthLimits struct {
    Download        BandwidthLimit  `json:"download,omitempty" form:"download,omitempty"` //Specifies settings for limiting the data transfer rate between
    Upload          BandwidthLimit  `json:"upload,omitempty" form:"upload,omitempty"` //Specifies settings for limiting the data transfer rate between
}

/*
 * Structure for the custom type VaultConfiguration
 */
type VaultConfiguration struct {
    Amazon          AmazonCloudCredentials `json:"amazon,omitempty" form:"amazon,omitempty"` //Specifies the cloud credentials to connect to a Amazon
    Azure           AzureCloudCredentials `json:"azure,omitempty" form:"azure,omitempty"` //Specifies the cloud credentials to connect to a Microsoft
    BucketName      *string         `json:"bucketName,omitempty" form:"bucketName,omitempty"` //Specifies the name of a storage location of the Vault,
    Google          GoogleCloudCredentials `json:"google,omitempty" form:"google,omitempty"` //Specifies the cloud credentials to connect to a Google service account.
    Nas             NASServerCredentials `json:"nas,omitempty" form:"nas,omitempty"` //Specifies the server credentials to connect to a NetApp server.
    Oracle          OracleCloudCredentials `json:"oracle,omitempty" form:"oracle,omitempty"` //Specifies the Oracle Cloud Credentials to connect to an Oracle S3 Compatible
    Qstar           QStarServerCredentials `json:"qstar,omitempty" form:"qstar,omitempty"` //Specifies the server credentials to connect to a QStar service
}

/*
 * Structure for the custom type VaultEncryptionKey
 */
type VaultEncryptionKey struct {
    ClusterName       *string         `json:"clusterName,omitempty" form:"clusterName,omitempty"` //Specifies the name of the source Cohesity Cluster
    EncryptionKeyData *string         `json:"encryptionKeyData,omitempty" form:"encryptionKeyData,omitempty"` //Specifies the encryption key data corresponding to the specified keyUid.
    KeyUid            UniqueGlobalId  `json:"keyUid,omitempty" form:"keyUid,omitempty"` //Specifies the universal id of the Data Encryption Key.
    VaultId           *int64          `json:"vaultId,omitempty" form:"vaultId,omitempty"` //Specifies the id of the Vault whose data is encrypted by
    VaultName         *string         `json:"vaultName,omitempty" form:"vaultName,omitempty"` //Specifies the name of the Vault whose data is encrypted by this key.
}

/*
 * Structure for the custom type View
 */
type View struct {
    AccessSids                      *[]string       `json:"accessSids,omitempty" form:"accessSids,omitempty"` //Array of Security Identifiers (SIDs)
    Aliases                         []*ViewAliasInformation `json:"aliases,omitempty" form:"aliases,omitempty"` //Aliases created for the view. A view alias allows a directory path inside
    AllSmbMountPaths                *[]string       `json:"allSmbMountPaths,omitempty" form:"allSmbMountPaths,omitempty"` //Array of SMB Paths.
    BasicMountPath                  *string         `json:"basicMountPath,omitempty" form:"basicMountPath,omitempty"` //Specifies the NFS mount path of the View (without the hostname
    CaseInsensitiveNamesEnabled     *bool           `json:"caseInsensitiveNamesEnabled,omitempty" form:"caseInsensitiveNamesEnabled,omitempty"` //Specifies whether to support case insensitive file/folder names. This
    CreateTimeMsecs                 *int64          `json:"createTimeMsecs,omitempty" form:"createTimeMsecs,omitempty"` //Specifies the time that the View was created in milliseconds.
    DataLockExpiryUsecs             *int64          `json:"dataLockExpiryUsecs,omitempty" form:"dataLockExpiryUsecs,omitempty"` //DataLock (Write Once Read Many) lock expiry epoch time in microseconds. If
    Description                     *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies an optional text description about the View.
    EnableFilerAuditLogging         *bool           `json:"enableFilerAuditLogging,omitempty" form:"enableFilerAuditLogging,omitempty"` //Specifies if Filer Audit Logging is enabled for this view.
    EnableMixedModePermissions      *bool           `json:"enableMixedModePermissions,omitempty" form:"enableMixedModePermissions,omitempty"` //If set, mixed mode (NFS and SMB) access is enabled for this view.
    EnableNfsViewDiscovery          *bool           `json:"enableNfsViewDiscovery,omitempty" form:"enableNfsViewDiscovery,omitempty"` //If set, it enables discovery of view for NFS.
    EnableSmbAccessBasedEnumeration *bool           `json:"enableSmbAccessBasedEnumeration,omitempty" form:"enableSmbAccessBasedEnumeration,omitempty"` //Specifies if access-based enumeration should be enabled.
    EnableSmbEncryption             *bool           `json:"enableSmbEncryption,omitempty" form:"enableSmbEncryption,omitempty"` //Specifies the SMB encryption for the View. If set, it enables the SMB
    EnableSmbViewDiscovery          *bool           `json:"enableSmbViewDiscovery,omitempty" form:"enableSmbViewDiscovery,omitempty"` //If set, it enables discovery of view for SMB.
    EnforceSmbEncryption            *bool           `json:"enforceSmbEncryption,omitempty" form:"enforceSmbEncryption,omitempty"` //Specifies the SMB encryption for all the sessions for the View.
    FileExtensionFilter             FileExtensionFilter `json:"fileExtensionFilter,omitempty" form:"fileExtensionFilter,omitempty"` //TODO: Write general description for this field
    FileLockConfig                  FileLevelDataLockConfigurations `json:"fileLockConfig,omitempty" form:"fileLockConfig,omitempty"` //Specifies a config to lock files in a view - to protect from malicious or
    LogicalQuota                    QuotaPolicy     `json:"logicalQuota,omitempty" form:"logicalQuota,omitempty"` //Specifies an optional logical quota limit (in bytes) for the usage allowed
    LogicalUsageBytes               *int64          `json:"logicalUsageBytes,omitempty" form:"logicalUsageBytes,omitempty"` //LogicalUsageBytes is the logical usage in bytes for the view.
    Name                            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of the View.
    NfsMountPath                    *string         `json:"nfsMountPath,omitempty" form:"nfsMountPath,omitempty"` //Specifies the path for mounting this View as an NFS share.
    ProtocolAccess                  ProtocolAccessEnum `json:"protocolAccess,omitempty" form:"protocolAccess,omitempty"` //Specifies the supported Protocols for the View.
    Qos                             QoS             `json:"qos,omitempty" form:"qos,omitempty"` //Specifies the Quality of Service (QoS) Policy for the View.
    S3AccessPath                    *string         `json:"s3AccessPath,omitempty" form:"s3AccessPath,omitempty"` //Specifies the path to access this View as an S3 share.
    SecurityMode                    SecurityModeEnum `json:"securityMode,omitempty" form:"securityMode,omitempty"` //Specifies the security mode used for this view.
    SmbMountPath                    *string         `json:"smbMountPath,omitempty" form:"smbMountPath,omitempty"` //Specifies the main path for mounting this View as an SMB share.
    SmbPermissionsInfo              SMBPermissionsInformation `json:"smbPermissionsInfo,omitempty" form:"smbPermissionsInfo,omitempty"` //Specifies information about SMB permissions.
    StoragePolicyOverride           StoragePolicyOverride `json:"storagePolicyOverride,omitempty" form:"storagePolicyOverride,omitempty"` //Specifies if inline deduplication and compression settings inherited from
    SubnetWhitelist                 []*Subnet       `json:"subnetWhitelist,omitempty" form:"subnetWhitelist,omitempty"` //Array of Subnets.
    TenantId                        *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Optional tenant id who has access to this View.
    ViewBoxId                       *int64          `json:"viewBoxId,omitempty" form:"viewBoxId,omitempty"` //Specifies the id of the Storage Domain (View Box) where the View is stored.
    ViewBoxName                     *string         `json:"viewBoxName,omitempty" form:"viewBoxName,omitempty"` //Specifies the name of the Storage Domain (View Box) where the View is stored.
    ViewId                          *int64          `json:"viewId,omitempty" form:"viewId,omitempty"` //Specifies an id of the View assigned by the Cohesity Cluster.
    ViewProtection                  ViewProtection  `json:"viewProtection,omitempty" form:"viewProtection,omitempty"` //Specifies information about the Protection Jobs that are protecting the
}

/*
 * Structure for the custom type ViewAlias
 */
type ViewAlias struct {
    AliasName       *string         `json:"aliasName,omitempty" form:"aliasName,omitempty"` //Alias name.
    ViewName        *string         `json:"viewName,omitempty" form:"viewName,omitempty"` //View name.
    ViewPath        *string         `json:"viewPath,omitempty" form:"viewPath,omitempty"` //View path for the alias.
}

/*
 * Structure for the custom type ViewAliasInformation
 */
type ViewAliasInformation struct {
    AliasName       *string         `json:"aliasName,omitempty" form:"aliasName,omitempty"` //Alias name.
    ViewPath        *string         `json:"viewPath,omitempty" form:"viewPath,omitempty"` //View path for the alias.
}

/*
 * Structure for the custom type DomainViewBox
 */
type DomainViewBox struct {
    AdDomainName                    *string         `json:"adDomainName,omitempty" form:"adDomainName,omitempty"` //Specifies an active directory domain that this view box is mapped to.
    ClientSubnetWhiteList           []*Subnet       `json:"clientSubnetWhiteList,omitempty" form:"clientSubnetWhiteList,omitempty"` //Array of Subnets.
    CloudDownWaterfallThresholdPct  *int64          `json:"cloudDownWaterfallThresholdPct,omitempty" form:"cloudDownWaterfallThresholdPct,omitempty"` //Specifies the cloud down water-fall threshold percentage. This indicates
    CloudDownWaterfallThresholdSecs *int64          `json:"cloudDownWaterfallThresholdSecs,omitempty" form:"cloudDownWaterfallThresholdSecs,omitempty"` //Specifies the cloud down water-fall threshold seconds. This indicates
    ClusterPartitionId              int64           `json:"clusterPartitionId" form:"clusterPartitionId"` //Specifies the Cluster Partition id where the Storage Domain (View Box) is
    ClusterPartitionName            *string         `json:"clusterPartitionName,omitempty" form:"clusterPartitionName,omitempty"` //Specifies the Cohesity Cluster name where the Storage Domain (View Box) is
    DefaultUserQuotaPolicy          QuotaPolicy     `json:"defaultUserQuotaPolicy,omitempty" form:"defaultUserQuotaPolicy,omitempty"` //Specifies an optional quota policy/limits that are inherited by all users
    DefaultViewQuotaPolicy          QuotaPolicy     `json:"defaultViewQuotaPolicy,omitempty" form:"defaultViewQuotaPolicy,omitempty"` //Specifies an optional default logical quota limit (in bytes)
    Id                              *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies the Id of the Storage Domain (View Box).
    LdapProviderId                  *int64          `json:"ldapProviderId,omitempty" form:"ldapProviderId,omitempty"` //When set, the following provides the LDAP provider the view box is
    Name                            string          `json:"name" form:"name"` //Specifies the name of the Storage Domain (View Box).
    PhysicalQuota                   QuotaPolicy     `json:"physicalQuota,omitempty" form:"physicalQuota,omitempty"` //Specifies an optional quota limit (in bytes) for the physical
    RemovalState                    RemovalState1Enum `json:"removalState,omitempty" form:"removalState,omitempty"` //Specifies the current removal state of the Storage Domain (View Box).
    S3BucketsAllowed                *bool           `json:"s3BucketsAllowed,omitempty" form:"s3BucketsAllowed,omitempty"` //Specifies whether creation of a S3 bucket is allowed in this
    Stats                           StorageDomainViewBoxStats `json:"stats,omitempty" form:"stats,omitempty"` //Provides statistics about the Storage Domain (View Box).
    StoragePolicy                   StoragePolicy   `json:"storagePolicy,omitempty" form:"storagePolicy,omitempty"` //Specifies the storage options applied to a Storage Domain (View Box).
    TenantIdVec                     *[]string       `json:"tenantIdVec,omitempty" form:"tenantIdVec,omitempty"` //Optional ids for the tenants that this view box belongs. This must be
    TreatFileSyncAsDataSync         *bool           `json:"treatFileSyncAsDataSync,omitempty" form:"treatFileSyncAsDataSync,omitempty"` //If 'true', when the Cohesity Cluster is writing to a file, the
}

/*
 * Structure for the custom type StorageDomainViewBoxPairing
 */
type StorageDomainViewBoxPairing struct {
    LocalViewBoxId    *int64          `json:"localViewBoxId,omitempty" form:"localViewBoxId,omitempty"` //Specifies the id of the Storage Domain (View Box) on the local Cluster.
    LocalViewBoxName  *string         `json:"localViewBoxName,omitempty" form:"localViewBoxName,omitempty"` //Specifies the name of the Storage Domain (View Box) on the local Cluster.
    RemoteViewBoxId   *int64          `json:"remoteViewBoxId,omitempty" form:"remoteViewBoxId,omitempty"` //Specifies the id of the Storage Domain (View Box) on the remote Cluster.
    RemoteViewBoxName *string         `json:"remoteViewBoxName,omitempty" form:"remoteViewBoxName,omitempty"` //Specifies the name of the Storage Domain (View Box) on the remote Cluster.
}

/*
 * Structure for the custom type StorageDomainViewBoxStats
 */
type StorageDomainViewBoxStats struct {
    CloudUsagePerfStats UsageAndPerformanceStatistics `json:"cloudUsagePerfStats,omitempty" form:"cloudUsagePerfStats,omitempty"` //Provides usage and performance statistics
    Id                  *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies the id of the Storage Domain (View Box).
    LocalUsagePerfStats UsageAndPerformanceStatistics `json:"localUsagePerfStats,omitempty" form:"localUsagePerfStats,omitempty"` //Provides usage and performance statistics
    LogicalStats        LogicalStatistics `json:"logicalStats,omitempty" form:"logicalStats,omitempty"` //Provides logical statistics for logical entities such as Clusters
    UsagePerfStats      UsageAndPerformanceStatistics `json:"usagePerfStats,omitempty" form:"usagePerfStats,omitempty"` //Provides usage and performance statistics
}

/*
 * Structure for the custom type ViewPrivileges
 */
type ViewPrivileges struct {
    PrivilegesType  PrivilegesTypeEnum `json:"privilegesType,omitempty" form:"privilegesType,omitempty"` //Specifies if all, none or specific views are allowed to be accessed.
    ViewIds         *[]int64        `json:"viewIds,omitempty" form:"viewIds,omitempty"` //Specifies the ids of the views which are allowed to be accessed in case
}

/*
 * Structure for the custom type ViewProtection
 */
type ViewProtection struct {
    Inactive        *bool           `json:"inactive,omitempty" form:"inactive,omitempty"` //Specifies if this View is an inactive View that was created on this
    MagnetoEntityId *int64          `json:"magnetoEntityId,omitempty" form:"magnetoEntityId,omitempty"` //Specifies the id of the Protection Source that is using this View.
    ProtectionJobs  []*ProtectionJobInfo `json:"protectionJobs,omitempty" form:"protectionJobs,omitempty"` //Array of Protection Jobs.
}

/*
 * Structure for the custom type ViewProtectionSource
 */
type ViewProtectionSource struct {
    Id              UniqueGlobalId  `json:"id,omitempty" form:"id,omitempty"` //Specifies a unique id of a Protection Source for a View.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies a human readable name of the Protection Source of a View.
    Type            Type18Enum      `json:"type,omitempty" form:"type,omitempty"` //Specifies the type of managed Object in a View Protection Source
}

/*
 * Structure for the custom type ViewUserQuotaParameters
 */
type ViewUserQuotaParameters struct {
    UserQuotaPolicy UserQuotaPolicy `json:"userQuotaPolicy,omitempty" form:"userQuotaPolicy,omitempty"` //Specifies the quota policy applied to a user.
    ViewName        *string         `json:"viewName,omitempty" form:"viewName,omitempty"` //View name of input view.
}

/*
 * Structure for the custom type ViewUserQuotas
 */
type ViewUserQuotas struct {
    Cookie                  *string         `json:"cookie,omitempty" form:"cookie,omitempty"` //This cookie can be used in the succeeding call to list user quotas and
    QuotaAndUsageInAllViews []*QuotaAndUsageInView `json:"quotaAndUsageInAllViews,omitempty" form:"quotaAndUsageInAllViews,omitempty"` //The quota and usage information for a user in all his views.
    SummaryForUser          UserQuotaSummaryForUser `json:"summaryForUser,omitempty" form:"summaryForUser,omitempty"` //Speifies the summary of quota information for a particular user.
    SummaryForView          UserQuotaSummaryForView `json:"summaryForView,omitempty" form:"summaryForView,omitempty"` //Specifies the user quota summary information/result for a view.
    UserQuotaSettings       UserQuotaSettings `json:"userQuotaSettings,omitempty" form:"userQuotaSettings,omitempty"` //Specifies the quota settings parameters for a particular user.
    UsersQuotaAndUsage      []*UserQuotaAndUsage `json:"usersQuotaAndUsage,omitempty" form:"usersQuotaAndUsage,omitempty"` //The list of user quota policies/overrides and usages.
}

/*
 * Structure for the custom type VirtualDiskId
 */
type VirtualDiskId struct {
    BusNumber       *int64          `json:"busNumber,omitempty" form:"busNumber,omitempty"` //Specifies the Id of the controller bus that controls the disk.
    ControllerType  *string         `json:"controllerType,omitempty" form:"controllerType,omitempty"` //Specifies the controller type like SCSI, or IDE etc.
    DiskId          *string         `json:"diskId,omitempty" form:"diskId,omitempty"` //Specfies the uuid of the virtual disk.
    UnitNumber      *int64          `json:"unitNumber,omitempty" form:"unitNumber,omitempty"` //Specifies the disk file name. This is the VMDK name and not the
}

/*
 * Structure for the custom type VirtualDiskInformation
 */
type VirtualDiskInformation struct {
    BusNumber       *int64          `json:"busNumber,omitempty" form:"busNumber,omitempty"` //Specifies the Id of the controller bus that controls the disk.
    ControllerType  *string         `json:"controllerType,omitempty" form:"controllerType,omitempty"` //Specifies the controller type like SCSI, or IDE etc.
    Filename        *string         `json:"filename,omitempty" form:"filename,omitempty"` //Specifies the host file name used as the virtual disk.
    UnitNumber      *int64          `json:"unitNumber,omitempty" form:"unitNumber,omitempty"` //Specifies the disk file name. This is the VMDK name and not the
}

/*
 * Structure for the custom type SpecifiesInformationAboutTheVirtualDisk
 */
type SpecifiesInformationAboutTheVirtualDisk struct {
    BusNumber       *int64          `json:"busNumber,omitempty" form:"busNumber,omitempty"` //Specifies the Id of the controller bus that controls the disk.
    ControllerType  *string         `json:"controllerType,omitempty" form:"controllerType,omitempty"` //Specifies the controller type like SCSI, or IDE etc.
    DiskId          *string         `json:"diskId,omitempty" form:"diskId,omitempty"` //Specifies original disk id. This is sufficient to identify the disk
    DiskLocation    ProtectionSource `json:"diskLocation,omitempty" form:"diskLocation,omitempty"` //Specifies a generic structure that represents a node
    DiskSizeInBytes *int64          `json:"diskSizeInBytes,omitempty" form:"diskSizeInBytes,omitempty"` //Specifies size of the virtual disk in bytes.
    FilePath        *string         `json:"filePath,omitempty" form:"filePath,omitempty"` //Specifies the original file path if applicable.
    MountPoints     *[]string       `json:"mountPoints,omitempty" form:"mountPoints,omitempty"` //Specifies the list of mount points.
    UnitNumber      *int64          `json:"unitNumber,omitempty" form:"unitNumber,omitempty"` //Specifies the disk file name. This is the VMDK name and not the
}

/*
 * Structure for the custom type VirtualDiskMapping
 */
type VirtualDiskMapping struct {
    DiskToOverwrite  VirtualDiskId   `json:"diskToOverwrite,omitempty" form:"diskToOverwrite,omitempty"` //Specifies information about virtual disk which includes disk uuid,
    SourceDisk       VirtualDiskId   `json:"sourceDisk,omitempty" form:"sourceDisk,omitempty"` //Specifies information about virtual disk which includes disk uuid,
    TargetLocationId *int64          `json:"targetLocationId,omitempty" form:"targetLocationId,omitempty"` //Specifies the target location information, for e.g. a datastore in
}

/*
 * Structure for the custom type VirtualDiskMappingResponse
 */
type VirtualDiskMappingResponse struct {
    DiskToOverwrite VirtualDiskId   `json:"diskToOverwrite,omitempty" form:"diskToOverwrite,omitempty"` //Specifies information about virtual disk which includes disk uuid,
    SourceDisk      VirtualDiskId   `json:"sourceDisk,omitempty" form:"sourceDisk,omitempty"` //Specifies information about virtual disk which includes disk uuid,
    TargetLocation  ProtectionSource `json:"targetLocation,omitempty" form:"targetLocation,omitempty"` //Specifies a generic structure that represents a node
}

/*
 * Structure for the custom type VirtualDiskRecoverState
 */
type VirtualDiskRecoverState struct {
    Error                      Error           `json:"error,omitempty" form:"error,omitempty"` //Details about the Error.
    IsInstantRecoveryFinished  *bool           `json:"isInstantRecoveryFinished,omitempty" form:"isInstantRecoveryFinished,omitempty"` //Specifies if instant recovery of the virtual disk is complete.
    TaskState                  TaskStateEnum   `json:"taskState,omitempty" form:"taskState,omitempty"` //Specifies the current state of the restore virtual disks task.
    VirtualDiskRestoreResponse VirtualDiskRestoreResponse `json:"virtualDiskRestoreResponse,omitempty" form:"virtualDiskRestoreResponse,omitempty"` //Specifies the parameters to recover virtual disks of a vm with full
}

/*
 * Structure for the custom type VirtualDiskRestoreParameters
 */
type VirtualDiskRestoreParameters struct {
    PowerOffVmBeforeRecovery *bool           `json:"powerOffVmBeforeRecovery,omitempty" form:"powerOffVmBeforeRecovery,omitempty"` //Specifies whether to power off the VM before recovering virtual disks.
    PowerOnVmAfterRecovery   *bool           `json:"powerOnVmAfterRecovery,omitempty" form:"powerOnVmAfterRecovery,omitempty"` //Specifies whether to power on the VM after recovering virtual disks.
    TargetSourceId           *int64          `json:"targetSourceId,omitempty" form:"targetSourceId,omitempty"` //Specifies the target entity to which the disks should be attached.
    VirtualDiskMappings      []*VirtualDiskMapping `json:"virtualDiskMappings,omitempty" form:"virtualDiskMappings,omitempty"` //Specifies the list of virtual disks mappings.
}

/*
 * Structure for the custom type VirtualDiskRestoreResponse
 */
type VirtualDiskRestoreResponse struct {
    PowerOffVmBeforeRecovery *bool           `json:"powerOffVmBeforeRecovery,omitempty" form:"powerOffVmBeforeRecovery,omitempty"` //Specifies whether to power off the VM before recovering virtual disks.
    PowerOnVmAfterRecovery   *bool           `json:"powerOnVmAfterRecovery,omitempty" form:"powerOnVmAfterRecovery,omitempty"` //Specifies whether to power on the VM after recovering virtual disks.
    TargetSource             ProtectionSource `json:"targetSource,omitempty" form:"targetSource,omitempty"` //Specifies a generic structure that represents a node
    VirtualDiskMappings      []*VirtualDiskMappingResponse `json:"virtualDiskMappings,omitempty" form:"virtualDiskMappings,omitempty"` //Specifies the list of virtual disks mappings.
}

/*
 * Structure for the custom type VLAN
 */
type VLAN struct {
    AddToClusterPartition *bool           `json:"addToClusterPartition,omitempty" form:"addToClusterPartition,omitempty"` //Specifies whether to add the VLAN IPs to the cluster partition
    Description           *string         `json:"description,omitempty" form:"description,omitempty"` //Specifies a description of the VLAN.
    Gateway               *string         `json:"gateway,omitempty" form:"gateway,omitempty"` //Specifies the Gateway of the VLAN.
    Hostname              *string         `json:"hostname,omitempty" form:"hostname,omitempty"` //Specifies the hostname of the VLAN.
    Id                    *int64          `json:"id,omitempty" form:"id,omitempty"` //Specifies the id of the VLAN.
    IfaceGroupName        *string         `json:"ifaceGroupName,omitempty" form:"ifaceGroupName,omitempty"` //Specifies the interface group name of the VLAN. It is in the format of
    InterfaceName         *string         `json:"interfaceName,omitempty" form:"interfaceName,omitempty"` //Specifies the interface name of the VLAN.
    Ips                   *[]string       `json:"ips,omitempty" form:"ips,omitempty"` //Array of IPs.
    Subnet                Subnet          `json:"subnet,omitempty" form:"subnet,omitempty"` //Specifies the subnet of the VLAN.
    TenantId              *string         `json:"tenantId,omitempty" form:"tenantId,omitempty"` //Optional tenant id that this vlan belongs to.
    VlanName              *string         `json:"vlanName,omitempty" form:"vlanName,omitempty"` //Specifies the VLAN name of the vlanId.
}

/*
 * Structure for the custom type VLANParameters
 */
type VLANParameters struct {
    DisableVlan     *bool           `json:"disableVlan,omitempty" form:"disableVlan,omitempty"` //Specifies whether to use the VIPs even when VLANs are configured on the
    InterfaceName   *string         `json:"interfaceName,omitempty" form:"interfaceName,omitempty"` //Specifies the interface name to use for mounting Cohesity's view on
    Vlan            *int64          `json:"vlan,omitempty" form:"vlan,omitempty"` //Specifies the VLAN to use for mounting Cohesity's view on the remote
}

/*
 * Structure for the custom type VmwareCloneParameters
 */
type VmwareCloneParameters struct {
    DatastoreFolderId *int64          `json:"datastoreFolderId,omitempty" form:"datastoreFolderId,omitempty"` //Specifies the folder where the restore datastore should be created.
    DetachNetwork     *bool           `json:"detachNetwork,omitempty" form:"detachNetwork,omitempty"` //Specifies whether the network should be detached from the
    DisableNetwork    *bool           `json:"disableNetwork,omitempty" form:"disableNetwork,omitempty"` //Specifies whether the network should be left in disabled state.
    NetworkId         *int64          `json:"networkId,omitempty" form:"networkId,omitempty"` //Specifies a network configuration to be attached to the cloned or
    PoweredOn         *bool           `json:"poweredOn,omitempty" form:"poweredOn,omitempty"` //Specifies the power state of the cloned or recovered objects.
    Prefix            *string         `json:"prefix,omitempty" form:"prefix,omitempty"` //Specifies a prefix to prepended to the source object name to derive a
    ResourcePoolId    *int64          `json:"resourcePoolId,omitempty" form:"resourcePoolId,omitempty"` //Specifies the resource pool where the cloned or recovered objects are
    Suffix            *string         `json:"suffix,omitempty" form:"suffix,omitempty"` //Specifies a suffix to appended to the original source object name
    VmFolderId        *int64          `json:"vmFolderId,omitempty" form:"vmFolderId,omitempty"` //Specifies a folder where the VMs should be restored. This is applicable
}

/*
 * Structure for the custom type VmwareEnvironmentJobParameters
 */
type VmwareEnvironmentJobParameters struct {
    ExcludedDisks             []*DiskUnit     `json:"excludedDisks,omitempty" form:"excludedDisks,omitempty"` //Specifies the list of Disks to be excluded from backing up. These disks
    FallbackToCrashConsistent *bool           `json:"fallbackToCrashConsistent,omitempty" form:"fallbackToCrashConsistent,omitempty"` //If true, takes a crash-consistent snapshot when app-consistent snapshot
    SkipPhysicalRdmDisks      *bool           `json:"skipPhysicalRdmDisks,omitempty" form:"skipPhysicalRdmDisks,omitempty"` //If true, skip physical RDM disks when backing up VMs. Otherwise, backup
}

/*
 * Structure for the custom type VmwareRestoreParameters
 */
type VmwareRestoreParameters struct {
    DatastoreFolderId *int64          `json:"datastoreFolderId,omitempty" form:"datastoreFolderId,omitempty"` //Specifies the folder where the restore datastore should be created.
    DatastoreId       *int64          `json:"datastoreId,omitempty" form:"datastoreId,omitempty"` //Specifies the datastore where the object's files should be
    DetachNetwork     *bool           `json:"detachNetwork,omitempty" form:"detachNetwork,omitempty"` //Specifies whether the network should be detached from the
    DisableNetwork    *bool           `json:"disableNetwork,omitempty" form:"disableNetwork,omitempty"` //Specifies whether the network should be left in disabled state.
    NetworkId         *int64          `json:"networkId,omitempty" form:"networkId,omitempty"` //Specifies a network configuration to be attached to the cloned or
    PoweredOn         *bool           `json:"poweredOn,omitempty" form:"poweredOn,omitempty"` //Specifies the power state of the cloned or recovered objects.
    Prefix            *string         `json:"prefix,omitempty" form:"prefix,omitempty"` //Specifies a prefix to prepended to the source object name to derive a
    ResourcePoolId    *int64          `json:"resourcePoolId,omitempty" form:"resourcePoolId,omitempty"` //Specifies the resource pool where the cloned or recovered objects are
    Suffix            *string         `json:"suffix,omitempty" form:"suffix,omitempty"` //Specifies a suffix to appended to the original source object name
    VmFolderId        *int64          `json:"vmFolderId,omitempty" form:"vmFolderId,omitempty"` //Specifies a folder where the VMs should be restored. This is applicable
}

/*
 * Structure for the custom type VmwareSourceSpecialJobParameters
 */
type VmwareSourceSpecialJobParameters struct {
    ApplicationParameters ApplicationSpecificParameters `json:"applicationParameters,omitempty" form:"applicationParameters,omitempty"` //TODO: Write general description for this field
    ExcludedDisks         []*DiskUnit     `json:"excludedDisks,omitempty" form:"excludedDisks,omitempty"` //Specifies the list of Disks to be excluded from backing up. These disks
    VmCredentials         Credentials     `json:"vmCredentials,omitempty" form:"vmCredentials,omitempty"` //Specifies the administrator credentials to log in to the
}

/*
 * Structure for the custom type VolumeSecurityInformation
 */
type VolumeSecurityInformation struct {
    GroupId         *int64          `json:"groupId,omitempty" form:"groupId,omitempty"` //Specifies the Unix group ID for this volume. 0 indicates the root id.
    Permissions     *string         `json:"permissions,omitempty" form:"permissions,omitempty"` //Specifies the Unix permission bits in octal string format.
    Style           StyleEnum       `json:"style,omitempty" form:"style,omitempty"` //Specifies the security style associated with this volume.
    UserId          *int64          `json:"userId,omitempty" form:"userId,omitempty"` //Specifies the Unix user id for this volume. 0 indicates the root id.
}

/*
 * Structure for the custom type VserverNetworkInterface
 */
type VserverNetworkInterface struct {
    DataProtocols   *[]DataProtocolEnum `json:"dataProtocols,omitempty" form:"dataProtocols,omitempty"` //Array of Data Protocols.
    IpAddress       *string         `json:"ipAddress,omitempty" form:"ipAddress,omitempty"` //Specifies the IP address of this interface.
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Specifies the name of this interface.
}

/*
 * Structure for the custom type WebHookDeliveryTarget
 */
type WebHookDeliveryTarget struct {
    CurlOptions     *string         `json:"curlOptions,omitempty" form:"curlOptions,omitempty"` //Specifies curl options used to invoke external api url defined above.
    ExternalApiUrl  *string         `json:"externalApiUrl,omitempty" form:"externalApiUrl,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type WindowsHostSnapshotParameters
 */
type WindowsHostSnapshotParameters struct {
    CopyOnlyBackup      *bool           `json:"copyOnlyBackup,omitempty" form:"copyOnlyBackup,omitempty"` //Specifies whether to backup regardless of the state of each file's
    DisableMetadata     *bool           `json:"disableMetadata,omitempty" form:"disableMetadata,omitempty"` //Specifies whether to disable fetching and storing of some metadata
    DisableNotification *bool           `json:"disableNotification,omitempty" form:"disableNotification,omitempty"` //Specifies whether to disable some notification steps when taking
    ExcludedVssWriters  *[]string       `json:"excludedVssWriters,omitempty" form:"excludedVssWriters,omitempty"` //Specifies a list of Windows VSS writers that are excluded from backups.
}

/*
 * Structure for the custom type WormRetentionProto
 */
type WormRetentionProto struct {
    PolicyType      *int64          `json:"policyType,omitempty" form:"policyType,omitempty"` //The type of WORM policy set on this run. This field is irrelevant
}
