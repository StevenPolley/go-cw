package connectwise

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//Checks for HTTP errors, and if all looks good, returns the body of the HTTP response as a byte slice
func getHTTPResponseBody(resp *http.Response) []byte {
	if resp.StatusCode != http.StatusOK {
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
func GetRequest(site *ConnectwiseSite, Url *url.URL) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", Url.String(), nil)
	check(err)
	req.Header.Set("Authorization", site.Auth)
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	check(err)
	defer response.Body.Close()

	return getHTTPResponseBody(response)
}
