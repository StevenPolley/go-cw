package connectwise

import (
	"encoding/json"
	"fmt"
	"time"
)

type Ticket struct {
	ID         int    `json:"id"`
	Summary    string `json:"summary"`
	RecordType string `json:"recordType"`
	Board      struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			BoardHref string `json:"board_href"`
		} `json:"_info"`
	} `json:"board"`
	Status struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			StatusHref string `json:"status_href"`
		} `json:"_info"`
	} `json:"status"`
	Company struct {
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
	AddressLine1        string `json:"addressLine1,omitempty"`
	City                string `json:"city,omitempty"`
	StateIdentifier     string `json:"stateIdentifier,omitempty"`
	Zip                 string `json:"zip,omitempty"`
	ContactName         string `json:"contactName"`
	ContactPhoneNumber  string `json:"contactPhoneNumber"`
	ContactEmailAddress string `json:"contactEmailAddress,omitempty"`
	Team                struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			TeamHref string `json:"team_href"`
		} `json:"_info"`
	} `json:"team"`
	Priority struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Sort int    `json:"sort"`
		Info struct {
			PriorityHref string `json:"priority_href"`
			ImageHref    string `json:"image_href"`
		} `json:"_info"`
	} `json:"priority"`
	ServiceLocation struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			LocationHref string `json:"location_href"`
		} `json:"_info"`
	} `json:"serviceLocation"`
	Source struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			SourceHref string `json:"source_href"`
		} `json:"_info"`
	} `json:"source"`
	Agreement struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AgreementHref string `json:"agreement_href"`
		} `json:"_info"`
	} `json:"agreement,omitempty"`
	Severity                   string    `json:"severity"`
	Impact                     string    `json:"impact"`
	AllowAllClientsPortalView  bool      `json:"allowAllClientsPortalView"`
	CustomerUpdatedFlag        bool      `json:"customerUpdatedFlag"`
	AutomaticEmailContactFlag  bool      `json:"automaticEmailContactFlag"`
	AutomaticEmailResourceFlag bool      `json:"automaticEmailResourceFlag"`
	AutomaticEmailCcFlag       bool      `json:"automaticEmailCcFlag"`
	ClosedDate                 time.Time `json:"closedDate"`
	ClosedBy                   string    `json:"closedBy"`
	ClosedFlag                 bool      `json:"closedFlag"`
	DateEntered                time.Time `json:"dateEntered"`
	EnteredBy                  string    `json:"enteredBy"`
	ActualHours                float64   `json:"actualHours,omitempty"`
	Approved                   bool      `json:"approved"`
	EstimatedExpenseCost       float64   `json:"estimatedExpenseCost"`
	EstimatedExpenseRevenue    float64   `json:"estimatedExpenseRevenue"`
	EstimatedProductCost       float64   `json:"estimatedProductCost"`
	EstimatedProductRevenue    float64   `json:"estimatedProductRevenue"`
	EstimatedTimeCost          float64   `json:"estimatedTimeCost"`
	EstimatedTimeRevenue       float64   `json:"estimatedTimeRevenue"`
	BillingMethod              string    `json:"billingMethod"`
	SubBillingMethod           string    `json:"subBillingMethod"`
	DateResolved               time.Time `json:"dateResolved"`
	DateResplan                time.Time `json:"dateResplan"`
	DateResponded              time.Time `json:"dateResponded"`
	ResolveMinutes             int       `json:"resolveMinutes"`
	ResPlanMinutes             int       `json:"resPlanMinutes"`
	RespondMinutes             int       `json:"respondMinutes"`
	IsInSLA                    bool      `json:"isInSla"`
	HasChildTicket             bool      `json:"hasChildTicket"`
	BillTime                   string    `json:"billTime"`
	BillExpenses               string    `json:"billExpenses"`
	BillProducts               string    `json:"billProducts"`
	LocationID                 int       `json:"locationId"`
	BusinessUnitID             int       `json:"businessUnitId"`
	MobileGUID                 string    `json:"mobileGuid"`
	SLA                        struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			SLAHref string `json:"sla_href"`
		} `json:"_info"`
	} `json:"sla"`
	Currency struct {
		ID      int    `json:"id"`
		Symbol  string `json:"symbol"`
		IsoCode string `json:"isoCode"`
		Name    string `json:"name"`
		Info    struct {
			CurrencyHref string `json:"currency_href"`
		} `json:"_info"`
	} `json:"currency"`
	Info struct {
		LastUpdated         time.Time `json:"lastUpdated"`
		UpdatedBy           string    `json:"updatedBy"`
		ActivitiesHref      string    `json:"activities_href"`
		ScheduleentriesHref string    `json:"scheduleentries_href"`
		DocumentsHref       string    `json:"documents_href"`
		ConfigurationsHref  string    `json:"configurations_href"`
		TasksHref           string    `json:"tasks_href"`
		NotesHref           string    `json:"notes_href"`
		ProductsHref        string    `json:"products_href"`
		TimeentriesHref     string    `json:"timeentries_href"`
		ExpenseEntriesHref  string    `json:"expenseEntries_href"`
	} `json:"_info"`
	CustomFields []struct {
		ID               int    `json:"id"`
		Caption          string `json:"caption"`
		Type             string `json:"type"`
		EntryMethod      string `json:"entryMethod"`
		NumberOfDecimals int    `json:"numberOfDecimals"`
	} `json:"customFields"`
	RequiredDate time.Time `json:"requiredDate,omitempty"`
	BudgetHours  float64   `json:"budgetHours,omitempty"`
	AddressLine2 string    `json:"addressLine2,omitempty"`
	Contact      struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			MobileGUID  string `json:"mobileGuid"`
			ContactHref string `json:"contact_href"`
		} `json:"_info"`
	} `json:"contact,omitempty"`
	ContactPhoneExtension string `json:"contactPhoneExtension,omitempty"`
}

type TimeEntryReference struct {
	ID   int
	Info struct {
		Notes    string `json:"notes"`
		TimeHref string `json:"time_href"`
	} `json:"_info"`
}

func (cw *ConnectwiseSite) GetTicketByID(ticketID int) *Ticket {

	Url := cw.BuildUrl(fmt.Sprintf("/service/tickets/%d", ticketID))

	body := cw.GetRequest(Url)
	ticket := Ticket{}
	check(json.Unmarshal(body, &ticket))

	return &ticket
}

func (cw *ConnectwiseSite) GetTicketTimeEntriesByID(ticketID int) *[]TimeEntryReference {

	Url := cw.BuildUrl(fmt.Sprintf("/service/tickets/%d/timeentries", ticketID))

	body := cw.GetRequest(Url)
	timeEntryReference := []TimeEntryReference{}
	check(json.Unmarshal(body, &timeEntryReference)) //  *[]TimeEntryReference

	return &timeEntryReference
}
