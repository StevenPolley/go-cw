package connectwise

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Checks for HTTP errors, and if all looks good, returns the body of the HTTP response as a byte slice
func getHTTPResponseBody(resp *http.Response) (body []byte) {
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
