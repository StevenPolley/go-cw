package connectwise

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func (cw *ConnectwiseSite) BuildUrl(restAction string) *url.URL {
	var Url *url.URL
	Url, err := url.Parse(cw.Site)
	check(err)
	Url.Path += restAction

	return Url
}

//Checks for HTTP errors, and if all looks good, returns the body of the HTTP response as a byte slice
//TBD: Needs to accept 201 and 204 (returned for Create and Delete operations)
func getHTTPResponseBody(resp *http.Response) []byte {
	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusCreated) && (resp.StatusCode != http.StatusNoContent) {
		out := fmt.Sprintf("CW API returned HTTP Status Code %s\n%s", resp.Status, resp.Body)
		log.Fatal(out)
		return make([]byte, 0)
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		check(err)

		return body
	}
}

//Takes a ConnectwiseSite and request URL, and returns the body of the response
func (cw *ConnectwiseSite) GetRequest(Url *url.URL) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", Url.String(), nil)
	check(err)
	req.Header.Set("Authorization", cw.Auth)
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	check(err)
	defer response.Body.Close()

	return getHTTPResponseBody(response)
}

//Takes a ConnectwiseSite and request URL, and returns the body of the response
func (cw *ConnectwiseSite) PostRequest(Url *url.URL, body io.Reader) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("POST", Url.String(), body)
	check(err)
	req.Header.Set("Authorization", cw.Auth)
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	check(err)
	defer response.Body.Close()

	return getHTTPResponseBody(response)
}

//Takes a ConnectwiseSite and request URL, and returns the body of the response
func (cw *ConnectwiseSite) DeleteRequest(Url *url.URL) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", Url.String(), nil)
	check(err)
	req.Header.Set("Authorization", cw.Auth)
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	check(err)
	defer response.Body.Close()

	return getHTTPResponseBody(response)
}
