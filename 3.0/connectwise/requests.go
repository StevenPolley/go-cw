package connectwise

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//Request is a struct which holds all information that is collected in order to make a request
type Request struct {
	CW         *Site
	RestAction string
	Method     string     //GET, POST, DELETE, etc
	Body       []byte     //In a GET request, this is an instance of the struct that the expected json data is to be unmarshaled into
	URLValues  url.Values //Parameters sent to CW for filtering by conditions that utilize strings
	Page       int
	PageSize   int
}

//NewRequest is a function which takes the mandatory fields to perform a request to the CW API and returns a pointer to a Request struct
func NewRequest(cw *Site, restAction, method string, body []byte) *Request {
	req := Request{CW: cw, RestAction: restAction, Method: method, Body: body}
	req.URLValues = url.Values{}
	return &req
}

//Do is a method of the Request struct which uses the data contained within the Request instance to perform an HTTP request to ConnectWise
func (req *Request) Do() error {
	cwurl, err := req.CW.BuildURL(req.RestAction)
	if err != nil {
		return fmt.Errorf("could not build url %s: %s", req.RestAction, err)
	}

	//If pagination parameters are not being specified, then don't include them in the URL - not all endpoints will accept page and pagesize, they will 400 - bad request
	if req.Page == 0 || req.PageSize == 0 {
		cwurl.RawQuery = req.URLValues.Encode()
	} else {
		cwurl.RawQuery = fmt.Sprintf("%s&page=%d&pageSize=%d", req.URLValues.Encode(), req.Page, req.PageSize)
	}

	client := &http.Client{}
	jsonBuffer := bytes.NewReader(req.Body)
	httpreq, err := http.NewRequest(req.Method, cwurl.String(), jsonBuffer)
	if err != nil {
		return fmt.Errorf("could not construct http request: %s", err)
	}
	httpreq.Header.Set("Authorization", req.CW.Auth)
	httpreq.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(httpreq)
	if err != nil {
		return fmt.Errorf("could not perform http %s request: %s", req.Method, err)
	}
	req.Body, err = getHTTPResponseBody(resp)
	if err != nil {
		return fmt.Errorf("failed to get http response body: %s", err)
	}

	return nil
}

//BuildURL will take a REST action such as "/companies/company/5" and then append the CW site to it and return a pointer to a url.URL
func (cw *Site) BuildURL(restAction string) (*url.URL, error) {
	var cwurl *url.URL
	cwurl, err := url.Parse(cw.Site)
	if err != nil {
		return nil, fmt.Errorf("could not %s as url: %s", cw.Site, err)
	}
	cwurl.Path += restAction

	return cwurl, nil
}

//Checks for HTTP errors, and if all looks good, returns the body of the HTTP response as a byte slice
//TBD: Needs to accept 201 and 204 (returned for Create and Delete operations)
func getHTTPResponseBody(resp *http.Response) ([]byte, error) {
	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusCreated) && (resp.StatusCode != http.StatusNoContent) {
		return nil, fmt.Errorf("cw api returned unexpected http status %d - %s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body of request")
	}

	return body, nil
}
