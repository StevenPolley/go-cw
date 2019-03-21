package connectwise

import (
	"encoding/json"
	"fmt"
)

//Callback is a struct to hold the unmarshaled JSON data when making a call to the System API
type Callback struct {
	ID           int    `json:"id"`
	Description  string `json:"description"`
	URL          string `json:"url"`
	ObjectID     int    `json:"objectId"`
	Type         string `json:"type"`
	Level        string `json:"level"`
	MemberID     int    `json:"memberId"`
	InactiveFlag bool   `json:"inactiveFlag"`
	Info         struct {
		LastUpdated string `json:"lastUpdated"`
		UpdatedBy   string `json:"updatedBy"`
	} `json:"_info"`
}

type Member struct {
	ID         int    `json:"id"`
	Identifier string `json:"identifier"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName,omitempty"`
	Title      string `json:"title,omitempty"`
	ReportCard struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			ReportCardHref string `json:"reportCard_href"`
		} `json:"_info"`
	} `json:"reportCard,omitempty"`
	LicenseClass      string `json:"licenseClass"`
	DisableOnlineFlag bool   `json:"disableOnlineFlag"`
	EnableMobileFlag  bool   `json:"enableMobileFlag,omitempty"`
	Type              struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"type,omitempty"`
	EmployeeIdentifer string `json:"employeeIdentifer,omitempty"`
	TimeZone          struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			TimeZoneSetupHref string `json:"timeZoneSetup_href"`
		} `json:"_info"`
	} `json:"timeZone,omitempty"`
	Country struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			CountryHref string `json:"country_href"`
		} `json:"_info"`
	} `json:"country,omitempty"`
	ServiceBoardTeamIds []int  `json:"serviceBoardTeamIds"`
	EnableMobileGpsFlag bool   `json:"enableMobileGpsFlag"`
	InactiveFlag        bool   `json:"inactiveFlag"`
	LastLogin           string `json:"lastLogin,omitempty"`
	Photo               struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			Filename             string `json:"filename"`
			DocumentHref         string `json:"document_href"`
			DocumentDownloadHref string `json:"documentDownload_href"`
		} `json:"_info"`
	} `json:"photo,omitempty"`
	OfficeEmail     string `json:"officeEmail,omitempty"`
	OfficePhone     string `json:"officePhone,omitempty"`
	OfficeExtension string `json:"officeExtension,omitempty"`
	MobilePhone     string `json:"mobilePhone,omitempty"`
	MobileExtension string `json:"mobileExtension,omitempty"`
	HomePhone       string `json:"homePhone,omitempty"`
	HomeExtension   string `json:"homeExtension,omitempty"`
	DefaultEmail    string `json:"defaultEmail"`
	DefaultPhone    string `json:"defaultPhone"`
	SecurityRole    struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"securityRole,omitempty"`
	AdminFlag      bool `json:"adminFlag"`
	StructureLevel struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
		} `json:"_info"`
	} `json:"structureLevel"`
	SecurityLocation struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			LocationHref string `json:"location_href"`
		} `json:"_info"`
	} `json:"securityLocation"`
	DefaultLocation struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			LocationHref string `json:"location_href"`
		} `json:"_info"`
	} `json:"defaultLocation,omitempty"`
	DefaultDepartment struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			DepartmentHref string `json:"department_href"`
		} `json:"_info"`
	} `json:"defaultDepartment,omitempty"`
	RestrictLocationFlag   bool `json:"restrictLocationFlag"`
	RestrictDepartmentFlag bool `json:"restrictDepartmentFlag"`
	WorkRole               struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			WorkRoleHref string `json:"workRole_href"`
		} `json:"_info"`
	} `json:"workRole,omitempty"`
	WorkType struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			WorkTypeHref string `json:"workType_href"`
		} `json:"_info"`
	} `json:"workType,omitempty"`
	TimeApprover struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			MemberHref string `json:"member_href"`
		} `json:"_info"`
	} `json:"timeApprover,omitempty"`
	ExpenseApprover struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			MemberHref string `json:"member_href"`
		} `json:"_info"`
	} `json:"expenseApprover,omitempty"`
	DailyCapacity                            float64 `json:"dailyCapacity,omitempty"`
	HourlyCost                               float64 `json:"hourlyCost"`
	IncludeInUtilizationReportingFlag        bool    `json:"includeInUtilizationReportingFlag"`
	RequireExpenseEntryFlag                  bool    `json:"requireExpenseEntryFlag"`
	RequireTimeSheetEntryFlag                bool    `json:"requireTimeSheetEntryFlag"`
	RequireStartAndEndTimeOnTimeEntryFlag    bool    `json:"requireStartAndEndTimeOnTimeEntryFlag"`
	AllowInCellEntryOnTimeSheet              bool    `json:"allowInCellEntryOnTimeSheet"`
	EnterTimeAgainstCompanyFlag              bool    `json:"enterTimeAgainstCompanyFlag"`
	AllowExpensesEnteredAgainstCompaniesFlag bool    `json:"allowExpensesEnteredAgainstCompaniesFlag"`
	TimeReminderEmailFlag                    bool    `json:"timeReminderEmailFlag"`
	DaysTolerance                            int     `json:"daysTolerance,omitempty"`
	MinimumHours                             float64 `json:"minimumHours,omitempty"`
	TimeSheetStartDate                       string  `json:"timeSheetStartDate"`
	HireDate                                 string  `json:"hireDate,omitempty"`
	ServiceDefaultLocation                   struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			LocationHref string `json:"location_href"`
		} `json:"_info"`
	} `json:"serviceDefaultLocation,omitempty"`
	ServiceDefaultDepartment struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			DepartmentHref string `json:"department_href"`
		} `json:"_info"`
	} `json:"serviceDefaultDepartment,omitempty"`
	ServiceDefaultBoard struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			BoardHref string `json:"board_href"`
		} `json:"_info"`
	} `json:"serviceDefaultBoard,omitempty"`
	RestrictServiceDefaultLocationFlag   bool  `json:"restrictServiceDefaultLocationFlag"`
	RestrictServiceDefaultDepartmentFlag bool  `json:"restrictServiceDefaultDepartmentFlag"`
	ExcludedServiceBoardIds              []int `json:"excludedServiceBoardIds"`
	ProjectDefaultLocation               struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			LocationHref string `json:"location_href"`
		} `json:"_info"`
	} `json:"projectDefaultLocation,omitempty"`
	ProjectDefaultDepartment struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			DepartmentHref string `json:"department_href"`
		} `json:"_info"`
	} `json:"projectDefaultDepartment,omitempty"`
	RestrictProjectDefaultLocationFlag   bool          `json:"restrictProjectDefaultLocationFlag"`
	RestrictProjectDefaultDepartmentFlag bool          `json:"restrictProjectDefaultDepartmentFlag"`
	ExcludedProjectBoardIds              []interface{} `json:"excludedProjectBoardIds"`
	ScheduleDefaultLocation              struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			LocationHref string `json:"location_href"`
		} `json:"_info"`
	} `json:"scheduleDefaultLocation,omitempty"`
	ScheduleDefaultDepartment struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			DepartmentHref string `json:"department_href"`
		} `json:"_info"`
	} `json:"scheduleDefaultDepartment,omitempty"`
	ScheduleCapacity               float64 `json:"scheduleCapacity,omitempty"`
	RestrictScheduleFlag           bool    `json:"restrictScheduleFlag"`
	HideMemberInDispatchPortalFlag bool    `json:"hideMemberInDispatchPortalFlag"`
	SalesDefaultLocation           struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			LocationHref string `json:"location_href"`
		} `json:"_info"`
	} `json:"salesDefaultLocation,omitempty"`
	RestrictDefaultSalesTerritoryFlag bool   `json:"restrictDefaultSalesTerritoryFlag"`
	RestrictDefaultWarehouseFlag      bool   `json:"restrictDefaultWarehouseFlag"`
	RestrictDefaultWarehouseBinFlag   bool   `json:"restrictDefaultWarehouseBinFlag"`
	MapiName                          string `json:"mapiName,omitempty"`
	CalendarSyncIntegrationFlag       bool   `json:"calendarSyncIntegrationFlag"`
	EnableLdapAuthenticationFlag      bool   `json:"enableLdapAuthenticationFlag"`
	CompanyActivityTabFormat          string `json:"companyActivityTabFormat,omitempty"`
	InvoiceTimeTabFormat              string `json:"invoiceTimeTabFormat,omitempty"`
	InvoiceScreenDefaultTabFormat     string `json:"invoiceScreenDefaultTabFormat"`
	InvoicingDisplayOptions           string `json:"invoicingDisplayOptions"`
	AgreementInvoicingDisplayOptions  string `json:"agreementInvoicingDisplayOptions"`
	TimebasedOneTimePasswordActivated bool   `json:"timebasedOneTimePasswordActivated"`
	Info                              struct {
		LastUpdated string `json:"lastUpdated"`
		UpdatedBy   string `json:"updatedBy"`
		ImageHref   string `json:"image_href"`
	} `json:"_info"`
	Notes               string  `json:"notes,omitempty"`
	InactiveDate        string  `json:"inactiveDate,omitempty"`
	BillableForecast    float64 `json:"billableForecast,omitempty"`
	ProjectDefaultBoard struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			BoardHref string `json:"board_href"`
		} `json:"_info"`
	} `json:"projectDefaultBoard,omitempty"`
	Calendar struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			CalendarHref string `json:"calendar_href"`
		} `json:"_info"`
	} `json:"calendar,omitempty"`
	HourlyRate        float64 `json:"hourlyRate,omitempty"`
	LdapConfiguration struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Server string `json:"server"`
		Info   struct {
			LdapConfigurationHref string `json:"ldapConfiguration_href"`
		} `json:"_info"`
	} `json:"ldapConfiguration,omitempty"`
	LdapUserName string `json:"ldapUserName,omitempty"`
}

//GetSystemMembers returns a slice of Member structs containing all the members (users) of connectwise
//TBD finish this, I don't have permissions with my API key to see the JSON data
func (cw *Site) GetSystemMembers() (*[]Member, error) {
	req := cw.NewRequest("/system/members", "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	members := &[]Member{}
	err = json.Unmarshal(req.Body, members)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return members, nil
}

func (cw *Site) GetSystemMemberByID(memberID int) (*Member, error) {
	req := cw.NewRequest(fmt.Sprintf("/system/members/%d", memberID), "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	member := &Member{}
	err = json.Unmarshal(req.Body, member)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return member, nil
}

func (cw *Site) GetSystemMemberByIdentifier(identifier string) (*Member, error) {
	req := cw.NewRequest(fmt.Sprintf("/system/members/%s", identifier), "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	member := &Member{}
	err = json.Unmarshal(req.Body, member)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return member, nil
}

func (cw *Site) SearchSystemMembersByIdentifier(identifier string) (*[]Member, error) {
	req := cw.NewRequest("/system/members", "GET", nil)
	req.URLValues.Add("conditions", "identifier contains \""+identifier+"\"")
	req.URLValues.Add("pageSize", "40")
	req.URLValues.Add("fields", "identifier,officeEmail")
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	members := &[]Member{}
	err = json.Unmarshal(req.Body, members)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return members, nil
}

//GetCallbacks returns a slice of Callback structs containing all the callbacks currently registered with ConnectWise
func (cw *Site) GetCallbacks() (*[]Callback, error) {
	req := cw.NewRequest("/system/callbacks", "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	callbacks := &[]Callback{}
	err = json.Unmarshal(req.Body, callbacks)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return callbacks, nil
}

//NewCallback expects a Callback struct and will register a new callback with Connectwise
//TBD: This should return something useful, response body??
func (cw *Site) NewCallback(callback *Callback) (*Callback, error) {
	jsonCallback, err := json.Marshal(callback)
	if err != nil {
		return nil, fmt.Errorf("could not marshal json data: %s", err)
	}

	req := cw.NewRequest("/system/callbacks", "POST", jsonCallback)
	err = req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	callback = &Callback{}
	err = json.Unmarshal(req.Body, callback)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return callback, nil
}

//DeleteCallback expects the ID of an existing callback and will unregister it with ConnectWise
//Does not return anything - CW gives an empty response body
func (cw *Site) DeleteCallback(callback int) error {
	req := cw.NewRequest(fmt.Sprintf("/system/callbacks/%d", callback), "DELETE", nil)
	err := req.Do()
	if err != nil {
		return fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	return nil
}
