package responses

import (
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

// type ProductGroupRes struct {
// 	ConnectionKey       string     `json:"connection_key"`
// 	Result              bool       `json:"result"`
// 	RedisKey            string     `json:"redis_key"`
// 	Filepath            string     `json:"filepath"`
// 	APIStatusCode       int        `json:"api_status_code"`
// 	RuntimeSessionID    string     `json:"runtime_session_id"`
// 	BusinessPartnerID   *int       `json:"business_partner"`
// 	ServiceLabel        string     `json:"service_label"`
// 	APIType             string     `json:"api_type"`
// 	Message             *PGMessage `json:"message"`
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

type PGMessage struct {
	ProductGroup     *ProductGroup     `json:"ProductGroup"`
	ProductGroupText *ProductGroupText `json:"ProductGroupText"`
}

type ProductGroupReads struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Deleted       string `json:"deleted"`
}

type ProductGroup struct {
	ProductGroup string `json:"ProductGroup"`
}

type ProductGroupText struct {
	ProductGroup     string  `json:"ProductGroup"`
	Language         string  `json:"Language"`
	ProductGroupName *string `json:"ProductGroupName"`
}

func CreateProductGroupRes(msg rabbitmq.RabbitmqMessage) (*ProductGroupRes, error) {
	res := ProductGroupRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
