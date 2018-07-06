package connectwise

import (
	"encoding/base64"
	"fmt"
)

//ConnectwiseSite is a stuct containing the URL of the site and the API authorization token in the format that CW expects it.
type ConnectwiseSite struct {
	Site string
	Auth string
}

//NewSite returns a pointer to a ConnectwiseSite struct with the site and auth string available for use in API requests
func NewSite(site string, publicKey string, privateKey string, company string) *ConnectwiseSite {
	//The auth string must be formatted in this way when used in requests to the API
	authString := fmt.Sprintf("%s+%s:%s", company, publicKey, privateKey)
	authString = base64.StdEncoding.EncodeToString([]byte(authString))
	authString = fmt.Sprintf("Basic %s", authString)

	cwSite := ConnectwiseSite{Site: site, Auth: authString}

	return &cwSite
}
