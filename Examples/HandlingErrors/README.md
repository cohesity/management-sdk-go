## Usage: 
This example is to show error handling with Cohesity Golang SDK.
```
go run protectionIdError.go
```

## Connect to the Cohesity Cluster

```
Username := "clusterUsername"
Password := "clusterPassword"
ClusterVip := "cluster.org.company.com"
Domain := "AD.org.company.com" //Initialize to 'LOCAL' for cluster user.

client, err := CohesityManagementSdk.NewCohesitySdkClient(ClusterVip, Username, Password, Domain)
if err != nil {
	fmt.Println(err) // Handle Error.
	return
}
```

## Error Avoidance.
```

1. Incorrect username/password: Response status code: 400, Response: {"errorCode":"KValidationError","message":"Could not authenticate active directory user. Username '{AD_USER}' not found on Active Directory domain '{DOMAIN}'."}
2. Incorrect Domain:  Response status code: 400, Response: {"errorCode":"KValidationError","message":"Domain '{DOMAIN}' does not exist"}
3. Incorrect IP: Post https://{junkIP}/irisservices/api/v1/public/accessTokens: dial tcp: lookup {junkIP}: no such host
4. Unreacable IP: Post https://{UnreachableIP}/irisservices/api/v1/public/accessTokens: dial tcp {UnreachableIP}: connect: operation timed out
5. Invalid URL : parse https://{InvalidURL}}/irisservices/api/v1/public/accessTokens: invalid URL escape "%3A"
6. Cluster Actions: Response status code: 404, Response: {"errorCode":"KEntityNotExistsError","message":"Protection Job Id {ID} not found."} 
```
