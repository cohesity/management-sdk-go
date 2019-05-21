## Usage: 
```
go run listProtectionJobs.go
```

## Connect to the Cohesity Cluster

```
Username := "clusterUsername"
Password := "clusterPassword"
ClusterVip := "cluster.org.company.com"
Domain := "AD.org.company.com" //Initialize to 'LOCAL' for cluster user.

client, _ := CohesityManagementSdk.NewCohesitySdkClient(ClusterVip, Username, Password, Domain)
```
## Call protection Jobs endpoint. 
```
result, err := protectionJobs.GetProtectionJobs(onlyReturnBasicSummary, includeRpoSnapshots, ids, names, policyIds, 
													isActive, includeLastRunAndStats, onlyReturnDataMigrationJobs, 
													tenantIds, allUnderHierarchy, environments, isDeleted)
```

## Sample Output
```
Protection Job Names:
1 . backing up cohesity-view1
2 . NAS protection-job
3 . s3 View protect
4 . Windows-block-based
5 . Windows-file-based
6 . ci-cd-protection-job
7 . automation-vcenter-vms
```
