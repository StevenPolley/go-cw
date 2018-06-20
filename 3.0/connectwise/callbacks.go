package connectwise

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

type Callback struct {
	Id           int
	Description  string
	Url          string
	ObjectId     int
	Type         string
	Level        string
	MemberId     int
	InactiveFlag bool
}

func GetCallbacks(site *ConnectwiseSite) {

	//Build the request URL
	var Url *url.URL
	Url, err := url.Parse(site.Site)
	check(err)
	Url.Path += "/system/callbacks"

	body := GetRequest(site, Url)
	fmt.Print(string(body))
	//	check(json.Unmarshal(body, &ticket))

}

func NewCallback(site *ConnectwiseSite, callback Callback) {

	var Url *url.URL
	Url, err := url.Parse(site.Site)
	check(err)
	Url.Path += "/system/callbacks"

	jsonCallback, err := json.Marshal(callback)
	check(err)

	jsonBuffer := bytes.NewReader(jsonCallback)

	body := PostRequest(site, Url, jsonBuffer)

	fmt.Print(string(body))

}

func DeleteCallback(site *ConnectwiseSite, callback int) {

	var Url *url.URL
	Url, err := url.Parse(site.Site)
	check(err)
	Url.Path += fmt.Sprintf("/system/callbacks/%d", callback)

	body := DeleteRequest(site, Url)

	fmt.Print(string(body))

}
