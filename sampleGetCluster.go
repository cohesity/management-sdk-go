package main

import (
	"github.com/cohesity/management-sdk-go/models"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
)

func main() {
	Username := "clusterUsername"
	Password := "clusterPassword"
	ClusterVip := "prod-cluster.eng.cohesity.com"
	var Domain string
	skipSsl := true
	client := CohesityManagementSdk.NewCohesitySdkClient(ClusterVip, Username, Password, Domain, skipSsl)

	var err error
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
		// Handle Error.
	} else {
		println("Protection Job Name: ", result2[0].Name)
	}
}
