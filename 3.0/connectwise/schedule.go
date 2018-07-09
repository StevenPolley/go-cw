package connectwise

import (
	"encoding/json"
	"fmt"
)

type Calendar struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	HolidayList struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			HolidayListHref string `json:"holidayList_href"`
		} `json:"_info"`
	} `json:"holidayList"`
	MondayStartTime    string `json:"mondayStartTime"`
	MondayEndTime      string `json:"mondayEndTime"`
	TuesdayStartTime   string `json:"tuesdayStartTime"`
	TuesdayEndTime     string `json:"tuesdayEndTime"`
	WednesdayStartTime string `json:"wednesdayStartTime"`
	WednesdayEndTime   string `json:"wednesdayEndTime"`
	ThursdayStartTime  string `json:"thursdayStartTime"`
	ThursdayEndTime    string `json:"thursdayEndTime"`
	FridayStartTime    string `json:"fridayStartTime"`
	FridayEndTime      string `json:"fridayEndTime"`
	Info               struct {
		LastUpdated string `json:"lastUpdated"`
		UpdatedBy   string `json:"updatedBy"`
	} `json:"_info"`
}

type ScheduleEntry struct {
	ID       int    `json:"id"`
	ObjectID int    `json:"objectId"`
	Name     string `json:"name"`
	Member   struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			MemberHref string `json:"member_href"`
		} `json:"_info"`
	} `json:"member,omitempty"`
	DateStart string `json:"dateStart,omitempty"`
	DateEnd   string `json:"dateEnd,omitempty"`
	Reminder  struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"reminder,omitempty"`
	Status struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			StatusHref string `json:"status_href"`
		} `json:"_info"`
	} `json:"status,omitempty"`
	Type struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Info       struct {
			TypeHref string `json:"type_href"`
		} `json:"_info"`
	} `json:"type"`
	DoneFlag         bool    `json:"doneFlag"`
	AcknowledgedFlag bool    `json:"acknowledgedFlag"`
	OwnerFlag        bool    `json:"ownerFlag"`
	MeetingFlag      bool    `json:"meetingFlag"`
	MobileGUID       string  `json:"mobileGuid"`
	Hours            float64 `json:"hours"`
	Info             struct {
		UpdatedBy        string `json:"updatedBy"`
		ObjectMobileGUID string `json:"objectMobileGuid"`
	} `json:"_info"`
}

//GetCalendars returns a list of global connectwise calendars - not used for member calendars
func (cw *Site) GetCalendars() (*[]Calendar, error) {
	req := cw.NewRequest("/schedule/calendars", "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	calendar := &[]Calendar{}
	err = json.Unmarshal(req.Body, calendar)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return calendar, nil
}

//ScheduleEntryCount returns the number of companies in ConnectWise
func (cw *Site) ScheduleEntryCount() (int, error) {
	req := cw.NewRequest("/schedule/entries/count", "GET", nil)
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

//GetScheduleEntriesOnDay expects a date in the format of "YYYY-MM-DD" and returns a pointer to a slice of ScheduleEntries on that day
func (cw *Site) GetScheduleEntriesOnDay(day string) (*[]ScheduleEntry, error) {
	req := cw.NewRequest("/schedule/entries", "GET", nil)
	//	req.URLValues.Add("orderBy", "dateStart desc")
	req.URLValues.Add("conditions", "datestart=2036-02-28T18:18:00Z")
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	scheduleEntry := &[]ScheduleEntry{}
	err = json.Unmarshal(req.Body, scheduleEntry)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return scheduleEntry, nil
}
