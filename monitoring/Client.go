// Copyright 2019 Cohesity Inc.
package monitoring


import(
	"errors"
	"fmt"
	"github.com/cohesity/management-sdk-go/unirest-go"
	"github.com/cohesity/management-sdk-go/apihelper"
	"github.com/cohesity/management-sdk-go/configuration"
)
/*
 * Client structure as interface implementation
 */
type MONITORING_IMPL struct {
     config configuration.CONFIGURATION
}

/**
 * Specifying parameters can alter the results that are returned.
 * @param    int64           startTime     parameter: Required
 * @param    int64           endTime       parameter: Required
 * @param    []string        jobTypes      parameter: Optional
 * @param    *int64          page          parameter: Optional
 * @param    *int64          pageSize      parameter: Optional
 * @return	Returns the  response from the API call
 */
func (me *MONITORING_IMPL) GetAllJobRuns (
            startTime int64,
            endTime int64,
            jobTypes []string,
            page *int64,
            pageSize *int64) (error) {
    //the endpoint path uri
    _pathUrl := "/public/monitoring/jobs"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "startTime" : startTime,
        "endTime" : endTime,
        "jobTypes" : jobTypes,
        "page" : page,
        "pageSize" : pageSize,
    })
    if err != nil {
        //error in query param handling
        return err
    }

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
        "user-agent" : "cohesity-Go-sdk-1.1.2",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
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

