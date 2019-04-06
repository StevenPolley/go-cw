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
	PhoneNumber string `json:"phoneNumber"`
	FaxNumber   string `json:"faxNumber"`
	Website     string `json:"website"`
	Territory   struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			LocationHref string `json:"location_href"`
		} `json:"_info"`
	} `json:"territory"`
	Market struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			MarketHref string `json:"Market_href"`
		} `json:"_info"`
	} `json:"market"`
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
	TimeZoneSetup     struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			TimeZoneSetupHref string `json:"timeZoneSetup_href"`
		} `json:"_info"`
	} `json:"timeZoneSetup"`
	LeadFlag          bool   `json:"leadFlag"`
	UnsubscribeFlag   bool   `json:"unsubscribeFlag"`
	UserDefinedField1 string `json:"userDefinedField1"`
	UserDefinedField2 string `json:"userDefinedField2"`
	UserDefinedField3 string `json:"userDefinedField3"`
	UserDefinedField4 string `json:"userDefinedField4"`
	UserDefinedField5 string `json:"userDefinedField5"`
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
	TypeIds               []int  `json:"typeIds"`
	Info                  struct {
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

type CompanySite struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
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
	AddressFormat        string  `json:"addressFormat"`
	PhoneNumber          string  `json:"phoneNumber"`
	FaxNumber            string  `json:"faxNumber"`
	TaxCodeID            int     `json:"taxCodeId,omitempty"`
	ExpenseReimbursement float64 `json:"expenseReimbursement"`
	PrimaryAddressFlag   bool    `json:"primaryAddressFlag"`
	DefaultShippingFlag  bool    `json:"defaultShippingFlag"`
	DefaultBillingFlag   bool    `json:"defaultBillingFlag"`
	DefaultMailingFlag   bool    `json:"defaultMailingFlag"`
	InactiveFlag         bool    `json:"inactiveFlag"`
	MobileGUID           string  `json:"mobileGuid"`
	TimeZone             struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			TimeZoneSetupHref string `json:"timeZoneSetup_href"`
		} `json:"_info"`
	} `json:"timeZone"`
	Company struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			CompanyHref string `json:"company_href"`
			MobileGUID  string `json:"mobileGuid"`
		} `json:"_info"`
	} `json:"company"`
	Info struct {
		LastUpdated time.Time `json:"lastUpdated"`
		UpdatedBy   string    `json:"updatedBy"`
		DateEntered time.Time `json:"dateEntered"`
		EnteredBy   string    `json:"enteredBy"`
	} `json:"_info"`
	Calendar struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			CalendarHref string `json:"calendar_href"`
		} `json:"_info"`
	} `json:"calendar,omitempty"`
}

type Contact struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName,omitempty"`
	Company   struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			CompanyHref string `json:"company_href"`
			MobileGUID  string `json:"mobileGuid"`
		} `json:"_info"`
	} `json:"company"`
	Site struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			SiteHref   string `json:"site_href"`
			MobileGUID string `json:"mobileGuid"`
		} `json:"_info"`
	} `json:"site,omitempty"`
	InactiveFlag           bool   `json:"inactiveFlag"`
	Title                  string `json:"title,omitempty"`
	MarriedFlag            bool   `json:"marriedFlag"`
	ChildrenFlag           bool   `json:"childrenFlag"`
	PortalSecurityLevel    int    `json:"portalSecurityLevel,omitempty"`
	DisablePortalLoginFlag bool   `json:"disablePortalLoginFlag,omitempty"`
	UnsubscribeFlag        bool   `json:"unsubscribeFlag"`
	MobileGUID             string `json:"mobileGuid"`
	DefaultBillingFlag     bool   `json:"defaultBillingFlag"`
	DefaultFlag            bool   `json:"defaultFlag"`
	CommunicationItems     []struct {
		ID   int `json:"id"`
		Type struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"type"`
		Value             string `json:"value"`
		Extension         string `json:"extension,omitempty"`
		DefaultFlag       bool   `json:"defaultFlag"`
		CommunicationType string `json:"communicationType"`
	} `json:"communicationItems,omitempty"`
	TypeIds []interface{} `json:"typeIds"`
	Info    struct {
		LastUpdated        string `json:"lastUpdated"`
		UpdatedBy          string `json:"updatedBy"`
		CommunicationsHref string `json:"communications_href"`
		NotesHref          string `json:"notes_href"`
		TracksHref         string `json:"tracks_href"`
		PortalSecurityHref string `json:"portalSecurity_href"`
		ActivitiesHref     string `json:"activities_href"`
		DocumentsHref      string `json:"documents_href"`
		ConfigurationsHref string `json:"configurations_href"`
		TicketsHref        string `json:"tickets_href"`
		OpportunitiesHref  string `json:"opportunities_href"`
		ProjectsHref       string `json:"projects_href"`
		ImageHref          string `json:"image_href"`
	} `json:"_info"`
	AddressLine1 string `json:"addressLine1,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Zip          string `json:"zip,omitempty"`
	Department   struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			DepartmentHref string `json:"department_href"`
		} `json:"_info"`
	} `json:"department,omitempty"`
	RelationshipOverride string `json:"relationshipOverride,omitempty"`
	School               string `json:"school,omitempty"`
	NickName             string `json:"nickName,omitempty"`
	SignificantOther     string `json:"significantOther,omitempty"`
	Relationship         struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			RelationshipHref string `json:"relationship_href"`
		} `json:"_info"`
	} `json:"relationship,omitempty"`
	Gender string `json:"gender,omitempty"`
}

type ContactCommunication struct {
	ID        int `json:"id"`
	ContactID int `json:"contactId"`
	Type      struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"type"`
	Value             string `json:"value"`
	DefaultFlag       bool   `json:"defaultFlag"`
	MobileGUID        string `json:"mobileGuid"`
	CommunicationType string `json:"communicationType"`
	Info              struct {
		LastUpdated       time.Time `json:"lastUpdated"`
		UpdatedBy         string    `json:"updatedBy"`
		ContactMobileGUID string    `json:"contactMobileGuid"`
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

	fmt.Println(string(req.Body))

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

func (cw *Site) GetCompanySites(companyID int) (*[]CompanySite, error) {
	req := cw.NewRequest(fmt.Sprintf("/company/companies/%d/sites", companyID), "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	companySite := &[]CompanySite{}
	err = json.Unmarshal(req.Body, companySite)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return companySite, nil

}

func (cw *Site) GetContactByID(contactID int) (*Contact, error) {
	req := cw.NewRequest(fmt.Sprintf("/company/contacts/%d", contactID), "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	contact := &Contact{}
	err = json.Unmarshal(req.Body, contact)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return contact, nil
}

func (cw *Site) NewContact(contact *Contact) (*Contact, error) {
	jsonContact, err := json.Marshal(contact)
	if err != nil {
		return nil, fmt.Errorf("could not marshal json data: %s", err)
	}

	req := cw.NewRequest("/company/contacts", "POST", jsonContact)
	err = req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	contact = &Contact{}
	err = json.Unmarshal(req.Body, contact)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return contact, nil
}

func (cw *Site) UpdateContact(contact *Contact) (*Contact, error) {
	jsonContact, err := json.Marshal(contact)
	if err != nil {
		return nil, fmt.Errorf("could not marshal json data: %s", err)
	}

	req := cw.NewRequest(fmt.Sprintf("/company/contacts/%d", contact.ID), "PUT", jsonContact)
	err = req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	contact = &Contact{}
	err = json.Unmarshal(req.Body, contact)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return contact, nil
}

func (cw *Site) GetContactCommunicationsByContactID(contactID, communicationID int) (*[]ContactCommunication, error) {
	req := cw.NewRequest(fmt.Sprintf("/company/contacts/%d/communications", contactID), "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	communications := &[]ContactCommunication{}
	err = json.Unmarshal(req.Body, communications)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return communications, nil
}

func (cw *Site) NewContactCommunication(communication *ContactCommunication) (*ContactCommunication, error) {
	jsonCommunication, err := json.Marshal(communication)
	if err != nil {
		return nil, fmt.Errorf("could not marshal json data: %s", err)
	}

	req := cw.NewRequest(fmt.Sprintf("/company/contacts/%d/communications", communication.ContactID), "POST", jsonCommunication)
	err = req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	communication = &ContactCommunication{}
	err = json.Unmarshal(req.Body, communication)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return communication, nil
}

func (cw *Site) SearchCompanyByName(name string) (*[]Company, error) {
	req := cw.NewRequest("/company/companies", "GET", nil)
	req.URLValues.Add("conditions", "name contains \""+name+"\"")
	req.URLValues.Add("pageSize", "40")
	req.URLValues.Add("fields", "name")
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	companies := &[]Company{}
	err = json.Unmarshal(req.Body, companies)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return companies, nil
}
