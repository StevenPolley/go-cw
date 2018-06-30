package connectwise

import (
	"encoding/json"
	"fmt"
	"net/url"
)

//Agreement is a struct to hold the unmarshaled JSON data when making a call to the Finance API
type Agreement struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			TypeHref string `json:"type_href"`
		} `json:"_info"`
	} `json:"type"`
	Company struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			CompanyHref string `json:"company_href"`
		} `json:"_info"`
	} `json:"company"`
	Contact struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			ContactHref string `json:"contact_href"`
		} `json:"_info"`
	} `json:"contact"`
	CustomerPO               string  `json:"customerPO"`
	LocationID               int     `json:"locationId"`
	BusinessUnitID           int     `json:"businessUnitId"`
	RestrictLocationFlag     bool    `json:"restrictLocationFlag"`
	RestrictDepartmentFlag   bool    `json:"restrictDepartmentFlag"`
	StartDate                string  `json:"startDate"`
	EndDate                  string  `json:"endDate,omitempty"`
	NoEndingDateFlag         bool    `json:"noEndingDateFlag"`
	CancelledFlag            bool    `json:"cancelledFlag"`
	ReasonCancelled          string  `json:"reasonCancelled"`
	WorkOrder                string  `json:"workOrder"`
	InternalNotes            string  `json:"internalNotes"`
	ApplicationUnits         string  `json:"applicationUnits"`
	ApplicationLimit         float64 `json:"applicationLimit"`
	ApplicationCycle         string  `json:"applicationCycle,omitempty"`
	ApplicationUnlimitedFlag bool    `json:"applicationUnlimitedFlag"`
	OneTimeFlag              bool    `json:"oneTimeFlag"`
	CoverAgreementTime       bool    `json:"coverAgreementTime"`
	CoverAgreementProduct    bool    `json:"coverAgreementProduct"`
	CoverAgreementExpense    bool    `json:"coverAgreementExpense"`
	CoverSalesTax            bool    `json:"coverSalesTax"`
	CarryOverUnused          bool    `json:"carryOverUnused"`
	AllowOverruns            bool    `json:"allowOverruns"`
	ExpiredDays              int     `json:"expiredDays,omitempty"`
	Limit                    int     `json:"limit,omitempty"`
	ExpireWhenZero           bool    `json:"expireWhenZero"`
	ChargeToFirm             bool    `json:"chargeToFirm"`
	EmployeeCompRate         string  `json:"employeeCompRate"`
	CompHourlyRate           float64 `json:"compHourlyRate"`
	CompLimitAmount          float64 `json:"compLimitAmount"`
	BillCycleID              int     `json:"billCycleId"`
	BillOneTimeFlag          bool    `json:"billOneTimeFlag"`
	BillTermsID              int     `json:"billTermsId"`
	InvoicingCycle           string  `json:"invoicingCycle"`
	BillAmount               float64 `json:"billAmount"`
	Taxable                  bool    `json:"taxable"`
	ProrateFirstBill         float64 `json:"prorateFirstBill,omitempty"`
	BillStartDate            string  `json:"billStartDate"`
	TaxCodeID                int     `json:"taxCodeId"`
	RestrictDownPayment      bool    `json:"restrictDownPayment"`
	ProrateFlag              bool    `json:"prorateFlag"`
	InvoiceDescription       string  `json:"invoiceDescription"`
	TopComment               bool    `json:"topComment"`
	BottomComment            bool    `json:"bottomComment"`
	BillTime                 string  `json:"billTime"`
	BillExpenses             string  `json:"billExpenses"`
	BillProducts             string  `json:"billProducts"`
	BillableTimeInvoice      bool    `json:"billableTimeInvoice"`
	BillableExpenseInvoice   bool    `json:"billableExpenseInvoice"`
	BillableProductInvoice   bool    `json:"billableProductInvoice"`
	Currency                 struct {
		ID      int    `json:"id"`
		Symbol  string `json:"symbol"`
		IsoCode string `json:"isoCode"`
		Name    string `json:"name"`
		Info    struct {
			CurrencyHref string `json:"currency_href"`
		} `json:"_info"`
	} `json:"currency"`
	Info struct {
		LastUpdated string `json:"lastUpdated"`
		UpdatedBy   string `json:"updatedBy"`
	} `json:"_info"`
	DateCancelled         string `json:"dateCancelled,omitempty"`
	SLAID                 int    `json:"slaId,omitempty"`
	EmployeeCompNotExceed string `json:"employeeCompNotExceed,omitempty"`
	BillToCompany         struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			CompanyHref string `json:"company_href"`
		} `json:"_info"`
	} `json:"billToCompany,omitempty"`
	BillToSite struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			SiteHref string `json:"site_href"`
		} `json:"_info"`
	} `json:"billToSite,omitempty"`
}

//GetAgreements returns a list of agreements, not paginated and currently not that usefule
//TBD: Pagination and filtering options still need to be considered
func (cw *ConnectwiseSite) GetAgreements() *[]Agreement {

	//Build the request URL
	cwurl := cw.BuildURL("/finance/agreements")

	body := cw.GetRequest(cwurl)
	agreements := []Agreement{}
	check(json.Unmarshal(body, &agreements))

	return &agreements

}

func (cw *ConnectwiseSite) GetAgreementsByCompanyName(companyName string) *[]Agreement {

	cwurl := cw.BuildURL("/finance/agreements")
	parameters := url.Values{}
	parameters.Add("conditions", "company/name=\""+companyName+"\"")
	cwurl.RawQuery = parameters.Encode()

	body := cw.GetRequest(cwurl)
	agreements := []Agreement{}
	check(json.Unmarshal(body, &agreements))

	return &agreements
}

//GetBillingCycles is not complete
//TBD: Finish this.
func (cw *ConnectwiseSite) GetBillingCycles() {

	cwurl := cw.BuildURL("/finance/billingCycles")

	body := cw.GetRequest(cwurl)
	fmt.Print(string(body))
	//	check(json.Unmarshal(body, &ticket))
}
