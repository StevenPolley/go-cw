package connectwise

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//Site is a stuct containing the URL of the site and the API authorization token in the format that CW expects it.
type Site struct {
	Site           string
	AuthAPIKey     string //Preferable authentication method
	CompanyName    string //Used for user impersonation, but collected for API key as well so it can be accessed publicly later on if required
	AuthUsername   string // User for user impersonation
	AuthMemberHash string //Used for user impersonation
	ClientID       string // Required as of 2019.02 https://developer.connectwise.com/ClientID

}

//Count is a struct used for unmarshalling JSON data when using the Count endpoints in Connectwise (eg: counting number of companies)
type Count struct {
	Count int `json:"count"`
}

//NewSite returns a pointer to a ConnectwiseSite struct with the site and auth string available for use in API requests
func NewSite(site, publicKey, privateKey, company, clientID string) *Site {
	//The auth string must be formatted in this way when used in requests to the API
	authString := fmt.Sprintf("%s+%s:%s", company, publicKey, privateKey)
	authString = base64.StdEncoding.EncodeToString([]byte(authString))
	authString = fmt.Sprintf("Basic %s", authString)

	cwSite := Site{Site: site, AuthAPIKey: authString, CompanyName: company, ClientID: clientID}

	return &cwSite
}

//NewSiteUserImpersonation is similar to NewSite but is used for user impersonation and instead of an API key takes the username and password
func NewSiteUserImpersonation(site, username, password, company, clientID string) (*Site, error) {

	//We must retrieve a user hash which is good for 4 hours
	authBaseURL := strings.TrimSuffix(site, "/apis/3.0")
	authURL, err := url.Parse(authBaseURL)
	if err != nil {
		return nil, fmt.Errorf("could not build url %s: %s", authBaseURL, err)
	}
	authURL.Path += "/login/login.aspx"

	form := url.Values{}
	form.Add("username", username)
	form.Add("password", password)
	form.Add("companyname", company)

	client := &http.Client{}

	authReq, err := http.NewRequest("POST", authURL.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("could not construct http request: %s", err)
	}

	authReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(authReq)
	if err != nil {
		return nil, fmt.Errorf("could not perform http authentication request: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cw api returned unexpected http status %d - %s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body of request")
	}

	if string(body) == "" {
		return nil, fmt.Errorf("could not authenticate with connectwise as %s: authentication request sent to connectwise, but response body was empty", username)
	}

	if string(body) == "CompanyFAIL" {
		return nil, fmt.Errorf("could not authenticate with connectwise as %s: connectwise response body is %s", username, string(body))
	}

	cwSite := Site{Site: site, AuthUsername: username, AuthMemberHash: string(body), CompanyName: company, ClientID: clientID}

	return &cwSite, nil
}
