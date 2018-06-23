package connectwise

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Callback struct {
	ID           int
	Description  string
	Url          string
	ObjectId     int
	Type         string
	Level        string
	MemberId     int
	InactiveFlag bool
}

func (cw *ConnectwiseSite) GetCallbacks() *[]Callback {

	Url := cw.BuildUrl("/system/callbacks")
	body := cw.GetRequest(Url)

	callbacks := []Callback{}
	check(json.Unmarshal(body, &callbacks))

	return &callbacks

}

//TBD: This should return something?
func (cw *ConnectwiseSite) NewCallback(callback Callback) {

	Url := cw.BuildUrl("/system/callbacks")
	jsonCallback, err := json.Marshal(callback)
	check(err)

	jsonBuffer := bytes.NewReader(jsonCallback)

	cw.PostRequest(Url, jsonBuffer)

}

func (cw *ConnectwiseSite) DeleteCallback(callback int) {

	Url := cw.BuildUrl(fmt.Sprintf("/system/callbacks/%d", callback))
	body := cw.DeleteRequest(Url)
	fmt.Print(string(body))

}
