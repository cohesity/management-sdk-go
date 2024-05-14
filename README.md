Cohesity Management SDK GO
=================

Note: This version of SDK is compatible with 6.3.1x only. We plan to release SDK with newer versions of the cluster soon. 
Do reach out to cohesity-api-sdks@cohesity.com for more details

# Overview

The *Cohesity Management SDK*  provides an easy-to-use language binding to 
harness the power of *Cohesity DataPlatform APIs* in your Golang applications
.These
 APIs are available for apps running on Cohesity Apps Infrastructure.


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
Username := "cluster_admin"
Password := "cluster_password"
ClusterVip := "prod-cluster.eng.cohesity.com"
var Domain string // Set for AD user only.
client := CohesityManagementSdk.NewCohesitySdkClient(ClusterVip, Username, Password, Domain)
```

Get Cluster Basic Info
```
cluster = client.Cluster()
cluster.GetBasicClusterInfo()
```

## Questions or Feedback :

We would love to hear from you. Please send your questions and feedback to: *support@cohesity.com*
