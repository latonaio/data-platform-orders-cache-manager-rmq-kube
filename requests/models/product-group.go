package models

type ProductGroupReq struct {
	ConnectionKey     string        `json:"connection_key"`
	Result            bool          `json:"result"`
	RedisKey          string        `json:"redis_key"`
	Filepath          string        `json:"filepath"`
	APIStatusCode     int           `json:"api_status_code"`
	RuntimeSessionID  string        `json:"runtime_session_id"`
	BusinessPartnerID *int          `json:"business_partner"`
	ServiceLabel      string        `json:"service_label"`
	APIType           string        `json:"api_type"`
	ProductGroup      *ProductGroup `json:"ProductGroup"`
	APISchema         string        `json:"api_schema"`
	Accepter          []string      `json:"accepter"`
	Deleted           bool          `json:"deleted"`
}

type ProductGroup struct {
	ProductGroup     string           `json:"ProductGroup"`
	ProductGroupText ProductGroupText `json:"ProductGroupText"`
}

type ProductGroupText struct {
	ProductGroup     string  `json:"ProductGroup"`
	Language         string  `json:"Language"`
	ProductGroupName *string `json:"ProductGroupName"`
}
