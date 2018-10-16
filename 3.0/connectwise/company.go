package connectwise

import (
	"encoding/json"
	"fmt"
	"time"
)

//Company is a struct to hold the unmarshaled JSON data when making a call to the Company API
type Company struct {
	ID         int    `json:"id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Status     struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			StatusHref string `json:"status_href"`
		} `json:"_info"`
	} `json:"status"`
	Type struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			TypeHref string `json:"type_href"`
		} `json:"_info"`
	} `json:"type"`
	AddressLine1 string `json:"addressLine1"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          string `json:"zip"`
	Country      struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			CountryHref string `json:"country_href"`
		} `json:"_info"`
	} `json:"country"`
	PhoneNumber    string `json:"phoneNumber"`
	FaxNumber      string `json:"faxNumber"`
	Website        string `json:"website"`
	TerritoryID    int    `json:"territoryId"`
	MarketID       int    `json:"marketId"`
	AccountNumber  string `json:"accountNumber"`
	DefaultContact struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			ContactHref string `json:"contact_href"`
		} `json:"_info"`
	} `json:"defaultContact"`
	DateAcquired      time.Time `json:"dateAcquired"`
	AnnualRevenue     float64   `json:"annualRevenue"`
	NumberOfEmployees int       `json:"numberOfEmployees"`
	TimeZone          struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			TimeZoneSetupHref string `json:"timeZoneSetup_href"`
		} `json:"_info"`
	} `json:"timeZone"`
	LeadFlag          bool   `json:"leadFlag"`
	UnsubscribeFlag   bool   `json:"unsubscribeFlag"`
	UserDefinedField1 string `json:"userDefinedField1"`
	UserDefinedField2 string `json:"userDefinedField2"`
	UserDefinedField3 string `json:"userDefinedField3"`
	UserDefinedField7 string `json:"userDefinedField7"`
	VendorIdentifier  string `json:"vendorIdentifier"`
	TaxIdentifier     string `json:"taxIdentifier"`
	TaxCode           struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			TaxCodeHref string `json:"taxCode_href"`
		} `json:"_info"`
	} `json:"taxCode"`
	BillingTerms struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"billingTerms"`
	BillToCompany struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			CompanyHref string `json:"company_href"`
		} `json:"_info"`
	} `json:"billToCompany"`
	BillingSite struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			SiteHref string `json:"site_href"`
		} `json:"_info"`
	} `json:"billingSite"`
	BillingContact struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			ContactHref string `json:"contact_href"`
		} `json:"_info"`
	} `json:"billingContact"`
	InvoiceDeliveryMethod struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"invoiceDeliveryMethod"`
	InvoiceToEmailAddress string `json:"invoiceToEmailAddress"`
	DeletedFlag           bool   `json:"deletedFlag"`
	MobileGUID            string `json:"mobileGuid"`
	Currency              struct {
		ID      int    `json:"id"`
		Symbol  string `json:"symbol"`
		IsoCode string `json:"isoCode"`
		Name    string `json:"name"`
		Info    struct {
			CurrencyHref string `json:"currency_href"`
		} `json:"_info"`
	} `json:"currency"`
	Info struct {
		LastUpdated        time.Time `json:"lastUpdated"`
		UpdatedBy          string    `json:"updatedBy"`
		DateEntered        time.Time `json:"dateEntered"`
		EnteredBy          string    `json:"enteredBy"`
		ContactsHref       string    `json:"contacts_href"`
		AgreementsHref     string    `json:"agreements_href"`
		TicketsHref        string    `json:"tickets_href"`
		OpportunitiesHref  string    `json:"opportunities_href"`
		ActivitiesHref     string    `json:"activities_href"`
		ProjectsHref       string    `json:"projects_href"`
		ConfigurationsHref string    `json:"configurations_href"`
		OrdersHref         string    `json:"orders_href"`
		DocumentsHref      string    `json:"documents_href"`
		SitesHref          string    `json:"sites_href"`
		TeamsHref          string    `json:"teams_href"`
		ReportsHref        string    `json:"reports_href"`
		NotesHref          string    `json:"notes_href"`
	} `json:"_info"`
	CustomFields []struct {
		ID               int    `json:"id"`
		Caption          string `json:"caption"`
		Type             string `json:"type"`
		EntryMethod      string `json:"entryMethod"`
		NumberOfDecimals int    `json:"numberOfDecimals"`
		Value            string `json:"value"`
	} `json:"customFields"`
}

type CompanyStatus struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	DefaultFlag          bool   `json:"defaultFlag"`
	InactiveFlag         bool   `json:"inactiveFlag"`
	NotifyFlag           bool   `json:"notifyFlag"`
	DisallowSavingFlag   bool   `json:"disallowSavingFlag"`
	NotificationMessage  string `json:"notificationMessage,omitempty"`
	CustomNoteFlag       bool   `json:"customNoteFlag"`
	CancelOpenTracksFlag bool   `json:"cancelOpenTracksFlag"`
	Info                 struct {
		LastUpdated time.Time `json:"lastUpdated"`
		UpdatedBy   string    `json:"updatedBy"`
	} `json:"_info"`
}

type CompanyTeamMember struct {
	ID      int `json:"id"`
	Company struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			CompanyHref string `json:"company_href"`
		} `json:"_info"`
	} `json:"company"`
	TeamRole struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			TeamroleHref string `json:"teamrole_href"`
		} `json:"_info"`
	} `json:"teamRole"`
	LocationID     int `json:"locationId"`
	BusinessUnitID int `json:"businessUnitId"`
	Member         struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			MemberHref string `json:"member_href"`
		} `json:"_info"`
	} `json:"member"`
	AccountManagerFlag bool `json:"accountManagerFlag"`
	TechFlag           bool `json:"techFlag"`
	SalesFlag          bool `json:"salesFlag"`
	Info               struct {
		LastUpdated time.Time `json:"lastUpdated"`
		UpdatedBy   string    `json:"updatedBy"`
		DateEntered time.Time `json:"dateEntered"`
		EnteredBy   string    `json:"enteredBy"`
	} `json:"_info"`
}

//CompanyCount returns the number of companies in ConnectWise
func (cw *Site) CompanyCount() (int, error) {
	req := cw.NewRequest("/company/companies/count", "GET", nil)
	err := req.Do()
	if err != nil {
		return 0, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	count := &Count{}
	err = json.Unmarshal(req.Body, count)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return count.Count, nil
}

//GetCompanyByName expects an exact match, perhaps an improvement could be made to support wildcard characters.
//Will return a pointer to a Company
func (cw *Site) GetCompanyByName(companyName string) (*Company, error) {
	req := cw.NewRequest("/company/companies", "GET", nil)
	req.URLValues.Add("conditions", "name=\""+companyName+"\"")

	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	co := &[]Company{}
	err = json.Unmarshal(req.Body, co)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}
	if len(*co) == 0 {
		return nil, fmt.Errorf("connectsise returned no results for %s", companyName)
	}

	//This endpoint always returns a JSON array, but given the condition we apply, we can safely just return the first and only item in the array
	return &(*co)[0], nil
}

//GetCompanyByIdentifier expects an exact match, perhaps an improvement could be made to support wildcard characters.
//Accepts a company identifier/shortname
//Will return a pointer to a Company
func (cw *Site) GetCompanyByIdentifier(companyIdentifier string) (*Company, error) {
	req := cw.NewRequest("/company/companies", "GET", nil)
	req.URLValues.Add("conditions", "identifier=\""+companyIdentifier+"\"")

	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	co := &[]Company{}
	err = json.Unmarshal(req.Body, co)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}
	if len(*co) == 0 {
		return nil, fmt.Errorf("connectsise returned no results for %s", companyIdentifier)
	}

	//This endpoint always returns a JSON array, but given the condition we apply, we can safely just return the first and only item in the array
	return &(*co)[0], nil
}

//GetCompanyByID expects the Connectwise Company ID and returns a pointer to a Company
//Does not return a slice like GetCompanyByName as the ID will only ever have one match
func (cw *Site) GetCompanyByID(companyID int) (*Company, error) {
	req := cw.NewRequest(fmt.Sprintf("/company/companies/%d", companyID), "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	co := &Company{}
	err = json.Unmarshal(req.Body, co)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return co, nil
}

func (cw *Site) GetCompanyStatuses() (*[]CompanyStatus, error) {
	req := cw.NewRequest("/company/companies/statuses", "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	costat := &[]CompanyStatus{}
	err = json.Unmarshal(req.Body, costat)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return costat, nil
}

func (cw *Site) GetCompanyTeamMembers(companyID int) (*[]CompanyTeamMember, error) {
	req := cw.NewRequest(fmt.Sprintf("/company/companies/%d/teams", companyID), "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	teamMembers := &[]CompanyTeamMember{}
	err = json.Unmarshal(req.Body, teamMembers)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return teamMembers, nil

}
