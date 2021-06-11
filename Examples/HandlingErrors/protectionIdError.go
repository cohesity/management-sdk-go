// Copyright 2019 Cohesity Inc.

//Usage:
//go run protectionIdError.go

package main

import (
	"fmt"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	"github.com/cohesity/management-sdk-go/models"
)

func main() {
	Username := "clusterUsername"
	Password := "clusterPassword"
	ClusterVip := "cluster.org.company.com"
	Domain := "AD.org.company.com" //Initialize to 'LOCAL' for cluster user.
	var err error
	client, err := CohesityManagementSdk.NewCohesitySdkClient(ClusterVip, Username, Password, Domain)
	if err != nil {
		fmt.Println(err)
		return
	}
	protectionJobs := client.ProtectionJobs()

	//Wrong ID, Protection Job with that ID does not exists.
	Id := int64(13377331)

	var result *models.ProtectionJob
	result, err = protectionJobs.GetProtectionJobById(Id) //This will return error.

	if err != nil {
		// Printing the Error Message with ErrorCode and ErrorMessage.
		fmt.Println(err)
		//Output:  Response status code: 404, Response: {"errorCode":"KEntityNotExistsError","message":"Protection Job Id 13377331 not found."}
	} else {
		println(result)
	}
}
