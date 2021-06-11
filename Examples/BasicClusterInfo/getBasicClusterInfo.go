// Copyright 2019 Cohesity Inc.

//Usage:
//go run getBasicClusterInfo.go

package main

import (
	"fmt"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
)

func main() {
	Username := "admin"
	Password := "admin"
	ClusterVip := "10.2.145.49"
	Domain := "LOCAL" //Initialize to 'LOCAL' for cluster user

	client, err := CohesityManagementSdk.NewCohesitySdkClient(ClusterVip, Username, Password, Domain)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cluster := client.Cluster()
	clusterInfo, err := cluster.GetBasicClusterInfo()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Cluster Name: ", *clusterInfo.Name)
	fmt.Println("Cluster Version: ", *clusterInfo.ClusterSoftwareVersion)

}
