Cohesity Management SDK GO
=================

# Overview

The Cohesity Management SDK provides an easy-to-use language binding to harness the power of Cohesity REST APIs in your GoLang applications.

# Getting Started

```bash
go get github.com/cohesity/management-sdk-go
```
Note: To update the package do:

```
 go get -u github.com/cohesity/management-sdk-go
```

### How to use
In order to setup authentication and initialization of the API client, you need the following information.

Import library
```
import CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
```

Initialize the App Client.
```
Username := "clusterUsername"
Password := "clusterPassword"
ClusterVip := "cluster.org.company.com"
Domain := "AD.org.company.com" //Initialize to 'LOCAL' for cluster user
client, err := CohesityManagementSdk.NewCohesitySdkClient(ClusterVip, Username, Password, Domain)
if err != nil {
    fmt.Println(err.Error())
    return
}
```

Get Cluster Basic Info
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

## Questions or Feedback :

We would love to hear from you. Please send your questions and feedback to: *cohesity-api-sdks@cohesity.com*
