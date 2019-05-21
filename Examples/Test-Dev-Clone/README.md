## Usage: 
This is an example for Cloning a VM for Test Dev workflow.
```
go run testDevCloneVM.go
```

## Connect to the Cohesity Cluster

```
Username := "clusterUsername"
Password := "clusterPassword"
ClusterVip := "cluster.org.company.com"
Domain := "AD.org.company.com" //Initialize to 'LOCAL' for cluster user.

client, err := CohesityManagementSdk.NewCohesitySdkClient(ClusterVip, Username, Password, Domain)
if err != nil {
	fmt.Println(err) // Handle Error.
	return
}
```

## Search Object

```
search := &vmToClone
environments = []models.EnvironmentSearchObjectsEnum{models.EnvironmentSearchObjects_KVMWARE}
result, err = restore.SearchObjects(startIndex, operatingSystems, application, ownerEntityId, 
	    jobIds, viewBoxIds, environments, endTimeUsecs, pageCount, search, registeredSourceIds, 
	    startTimeUsecs)
	    
```

## Clone VM
```
body := &models.CloneTaskRequest{}
body.Name = cloneTaskName
body.Type = models.TypeCloneTaskRequest_KCLONEVMS
body.TargetViewName = &targetViewName
body.VmwareParameters = models.VmwareCloneParameters{}
body.VmwareParameters.ResourcePoolId = &resourcePoolId
body.VmwareParameters.PoweredOn = &isPoweredOn
body.VmwareParameters.Prefix = &prefix

body.Objects = make([]*models.RestoreObjectDetails, 1)
body.Objects[0] = &models.RestoreObjectDetails{}
body.Objects[0].JobId = jobId
body.Objects[0].SourceName = &vmToClone
body.Objects[0].Environment = models.EnvironmentRestoreObjectDetails_KVMWARE

cloneTaskResult, err := restore.CreateCloneTask(body)
```

## Poll for Clone Task
```
for retryTimeout > 0 {
	result, _ := restore.GetRestoreTaskById(taskID)
	if result.Status == models.StatusRestoreTask_KFINISHED {
		fmt.Println("Task Clone VM Completed.")
	} else {
		time.Sleep(time.Duration(interval) * time.Second)
		increment += interval
		fmt.Printf("Task status: %v , Time elapsed retrying: %v\n", result.Status, increment)
		}
	}

```

## Destroy Clone VM
```
restore.DeletePublicDestroyCloneTask(*cloneTaskResult.Id)
fmt.Printf("Cloned VM: %v destroyed\n", vmToClone)
```


