# go-cw
Go structs and methods for the ConnectWise REST API

Note: This is far from complete, I'm simply adding structs and methods as I have an actual requirement for them. If you add to this, please feel free to send a pull request.

#Installation
```
go get github.com/StevenPolley/go-cw
```

#Usage
```
package main

import (
	"github.com/StevenPolley/go-cw/3.0/connectwise"
	"fmt"
)

const (
	cwSite		= "https://yourconnectwisesite.com/v4_6_release/apis/3.0"
	cwAPIKeyPrivate = "ASDLFK4ah89ad"
	cwAPIKey	= "ASLDFKJ2342kl"
	cwCompany	= "yourcompanyname"
)

func main() {
	cw := connectwise.NewSite(cwSite, cwAPIKey, cwAPIKeyPrivate, cwCompany)
	companyDataByID := connectwise.GetCompaniesByID(cw, 2) //Retrieves company ID 2 from CW and returns pointer to struct struct
	fmt.Println(*companyDataByID)
}
```
