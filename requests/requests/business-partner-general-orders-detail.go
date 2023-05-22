package requests

import (
	dpfm_api_input_reader "data-platform-orders-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-orders-cache-manager-rmq-kube/requests/models"
	"data-platform-orders-cache-manager-rmq-kube/responses"
	"data-platform-orders-cache-manager-rmq-kube/parameters"
)

func ReadBusinessPartnerGeneralOrdersDetail(param *dpfm_api_input_reader.OrdersDetailParams, *parameters.Parameters, productRes *responses.ProductMasterRes, oiRes *responses.OrdersRes, sID string) *models.BusinessPartner {
	return &models.BusinessPartner{
		General: models.BPGeneral{
			BusinessPartner: *(*oiRes.Message.Item)[0].ProductionPlantBusinessPartner,
			Language:        &param.Language,
		},
		Accepter: []string{
			"General",
		},
		RuntimeSessionID: sID,
	}
}
