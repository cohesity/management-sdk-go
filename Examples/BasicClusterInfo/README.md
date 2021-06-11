
## Usage: 

Hello World example of getting the cluster information.

```
go run getBasicClusterInfo.go
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

## Get Cluster Info
```
cluster := client.Cluster()
clusterInfo, err := cluster.GetBasicClusterInfo()
if err != nil {
	fmt.Println(err.Error())
	return
}
fmt.Println("Cluster Name: ", *clusterInfo.Name)
fmt.Println("Cluster Version: ", *clusterInfo.ClusterSoftwareVersion)
```

## Sample Output
```
Cluster Name:  CDN4-VirtualEdition-Cluster
Cluster Version:  6.2
```