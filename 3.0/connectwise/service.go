package connectwise

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

//ServiceTeam is a struct to hold the unmarshaled JSON data when making a call to the service API
type ServiceTeam struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Leader struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			MemberHref string `json:"member_href"`
		} `json:"_info"`
	} `json:"leader"`
	Location struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			LocationHref string `json:"location_href"`
		} `json:"_info"`
	} `json:"location"`
	Department struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			DepartmentHref string `json:"department_href"`
		} `json:"_info"`
	} `json:"department"`
	DeleteNotifyFlag bool `json:"deleteNotifyFlag"`
	Info             struct {
		LastUpdated time.Time `json:"lastUpdated"`
		UpdatedBy   string    `json:"updatedBy"`
		DateEntered time.Time `json:"dateEntered"`
		EnteredBy   string    `json:"enteredBy"`
	} `json:"_info"`
}

//BoardTeam is a struct to hold the unmarshaled JSON data when making a call to the service API
type BoardTeam struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	TeamLeader struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			MemberHref string `json:"member_href"`
		} `json:"_info"`
	} `json:"teamLeader"`
	Members              []int `json:"members"`
	DefaultFlag          bool  `json:"defaultFlag"`
	NotifyOnTicketDelete bool  `json:"notifyOnTicketDelete"`
	BoardID              int   `json:"boardId"`
	LocationID           int   `json:"locationId"`
	BusinessUnitID       int   `json:"businessUnitId"`
	Info                 struct {
		LastUpdated time.Time `json:"lastUpdated"`
		UpdatedBy   string    `json:"updatedBy"`
	} `json:"_info"`
}

type BoardStatus struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Board struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			BoardHref string `json:"board_href"`
		} `json:"_info"`
	} `json:"board"`
	ExternalIntegrationXref struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			StatusExternalIntegrationHref string `json:"statusExternalIntegration_href"`
		} `json:"_info"`
	} `json:"externalIntegrationXref,omitempty"`
	SortOrder                 int    `json:"sortOrder"`
	DisplayOnBoard            bool   `json:"displayOnBoard"`
	Inactive                  bool   `json:"inactive"`
	ClosedStatus              bool   `json:"closedStatus"`
	TimeEntryNotAllowed       bool   `json:"timeEntryNotAllowed"`
	DefaultFlag               bool   `json:"defaultFlag"`
	EscalationStatus          string `json:"escalationStatus"`
	CustomerPortalDescription string `json:"customerPortalDescription,omitempty"`
	CustomerPortalFlag        bool   `json:"customerPortalFlag"`
	EmailTemplate             struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Info       struct {
			EmailTemplateHref string `json:"emailTemplate_href"`
		} `json:"_info"`
	} `json:"emailTemplate,omitempty"`
	Info struct {
		LastUpdated time.Time `json:"lastUpdated"`
		UpdatedBy   string    `json:"updatedBy"`
		DateEntered time.Time `json:"dateEntered"`
		EnteredBy   string    `json:"enteredBy"`
	} `json:"_info"`
}

//Ticket is a struct to hold the unmarshaled JSON data when making a call to the Service API
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
		//It's not always a string.  I need to somehow *puts on sunglasses*... reflect on this as this is a dynamic type.
		Value string `json:"value"`
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

//Board is a struct to hold the unmarshaled JSON data when making a call to the Service API
type Board struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			LocationHref string `json:"location_href"`
		} `json:"_info"`
	} `json:"location"`
	Department struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			DepartmentHref string `json:"department_href"`
		} `json:"_info"`
	} `json:"department"`
	InactiveFlag    bool `json:"inactiveFlag"`
	SignOffTemplate struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"signOffTemplate"`
	SendToContactFlag  bool `json:"sendToContactFlag"`
	SendToResourceFlag bool `json:"sendToResourceFlag"`
	ProjectFlag        bool `json:"projectFlag"`
	BoardIcon          struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			Filename             string `json:"filename"`
			DocumentHref         string `json:"document_href"`
			DocumentDownloadHref string `json:"documentDownload_href"`
		} `json:"_info"`
	} `json:"boardIcon,omitempty"`
	BillTicketsAfterClosedFlag    bool `json:"billTicketsAfterClosedFlag"`
	BillTicketSeparatelyFlag      bool `json:"billTicketSeparatelyFlag"`
	BillUnapprovedTimeExpenseFlag bool `json:"billUnapprovedTimeExpenseFlag"`
	OverrideBillingSetupFlag      bool `json:"overrideBillingSetupFlag"`
	DispatchMember                struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			MemberHref string `json:"member_href"`
		} `json:"_info"`
	} `json:"dispatchMember,omitempty"`
	ServiceManagerMember struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			MemberHref string `json:"member_href"`
		} `json:"_info"`
	} `json:"serviceManagerMember,omitempty"`
	DutyManagerMember struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			MemberHref string `json:"member_href"`
		} `json:"_info"`
	} `json:"dutyManagerMember,omitempty"`
	OncallMember struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			MemberHref string `json:"member_href"`
		} `json:"_info"`
	} `json:"oncallMember,omitempty"`
	WorkRole struct {
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
	BillTime                            string `json:"billTime"`
	BillExpense                         string `json:"billExpense"`
	BillProduct                         string `json:"billProduct"`
	AutoAssignNewTicketsFlag            bool   `json:"autoAssignNewTicketsFlag"`
	AutoAssignNewECTicketsFlag          bool   `json:"autoAssignNewECTicketsFlag"`
	AutoAssignNewPortalTicketsFlag      bool   `json:"autoAssignNewPortalTicketsFlag"`
	DiscussionsLockedFlag               bool   `json:"discussionsLockedFlag"`
	TimeEntryLockedFlag                 bool   `json:"timeEntryLockedFlag"`
	NotifyEmailFrom                     string `json:"notifyEmailFrom,omitempty"`
	NotifyEmailFromName                 string `json:"notifyEmailFromName"`
	ClosedLoopDiscussionsFlag           bool   `json:"closedLoopDiscussionsFlag"`
	ClosedLoopResolutionFlag            bool   `json:"closedLoopResolutionFlag"`
	ClosedLoopInternalAnalysisFlag      bool   `json:"closedLoopInternalAnalysisFlag"`
	TimeEntryDiscussionFlag             bool   `json:"timeEntryDiscussionFlag"`
	TimeEntryResolutionFlag             bool   `json:"timeEntryResolutionFlag"`
	TimeEntryInternalAnalysisFlag       bool   `json:"timeEntryInternalAnalysisFlag"`
	ProblemSort                         string `json:"problemSort"`
	ResolutionSort                      string `json:"resolutionSort"`
	InternalAnalysisSort                string `json:"internalAnalysisSort"`
	EmailConnectorAllowReopenClosedFlag bool   `json:"emailConnectorAllowReopenClosedFlag"`
	EmailConnectorReopenStatus          struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"emailConnectorReopenStatus,omitempty"`
	EmailConnectorReopenResourcesFlag   bool   `json:"emailConnectorReopenResourcesFlag"`
	EmailConnectorNewTicketNoMatchFlag  bool   `json:"emailConnectorNewTicketNoMatchFlag"`
	EmailConnectorNeverReopenByDaysFlag bool   `json:"emailConnectorNeverReopenByDaysFlag"`
	EmailConnectorReopenDaysLimit       int    `json:"emailConnectorReopenDaysLimit"`
	UseMemberDisplayNameFlag            bool   `json:"useMemberDisplayNameFlag"`
	SendToCCFlag                        bool   `json:"sendToCCFlag"`
	AutoAssignTicketOwnerFlag           bool   `json:"autoAssignTicketOwnerFlag"`
	ClosedLoopAllFlag                   bool   `json:"closedLoopAllFlag"`
	AllSort                             string `json:"allSort"`
	Info                                struct {
		LastUpdated time.Time `json:"lastUpdated"`
		UpdatedBy   string    `json:"updatedBy"`
	} `json:"_info"`
	ShowDependenciesFlag bool `json:"showDependenciesFlag,omitempty"`
	ShowEstimatesFlag    bool `json:"showEstimatesFlag,omitempty"`
	AutoCloseStatus      struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"autoCloseStatus,omitempty"`
}

type Source struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DefaultFlag bool   `json:"defaultFlag"`
	Info        struct {
		LastUpdated time.Time `json:"lastUpdated"`
		UpdatedBy   string    `json:"updatedBy"`
	} `json:"_info"`
	EnteredBy   string    `json:"enteredBy"`
	DateEntered time.Time `json:"dateEntered"`
}

//TimeEntryReference is a struct to hold the unmarshaled JSON data when making a call to the Service API
//TBD: For some reason the Info struct contained in TimeEntryReference does get data when the JSON is unmarshaled into this struct.  The ID works fine
type TimeEntryReference struct {
	ID   int
	Info struct {
		Notes    string `json:"notes"`
		TimeHref string `json:"time_href"`
	} `json:"_info"`
}

//ConfigurationReference is a struct to hold a reference to Configuration data in the Connectwise API
//TBD: Do these "reference" types really need to be exported???
type ConfigurationReference struct {
	ID               int    `json:"id"`
	DeviceIdentifier string `json:"deviceIdentifier"`
	Info             struct {
		Name              string `json:"name"`
		ConfigurationHref string `json:"configuration_href"`
	} `json:"_info"`
}

//GetTicketByID expects a ticket ID and returns a pointer to a Ticket struct
func (cw *Site) GetTicketByID(ticketID int) (*Ticket, error) {
	req := cw.NewRequest(fmt.Sprintf("/service/tickets/%d", ticketID), "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	ticket := &Ticket{}
	err = json.Unmarshal(req.Body, ticket)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return ticket, nil
}

//GetTicketTimeEntriesByID expects a ticket ID and returns a pointer a to a slice of TimeEntryReference's, all the time entries attached to that ticket
func (cw *Site) GetTicketTimeEntriesByID(ticketID int) (*[]TimeEntryReference, error) {

	req := cw.NewRequest(fmt.Sprintf("/service/tickets/%d/timeentries", ticketID), "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	timeEntryReference := &[]TimeEntryReference{}
	err = json.Unmarshal(req.Body, timeEntryReference)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return timeEntryReference, nil
}

//TicketTimeEntryCount returns the number of time entries on a ticket in ConnectWise
func (cw *Site) TicketTimeEntryCount(ticketID int) (int, error) {
	req := cw.NewRequest(fmt.Sprintf("/service/tickets/%d/timeentries/count", ticketID), "GET", nil)
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

//GetTicketConfigurationsByID expects a ticket ID and returns a pointer to a slice of the configurations attached to the ticket
func (cw *Site) GetTicketConfigurationsByID(ticketID int) (*[]ConfigurationReference, error) {
	req := cw.NewRequest(fmt.Sprintf("/service/tickets/%d/configurations", ticketID), "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	configurationReference := &[]ConfigurationReference{}
	err = json.Unmarshal(req.Body, configurationReference)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return configurationReference, nil
}

//GetServiceTeams returns a pointer to a slice of the service teams in Connectwise
func (cw *Site) GetServiceTeams() (*[]ServiceTeam, error) {
	req := cw.NewRequest("/service/teams", "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	serviceTeam := &[]ServiceTeam{}
	err = json.Unmarshal(req.Body, serviceTeam)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return serviceTeam, nil
}

//GetBoardTeams returns a pointer to a slice of the service teams in Connectwise
func (cw *Site) GetBoardTeams(boardID int) (*[]BoardTeam, error) {
	req := cw.NewRequest(fmt.Sprintf("/service/boards/%d/teams", boardID), "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	boardTeam := &[]BoardTeam{}
	err = json.Unmarshal(req.Body, boardTeam)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return boardTeam, nil
}

func (cw *Site) GetBoardStatuses(boardID int) (*[]BoardStatus, error) {
	req := cw.NewRequest(fmt.Sprintf("/service/boards/%d/statuses", boardID), "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	boardStatus := &[]BoardStatus{}
	err = json.Unmarshal(req.Body, boardStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return boardStatus, nil

}

//GetBoardTeamByName returns a pointer to a board team in Connectwise
func (cw *Site) GetBoardTeamByName(boardID int, teamName string) (*BoardTeam, error) {
	req := cw.NewRequest(fmt.Sprintf("/service/boards/%d/teams", boardID), "GET", nil)
	req.URLValues.Add("conditions", "name=\""+teamName+"\"")
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	boardTeam := &[]BoardTeam{}
	err = json.Unmarshal(req.Body, boardTeam)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	if len(*boardTeam) == 0 {
		return nil, fmt.Errorf("connectwise returned no results for %d/%s", boardID, teamName)
	}

	return &(*boardTeam)[0], nil
}

//GetBoards returns a pointer to a slice of service boards
func (cw *Site) GetBoards() (*[]Board, error) {
	req := cw.NewRequest("/service/boards", "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	board := &[]Board{}
	err = json.Unmarshal(req.Body, board)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return board, nil
}

//AssignTicketToTeam will set the team/id of a ticket
func (cw *Site) AssignTicketToTeam(ticketID, teamID int) (*Ticket, error) {
	patches := &[]PatchString{}
	patch := &PatchString{
		Op:    "replace",
		Path:  "team/id",
		Value: strconv.Itoa(teamID)}
	*patches = append(*patches, *patch)

	patchBody, err := json.Marshal(patches)
	if err != nil {
		return nil, fmt.Errorf("could not marhsal patch json to byte slice: %s", err)
	}

	req := cw.NewRequest(fmt.Sprintf("/service/tickets/%d", ticketID), "PATCH", patchBody)
	err = req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	ticket := &Ticket{}
	err = json.Unmarshal(req.Body, ticket)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return ticket, nil
}

//GetSources will return a pointer to a slice of sources.  Sources are referring to Ticket.Source field
func (cw *Site) GetSources() (*[]Source, error) {
	req := cw.NewRequest("/service/sources", "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	source := &[]Source{}
	err = json.Unmarshal(req.Body, source)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return source, nil
}
