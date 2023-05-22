package models

type ProductStockReq struct {
	ConnectionKey     string       `json:"connection_key"`
	Result            bool         `json:"result"`
	RedisKey          string       `json:"redis_key"`
	Filepath          string       `json:"filepath"`
	APIStatusCode     int          `json:"api_status_code"`
	RuntimeSessionID  string       `json:"runtime_session_id"`
	BusinessPartnerID *int         `json:"business_partner"`
	ServiceLabel      string       `json:"service_label"`
	APIType           string       `json:"api_type"`
	ProductStock      ProductStock `json:"ProductStock"`
	APISchema         string       `json:"api_schema"`
	Accepter          []string     `json:"accepter"`
	Deleted           bool         `json:"deleted"`
}

type ProductStock struct {
	BusinessPartner           int                      `json:"BusinessPartner"`
	Product                   string                   `json:"Product"`
	Plant                     string                   `json:"Plant"`
	StorageLocation           *string                  `json:"StorageLocation"`
	Batch                     *string                  `json:"Batch"`
	OrderID                   *int                     `json:"OrderID"`
	OrderItem                 *int                     `json:"OrderItem"`
	Project                   *string                  `json:"Project"`
	InventoryStockType        *string                  `json:"InventoryStockType"`
	InventorySpecialStockType *string                  `json:"InventorySpecialStockType"`
	ProductStock              *float32                 `json:"ProductStock"`
	ProductStockAvailability  ProductStockAvailability `json:"ProductStockAvailability"`
}

type ProductStockAvailability struct {
	BusinessPartner              int      `json:"BusinessPartner"`
	Product                      string   `json:"Product"`
	Plant                        string   `json:"Plant"`
	Batch                        *string  `json:"Batch"`
	BatchValidityEndDate         *string  `json:"BatchValidityEndDate"`
	OrderID                      *int     `json:"OrderID"`
	OrderItem                    *int     `json:"OrderItem"`
	Project                      *string  `json:"Project"`
	InventoryStockType           *string  `json:"InventoryStockType"`
	InventorySpecialStockType    *string  `json:"InventorySpecialStockType"`
	ProductStockAvailabilityDate string   `json:"ProductStockAvailabilityDate"`
	AvailableProductStock        *float32 `json:"AvailableProductStock"`
}
