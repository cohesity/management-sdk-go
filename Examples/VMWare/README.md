## Usage: 
This is an example to add vCenter as a Protection Source. 
```
go run registerVCenter.go
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

## Register vCenter

```
Body := &models.RegisterProtectionSourceParameters{}
Body.Endpoint = &vcenterIP
Body.Password = &vcenterPassword
Body.Username = &vcenterUsername
Body.VmwareType = models.VmwareType_KVCENTER
Body.Environment = models.EnvironmentRegisterProtectionSourceParameters_KVMWARE

result, err := client.ProtectionSources().CreateRegisterProtectionSource(Body)
```