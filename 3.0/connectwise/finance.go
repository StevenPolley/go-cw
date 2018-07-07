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

//GetAgreements returns a list of agreements, not paginated and currently not that useful
//TBD: Pagination and filtering options still need to be considered
func (cw *Site) GetAgreements() (*[]Agreement, error) {
	restAction := "/finance/agreements"
	cwurl, err := cw.BuildURL(restAction)
	if err != nil {
		return nil, fmt.Errorf("could not build url %s: %s", restAction, err)
	}

	body, err := cw.GetRequest(cwurl)
	if err != nil {
		return nil, fmt.Errorf("could not get request %s: %s", cwurl, err)
	}
	agreements := []Agreement{}
	err = json.Unmarshal(body, &agreements)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return &agreements, nil

}

//GetAgreementsByCompanyName returns a list of agreements that belong to an exact matching company name
//TBD: Pagination and filtering options still need to be considered
func (cw *Site) GetAgreementsByCompanyName(companyName string) (*[]Agreement, error) {
	restAction := "/finance/agreements"
	cwurl, err := cw.BuildURL(restAction)
	if err != nil {
		return nil, fmt.Errorf("could not build url %s: %s", restAction, err)
	}

	parameters := url.Values{}
	parameters.Add("conditions", "company/name=\""+companyName+"\"")
	cwurl.RawQuery = parameters.Encode()

	body, err := cw.GetRequest(cwurl)
	if err != nil {
		return nil, fmt.Errorf("could not get request %s: %s", cwurl, err)
	}
	agreements := []Agreement{}
	err = json.Unmarshal(body, &agreements)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into struct: %s", err)
	}

	return &agreements, nil
}

//GetBillingCycles is not complete
//TBD: Finish this.
/*
func (cw *Site) GetBillingCycles() {
	restAction := "/finance/billingCycles"
	cwurl, err := cw.BuildURL(restAction)
	if err != nil {
		return nil, fmt.Errorf("could not build url %s: %g", restAction, err)
	}

	body, err := cw.GetRequest(cwurl)
	if err != nil {
		return nil, fmt.Errorf("could not get request %s: %g", cwurl, err)
	}
	fmt.Print(string(body))
	//	check(json.Unmarshal(body, &ticket))
}
*/
