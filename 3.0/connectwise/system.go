package connectwise

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
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

func GetCallbacks(site *ConnectwiseSite) *[]Callback {

	Url := BuildUrl(site, "/system/callbacks")
	body := GetRequest(site, Url)

	callbacks := []Callback{}
	check(json.Unmarshal(body, &callbacks))

	return &callbacks

}

//TBD: This should return something?
func NewCallback(site *ConnectwiseSite, callback Callback) {

	Url := BuildUrl(site, "/system/callbacks")
	jsonCallback, err := json.Marshal(callback)
	check(err)

	jsonBuffer := bytes.NewReader(jsonCallback)

	//body := PostRequest(site, Url, jsonBuffer)
	PostRequest(site, Url, jsonBuffer)

}

func DeleteCallback(site *ConnectwiseSite, callback int) {

	Url := BuildUrl(site, fmt.Sprintf("/system/callbacks/%d", callback))
	body := DeleteRequest(site, Url)
	fmt.Print(string(body))

}
