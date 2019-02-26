package protectionruns


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
type PROTECTIONRUNS_IMPL struct {
     config configuration.CONFIGURATION
}

/**
 * Cancel a Protection Job run.
 * @param    int64                                      id       parameter: Required
 * @param    *models.CancelAProtectionJobRun        body     parameter: Optional
 * @return	Returns the  response from the API call
 */
func (me *PROTECTIONRUNS_IMPL) CreateCancelProtectionJobRun (
            id int64,
            body *models.CancelAProtectionJobRun) (error) {
    //the endpoint path uri
    _pathUrl := "/public/protectionRuns/cancel/{id}"

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
        "content-type" : "application/json; charset=utf-8",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, body)
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
 * If no parameters are specified, Job Runs currently
 * on the Cohesity Cluster are returned. Both running and completed Job Runs
 * are reported.
 * Specifying parameters filters the results that are returned.
 * @param    *int64          startTimeUsecs                parameter: Optional
 * @param    *bool           excludeTasks                  parameter: Optional
 * @param    *int64          sourceId                      parameter: Optional
 * @param    *int64          jobId                         parameter: Optional
 * @param    *int64          endTimeUsecs                  parameter: Optional
 * @param    *int64          numRuns                       parameter: Optional
 * @param    []string        runTypes                      parameter: Optional
 * @param    *bool           excludeErrorRuns              parameter: Optional
 * @param    *bool           excludeNonRestoreableRuns     parameter: Optional
 * @param    *int64          startedTimeUsecs              parameter: Optional
 * @return	Returns the []*models.ProtectionJobRunInstance response from the API call
 */
func (me *PROTECTIONRUNS_IMPL) GetProtectionRuns (
            startTimeUsecs *int64,
            excludeTasks *bool,
            sourceId *int64,
            jobId *int64,
            endTimeUsecs *int64,
            numRuns *int64,
            runTypes []string,
            excludeErrorRuns *bool,
            excludeNonRestoreableRuns *bool,
            startedTimeUsecs *int64) ([]*models.ProtectionJobRunInstance, error) {
    //the endpoint path uri
    _pathUrl := "/public/protectionRuns"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "startTimeUsecs" : startTimeUsecs,
        "excludeTasks" : excludeTasks,
        "sourceId" : sourceId,
        "jobId" : jobId,
        "endTimeUsecs" : endTimeUsecs,
        "numRuns" : numRuns,
        "runTypes" : runTypes,
        "excludeErrorRuns" : excludeErrorRuns,
        "excludeNonRestoreableRuns" : excludeNonRestoreableRuns,
        "startedTimeUsecs" : startedTimeUsecs,
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
    var retVal []*models.ProtectionJobRunInstance
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Update the expiration date (retention period) for the specified Protection
 * Job Runs and their snapshots.
 * After an expiration time is reached, the Job Run and its snapshots are deleted.
 * If an expiration time of 0 is specified, a Job Run and its snapshots
 * are immediately deleted.
 * @param    *models.UpdateProtectionJobRunsParameters        body     parameter: Required
 * @return	Returns the  response from the API call
 */
func (me *PROTECTIONRUNS_IMPL) UpdateProtectionRuns (
            body *models.UpdateProtectionJobRunsParameters) (error) {
//validating required parameters
    if (body == nil){
        return errors.New("The parameter 'body' is a required parameter and cannot be nil.")
}     //the endpoint path uri
    _pathUrl := "/public/protectionRuns"

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
        return err
    }
     if me.config.AccessToken() == nil {
        return errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-6.2.0",
        "content-type" : "application/json; charset=utf-8",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Put(_queryBuilder, headers, body)
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

