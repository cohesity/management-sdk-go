// Copyright 2019 Cohesity Inc.
package restoretasks


import(
	"errors"
	"fmt"
	"encoding/json"
	"github.com/cohesity/management-sdk-go/unirest-go"
	"github.com/cohesity/management-sdk-go/apihelper"
	"github.com/cohesity/management-sdk-go/configuration"
	"github.com/cohesity/management-sdk-go/models"
)
/*
 * Client structure as interface implementation
 */
type RESTORETASKS_IMPL struct {
     config configuration.CONFIGURATION
}

/**
 * Returns the created Restore Task.
 * @param    *models.ApplicationsRestoreTaskRequest        body     parameter: Required
 * @return	Returns the *models.RestoreTask response from the API call
 */
func (me *RESTORETASKS_IMPL) CreateApplicationsCloneTask (
            body *models.ApplicationsRestoreTaskRequest) (*models.RestoreTask, error) {
//validating required parameters
    if (body == nil){
        return nil,errors.New("The parameter 'body' is a required parameter and cannot be nil.")
}     //the endpoint path uri
    _pathUrl := "/public/restore/applicationsClone"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.0",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, body)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models.RestoreTask = &models.RestoreTask{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Returns the created Restore Task.
 * @param    *models.ApplicationsRestoreTaskRequest        body     parameter: Required
 * @return	Returns the *models.RestoreTask response from the API call
 */
func (me *RESTORETASKS_IMPL) CreateApplicationsRecoverTask (
            body *models.ApplicationsRestoreTaskRequest) (*models.RestoreTask, error) {
//validating required parameters
    if (body == nil){
        return nil,errors.New("The parameter 'body' is a required parameter and cannot be nil.")
}     //the endpoint path uri
    _pathUrl := "/public/restore/applicationsRecover"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.0",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, body)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models.RestoreTask = &models.RestoreTask{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Returns the created Restore Task that clones VMs or a View.
 * @param    *models.CloneTaskRequest        body     parameter: Required
 * @return	Returns the *models.RestoreTask response from the API call
 */
func (me *RESTORETASKS_IMPL) CreateCloneTask (
            body *models.CloneTaskRequest) (*models.RestoreTask, error) {
//validating required parameters
    if (body == nil){
        return nil,errors.New("The parameter 'body' is a required parameter and cannot be nil.")
}     //the endpoint path uri
    _pathUrl := "/public/restore/clone"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.0",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, body)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models.RestoreTask = &models.RestoreTask{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Destroy a clone task with specified id.
 * @param    int64        id     parameter: Required
 * @return	Returns the  response from the API call
 */
func (me *RESTORETASKS_IMPL) DeletePublicDestroyCloneTask (
            id int64) (error) {
    //the endpoint path uri
    _pathUrl := "/public/restore/clone/{id}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "id" : id,
    })
    if err != nil {
        //error in template param handling
        return err
    }

    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return err
    }
     if me.config.AccessToken() == nil {
        return errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.0",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Delete(_queryBuilder, headers, nil)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return err
    }

    //returning the response
    return nil

}

/**
 * Returns the created Restore Task that deploys VMs on cloud. This operation
 * returns the target where cloud is deployed. Currently, VMs can be deployed
 * in either AWS target or Azure target.
 * @param    *models.DeployTaskRequest        body     parameter: Required
 * @return	Returns the *models.RestoreTask response from the API call
 */
func (me *RESTORETASKS_IMPL) CreateDeployTask (
            body *models.DeployTaskRequest) (*models.RestoreTask, error) {
//validating required parameters
    if (body == nil){
        return nil,errors.New("The parameter 'body' is a required parameter and cannot be nil.")
}     //the endpoint path uri
    _pathUrl := "/public/restore/deploy"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.0",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, body)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models.RestoreTask = &models.RestoreTask{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Returns the created download Task information that downloads files and folders.
 * @param    *models.DownloadFilesAndFoldersParams        body     parameter: Required
 * @return	Returns the *models.RestoreTask response from the API call
 */
func (me *RESTORETASKS_IMPL) CreateDownloadFilesAndFolders (
            body *models.DownloadFilesAndFoldersParams) (*models.RestoreTask, error) {
//validating required parameters
    if (body == nil){
        return nil,errors.New("The parameter 'body' is a required parameter and cannot be nil.")
}     //the endpoint path uri
    _pathUrl := "/public/restore/downloadFilesAndFolders"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.0",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, body)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models.RestoreTask = &models.RestoreTask{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Use the files and folders returned by this operation to populate the
 * list of files and folders to recover in the
 * POST /public/restore/files operation.
 * If no search pattern or filter parameters are specified, all files
 * and folders currently found on the Cohesity Cluster are returned.
 * Specify a search pattern or parameters to filter the results that
 * are returned.
 * The term "items" below refers to files and folders that are found
 * in the source objects (such as VMs).
 * @param    []int64                                                sourceIds               parameter: Optional
 * @param    *bool                                                  folderOnly              parameter: Optional
 * @param    []int64                                                jobIds                  parameter: Optional
 * @param    []int64                                                viewBoxIds              parameter: Optional
 * @param    *int64                                                 startTimeUsecs          parameter: Optional
 * @param    *int64                                                 startIndex              parameter: Optional
 * @param    *int64                                                 pageCount               parameter: Optional
 * @param    *string                                                search                  parameter: Optional
 * @param    []int64                                                registeredSourceIds     parameter: Optional
 * @param    []models.EnvironmentSearchRestoredFilesEnum        environments            parameter: Optional
 * @param    *int64                                                 endTimeUsecs            parameter: Optional
 * @return	Returns the *models.FileSearchResults response from the API call
 */
func (me *RESTORETASKS_IMPL) SearchRestoredFiles (
            sourceIds []int64,
            folderOnly *bool,
            jobIds []int64,
            viewBoxIds []int64,
            startTimeUsecs *int64,
            startIndex *int64,
            pageCount *int64,
            search *string,
            registeredSourceIds []int64,
            environments []models.EnvironmentSearchRestoredFilesEnum,
            endTimeUsecs *int64) (*models.FileSearchResults, error) {
    //the endpoint path uri
    _pathUrl := "/public/restore/files"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "sourceIds" : sourceIds,
        "folderOnly" : folderOnly,
        "jobIds" : jobIds,
        "viewBoxIds" : viewBoxIds,
        "startTimeUsecs" : startTimeUsecs,
        "startIndex" : startIndex,
        "pageCount" : pageCount,
        "search" : search,
        "registeredSourceIds" : registeredSourceIds,
        "environments" : models.EnvironmentSearchRestoredFilesEnumArrayToValue(environments),
        "endTimeUsecs" : endTimeUsecs,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.0",
        "accept" : "application/json",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models.FileSearchResults = &models.FileSearchResults{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Returns the created Restore Task that recovers files and folders.
 * @param    *models.RestoreFilesTaskRequest        body     parameter: Required
 * @return	Returns the *models.RestoreTask response from the API call
 */
func (me *RESTORETASKS_IMPL) CreateRestoreFilesTask (
            body *models.RestoreFilesTaskRequest) (*models.RestoreTask, error) {
//validating required parameters
    if (body == nil){
        return nil,errors.New("The parameter 'body' is a required parameter and cannot be nil.")
}     //the endpoint path uri
    _pathUrl := "/public/restore/files"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.0",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, body)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models.RestoreTask = &models.RestoreTask{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Get the information about snapshots that contain the specified file or folder. In addition, information about the file or folder is provided.
 * @param    int64         jobId                    parameter: Required
 * @param    int64         clusterId                parameter: Required
 * @param    int64         clusterIncarnationId     parameter: Required
 * @param    int64         sourceId                 parameter: Required
 * @param    string        filename                 parameter: Required
 * @return	Returns the []*models.FileSnapshotInformation response from the API call
 */
func (me *RESTORETASKS_IMPL) GetFileSnapshotsInformation (
            jobId int64,
            clusterId int64,
            clusterIncarnationId int64,
            sourceId int64,
            filename string) ([]*models.FileSnapshotInformation, error) {
    //the endpoint path uri
    _pathUrl := "/public/restore/files/snapshotsInformation"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "jobId" : jobId,
        "clusterId" : clusterId,
        "clusterIncarnationId" : clusterIncarnationId,
        "sourceId" : sourceId,
        "filename" : filename,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.0",
        "accept" : "application/json",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal []*models.FileSnapshotInformation
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * If no search pattern or filter parameters are specified, all backup objects
 * currently found on the Cohesity Cluster are returned.
 * Only leaf objects that have been protected by a Job are returned such as VMs,
 * Views and databases.
 * Specify a search pattern or parameters to filter the results that
 * are returned.
 * The term "items" below refers to leaf backup objects such as VMs,
 * Views and databases.
 * @param    *int64                                           startIndex              parameter: Optional
 * @param    []string                                         operatingSystems        parameter: Optional
 * @param    *string                                          application             parameter: Optional
 * @param    *int64                                           ownerEntityId           parameter: Optional
 * @param    []int64                                          jobIds                  parameter: Optional
 * @param    []int64                                          viewBoxIds              parameter: Optional
 * @param    []models.EnvironmentSearchObjectsEnum        environments            parameter: Optional
 * @param    *int64                                           endTimeUsecs            parameter: Optional
 * @param    *int64                                           pageCount               parameter: Optional
 * @param    *string                                          search                  parameter: Optional
 * @param    []int64                                          registeredSourceIds     parameter: Optional
 * @param    *int64                                           startTimeUsecs          parameter: Optional
 * @return	Returns the *models.ObjectSearchResults response from the API call
 */
func (me *RESTORETASKS_IMPL) SearchObjects (
            startIndex *int64,
            operatingSystems []string,
            application *string,
            ownerEntityId *int64,
            jobIds []int64,
            viewBoxIds []int64,
            environments []models.EnvironmentSearchObjectsEnum,
            endTimeUsecs *int64,
            pageCount *int64,
            search *string,
            registeredSourceIds []int64,
            startTimeUsecs *int64) (*models.ObjectSearchResults, error) {
    //the endpoint path uri
    _pathUrl := "/public/restore/objects"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "startIndex" : startIndex,
        "operatingSystems" : operatingSystems,
        "application" : application,
        "ownerEntityId" : ownerEntityId,
        "jobIds" : jobIds,
        "viewBoxIds" : viewBoxIds,
        "environments" : models.EnvironmentSearchObjectsEnumArrayToValue(environments),
        "endTimeUsecs" : endTimeUsecs,
        "pageCount" : pageCount,
        "search" : search,
        "registeredSourceIds" : registeredSourceIds,
        "startTimeUsecs" : startTimeUsecs,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.0",
        "accept" : "application/json",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models.ObjectSearchResults = &models.ObjectSearchResults{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Search for Emails and Emails' folders to recover that match the specified
 * search and filter criterias on the Cohesity cluster.
 * @param    []int64         mailboxIds                parameter: Optional
 * @param    []int64         protectionJobIds          parameter: Optional
 * @param    []string        ccRecipientAddresses      parameter: Optional
 * @param    []string        bccRecipientAddresses     parameter: Optional
 * @param    *int64          receivedTimeSeconds       parameter: Optional
 * @param    *int64          receivedStartTime         parameter: Optional
 * @param    *string         folderName                parameter: Optional
 * @param    []int64         domainIds                 parameter: Optional
 * @param    []string        recipientAddresses        parameter: Optional
 * @param    *bool           hasAttachments            parameter: Optional
 * @param    *int64          receivedEndTime           parameter: Optional
 * @param    *string         emailSubject              parameter: Optional
 * @param    *string         senderAddress             parameter: Optional
 * @param    *int64          sentTimeSeconds           parameter: Optional
 * @param    *bool           showOnlyEmailFolders      parameter: Optional
 * @return	Returns the *models.FileSearchResults response from the API call
 */
func (me *RESTORETASKS_IMPL) GetOutlookEmails (
            mailboxIds []int64,
            protectionJobIds []int64,
            ccRecipientAddresses []string,
            bccRecipientAddresses []string,
            receivedTimeSeconds *int64,
            receivedStartTime *int64,
            folderName *string,
            domainIds []int64,
            recipientAddresses []string,
            hasAttachments *bool,
            receivedEndTime *int64,
            emailSubject *string,
            senderAddress *string,
            sentTimeSeconds *int64,
            showOnlyEmailFolders *bool) (*models.FileSearchResults, error) {
    //the endpoint path uri
    _pathUrl := "/public/restore/office365/outlook/emails"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "mailboxIds" : mailboxIds,
        "protectionJobIds" : protectionJobIds,
        "ccRecipientAddresses" : ccRecipientAddresses,
        "bccRecipientAddresses" : bccRecipientAddresses,
        "receivedTimeSeconds" : receivedTimeSeconds,
        "receivedStartTime" : receivedStartTime,
        "folderName" : folderName,
        "domainIds" : domainIds,
        "recipientAddresses" : recipientAddresses,
        "hasAttachments" : hasAttachments,
        "receivedEndTime" : receivedEndTime,
        "emailSubject" : emailSubject,
        "senderAddress" : senderAddress,
        "sentTimeSeconds" : sentTimeSeconds,
        "showOnlyEmailFolders" : showOnlyEmailFolders,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.0",
        "accept" : "application/json",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models.FileSearchResults = &models.FileSearchResults{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Returns the created Restore Task. This operation returns the following
 * types of Restore Tasks: 1) A Restore Task that recovers VMs back to the
 * original location or a new location. 2) A Restore Task that mounts the
 * volumes of a Server (such as a VM or Physical Server) onto the specified
 * target system. The Snapshots of the Server that contains the volumes
 * that are mounted is determined by Array of Objects.
 * The content of the Server is available from the mount point
 * for the Granular Level Recovery (GLR) of application data. For example
 * recovering Microsoft Exchange data using Kroll Ontrack® PowerControls™.
 * NOTE: Volumes are mounted "instantly" if the Snapshot is stored locally on the
 * Cohesity Cluster. If the Snapshot is archival target, it will take longer
 * because it must be retrieved.
 * @param    *models.RecoverTaskRequest        body     parameter: Required
 * @return	Returns the *models.RestoreTask response from the API call
 */
func (me *RESTORETASKS_IMPL) CreateRecoverTask (
            body *models.RecoverTaskRequest) (*models.RestoreTask, error) {
//validating required parameters
    if (body == nil){
        return nil,errors.New("The parameter 'body' is a required parameter and cannot be nil.")
}     //the endpoint path uri
    _pathUrl := "/public/restore/recover"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.0",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, body)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models.RestoreTask = &models.RestoreTask{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Updates an existing restore task with additional params which are needed for
 * features like Hot-Standby.
 * @param    *models.UpdateRestoreTaskParams        body     parameter: Required
 * @return	Returns the *models.RestoreTask response from the API call
 */
func (me *RESTORETASKS_IMPL) UpdateRestoreTask (
            body *models.UpdateRestoreTaskParams) (*models.RestoreTask, error) {
//validating required parameters
    if (body == nil){
        return nil,errors.New("The parameter 'body' is a required parameter and cannot be nil.")
}     //the endpoint path uri
    _pathUrl := "/public/restore/recover"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.0",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Put(_queryBuilder, headers, body)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models.RestoreTask = &models.RestoreTask{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * If no parameters are specified, all Restore Tasks found
 * on the Cohesity Cluster are returned. Both running and completed
 * Restore Tasks are reported.
 * Specifying parameters filters the results that are returned.
 * @param    []int64                                          taskIds            parameter: Optional
 * @param    *int64                                           startTimeUsecs     parameter: Optional
 * @param    *int64                                           endTimeUsecs       parameter: Optional
 * @param    *int64                                           pageCount          parameter: Optional
 * @param    []string                                         taskTypes          parameter: Optional
 * @param    models.EnvironmentGetRestoreTasksEnum        environment        parameter: Optional
 * @return	Returns the []*models.RestoreTask response from the API call
 */
func (me *RESTORETASKS_IMPL) GetRestoreTasks (
            taskIds []int64,
            startTimeUsecs *int64,
            endTimeUsecs *int64,
            pageCount *int64,
            taskTypes []string,
            environment models.EnvironmentGetRestoreTasksEnum) ([]*models.RestoreTask, error) {
    //the endpoint path uri
    _pathUrl := "/public/restore/tasks"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "taskIds" : taskIds,
        "startTimeUsecs" : startTimeUsecs,
        "endTimeUsecs" : endTimeUsecs,
        "pageCount" : pageCount,
        "taskTypes" : taskTypes,
        "environment" : models.EnvironmentGetRestoreTasksEnumToValue(environment),
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.0",
        "accept" : "application/json",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal []*models.RestoreTask
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Cancel a recover or clone task with specified id.
 * @param    int64        id     parameter: Required
 * @return	Returns the  response from the API call
 */
func (me *RESTORETASKS_IMPL) UpdateCancelRestoreTask (
            id int64) (error) {
    //the endpoint path uri
    _pathUrl := "/public/restore/tasks/cancel/{id}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "id" : id,
    })
    if err != nil {
        //error in template param handling
        return err
    }

    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return err
    }
     if me.config.AccessToken() == nil {
        return errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.0",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Put(_queryBuilder, headers, nil)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return err
    }

    //returning the response
    return nil

}

/**
 * Returns the Restore Task corresponding to the specified Restore Task id.
 * @param    int64        id     parameter: Required
 * @return	Returns the *models.RestoreTask response from the API call
 */
func (me *RESTORETASKS_IMPL) GetRestoreTaskById (
            id int64) (*models.RestoreTask, error) {
    //the endpoint path uri
    _pathUrl := "/public/restore/tasks/{id}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "id" : id,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.0",
        "accept" : "application/json",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models.RestoreTask = &models.RestoreTask{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Fetches information of virtual disk.
 * @param    int64        jobRunId                 parameter: Required
 * @param    int64        startTimeUsecs           parameter: Required
 * @param    int64        sourceId                 parameter: Required
 * @param    int64        clusterId                parameter: Required
 * @param    int64        clusterIncarnationId     parameter: Required
 * @param    int64        jobId                    parameter: Required
 * @return	Returns the []*models.VirtualDiskInformation response from the API call
 */
func (me *RESTORETASKS_IMPL) GetVirtualDiskInformation (
            jobRunId int64,
            startTimeUsecs int64,
            sourceId int64,
            clusterId int64,
            clusterIncarnationId int64,
            jobId int64) ([]*models.VirtualDiskInformation, error) {
    //the endpoint path uri
    _pathUrl := "/public/restore/virtualDiskInformation"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "jobRunId" : jobRunId,
        "startTimeUsecs" : startTimeUsecs,
        "sourceId" : sourceId,
        "clusterId" : clusterId,
        "clusterIncarnationId" : clusterIncarnationId,
        "jobId" : jobId,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.0",
        "accept" : "application/json",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal []*models.VirtualDiskInformation
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * All fields must be specified for this operation.
 * To get values for these fields, invoke the GET /public/restore/objects
 * operation.
 * A specific Job Run is defined by the jobRunId, startedTimeUsecs, and
 * attemptNumber fields.
 * @param    *int64        jobId                    parameter: Optional
 * @param    *int64        clusterId                parameter: Optional
 * @param    *int64        clusterIncarnationId     parameter: Optional
 * @param    *int64        attemptNumber            parameter: Optional
 * @param    *int64        sourceId                 parameter: Optional
 * @param    *int64        jobRunId                 parameter: Optional
 * @param    *int64        startedTimeUsecs         parameter: Optional
 * @param    *int64        originalJobId            parameter: Optional
 * @param    *bool         supportedVolumesOnly     parameter: Optional
 * @return	Returns the []*models.VmVolumesInformation response from the API call
 */
func (me *RESTORETASKS_IMPL) GetVmVolumesInformation (
            jobId *int64,
            clusterId *int64,
            clusterIncarnationId *int64,
            attemptNumber *int64,
            sourceId *int64,
            jobRunId *int64,
            startedTimeUsecs *int64,
            originalJobId *int64,
            supportedVolumesOnly *bool) ([]*models.VmVolumesInformation, error) {
    //the endpoint path uri
    _pathUrl := "/public/restore/vms/volumesInformation"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "jobId" : jobId,
        "clusterId" : clusterId,
        "clusterIncarnationId" : clusterIncarnationId,
        "attemptNumber" : attemptNumber,
        "sourceId" : sourceId,
        "jobRunId" : jobRunId,
        "startedTimeUsecs" : startedTimeUsecs,
        "originalJobId" : originalJobId,
        "supportedVolumesOnly" : supportedVolumesOnly,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.0",
        "accept" : "application/json",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal []*models.VmVolumesInformation
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

