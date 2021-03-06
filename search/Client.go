package search


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
type SEARCH_IMPL struct {
     config configuration.CONFIGURATION
}

/**
 * Returns list of all the objects along with the protection status information.
 * @param    []int64                                         lastProtectionJobRunStatus     parameter: Optional
 * @param    []models.PhysicalServerHostTypesEnum        physicalServerHostTypes        parameter: Optional
 * @param    []string                                        registeredSourceUuids          parameter: Optional
 * @param    *int64                                          startIndex                     parameter: Optional
 * @param    *int64                                          pageCount                      parameter: Optional
 * @param    *string                                         searchString                   parameter: Optional
 * @param    []string                                        protectionStatus               parameter: Optional
 * @param    []models.Environments7Enum                  environments                   parameter: Optional
 * @return	Returns the []*models.ProtectionSourceResponse response from the API call
 */
func (me *SEARCH_IMPL) SearchProtectionSources (
            lastProtectionJobRunStatus []int64,
            physicalServerHostTypes []models.PhysicalServerHostTypesEnum,
            registeredSourceUuids []string,
            startIndex *int64,
            pageCount *int64,
            searchString *string,
            protectionStatus []string,
            environments []models.Environments7Enum) ([]*models.ProtectionSourceResponse, error) {
    //the endpoint path uri
    _pathUrl := "/public/search/protectionSources"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "lastProtectionJobRunStatus" : lastProtectionJobRunStatus,
        "physicalServerHostTypes" : models.PhysicalServerHostTypesEnumArrayToValue(physicalServerHostTypes),
        "registeredSourceUuids" : registeredSourceUuids,
        "startIndex" : startIndex,
        "pageCount" : pageCount,
        "searchString" : searchString,
        "protectionStatus" : protectionStatus,
        "environments" : models.Environments7EnumArrayToValue(environments),
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
    var retVal []*models.ProtectionSourceResponse
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Returns the information of latest snapshot of a particular object across
 * all snapshot target locations.
 * @param    string        muuid    parameter: Required
 * @return	Returns the *models.ProtectionRunResponse response from the API call
 */
func (me *SEARCH_IMPL) SearchProtectionRuns (
            muuid string) (*models.ProtectionRunResponse, error) {
    //the endpoint path uri
    _pathUrl := "/public/search/protectionRuns"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "uuid" : muuid,
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
    var retVal *models.ProtectionRunResponse = &models.ProtectionRunResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

