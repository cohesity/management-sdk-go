// Copyright 2019 Cohesity Inc.

//Usage:
//go run unresolvedAlertsWorkflow.go
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

	client, err := CohesityManagementSdk.NewCohesitySdkClient(ClusterVip, Username, Password, Domain)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	idList := RecentUnresolvedAlertList(client)
	print("Specify one among AlertID to Resolve: ")
	var inpurtAlertID string
	fmt.Scanf("%s", &inpurtAlertID)
	if !CheckIDexits(inpurtAlertID, idList) {
		println("The above input ID does not exists in the cluster.")
		return
	}
	ResolveAlert(client, inpurtAlertID)
}

func RecentUnresolvedAlertList(client CohesityManagementSdk.COHESITYMANAGEMENTSDK) []*string {
	alerts := client.Alerts()
	var MaxAlerts int64 = 100
	var allUnderHierarchy *bool
	var alertIDList []string
	var propertyValue *string
	var alertStateList []models.AlertStateListEnum
	var resolutionIDList []int64
	var tenantIds []string
	var alertTypeList []int64
	var propertyKey *string
	var startDateUsecs *int64
	var endDateUsecs *int64
	var idList []*string
	var alertCategoryList []models.AlertCategoryListGetAlertsEnum
	var alertSeverityList []models.AlertSeverityListEnum
	unresolvedAlertCount := 10

	result, err := alerts.GetAlerts(MaxAlerts, allUnderHierarchy, alertIDList, alertCategoryList, propertyKey, propertyValue,
		endDateUsecs, alertSeverityList, tenantIds, alertTypeList, startDateUsecs, alertStateList, resolutionIDList)
	if err != nil {
		println(err.Error())
	} else {
		if len(result) == 0 {
			println("There are no Alerts in the cluster.")
		} else {
			for _, alert := range result {
				if unresolvedAlertCount == 0 {
					break
				}
				if alert.AlertState == models.AlertState_KOPEN {
					id := string(*alert.Id)
					idList = append(idList, &id)
					fmt.Printf("Unresolved AlertID: %v , AlertName: %v \n", id, models.AlertCategoryEnumToValue(alert.AlertCategory))
					unresolvedAlertCount--
				}

			}
		}
	}
	if len(idList) == 0 {
		fmt.Printf("All Alerts in the cluster are resolved!!")
	}
	return idList
}

func ResolveAlert(client CohesityManagementSdk.COHESITYMANAGEMENTSDK, alertID string) {

	Body := &models.AlertResolutionRequest{}
	Body.AlertIdList = &[]string{alertID}
	Body.ResolutionDetails = models.AlertResolutionInfo{}
	resolutionSummary := "Resolving this Alert through GoLang Automation."
	Body.ResolutionDetails.ResolutionSummary = &resolutionSummary
	alerts := client.Alerts()
	var err error
	var result *models.AlertResolution
	result, err = alerts.CreateResolution(Body)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Printf("Alert ID %v is resolved with resolution ID %v \n", alertID, *result.ResolutionDetails.ResolutionId)
	}

}
func CheckIDexits(id string, idList []*string) bool {
	for _, ele := range idList {
		if id == *ele {
			return true
		}
	}
	return false
}
