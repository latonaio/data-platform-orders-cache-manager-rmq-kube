package requests

import (
	dpfm_api_input_reader "data-platform-orders-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-orders-cache-manager-rmq-kube/requests/models"
)

func ReadOrdersHeader(param *dpfm_api_input_reader.OrdersListParams, sID string) *models.Orders {
	req := &models.Orders{
		Header: &models.OrdersHeader{
			Buyer:                     param.Buyer,
			Seller:                    param.Seller,
			HeaderDeliveryBlockStatus: &param.HeaderDeliveryBlockStatus,
			// HeaderDeliveryStatus:            &param.HeaderDeliveryStatus,
			HeaderCompleteDeliveryIsDefined: &param.HeaderCompleteDeliveryIsDefined,
			HeaderIsCanceled:                param.HeaderIsCanceled,
			HeaderIsDeleted:                 param.HeaderIsDeleted,
		},
		Accepter: []string{
			"Headers",
		},
		RuntimeSessionID: sID,
	}

	// if param.User == "Buyer" {
	// 	req.Header.Buyer = &param.BusinessPartner
	// } else if param.User == "Seller" {
	// 	req.Header.Seller = &param.BusinessPartner
	// }
	return req
}
