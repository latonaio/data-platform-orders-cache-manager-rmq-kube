package dpfm_api_output_formatter

type OrdersDetailList struct {
	Header  OrdersDetailHeader  `json:"Header"`
	Details []OrdersItemSummary `json:"Details"`
}
type OrdersItemSummary struct {
	OrderItem                   int    `json:"OrderItem"`
	Product                     string `json:"Product"`
	OrderItemTextByBuyer        string `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller       string `json:"OrderItemTextBySeller"`
	OrderQuantityInDeliveryUnit string `json:"OrderQuantityInDeliveryUnit"`
	DeliveryUnit                string `json:"DeliveryUnit"`
	ConditionRateValue          string `json:"ConditionRateValue"`
	RequestedDeliveryDate       string `json:"RequestedDeliveryDate"`
	NetAmount                   string `json:"NetAmount"`
	IsCanceled                  bool   `json:"IsCanceled"`
}

type OrdersDetailHeader struct {
	Index int    `json:"Index"`
	Key   string `json:"Key"`
}
