package dpfm_api_input_reader

import (
	"encoding/json"
	"fmt"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type OrdersList struct {
	UIKeyGeneralUserID          string           `json:"ui_key_general_user_id"`
	UIKeyGeneralUserLanguage    string           `json:"ui_key_general_user_language"`
	UIKeyGeneralBusinessPartner string           `json:"ui_key_general_business_partner"`
	UIFunction                  string           `json:"ui_function"`
	UIKeyFunctionURL            string           `json:"ui_key_function_url"`
	RuntimeSessionID            string           `json:"runtime_session_id"`
	Params                      OrdersListParams `json:"Params"`
}

type OrdersListParams struct {
	OrderID                         int    `json:"OrderId"`
	User                            string `json:"User"`
	HeaderCompleteDeliveryIsDefined bool   `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryBlockStatus       bool   `json:"HeaderDeliveryBlockStatus"`
	HeaderDeliveryStatus            string `json:"HeaderDeliveryStatus"`
	HeaderIsCanceled                bool   `json:"HeaderIsCanceled"`
	HeaderIsDeleted                 bool   `json:"HeaderIsDeleted"`
	Language                        string `json:"Language"`
	Buyer                           *int   `json:"Buyer"`
	Seller                          *int   `json:"Seller"`
	// BusinessPartner                 int    `json:"BusinessPartner"`
}

func ReadOrdersList(msg rabbitmq.RabbitmqMessage) *OrdersList {
	d := OrdersList{}
	err := json.Unmarshal(msg.Raw(), &d)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	lang := d.UIKeyGeneralUserLanguage[len("orders/list/language="):]
	d.UIKeyGeneralUserID = d.UIKeyGeneralUserID[len("orders/list/userID="):]
	// bp, err := strconv.Atoi(d.UIKeyGeneralBusinessPartner[len("orders/list/businessPartnerID="):])
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// 	return nil
	// }
	d.Params.Language = lang
	// d.Params.BusinessPartner = bp
	return &d
}
