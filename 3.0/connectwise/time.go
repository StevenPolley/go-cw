package connectwise

import (
	"encoding/json"
	"fmt"
)

//TimeEntry is a struct to hold the unmarshaled JSON data when making a call to the Time API
type TimeEntry struct {
	ID      int `json:"id,omitempty"`
	Company struct {
		ID         int    `json:"id,omitempty"`
		Identifier string `json:"identifier,omitempty"`
		Name       string `json:"name,omitempty"`
		Info       struct {
			CompanyHref string `json:"company_href,omitempty"`
			MobileGUID  string `json:"mobileGuid,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"company,omitempty"`
	ChargeToID   int    `json:"chargeToId,omitempty"`
	ChargeToType string `json:"chargeToType,omitempty"`
	Member       struct {
		ID         int    `json:"id,omitempty"`
		Identifier string `json:"identifier,omitempty"`
		Name       string `json:"name,omitempty"`
		Info       struct {
			MemberHref string `json:"member_href,omitempty"`
			ImageHref  string `json:"image_href,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"member,omitempty"`
	LocationID     int `json:"locationId,omitempty"`
	BusinessUnitID int `json:"businessUnitId,omitempty"`
	WorkType       struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			WorkTypeHref string `json:"workType_href,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"workType,omitempty"`
	WorkRole struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			WorkRoleHref string `json:"workRole_href,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"workRole,omitempty"`
	TimeStart                  string  `json:"timeStart,omitempty"`
	TimeEnd                    string  `json:"timeEnd,omitempty"`
	HoursDeduct                float64 `json:"hoursDeduct,omitempty"`
	ActualHours                float64 `json:"actualHours,omitempty"`
	BillableOption             string  `json:"billableOption,omitempty"`
	Notes                      string  `json:"notes,omitempty"`
	AddToDetailDescriptionFlag bool    `json:"addToDetailDescriptionFlag,omitempty"`
	AddToInternalAnalysisFlag  bool    `json:"addToInternalAnalysisFlag,omitempty"`
	AddToResolutionFlag        bool    `json:"addToResolutionFlag,omitempty"`
	EmailResourceFlag          bool    `json:"emailResourceFlag,omitempty"`
	EmailContactFlag           bool    `json:"emailContactFlag,omitempty"`
	EmailCcFlag                bool    `json:"emailCcFlag,omitempty"`
	HoursBilled                float64 `json:"hoursBilled,omitempty"`
	EnteredBy                  string  `json:"enteredBy,omitempty"`
	DateEntered                string  `json:"dateEntered,omitempty"`
	MobileGUID                 string  `json:"mobileGuid,omitempty"`
	HourlyRate                 float64 `json:"hourlyRate,omitempty"`
	TimeSheet                  struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			TimeSheetHref string `json:"timeSheet_href,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"timeSheet,omitempty"`
	Status string `json:"status,omitempty"`
	Info   struct {
		LastUpdated        string `json:"lastUpdated,omitempty"`
		UpdatedBy          string `json:"updatedBy,omitempty"`
		ChargeToMobileGUID string `json:"chargeToMobileGuid,omitempty"`
	} `json:"_info,omitempty"`
	CustomFields []struct {
		ID               int
		Caption          string
		Type             string
		EntryMethod      string
		NumberOfDecimals int
		Value            string
	} `json:",omitempty"`
}

//TimeEntryPost - We need a special struct for posting anything because Golang will not omit an empty struct, and CW doesn't like that.  FFS KEN THOMPSON
type TimeEntryPost struct {
	ID int `json:"id,omitempty"`
	/*
		Company struct {
			ID         int    `json:"id,omitempty"`
			Identifier string `json:"identifier,omitempty"`
			Name       string `json:"name,omitempty"`
		} `json:"company,omitempty"`*/
	ChargeToID   int    `json:"chargeToId,omitempty"`
	ChargeToType string `json:"chargeToType,omitempty"`
	/*Member       struct {
		ID         int    `json:"id,omitempty"`
		Identifier string `json:"identifier,omitempty"`
		Name       string `json:"name,omitempty"`
	} `json:"member,omitempty"`
	*/
	LocationID     int `json:"locationId,omitempty"`
	BusinessUnitID int `json:"businessUnitId,omitempty"`
	/*	WorkType       struct {
			ID   int    `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
			Info struct {
				WorkTypeHref string `json:"workType_href,omitempty"`
			} `json:"_info,omitempty"`
		} `json:"workType,omitempty"`
	*/
	TimeStart                  string  `json:"timeStart,omitempty"`
	TimeEnd                    string  `json:"timeEnd,omitempty"`
	HoursDeduct                float64 `json:"hoursDeduct,omitempty"`
	ActualHours                float64 `json:"actualHours,omitempty"`
	BillableOption             string  `json:"billableOption,omitempty"`
	Notes                      string  `json:"notes,omitempty"`
	AddToDetailDescriptionFlag bool    `json:"addToDetailDescriptionFlag,omitempty"`
	AddToInternalAnalysisFlag  bool    `json:"addToInternalAnalysisFlag,omitempty"`
	AddToResolutionFlag        bool    `json:"addToResolutionFlag,omitempty"`
	EmailResourceFlag          bool    `json:"emailResourceFlag,omitempty"`
	EmailContactFlag           bool    `json:"emailContactFlag,omitempty"`
	EmailCcFlag                bool    `json:"emailCcFlag,omitempty"`
	HoursBilled                float64 `json:"hoursBilled,omitempty"`
	EnteredBy                  string  `json:"enteredBy,omitempty"`
	DateEntered                string  `json:"dateEntered,omitempty"`
	MobileGUID                 string  `json:"mobileGuid,omitempty"`
	HourlyRate                 float64 `json:"hourlyRate,omitempty"`
	/*TimeSheet                  struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Info struct {
			TimeSheetHref string `json:"timeSheet_href,omitempty"`
		} `json:"_info,omitempty"`
	} `json:"timeSheet,omitempty"`
	*/
	Status string `json:"status,omitempty"`
}

//GetTimeEntryByID expects a time entry ID and will return a pointer to a TimeEntry struct
func (cw *Site) GetTimeEntryByID(timeEntryID int) (*TimeEntry, error) {
	req := cw.NewRequest(fmt.Sprintf("/time/entries/%d", timeEntryID), "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	timeEntry := &TimeEntry{}
	err = json.Unmarshal(req.Body, timeEntry)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return timeEntry, nil
}

func (cw *Site) GetTimeEntriesByMember(memberIdentifier string) (*[]TimeEntry, error) {
	req := cw.NewRequest("/time/entries", "GET", nil)
	req.URLValues.Add("conditions", "member/identifier=\""+memberIdentifier+"\"")
	req.URLValues.Add("orderBy", "id desc")
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	timeEntries := &[]TimeEntry{}
	err = json.Unmarshal(req.Body, timeEntries)
	if err != nil {
		fmt.Println(string(req.Body))
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)

	}

	return timeEntries, nil
}

//PostTimeEntry accepts a preconfigured pointer to connectwise.TimeEntry struct with all the data that is to be posted.
func (cw *Site) PostTimeEntry(timeEntry *TimeEntryPost) (*TimeEntry, error) {
	js, err := json.Marshal(timeEntry)
	if err != nil {
		return nil, fmt.Errorf("could not marshal timeEntry struct to json bytes: %s", err)
	}

	req := cw.NewRequest("/time/entries", "POST", js)
	err = req.Do()
	if err != nil {
		return nil, fmt.Errorf("post request failed for %s: %s", req.RestAction, err)
	}

	te := &TimeEntry{}
	err = json.Unmarshal(req.Body, te)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return te, nil
}
