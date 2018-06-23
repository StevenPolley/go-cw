package connectwise

import (
	"fmt"
)

func (cw *ConnectwiseSite) GetTimeEntryByID(timeEntryID int) {

	Url := cw.BuildUrl(fmt.Sprintf("/time/entries/%d", timeEntryID))

	body := cw.GetRequest(Url)
	fmt.Print(string(body))

	//timeEntry := TimeEntry{}
	//	check(json.Unmarshal(body, &timeEntry))

	//	return &timeEntry
}
