## Usage: 

```
go run unresolvedAlertsWorkflow.go
```
## Workflow
```
Alerts workflow:
1. Lists 10 Recent Unresolved Alerts.(Number 10 is tweakable in the code)
2. Accepts Alert ID to resolve as user input. 
3. Checks if the Alert ID exists in the list specified.
4. Resolves the Alert and gives back Notification ID.
```
## Connect to the Cohesity Cluster

```
Username := "clusterUsername"
Password := "clusterPassword"
ClusterVip := "cluster.org.company.com"
Domain := "AD.org.company.com" //Initialize to 'LOCAL' for cluster user.

client := CohesityManagementSdk.NewCohesitySdkClient(ClusterVip, Username, Password, Domain)
```
## List 10 Alerts. 
```
MaxAlerts := 10
result, err := alerts.GetAlerts(MaxAlerts, allUnderHierarchy, alertIDList, alertCategoryList, propertyKey, propertyValue, 
		endDateUsecs, alertSeverityList, tenantIds, alertTypeList, startDateUsecs, alertStateList, resolutionIDList)
```

## Resolve the Alert specified. 
```
Body := &models.AlertResolutionRequest{}
Body.AlertIdList = &[]string{alertID}
Body.ResolutionDetails = models.AlertResolutionInfo{}
resolutionSummary := "Resolving this Alert through GoLang Automation."
Body.ResolutionDetails.ResolutionSummary = &resolutionSummary
alerts := client.Alerts()
var err error
var result *models.AlertResolution
result, err = alerts.CreateResolution(Body)
```

## Sample Output
```
Unresolved AlertID: 2620510:1556920180603842 , AlertName: kNodeHealth 
Unresolved AlertID: 2594229:1556675692415861 , AlertName: kNode 
Unresolved AlertID: 2526977:1556148510919186 , AlertName: kBackupRestore 
Unresolved AlertID: 2525935:1556144911802941 , AlertName: kBackupRestore 
Unresolved AlertID: 2525918:1556141307008650 , AlertName: kNodeHealth 
Unresolved AlertID: 2524820:1556137707649518 , AlertName: kBackupRestore 
Unresolved AlertID: 2524803:1556134104671392 , AlertName: kNodeHealth 
Unresolved AlertID: 2523761:1556130507672593 , AlertName: kBackupRestore 
Unresolved AlertID: 2523744:1556126904514742 , AlertName: kNode 
Unresolved AlertID: 2523727:1556123307494704 , AlertName: kBackupRestore 
Specify one among AlertID to Resolve: 2525935:1556144911802941
Alert ID 2525935:1556144911802941 is resolved with resolution ID 2622618 
```
