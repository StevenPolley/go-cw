package connectwise

import (
	"encoding/json"
	"fmt"
)

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
	CustomFields struct {
		ID               int
		Caption          string
		Type             string
		EntryMethod      string
		NumberOfDecimals int
		Value            string
	}
}

func (cw *ConnectwiseSite) GetTimeEntryByID(timeEntryID int) *TimeEntry {

	Url := cw.BuildUrl(fmt.Sprintf("/time/entries/%d", timeEntryID))

	body := cw.GetRequest(Url)
	fmt.Print(string(body))

	timeEntry := TimeEntry{}
	check(json.Unmarshal(body, &timeEntry))

	return &timeEntry
}
