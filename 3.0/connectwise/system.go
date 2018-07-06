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

func (cw *ConnectwiseSite) GetCallbacks() *[]Callback {

	URL := cw.BuildURL("/system/callbacks")
	body := cw.GetRequest(URL)

	callbacks := []Callback{}
	check(json.Unmarshal(body, &callbacks))

	return &callbacks

}

//TBD: This should return something?
func (cw *ConnectwiseSite) NewCallback(callback Callback) {

	URL := cw.BuildURL("/system/callbacks")
	jsonCallback, err := json.Marshal(callback)
	check(err)

	jsonBuffer := bytes.NewReader(jsonCallback)

	cw.PostRequest(URL, jsonBuffer)

}

func (cw *ConnectwiseSite) DeleteCallback(callback int) {

	URL := cw.BuildURL(fmt.Sprintf("/system/callbacks/%d", callback))
	body := cw.DeleteRequest(URL)
	fmt.Print(string(body))

}
