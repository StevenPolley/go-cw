package connectwise

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Callback struct {
	ID           int
	Description  string
	URL          string
	ObjectId     int
	Type         string
	Level        string
	MemberId     int
	InactiveFlag bool
}

func (cw *ConnectwiseSite) GetCallbacks() (*[]Callback, error) {
	restAction := "/system/callbacks"
	cwurl, err := cw.BuildURL(restAction)
	if err != nil {
		return nil, fmt.Errorf("could not build url %s: %s", restAction, err)
	}

	body, err := cw.GetRequest(cwurl)
	if err != nil {
		return nil, fmt.Errorf("could not get request %s: %s", cwurl, err)
	}
	callbacks := []Callback{}
	err = json.Unmarshal(body, &callbacks)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return &callbacks, nil

}

//TBD: This should return something useful, response body??
func (cw *ConnectwiseSite) NewCallback(callback Callback) error {
	restAction := "/system/callbacks"
	cwurl, err := cw.BuildURL(restAction)
	if err != nil {
		return fmt.Errorf("could not build url %s: %s", restAction, err)
	}

	jsonCallback, err := json.Marshal(callback)
	if err != nil {
		return fmt.Errorf("could not marshal json data: %s", err)
	}

	jsonBuffer := bytes.NewReader(jsonCallback)

	_, err = cw.PostRequest(cwurl, jsonBuffer)
	if err != nil {
		return fmt.Errorf("could not post request %s: %s", cwurl, err)
	}

	return nil
}

//TBD: This should return something useful, response body??
func (cw *ConnectwiseSite) DeleteCallback(callback int) error {
	restAction := fmt.Sprintf("/system/callbacks/%d", callback)
	cwurl, err := cw.BuildURL(restAction)
	if err != nil {
		return fmt.Errorf("could not build url %s: %s", restAction, err)
	}
	_, err = cw.DeleteRequest(cwurl)
	if err != nil {
		return fmt.Errorf("could not delete request %s: %s", cwurl, err)
	}

	return nil
}
