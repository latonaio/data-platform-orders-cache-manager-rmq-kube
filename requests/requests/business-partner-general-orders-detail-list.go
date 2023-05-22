package requests

import (
	dpfm_api_input_reader "data-platform-orders-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-orders-cache-manager-rmq-kube/requests/models"
	"data-platform-orders-cache-manager-rmq-kube/responses"
)

func ReadBusinessPartnerGeneralOrdersDetailList(param *dpfm_api_input_reader.OrdersListParams, productRes *responses.ProductMasterRes, oiRes *responses.OrdersRes, sID string) *models.BusinessPartner {
	return &models.BusinessPartner{
		General: models.BPGeneral{
			BusinessPartner: *(*oiRes.Message.Item)[0].ProductionPlantBusinessPartner,
			Language:        &param.Language,
		},
		Accepter:         []string{},
		RuntimeSessionID: sID,
	}
}
