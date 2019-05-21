// Copyright 2019 Cohesity Inc.

//Usage:
//go run listProtectionJobs.go

package main

import (
	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	"github.com/cohesity/management-sdk-go/models"
)

func main() {

	Username := "clusterUsername"
	Password := "clusterPassword"
	ClusterVip := "cluster.org.company.com"
	Domain := "AD.org.company.com" //Initialize to 'LOCAL' for cluster user.

	var err2 error
	client, err2 := CohesityManagementSdk.NewCohesitySdkClient(ClusterVip, Username, Password, Domain)
	if err2 != nil {
		println(err2.Error())
		return
	}
	var err error
	protectionJobs := client.ProtectionJobs()
	/*
		var result []*models.ProtectionJob
		var includeLastRunAndStats bool
		var policyIds []string
		var isActive bool = true
		var isDeleted bool = false
		var onlyReturnBasicSummary bool = true
		Environments := []models.EnvironmentsEnum{}
		var tenantIds []string
		var allUnderHierarchy bool
		var ids []int64
		var names []string
	*/
	var onlyReturnBasicSummary *bool
	var includeRpoSnapshots *bool
	var ids []int64
	var names []string
	var policyIds []string
	var isActive *bool
	var includeLastRunAndStats *bool
	var onlyReturnDataMigrationJobs *bool
	var tenantIds []string
	var allUnderHierarchy *bool
	var environments []models.EnvironmentGetProtectionJobsEnum
	var isDeleted *bool
	result, err := protectionJobs.GetProtectionJobs(onlyReturnBasicSummary, includeRpoSnapshots, ids, names, policyIds,
		isActive, includeLastRunAndStats, onlyReturnDataMigrationJobs,
		tenantIds, allUnderHierarchy, environments, isDeleted)

	if err != nil {
		println(err.Error())
	} else {
		if len(result) == 0 {
			println("There are no Protection Jobs in the cluster.")
		} else {
			println("Protection Job Names:  ")
			for i, pJob := range result {
				println(i+1, ".", pJob.Name)
			}
		}
	}
}
