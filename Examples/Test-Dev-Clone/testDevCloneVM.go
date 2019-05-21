// Copyright 2019 Cohesity Inc.

//Usage:
//go run testDevCloneVM.go

package main

import (
	"fmt"
	"strings"
	"time"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	"github.com/cohesity/management-sdk-go/models"
	"github.com/cohesity/management-sdk-go/restoretasks"
)

func main() {
	Username := "clusterUsername"
	Password := "clusterPassword"
	ClusterVip := "cluster.org.company.com"
	Domain := "AD.org.company.com" //Initialize to 'LOCAL' for cluster user
	vmToClone := "vmName"
	var err error
	client, err := CohesityManagementSdk.NewCohesitySdkClient(ClusterVip, Username, Password, Domain)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	restore := client.RestoreTasks()

	// Get the JobID of the view.
	var result *models.ObjectSearchResults
	var jobIds []int64
	var registeredSourceIds []int64
	var environments []models.EnvironmentSearchObjectsEnum
	var startTimeUsecs *int64
	var application *string
	var ownerEntityId *int64
	var viewBoxIds []int64
	var endTimeUsecs *int64
	var startIndex *int64
	var pageCount *int64
	var operatingSystems []string

	// Search the Protected VM.
	search := &vmToClone
	environments = []models.EnvironmentSearchObjectsEnum{models.EnvironmentSearchObjects_KVMWARE}
	result, err = restore.SearchObjects(startIndex, operatingSystems, application, ownerEntityId, jobIds, viewBoxIds, environments, endTimeUsecs, pageCount, search, registeredSourceIds, startTimeUsecs)

	if err != nil {
		// Printing the Error Message with ErrorCode and ErrorMessage.
		fmt.Println(err)
	}

	//Clone Params.
	jobId := result.ObjectSnapshotInfo[0].JobId
	resourcePoolId := int64(100)
	targetViewName := "view-name"
	cloneTaskName := concat("CloneTask-", time.Now().Format(time.RFC3339))
	prefix := "Clone-"
	isPoweredOn := true

	body := &models.CloneTaskRequest{}
	body.Name = cloneTaskName
	body.Type = models.TypeCloneTaskRequest_KCLONEVMS
	body.TargetViewName = &targetViewName
	body.VmwareParameters = models.VmwareCloneParameters{}
	body.VmwareParameters.ResourcePoolId = &resourcePoolId
	body.VmwareParameters.PoweredOn = &isPoweredOn
	body.VmwareParameters.Prefix = &prefix

	body.Objects = make([]*models.RestoreObjectDetails, 1)
	body.Objects[0] = &models.RestoreObjectDetails{}
	body.Objects[0].JobId = jobId
	body.Objects[0].SourceName = &vmToClone
	body.Objects[0].Environment = models.EnvironmentRestoreObjectDetails_KVMWARE

	cloneTaskResult, err := restore.CreateCloneTask(body)
	if err != nil {
		fmt.Println(err)
	} else {
		pollForTaskComplete(restore, *cloneTaskResult.Id)
	}
	var input string
	fmt.Scanf("Clone VM created, Ready to destory? %s", &input)
	// Destroy Clone VM
	restore.DeletePublicDestroyCloneTask(*cloneTaskResult.Id)
	fmt.Printf("Cloned VM: %v destroyed\n", vmToClone)
}

func pollForTaskComplete(restore restoretasks.RESTORETASKS, taskID int64) {
	interval := 5
	retryTimeout := 600 //sec
	increment := 0

	for retryTimeout > 0 {
		result, _ := restore.GetRestoreTaskById(taskID)
		if result.Status == models.StatusRestoreTask_KFINISHED {
			fmt.Println("Task Clone VM Completed.")
		} else {
			time.Sleep(time.Duration(interval) * time.Second)
			increment += interval
			fmt.Printf("Task status: %v , Time elapsed retrying: %v\n", result.Status, increment)
		}
	}
}
func concat(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}
