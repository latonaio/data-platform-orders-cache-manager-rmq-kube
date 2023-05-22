package requests

import (
	dpfm_api_input_reader "data-platform-orders-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-orders-cache-manager-rmq-kube/requests/models"
)

func ReadOrdersItemOrdersDetailList(param *dpfm_api_input_reader.OrdersDetailListParams, sID string) *models.OrdersItem {
	return &models.OrdersItem{
		Header: &models.OrdersHeader{
			OrderID: *param.OrderID,
			Buyer:   param.Buyer,
			Seller:  param.Seller,
			Item: []models.OrdersItem{
				{
					OrderID:                 *param.OrderID,
					ItemDeliveryBlockStatus: param.ItemDeliveryBlockStatus,
					// ItemDeliveryStatus:            param.ItemDeliveryStatus,
					ItemCompleteDeliveryIsDefined: param.ItemDeliveryBlockStatus,
					ItemIsCanceled:                param.ItemIsCanceled,
					ItemIsDeleted:                 param.ItemIsDeleted,
				},
			},
		},
		Accepter: []string{
			"Header",
			"Items",
			"ItemPricingElements",
		},
		RuntimeSessionID: sID,
	}
}
