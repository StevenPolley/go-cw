# go-cw
Go structs and methods for the ConnectWise REST API

Note: This is far from complete, I'm simply adding structs and methods as I have an actual requirement for them. If you add to this, please feel free to send a pull request.

#Installation
```
go get deadbeef.codes/steven/go-cw
```

#Usage
```
package main

import (
	"deadbeef.codes/steven/go-cw/3.0/connectwise"
	"fmt"
)

const (
	cwSite		= "https://yourconnectwisesite.com/v4_6_release/apis/3.0"
	cwAPIKeyPrivate = "ASDLFK4ah89ad" //Put in either your private API key or account password if using user impersonation
	cwAPIKey	= "ASLDFKJ2342kl" //Put in either your public API key or account username if using user impersonation
	cwCompany	= "yourcompanyname" //The connectwise company name
)

func main() {
	cw := connectwise.NewSite(cwSite, cwAPIKey, cwAPIKeyPrivate, cwCompany)
	companyDataByID := cw.GetCompanyByID(2) //Retrieves company ID 2 from CW and returns type pointer a Company
	fmt.Println(*companyDataByID.Name)
}
```
