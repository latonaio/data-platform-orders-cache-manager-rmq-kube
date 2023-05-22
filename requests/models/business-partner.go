package models

type BusinessPartnerReq struct {
	ConnectionKey     string     `json:"connection_key"`
	Result            bool       `json:"result"`
	RedisKey          string     `json:"redis_key"`
	Filepath          string     `json:"filepath"`
	APIStatusCode     int        `json:"api_status_code"`
	RuntimeSessionID  string     `json:"runtime_session_id"`
	BusinessPartnerID *int       `json:"business_partner"`
	ServiceLabel      string     `json:"service_label"`
	APIType           string     `json:"api_type"`
	General           BPGeneral  `json:"BusinessPartnerGeneral"`
	Generals          BPGenerals `json:"BusinessPartnerGenerals"`
	APISchema         string     `json:"api_schema"`
	Accepter          []string   `json:"accepter"`
	Deleted           bool       `json:"deleted"`
}

type BPGeneral struct {
	BusinessPartner               int          `json:"BusinessPartner"`
	BusinessPartnerFullName       *string      `json:"BusinessPartnerFullName"`
	BusinessPartnerName           *string      `json:"BusinessPartnerName"`
	CreationDate                  *string      `json:"CreationDate"`
	CreationTime                  *string      `json:"CreationTime"`
	Industry                      *string      `json:"Industry"`
	LegalEntityRegistration       *string      `json:"LegalEntityRegistration"`
	Country                       *string      `json:"Country"`
	Language                      *string      `json:"Language"`
	Currency                      *string      `json:"Currency"`
	LastChangeDate                *string      `json:"LastChangeDate"`
	LastChangeTime                *string      `json:"LastChangeTime"`
	OrganizationBPName1           *string      `json:"OrganizationBPName1"`
	OrganizationBPName2           *string      `json:"OrganizationBPName2"`
	OrganizationBPName3           *string      `json:"OrganizationBPName3"`
	OrganizationBPName4           *string      `json:"OrganizationBPName4"`
	BPTag1                        *string      `json:"BPTag1"`
	BPTag2                        *string      `json:"BPTag2"`
	BPTag3                        *string      `json:"BPTag3"`
	BPTag4                        *string      `json:"BPTag4"`
	OrganizationFoundationDate    *string      `json:"OrganizationFoundationDate"`
	OrganizationLiquidationDate   *string      `json:"OrganizationLiquidationDate"`
	BusinessPartnerBirthplaceName *string      `json:"BusinessPartnerBirthplaceName"`
	BusinessPartnerDeathDate      *string      `json:"BusinessPartnerDeathDate"`
	BusinessPartnerIsBlocked      *bool        `json:"BusinessPartnerIsBlocked"`
	GroupBusinessPartnerName1     *string      `json:"GroupBusinessPartnerName1"`
	GroupBusinessPartnerName2     *string      `json:"GroupBusinessPartnerName2"`
	AddressID                     *int         `json:"AddressID"`
	BusinessPartnerIDByExtSystem  *string      `json:"BusinessPartnerIDByExtSystem"`
	IsMarkedForDeletion           *bool        `json:"IsMarkedForDeletion"`
	GeneralPDF                    BPGeneralPDF `json:"GeneralPDF"`
	Role                          Role         `json:"Role"`
	FinInst                       FinInst      `json:"FinInst"`
	Relationship                  Relationship `json:"Relationship"`
	Accounting                    BPAccounting `json:"Accounting"`
}
type BPGenerals struct {
	BusinessPartners []int `json:"BusinessPartners"`
}

type BPGeneralPDF struct {
	BusinessPartner          int     `json:"BusinessPartner"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FileName                 *string `json:"FileName"`
}

type Role struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	BusinessPartnerRole string `json:"BusinessPartnerRole"`
	ValidityEndDate     string `json:"ValidityEndDate"`
	ValidityStartDate   string `json:"ValidityStartDate"`
}

type FinInst struct {
	BusinessPartner           int     `json:"BusinessPartner"`
	FinInstIdentification     int     `json:"FinInstIdentification"`
	ValidityEndDate           string  `json:"ValidityEndDate"`
	ValidityStartDate         string  `json:"ValidityStartDate"`
	FinInstCountry            *string `json:"FinInstCountry"`
	FinInstCode               *string `json:"FinInstCode"`
	FinInstBranchCode         *string `json:"FinInstBranchCode"`
	FinInstFullCode           *string `json:"FinInstFullCode"`
	FinInstName               *string `json:"FinInstName"`
	FinInstBranchName         *string `json:"FinInstBranchName"`
	SWIFTCode                 *string `json:"SWIFTCode"`
	InternalFinInstCustomerID *int    `json:"InternalFinInstCustomerID"`
	InternalFinInstAccountID  *int    `json:"InternalFinInstAccountID"`
	FinInstControlKey         *string `json:"FinInstControlKey"`
	FinInstAccountName        *string `json:"FinInstAccountName"`
	FinInstAccount            *string `json:"FinInstAccount"`
	HouseBank                 *string `json:"HouseBank"`
	HouseBankAccount          *string `json:"HouseBankAccount"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}

type Relationship struct {
	BusinessPartner             int     `json:"BusinessPartner"`
	RelationshipNumber          int     `json:"RelationshipNumber"`
	ValidityEndDate             string  `json:"ValidityEndDate"`
	ValidityStartDate           string  `json:"ValidityStartDate"`
	RelationshipCategory        *string `json:"RelationshipCategory"`
	RelationshipBusinessPartner *int    `json:"RelationshipBusinessPartner"`
	BusinessPartnerPerson       *string `json:"BusinessPartnerPerson"`
	IsStandardRelationship      *bool   `json:"IsStandardRelationship"`
	IsMarkedForDeletion         *bool   `json:"IsMarkedForDeletion"`
}

type BPAccounting struct {
	BusinessPartner     int     `json:"BusinessPartner"`
	ChartOfAccounts     *string `json:"ChartOfAccounts"`
	FiscalYearVariant   *string `json:"FiscalYearVariant"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}
