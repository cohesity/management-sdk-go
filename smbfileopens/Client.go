package smbfileopens


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
type SMBFILEOPENS_IMPL struct {
     config configuration.CONFIGURATION
}

/**
 * If no parameters are specified, all active SMB file opens currently on the
 * Cohesity Cluster are returned. Specifying parameters filters the results that
 * are returned.
 * @param    *string        cookie        parameter: Optional
 * @param    *string        filePath      parameter: Optional
 * @param    *string        viewName      parameter: Optional
 * @param    *int64         pageCount     parameter: Optional
 * @return	Returns the *models.SMBActiveFileOpenResponse response from the API call
 */
func (me *SMBFILEOPENS_IMPL) GetSmbFileOpens (
            cookie *string,
            filePath *string,
            viewName *string,
            pageCount *int64) (*models.SMBActiveFileOpenResponse, error) {
    //the endpoint path uri
    _pathUrl := "/public/smbFileOpens"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "cookie" : cookie,
        "filePath" : filePath,
        "viewName" : viewName,
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
    var retVal *models.SMBActiveFileOpenResponse = &models.SMBActiveFileOpenResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Returns nothing upon success.
 * @param    *models.CloseSMBFileOpenParameters        body     parameter: Required
 * @return	Returns the  response from the API call
 */
func (me *SMBFILEOPENS_IMPL) CreateCloseSmbFileOpen (
            body *models.CloseSMBFileOpenParameters) (error) {
//validating required parameters
    if (body == nil){
        return errors.New("The parameter 'body' is a required parameter and cannot be nil.")
}     //the endpoint path uri
    _pathUrl := "/public/smbFileOpens"

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

