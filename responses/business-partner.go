package responses

import (
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

// type BusinessPartnerRes struct {
// 	ConnectionKey       string     `json:"connection_key"`
// 	Result              bool       `json:"result"`
// 	RedisKey            string     `json:"redis_key"`
// 	Filepath            string     `json:"filepath"`
// 	APIStatusCode       int        `json:"api_status_code"`
// 	RuntimeSessionID    string     `json:"runtime_session_id"`
// 	BusinessPartnerID   *int       `json:"business_partner"`
// 	ServiceLabel        string     `json:"service_label"`
// 	APIType             string     `json:"api_type"`
// 	Message             *BPMessage `json:"message"`
// 	APISchema           string     `json:"api_schema"`
// 	Accepter            []string   `json:"accepter"`
// 	Deleted             bool       `json:"deleted"`
// 	SQLUpdateResult     *bool      `json:"sql_update_result"`
// 	SQLUpdateError      string     `json:"sql_update_error"`
// 	SubfuncResult       *bool      `json:"subfunc_result"`
// 	SubfuncError        string     `json:"subfunc_error"`
// 	ExconfResult        *bool      `json:"exconf_result"`
// 	ExconfError         string     `json:"exconf_error"`
// 	APIProcessingResult *bool      `json:"api_processing_result"`
// 	APIProcessingError  string     `json:"api_processing_error"`
// }

type BPMessage struct {
	General      *BPGeneral    `json:"General"`
	Generals     *[]BPGeneral  `json:"Generals"`
	Role         *Role         `json:"Role"`
	FinInst      *FinInst      `json:"FinInst"`
	Relationship *Relationship `json:"Relationship"`
	Accounting   *Accounting   `json:"Accounting"`
}

type BusinessPartnerGeneral struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Deleted       string `json:"deleted"`
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
	GeneralPDF                    GeneralPDF   `json:"GeneralPDF"`
	Role                          Role         `json:"Role"`
	FinInst                       FinInst      `json:"FinInst"`
	Relationship                  Relationship `json:"Relationship"`
	Accounting                    Accounting   `json:"Accounting"`
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

func CreateBusinessPartnerRes(msg rabbitmq.RabbitmqMessage) (*BusinessPartnerRes, error) {
	res := BusinessPartnerRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
