package audit


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
type AUDIT_IMPL struct {
     config configuration.CONFIGURATION
}

/**
 * When actions (such as a login or a Job being paused) occur on the
 * Cohesity Cluster, the Cluster generates Audit Logs.
 * If no parameters are specified, all logs currently on the Cohesity Cluster
 * are returned. Specifying parameters filters the results that are returned.
 * @param    *bool           allUnderHierarchy     parameter: Optional
 * @param    []string        domains               parameter: Optional
 * @param    *string         search                parameter: Optional
 * @param    *int64          startIndex            parameter: Optional
 * @param    *string         outputFormat          parameter: Optional
 * @param    *string         tenantId              parameter: Optional
 * @param    []string        userNames             parameter: Optional
 * @param    []string        entityTypes           parameter: Optional
 * @param    []string        actions               parameter: Optional
 * @param    *int64          startTimeUsecs        parameter: Optional
 * @param    *int64          endTimeUsecs          parameter: Optional
 * @param    *int64          pageCount             parameter: Optional
 * @return	Returns the *models.ClusterAuditLogFilterResult response from the API call
 */
func (me *AUDIT_IMPL) SearchClusterAuditLogs (
            allUnderHierarchy *bool,
            domains []string,
            search *string,
            startIndex *int64,
            outputFormat *string,
            tenantId *string,
            userNames []string,
            entityTypes []string,
            actions []string,
            startTimeUsecs *int64,
            endTimeUsecs *int64,
            pageCount *int64) (*models.ClusterAuditLogFilterResult, error) {
    //the endpoint path uri
    _pathUrl := "/public/auditLogs/cluster"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "allUnderHierarchy" : allUnderHierarchy,
        "domains" : domains,
        "search" : search,
        "startIndex" : startIndex,
        "outputFormat" : outputFormat,
        "tenantId" : tenantId,
        "userNames" : userNames,
        "entityTypes" : entityTypes,
        "actions" : actions,
        "startTimeUsecs" : startTimeUsecs,
        "endTimeUsecs" : endTimeUsecs,
        "pageCount" : pageCount,
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
    var retVal *models.ClusterAuditLogFilterResult = &models.ClusterAuditLogFilterResult{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

