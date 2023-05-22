package dpfm_api_output_formatter

type OrdersList struct {
	Orders []Orders `json:"Orders"`
}
type Orders struct {
	OrderID        int     `json:"OrderID"`
	SellerName     string  `json:"SellerName"`
	BuyerName      string  `json:"BuyerName"`
	DeliveryStatus *string `json:"DeliveryStatus"`

	OrderDate           *string `json:"OrderDate"`
	PaymentTerms        *string `json:"PaymentTerms"`
	PaymentMethod       *string `json:"PaymentMethod"`
	TransactionCurrency *string `json:"TransactionCurrency"`
	OrderType           *string `json:"OrderType"`
}
