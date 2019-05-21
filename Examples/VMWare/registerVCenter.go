// Copyright 2019 Cohesity Inc.

//Usage:
//go run registerVCenter.go
package main

import (
	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	"github.com/cohesity/management-sdk-go/models"
)

func main() {

	// Cohesity Credentials
	Username := "clusterUsername"
	Password := "clusterPassword"
	ClusterVip := "cluster.org.company.com"
	Domain := "AD.org.company.com" //Initialize to 'LOCAL' for cluster user

	// vCenter Credentials
	vcenterIP := "vcenter-fqdn.org.company.com"
	vcenterUsername := "vcenterUsername"
	vcenterPassword := "vcenterPassword"

	client, err := CohesityManagementSdk.NewCohesitySdkClient(ClusterVip, Username, Password, Domain)
	if err != nil {
		println(err.Error())
		return
	}

	Body := &models.RegisterProtectionSourceParameters{}
	Body.Endpoint = &vcenterIP
	Body.Password = &vcenterPassword
	Body.Username = &vcenterUsername
	Body.VmwareType = models.VmwareType_KVCENTER
	Body.Environment = models.EnvironmentRegisterProtectionSourceParameters_KVMWARE

	result, err := client.ProtectionSources().CreateRegisterProtectionSource(Body)
	if err != nil {
		println(err.Error())
	} else {
		println(result)
	}
}
