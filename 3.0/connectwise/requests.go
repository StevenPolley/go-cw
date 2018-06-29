package connectwise

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//BuildURL will take a REST action such as "/companies/company/5" and then append the CW site to it and return a pointer to a url.URL
func (cw *ConnectwiseSite) BuildURL(restAction string) *url.URL {
	var cwurl *url.URL
	cwurl, err := url.Parse(cw.Site)
	check(err)
	cwurl.Path += restAction

	return cwurl
}

//Checks for HTTP errors, and if all looks good, returns the body of the HTTP response as a byte slice
//TBD: Needs to accept 201 and 204 (returned for Create and Delete operations)
func getHTTPResponseBody(resp *http.Response) []byte {
	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusCreated) && (resp.StatusCode != http.StatusNoContent) {
		out := fmt.Sprintf("CW API returned HTTP Status Code %s\n%s", resp.Status, resp.Body)
		log.Fatal(out)
		return make([]byte, 0)
	}

	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	return body
}

//GetRequest takes a ConnectwiseSite and request URL, and returns the body of the response
func (cw *ConnectwiseSite) GetRequest(cwurl *url.URL) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", cwurl.String(), nil)
	check(err)
	req.Header.Set("Authorization", cw.Auth)
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	check(err)
	defer response.Body.Close()

	return getHTTPResponseBody(response)
}

//PostRequest takes a ConnectwiseSite and request URL, and returns the body of the response
func (cw *ConnectwiseSite) PostRequest(cwurl *url.URL, body io.Reader) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("POST", cwurl.String(), body)
	check(err)
	req.Header.Set("Authorization", cw.Auth)
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	check(err)
	defer response.Body.Close()

	return getHTTPResponseBody(response)
}

//DeleteRequest takes a ConnectwiseSite and request URL, and returns the body of the response
func (cw *ConnectwiseSite) DeleteRequest(cwurl *url.URL) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", cwurl.String(), nil)
	check(err)
	req.Header.Set("Authorization", cw.Auth)
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	check(err)
	defer response.Body.Close()

	return getHTTPResponseBody(response)
}
