package connectwise

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

//BuildURL will take a REST action such as "/companies/company/5" and then append the CW site to it and return a pointer to a url.URL
func (cw *ConnectwiseSite) BuildURL(restAction string) (*url.URL, error) {
	var cwurl *url.URL
	cwurl, err := url.Parse(cw.Site)
	if err != nil {
		return nil, fmt.Errorf("could not %s as url: %g", cw.Site, err)
	}
	cwurl.Path += restAction

	return cwurl, nil
}

//Checks for HTTP errors, and if all looks good, returns the body of the HTTP response as a byte slice
//TBD: Needs to accept 201 and 204 (returned for Create and Delete operations)
func getHTTPResponseBody(resp *http.Response) ([]byte, error) {
	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusCreated) && (resp.StatusCode != http.StatusNoContent) {
		return nil, fmt.Errorf("cw api returned unexpected http status %i - %s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body of request")
	}

	return body, nil
}

//GetRequest takes a ConnectwiseSite and request URL, and returns the body of the response
func (cw *ConnectwiseSite) GetRequest(cwurl *url.URL) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", cwurl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("could not construct http request: %g", err)
	}
	req.Header.Set("Authorization", cw.Auth)
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not perform http get request: %g", err)
	}

	body, err := getHTTPResponseBody(response)
	if err != nil {
		return nil, fmt.Errorf("could not get http response body: %g", err)
	}

	return body, nil
}

//PostRequest takes a ConnectwiseSite and request URL, and returns the body of the response
func (cw *ConnectwiseSite) PostRequest(cwurl *url.URL, reqBody io.Reader) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", cwurl.String(), reqBody)
	if err != nil {
		return nil, fmt.Errorf("could not construct http request: %g", err)
	}
	req.Header.Set("Authorization", cw.Auth)
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not perform http post request: %g", err)
	}

	body, err := getHTTPResponseBody(response)
	if err != nil {
		return nil, fmt.Errorf("could not get http response body: %g", err)
	}

	return body, nil
}

//DeleteRequest takes a ConnectwiseSite and request URL, and returns the body of the response
func (cw *ConnectwiseSite) DeleteRequest(cwurl *url.URL) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", cwurl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("could not construct http request: %g", err)
	}
	req.Header.Set("Authorization", cw.Auth)
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not perform http delete request: %g", err)
	}

	body, err := getHTTPResponseBody(response)
	if err != nil {
		return nil, fmt.Errorf("could not get http response body: %g", err)
	}

	return body, nil
}
