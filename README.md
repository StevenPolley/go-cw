# go-cw
Go structs and methods for the ConnectWise REST API

Note: This is far from complete, I'm simply adding structs and methods as I have an actual requirement for them. If you add to this, please feel free to send a pull request.

# Installation
```
go get deadbeef.codes/steven/go-cw
```

# Usage
```
package main

import (
	"fmt"
	"log"
	"os"

	"deadbeef.codes/steven/go-cw/3.0/connectwise"
)

var cw *connectwise.ConnectwiseSite

func init() {
	publicKey := os.Getenv("gocwpublickey")
	privateKey := os.Getenv("gocwprivatekey")
	site := os.Getenv("gocwsite")
	company := os.Getenv("gocwcompany")

	cw = connectwise.NewSite(site, publicKey, privateKey, company)
}

func main() {

	co, err := cw.GetCompanyByID(2)
	if err != nil {
		log.Fatal("could not get company 2: %g", err)
	}

	//Refer to the Connectwise API documentation to see what fields are available
	fmt.Println(co.Name)
	fmt.Println(co.DefaultContact.Name)
}

```
