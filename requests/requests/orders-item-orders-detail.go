package requests

import (
	dpfm_api_input_reader "data-platform-orders-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-orders-cache-manager-rmq-kube/requests/models"
)

func ReadOrdersItemOrdersDetail(param *dpfm_api_input_reader.OrdersDetailParams, ppc int, sID string) *models.Orders {
	return &models.Orders{
		Header: &models.OrdersHeader{
			OrderID: param.OrderID,
			Item: []models.Item{
				{
					OrderItem: param.OrderItem,
				},
			},
		},
		Accepter: []string{
			"Item",
			"ItemPricingElement",
		},
		RuntimeSessionID: sID,
	}
}
