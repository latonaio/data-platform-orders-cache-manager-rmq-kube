package requests

import (
	dpfm_api_input_reader "data-platform-orders-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-orders-cache-manager-rmq-kube/requests/models"
)

func ReadProductTag(param *dpfm_api_input_reader.OrdersDetailParams, sID string) *models.ProductTag {
	return &models.ProductTag{
		ProductTag: models.ProductTag{
			Product: param.Product,
		},
		Accepter: []string{
			"ProductTag",
		},
		RuntimeSessionID: sID,
	}
}
