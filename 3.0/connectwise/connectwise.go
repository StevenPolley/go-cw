package connectwise

import (
	"encoding/base64"
	"fmt"
	"log"
)

type ConnectwiseSite struct {
	Site string
	Auth string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Returns a ConnectwiseSite struct with the site and auth string available for use in API requests
func NewSite(site string, publicKey string, privateKey string, company string) *ConnectwiseSite {
	authString := fmt.Sprintf("%s+%s:%s", company, publicKey, privateKey)
	authString = base64.StdEncoding.EncodeToString([]byte(authString))
	authString = fmt.Sprintf("Basic %s", authString)

	cwSite := ConnectwiseSite{Site: site, Auth: authString}

	return &cwSite
}
