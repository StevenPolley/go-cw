package cwautomate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	ID          string `json:"Id"`
	Name        string `json:"Name"`
	Company     string `json:"Company"`
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Address1    string `json:"Address1"`
	Address2    string `json:"Address2"`
	City        string `json:"City"`
	State       string `json:"State"`
	ZipCode     string `json:"ZipCode"`
	PhoneNumber string `json:"PhoneNumber"`
	FaxNumber   string `json:"FaxNumber"`
	Comment     string `json:"Comment"`
	Country     string `json:"Country"`
	ExternalID  string `json:"ExternalId"`
}

func (site Site) GetClientByID(clientID string) (*Client, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/cwa/api/v1/Clients/%s", site.Site, clientID), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create http get request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", site.CurrentSession.TokenType, site.CurrentSession.AccessToken))
	req.Header.Set("Content-Type", "application/json")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body of request: %v", err)
	}

	client := &Client{}
	err = json.Unmarshal(respBody, client)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response to client struct: %v", err)
	}

	return client, nil
}
