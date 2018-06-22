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

	callbacks := []Callback{}

	//Build the request URL
	var Url *url.URL
	Url, err := url.Parse(site.Site)
	check(err)
	Url.Path += "/system/callbacks"

	body := GetRequest(site, Url)
	check(json.Unmarshal(body, &callbacks))

	return &callbacks

}

func NewCallback(site *ConnectwiseSite, callback Callback) {

	var Url *url.URL
	Url, err := url.Parse(site.Site)
	check(err)
	Url.Path += "/system/callbacks"

	jsonCallback, err := json.Marshal(callback)
	check(err)

	jsonBuffer := bytes.NewReader(jsonCallback)

	//body := PostRequest(site, Url, jsonBuffer)
	PostRequest(site, Url, jsonBuffer)

}

func DeleteCallback(site *ConnectwiseSite, callback int) {

	var Url *url.URL
	Url, err := url.Parse(site.Site)
	check(err)
	Url.Path += fmt.Sprintf("/system/callbacks/%d", callback)

	body := DeleteRequest(site, Url)

	fmt.Print(string(body))

}
