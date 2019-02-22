package main

import (
	"github.com/cohesity/management-sdk-go/models"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
)

func main() {
	Username := "admin"
	Password := "admin"
	ClusterVip := "10.2.145.49"
	var Domain string
	client := CohesityManagementSdk.NewCohesitySdkClient(Username, Password, Domain)
	client.Configuration().SetClusterVip(&ClusterVip)
	client.Configuration().SetSkipSSL(true)

	client.Authorize()
	//cluster := client.Cluster()

	var err error
	//var result *models.BasicCohesityClusterInformation
	// var result *models.CohesityCluster
	// boolvalue := true
	// result, err = cluster.GetCluster(&boolvalue)

	// if err != nil {
	// 	println(err)
	// } else {
	// 	//TODO: Use result variable here
	// 	println(result)
	// }

	protectionJobs := client.ProtectionJobs()

	var result2 []*models.ProtectionJob
	var includeLastRunAndStats bool
	var policyIds []string
	var isActive bool = true
	var isDeleted bool = false
	var onlyReturnBasicSummary bool = true
	Environments := []models.EnvironmentsEnum{models.Environments_KVIEW}
	var tenantIds []string
	var allUnderHierarchy bool
	var ids []int64
	var names []string

	result2, err = protectionJobs.GetProtectionJobs(&includeLastRunAndStats, policyIds, &isActive, &isDeleted, &onlyReturnBasicSummary, Environments, tenantIds, &allUnderHierarchy, ids, names)
	if err != nil {

	} else {
		println("Protection Job Name: ", result2[0].Name)
	}
}
