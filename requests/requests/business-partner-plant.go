package requests

import (
	dpfm_api_input_reader "data-platform-orders-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-orders-cache-manager-rmq-kube/requests/models"
	"data-platform-orders-cache-manager-rmq-kube/responses"
	"data-platform-orders-cache-manager-rmq-kube/parameters"
)

func ReadBPPlant(param *dpfm_api_input_reader.OrdersDetailParams, *parameters.Parameters, productRes *responses.ProductMasterRes, sID string) *models.ProductMaster {
	return &models.ProductMaster{
		General: models.PMGeneral{
			Product: param.Product,
			BusinessPartner: models.PMBusinessPartner{
				BusinessPartner: productRes.Message.BusinessPartner.BusinessPartner,
			},
		},
		Accepter: []string{
			"BPPlant",
		},
		RuntimeSessionID: sID,
	}
}
