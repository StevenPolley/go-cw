package cwautomate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SiteCredentials struct {
	Username string
	Password string
}

type Site struct {
	Site           string
	Credentials    SiteCredentials
	CurrentSession *Session
}

type Session struct {
	AccessToken         string `json:"AccessToken"`
	TokenType           string `json:"TokenType"`
	ExpirationDate      string `json:"ExpirationDate"`
	UserID              string `json:"UserId"`
	IsTwoFactorRequired bool   `json:"IsTwoFactorRequired"`
}

func NewSite(site, username, password string) (*Site, error) {
	credentials := SiteCredentials{Username: username, Password: password}
	currentSession := &Session{}
	automateSite := Site{Site: site, Credentials: credentials, CurrentSession: currentSession}

	err := automateSite.RefreshToken()
	if err != nil {
		return nil, fmt.Errorf("failed to obtain login token using credentials and site provided: %v", err)
	}

	return &automateSite, nil
}

// RefreshToken is responsible for updating the API token which expires every hour
func (site *Site) RefreshToken() error {
	reqBody, err := json.Marshal(site.Credentials)
	if err != nil {
		return fmt.Errorf("failed to marshal site credentials - username '%s': %v", site.Credentials.Username, err)
	}

	resp, err := http.Post(fmt.Sprintf("%s/cwa/api/v1/apitoken", site.Site), "application/json", bytes.NewReader(reqBody))
	if resp.StatusCode != 200 {
		return fmt.Errorf("expected http status '200', got '%v'", resp.StatusCode)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response body of request")
	}

	session := &Session{}
	err = json.Unmarshal(respBody, session)
	if err != nil {
		return fmt.Errorf("could not unmarshal response body into Session struct: %v", err)
	}

	site.CurrentSession = session

	return nil
}
