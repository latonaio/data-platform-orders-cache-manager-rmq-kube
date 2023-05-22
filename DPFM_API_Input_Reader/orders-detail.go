package dpfm_api_input_reader

import (
	"encoding/json"
	"fmt"
	"strconv"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type OrdersDetail struct {
	UIKeyGeneralUserID          string             `json:"ui_key_general_user_id"`
	UIKeyGeneralUserLanguage    string             `json:"ui_key_general_user_language"`
	UIKeyGeneralBusinessPartner string             `json:"ui_key_general_business_partner"`
	UIFunction                  string             `json:"ui_function"`
	UIKeyFunctionURL            string             `json:"ui_key_function_url"`
	RuntimeSessionID            string             `json:"runtime_session_id"`
	Params                      OrdersDetailParams `json:"Params"`
}
type OrdersDetailParams struct {
	OrderID         int    `json:"OrderId"`
	OrderItem       int    `json:"OrderItem"`
	Product         string `json:"Product"`
	Language        string `json:"Language"`
	BusinessPartner int    `json:"BusinessPartner"`
}

func ReadOrdersDetail(msg rabbitmq.RabbitmqMessage) *OrdersDetail {
	d := OrdersDetail{}
	err := json.Unmarshal(msg.Raw(), &d)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	d.UIKeyGeneralUserID = d.UIKeyGeneralUserID[len("orders/detail/userID="):]
	d.UIKeyGeneralUserLanguage = d.UIKeyGeneralUserLanguage[len("orders/detail/language="):]
	bp, err := strconv.Atoi(d.UIKeyGeneralBusinessPartner[len("orders/detail/businessPartnerID="):])
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	d.Params.BusinessPartner = bp
	d.Params.Language = d.UIKeyGeneralUserLanguage
	return &d
}
