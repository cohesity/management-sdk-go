package restoretasks


import(
	"errors"
	"fmt"
	"encoding/json"
	"github.com/cohesity/management-sdk-go/models"
	"github.com/cohesity/management-sdk-go/unirest-go"
	"github.com/cohesity/management-sdk-go/apihelper"
	"github.com/cohesity/management-sdk-go/configuration"
)
/*
 * Client structure as interface implementation
 */
type RESTORETASKS_IMPL struct {
     config configuration.CONFIGURATION
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
 * @param    *string                               search                  parameter: Optional
 * @param    []int64                               jobIds                  parameter: Optional
 * @param    []int64                               registeredSourceIds     parameter: Optional
 * @param    []models.Environments1Enum        environments            parameter: Optional
 * @param    *int64                                startTimeUsecs          parameter: Optional
 * @param    *string                               application             parameter: Optional
 * @param    *int64                                ownerEntityId           parameter: Optional
 * @param    []int64                               viewBoxIds              parameter: Optional
 * @param    *int64                                endTimeUsecs            parameter: Optional
 * @param    *int64                                startIndex              parameter: Optional
 * @param    *int64                                pageCount               parameter: Optional
 * @param    []string                              operatingSystems        parameter: Optional
 * @return	Returns the *models.ObjectSearchResults response from the API call
 */
func (me *RESTORETASKS_IMPL) SearchObjects (
            search *string,
            jobIds []int64,
            registeredSourceIds []int64,
            environments []models.Environments1Enum,
            startTimeUsecs *int64,
            application *string,
            ownerEntityId *int64,
            viewBoxIds []int64,
            endTimeUsecs *int64,
            startIndex *int64,
            pageCount *int64,
            operatingSystems []string) (*models.ObjectSearchResults, error) {
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
        "search" : search,
        "jobIds" : jobIds,
        "registeredSourceIds" : registeredSourceIds,
        "environments" : models.Environments1EnumArrayToValue(environments),
        "startTimeUsecs" : startTimeUsecs,
        "application" : application,
        "ownerEntityId" : ownerEntityId,
        "viewBoxIds" : viewBoxIds,
        "endTimeUsecs" : endTimeUsecs,
        "startIndex" : startIndex,
        "pageCount" : pageCount,
        "operatingSystems" : operatingSystems,
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
        "user-agent" : "cohesity-Go-sdk-6.2.0",
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
 * All fields must be specified for this operation.
 * To get values for these fields, invoke the GET /public/restore/objects
 * operation.
 * A specific Job Run is defined by the jobRunId, startedTimeUsecs, and
 * attemptNumber fields.
 * @param    *int64        startedTimeUsecs         parameter: Optional
 * @param    *int64        attemptNumber            parameter: Optional
 * @param    *bool         supportedVolumesOnly     parameter: Optional
 * @param    *int64        jobId                    parameter: Optional
 * @param    *int64        jobRunId                 parameter: Optional
 * @param    *int64        sourceId                 parameter: Optional
 * @param    *int64        originalJobId            parameter: Optional
 * @param    *int64        clusterId                parameter: Optional
 * @param    *int64        clusterIncarnationId     parameter: Optional
 * @return	Returns the []*models.VMVolumeInformation response from the API call
 */
func (me *RESTORETASKS_IMPL) GetVmVolumesInformation (
            startedTimeUsecs *int64,
            attemptNumber *int64,
            supportedVolumesOnly *bool,
            jobId *int64,
            jobRunId *int64,
            sourceId *int64,
            originalJobId *int64,
            clusterId *int64,
            clusterIncarnationId *int64) ([]*models.VMVolumeInformation, error) {
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
        "startedTimeUsecs" : startedTimeUsecs,
        "attemptNumber" : attemptNumber,
        "supportedVolumesOnly" : supportedVolumesOnly,
        "jobId" : jobId,
        "jobRunId" : jobRunId,
        "sourceId" : sourceId,
        "originalJobId" : originalJobId,
        "clusterId" : clusterId,
        "clusterIncarnationId" : clusterIncarnationId,
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
        "user-agent" : "cohesity-Go-sdk-6.2.0",
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
    var retVal []*models.VMVolumeInformation
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
        "user-agent" : "cohesity-Go-sdk-6.2.0",
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
 * @return	Returns the *models.RestoreTask1 response from the API call
 */
func (me *RESTORETASKS_IMPL) GetRestoreTaskById (
            id int64) (*models.RestoreTask1, error) {
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
        "user-agent" : "cohesity-Go-sdk-6.2.0",
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
    var retVal *models.RestoreTask1 = &models.RestoreTask1{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Fetches information of virtual disk.
 * @param    int64        sourceId                 parameter: Required
 * @param    int64        clusterId                parameter: Required
 * @param    int64        clusterIncarnationId     parameter: Required
 * @param    int64        jobId                    parameter: Required
 * @param    int64        jobRunId                 parameter: Required
 * @param    int64        startTimeUsecs           parameter: Required
 * @return	Returns the []*models.SpecifiesInformationAboutTheVirtualDisk response from the API call
 */
func (me *RESTORETASKS_IMPL) GetVirtualDiskInformation (
            sourceId int64,
            clusterId int64,
            clusterIncarnationId int64,
            jobId int64,
            jobRunId int64,
            startTimeUsecs int64) ([]*models.SpecifiesInformationAboutTheVirtualDisk, error) {
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
        "sourceId" : sourceId,
        "clusterId" : clusterId,
        "clusterIncarnationId" : clusterIncarnationId,
        "jobId" : jobId,
        "jobRunId" : jobRunId,
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
        "user-agent" : "cohesity-Go-sdk-6.2.0",
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
    var retVal []*models.SpecifiesInformationAboutTheVirtualDisk
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
 * @param    []int64                            taskIds            parameter: Optional
 * @param    *int64                             startTimeUsecs     parameter: Optional
 * @param    *int64                             endTimeUsecs       parameter: Optional
 * @param    *int64                             pageCount          parameter: Optional
 * @param    []string                           taskTypes          parameter: Optional
 * @param    models.Environment2Enum        environment        parameter: Optional
 * @return	Returns the []*models.RestoreTask1 response from the API call
 */
func (me *RESTORETASKS_IMPL) GetRestoreTasks (
            taskIds []int64,
            startTimeUsecs *int64,
            endTimeUsecs *int64,
            pageCount *int64,
            taskTypes []string,
            environment models.Environment2Enum) ([]*models.RestoreTask1, error) {
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
        "environment" : models.Environment2EnumToValue(environment),
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
        "user-agent" : "cohesity-Go-sdk-6.2.0",
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
    var retVal []*models.RestoreTask1
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
 * @param    *int64          receivedTimeSeconds       parameter: Optional
 * @param    *bool           showOnlyEmailFolders      parameter: Optional
 * @param    *int64          receivedEndTime           parameter: Optional
 * @param    *string         folderName                parameter: Optional
 * @param    []int64         domainIds                 parameter: Optional
 * @param    []int64         mailboxIds                parameter: Optional
 * @param    []int64         protectionJobIds          parameter: Optional
 * @param    *string         senderAddress             parameter: Optional
 * @param    *int64          sentTimeSeconds           parameter: Optional
 * @param    *string         emailSubject              parameter: Optional
 * @param    []string        recipientAddresses        parameter: Optional
 * @param    *int64          receivedStartTime         parameter: Optional
 * @param    []string        bccRecipientAddresses     parameter: Optional
 * @param    *bool           hasAttachments            parameter: Optional
 * @param    []string        ccRecipientAddresses      parameter: Optional
 * @return	Returns the *models.FileFolderSearchResult1 response from the API call
 */
func (me *RESTORETASKS_IMPL) GetOutlookEmails (
            receivedTimeSeconds *int64,
            showOnlyEmailFolders *bool,
            receivedEndTime *int64,
            folderName *string,
            domainIds []int64,
            mailboxIds []int64,
            protectionJobIds []int64,
            senderAddress *string,
            sentTimeSeconds *int64,
            emailSubject *string,
            recipientAddresses []string,
            receivedStartTime *int64,
            bccRecipientAddresses []string,
            hasAttachments *bool,
            ccRecipientAddresses []string) (*models.FileFolderSearchResult1, error) {
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
        "receivedTimeSeconds" : receivedTimeSeconds,
        "showOnlyEmailFolders" : showOnlyEmailFolders,
        "receivedEndTime" : receivedEndTime,
        "folderName" : folderName,
        "domainIds" : domainIds,
        "mailboxIds" : mailboxIds,
        "protectionJobIds" : protectionJobIds,
        "senderAddress" : senderAddress,
        "sentTimeSeconds" : sentTimeSeconds,
        "emailSubject" : emailSubject,
        "recipientAddresses" : recipientAddresses,
        "receivedStartTime" : receivedStartTime,
        "bccRecipientAddresses" : bccRecipientAddresses,
        "hasAttachments" : hasAttachments,
        "ccRecipientAddresses" : ccRecipientAddresses,
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
        "user-agent" : "cohesity-Go-sdk-6.2.0",
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
    var retVal *models.FileFolderSearchResult1 = &models.FileFolderSearchResult1{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Get the information about snapshots that contain the specified file or folder. In addition, information about the file or folder is provided.
 * @param    int64         clusterId                parameter: Required
 * @param    int64         clusterIncarnationId     parameter: Required
 * @param    int64         sourceId                 parameter: Required
 * @param    string        filename                 parameter: Required
 * @param    int64         jobId                    parameter: Required
 * @return	Returns the []*models.FileFolderSnapshotInformation response from the API call
 */
func (me *RESTORETASKS_IMPL) GetFileSnapshotsInformation (
            clusterId int64,
            clusterIncarnationId int64,
            sourceId int64,
            filename string,
            jobId int64) ([]*models.FileFolderSnapshotInformation, error) {
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
        "clusterId" : clusterId,
        "clusterIncarnationId" : clusterIncarnationId,
        "sourceId" : sourceId,
        "filename" : filename,
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
        "user-agent" : "cohesity-Go-sdk-6.2.0",
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
    var retVal []*models.FileFolderSnapshotInformation
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
 * @param    *models.CreateRestoreTaskRequest        body     parameter: Required
 * @return	Returns the *models.RestoreTask1 response from the API call
 */
func (me *RESTORETASKS_IMPL) CreateRecoverTask (
            body *models.CreateRestoreTaskRequest) (*models.RestoreTask1, error) {
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
        "user-agent" : "cohesity-Go-sdk-6.2.0",
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
    var retVal *models.RestoreTask1 = &models.RestoreTask1{}
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
 * @param    []int64                               jobIds                  parameter: Optional
 * @param    *int64                                endTimeUsecs            parameter: Optional
 * @param    *int64                                startIndex              parameter: Optional
 * @param    []int64                               sourceIds               parameter: Optional
 * @param    *int64                                startTimeUsecs          parameter: Optional
 * @param    *int64                                pageCount               parameter: Optional
 * @param    *bool                                 folderOnly              parameter: Optional
 * @param    *string                               search                  parameter: Optional
 * @param    []int64                               registeredSourceIds     parameter: Optional
 * @param    []int64                               viewBoxIds              parameter: Optional
 * @param    []models.Environments1Enum        environments            parameter: Optional
 * @return	Returns the *models.FileFolderSearchResult1 response from the API call
 */
func (me *RESTORETASKS_IMPL) SearchRestoredFiles (
            jobIds []int64,
            endTimeUsecs *int64,
            startIndex *int64,
            sourceIds []int64,
            startTimeUsecs *int64,
            pageCount *int64,
            folderOnly *bool,
            search *string,
            registeredSourceIds []int64,
            viewBoxIds []int64,
            environments []models.Environments1Enum) (*models.FileFolderSearchResult1, error) {
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
        "jobIds" : jobIds,
        "endTimeUsecs" : endTimeUsecs,
        "startIndex" : startIndex,
        "sourceIds" : sourceIds,
        "startTimeUsecs" : startTimeUsecs,
        "pageCount" : pageCount,
        "folderOnly" : folderOnly,
        "search" : search,
        "registeredSourceIds" : registeredSourceIds,
        "viewBoxIds" : viewBoxIds,
        "environments" : models.Environments1EnumArrayToValue(environments),
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
        "user-agent" : "cohesity-Go-sdk-6.2.0",
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
    var retVal *models.FileFolderSearchResult1 = &models.FileFolderSearchResult1{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Returns the created Restore Task that deploys VMs on cloud. This operation
 * returns the target where cloud is deployed. Currently, VMs can be deployed
 * in either AWS target or Azure target.
 * @param    *models.DeployTaskRequest        body     parameter: Required
 * @return	Returns the *models.RestoreTask1 response from the API call
 */
func (me *RESTORETASKS_IMPL) CreateDeployTask (
            body *models.DeployTaskRequest) (*models.RestoreTask1, error) {
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
        "user-agent" : "cohesity-Go-sdk-6.2.0",
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
    var retVal *models.RestoreTask1 = &models.RestoreTask1{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Returns the created download Task information that downloads files and folders.
 * @param    *models.DownloadFilesAndFoldersParameters        body     parameter: Required
 * @return	Returns the *models.RestoreTask1 response from the API call
 */
func (me *RESTORETASKS_IMPL) CreateDownloadFilesAndFolders (
            body *models.DownloadFilesAndFoldersParameters) (*models.RestoreTask1, error) {
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
        "user-agent" : "cohesity-Go-sdk-6.2.0",
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
    var retVal *models.RestoreTask1 = &models.RestoreTask1{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Returns the created Restore Task that recovers files and folders.
 * @param    *models.RestoreTask        body     parameter: Required
 * @return	Returns the *models.RestoreTask1 response from the API call
 */
func (me *RESTORETASKS_IMPL) CreateRestoreFilesTask (
            body *models.RestoreTask) (*models.RestoreTask1, error) {
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
        "user-agent" : "cohesity-Go-sdk-6.2.0",
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
    var retVal *models.RestoreTask1 = &models.RestoreTask1{}
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
 * @return	Returns the *models.RestoreTask1 response from the API call
 */
func (me *RESTORETASKS_IMPL) UpdateRestoreTask (
            body *models.UpdateRestoreTaskParams) (*models.RestoreTask1, error) {
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
        "user-agent" : "cohesity-Go-sdk-6.2.0",
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
    var retVal *models.RestoreTask1 = &models.RestoreTask1{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Returns the created Restore Task.
 * @param    *models.CreateApplicationsRestoreTaskRequest        body     parameter: Required
 * @return	Returns the *models.RestoreTask1 response from the API call
 */
func (me *RESTORETASKS_IMPL) CreateApplicationsCloneTask (
            body *models.CreateApplicationsRestoreTaskRequest) (*models.RestoreTask1, error) {
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
        "user-agent" : "cohesity-Go-sdk-6.2.0",
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
    var retVal *models.RestoreTask1 = &models.RestoreTask1{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Returns the created Restore Task.
 * @param    *models.CreateApplicationsRestoreTaskRequest        body     parameter: Required
 * @return	Returns the *models.RestoreTask1 response from the API call
 */
func (me *RESTORETASKS_IMPL) CreateApplicationsRecoverTask (
            body *models.CreateApplicationsRestoreTaskRequest) (*models.RestoreTask1, error) {
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
        "user-agent" : "cohesity-Go-sdk-6.2.0",
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
    var retVal *models.RestoreTask1 = &models.RestoreTask1{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Returns the created Restore Task that clones VMs or a View.
 * @param    *models.CloneRestoreTaskRequest        body     parameter: Required
 * @return	Returns the *models.RestoreTask1 response from the API call
 */
func (me *RESTORETASKS_IMPL) CreateCloneTask (
            body *models.CloneRestoreTaskRequest) (*models.RestoreTask1, error) {
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
        "user-agent" : "cohesity-Go-sdk-6.2.0",
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
    var retVal *models.RestoreTask1 = &models.RestoreTask1{}
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
func (me *RESTORETASKS_IMPL) DestroyCloneTask (
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
        "user-agent" : "cohesity-Go-sdk-6.2.0",
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

