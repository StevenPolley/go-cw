package connectwise

import (
	"encoding/json"
	"fmt"
)

//TimeEntry is a struct to hold the unmarshaled JSON data when making a call to the Time API
type TimeEntry struct {
	ID      int `json:"id"`
	Company struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			CompanyHref string `json:"company_href"`
			MobileGUID  string `json:"mobileGuid"`
		} `json:"_info"`
	} `json:"company"`
	ChargeToID   int    `json:"chargeToId"`
	ChargeToType string `json:"chargeToType"`
	Member       struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			MemberHref string `json:"member_href"`
			ImageHref  string `json:"image_href"`
		} `json:"_info"`
	} `json:"member"`
	LocationID     int `json:"locationId"`
	BusinessUnitID int `json:"businessUnitId"`
	WorkType       struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			WorkTypeHref string `json:"workType_href"`
		} `json:"_info"`
	} `json:"workType"`
	WorkRole struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			WorkRoleHref string `json:"workRole_href"`
		} `json:"_info"`
	} `json:"workRole"`
	TimeStart                  string  `json:"timeStart"`
	TimeEnd                    string  `json:"timeEnd"`
	HoursDeduct                float64 `json:"hoursDeduct"`
	ActualHours                float64 `json:"actualHours"`
	BillableOption             string  `json:"billableOption"`
	Notes                      string  `json:"notes"`
	AddToDetailDescriptionFlag bool    `json:"addToDetailDescriptionFlag"`
	AddToInternalAnalysisFlag  bool    `json:"addToInternalAnalysisFlag"`
	AddToResolutionFlag        bool    `json:"addToResolutionFlag"`
	EmailResourceFlag          bool    `json:"emailResourceFlag"`
	EmailContactFlag           bool    `json:"emailContactFlag"`
	EmailCcFlag                bool    `json:"emailCcFlag"`
	HoursBilled                float64 `json:"hoursBilled"`
	EnteredBy                  string  `json:"enteredBy"`
	DateEntered                string  `json:"dateEntered"`
	MobileGUID                 string  `json:"mobileGuid"`
	HourlyRate                 float64 `json:"hourlyRate"`
	TimeSheet                  struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			TimeSheetHref string `json:"timeSheet_href"`
		} `json:"_info"`
	} `json:"timeSheet"`
	Status string `json:"status"`
	Info   struct {
		LastUpdated        string `json:"lastUpdated"`
		UpdatedBy          string `json:"updatedBy"`
		ChargeToMobileGUID string `json:"chargeToMobileGuid"`
	} `json:"_info"`
	CustomFields []struct {
		ID               int
		Caption          string
		Type             string
		EntryMethod      string
		NumberOfDecimals int
		Value            string
	}
}

//GetTimeEntryByID expects a time entry ID and will return a pointer to a TimeEntry struct
func (cw *ConnectwiseSite) GetTimeEntryByID(timeEntryID int) (*TimeEntry, error) {
	restAction := fmt.Sprintf("/time/entries/%d", timeEntryID)
	cwurl, err := cw.BuildURL(restAction)
	if err != nil {
		return nil, fmt.Errorf("could not build url %s: %s", restAction, err)
	}

	body, err := cw.GetRequest(cwurl)
	if err != nil {
		return nil, fmt.Errorf("could not get request %s: %s", cwurl, err)
	}

	timeEntry := TimeEntry{}
	err = json.Unmarshal(body, &timeEntry)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return &timeEntry, nil
}
