package dpfm_api_input_reader

import (
	"encoding/json"
	"fmt"
	"strconv"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type OrdersDetailList struct {
	UIKeyGeneralUserID          string                 `json:"ui_key_general_user_id"`
	UIKeyGeneralUserLanguage    string                 `json:"ui_key_general_user_language"`
	UIKeyGeneralBusinessPartner string                 `json:"ui_key_general_business_partner"`
	UIFunction                  string                 `json:"ui_function"`
	UIKeyFunctionURL            string                 `json:"ui_key_function_url"`
	RuntimeSessionID            string                 `json:"runtime_session_id"`
	Params                      OrdersDetailListParams `json:"Params"`
}
type OrdersDetailListParams struct {
	User                            *string `json:"User"`
	OrderID                         *int    `json:"OrderId"`
	ItemCompleteDeliveryIsDefined   *bool   `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus              *string `json:"ItemDeliveryStatus"`
	ItemDeliveryBlockStatus         *bool   `json:"ItemDeliveryBlockStatus"`
	ItemIsCanceled                  *bool   `json:"ItemIsCanceled"`
	ItemIsDeleted                   *bool   `json:"ItemIsDeleted"`
	HeaderCompleteDeliveryIsDefined bool    `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryBlockStatus       bool    `json:"HeaderDeliveryBlockStatus"`
	HeaderDeliveryStatus            string  `json:"HeaderDeliveryStatus"`
	HeaderIsCanceled                bool    `json:"HeaderIsCanceled"`
	HeaderIsDeleted                 bool    `json:"HeaderIsDeleted"`
	Buyer                           *int    `json:"Buyer"`
	Seller                          *int    `json:"Seller"`
	Language                        string  `json:"Language"`
	BusinessPartner                 int     `json:"BusinessPartner"`
}

func ReadOrdersDetailList(msg rabbitmq.RabbitmqMessage) *OrdersDetailList {
	d := OrdersDetailList{}
	err := json.Unmarshal(msg.Raw(), &d)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	d.UIKeyGeneralUserID = d.UIKeyGeneralUserID[len("orders/detail/list/userID="):]
	lang := d.UIKeyGeneralUserLanguage[len("orders/detail/list/language="):]
	bp, err := strconv.Atoi(d.UIKeyGeneralBusinessPartner[len("orders/detail/list/businessPartnerID="):])
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	d.Params.Language = lang
	d.Params.BusinessPartner = bp
	return &d
}
