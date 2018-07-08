package connectwise

import (
	"encoding/json"
	"fmt"
)

//Callback is a struct to hold the unmarshaled JSON data when making a call to the System API
type Callback struct {
	ID           int    `json:"id"`
	Description  string `json:"description"`
	URL          string `json:"url"`
	ObjectID     int    `json:"objectId"`
	Type         string `json:"type"`
	Level        string `json:"level"`
	MemberID     int    `json:"memberId"`
	InactiveFlag bool   `json:"inactiveFlag"`
	Info         struct {
		LastUpdated string `json:"lastUpdated"`
		UpdatedBy   string `json:"updatedBy"`
	} `json:"_info"`
}

//GetCallbacks returns a slice of Callback structs containing all the callbacks currently registered with ConnectWise
func (cw *Site) GetCallbacks() (*[]Callback, error) {
	req := cw.NewRequest("/system/callbacks", "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	callbacks := &[]Callback{}
	err = json.Unmarshal(req.Body, callbacks)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return callbacks, nil
}

//NewCallback expects a Callback struct and will register a new callback with Connectwise
//TBD: This should return something useful, response body??
func (cw *Site) NewCallback(callback *Callback) (*Callback, error) {
	jsonCallback, err := json.Marshal(callback)
	if err != nil {
		return nil, fmt.Errorf("could not marshal json data: %s", err)
	}

	req := cw.NewRequest("/system/callbacks", "POST", jsonCallback)
	err = req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	callback = &Callback{}
	err = json.Unmarshal(req.Body, callback)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return callback, nil
}

//DeleteCallback expects the ID of an existing callback and will unregister it with ConnectWise
//Does not return anything - CW gives an empty response body
func (cw *Site) DeleteCallback(callback int) error {
	req := cw.NewRequest(fmt.Sprintf("/system/callbacks/%d", callback), "DELETE", nil)
	err := req.Do()
	if err != nil {
		return fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	return nil
}
