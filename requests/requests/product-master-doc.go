package requests

import (
	dpfm_api_input_reader "data-platform-orders-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-orders-cache-manager-rmq-kube/requests/models"
	"data-platform-orders-cache-manager-rmq-kube/responses"
	"data-platform-orders-cache-manager-rmq-kube/parameters"
)

func ReadProductMasterDoc(param *dpfm_api_input_reader.OrdersDetailParams, *parameters.Parameters, productRes *responses.ProductMasterRes, sID string) *models.ProductMasterDoc {
	return &models.ProductMasterDoc{
		Product: models.PMDProduct{
			Product: param.Product,
			HeaderDoc: models.HeaderDoc{
				DocType:                  "IMAGE",
				DocIssuerBusinessPartner: productRes.Message.BusinessPartner.BusinessPartner,
			},
		},
		Accepter:         []string{},
		RuntimeSessionID: sID,
	}
}
