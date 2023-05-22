package requests

import (
	dpfm_api_input_reader "data-platform-orders-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-orders-cache-manager-rmq-kube/requests/models"
	"data-platform-orders-cache-manager-rmq-kube/responses"
	"data-platform-orders-cache-manager-rmq-kube/parameters"
)

func ReadDeliverToPlant(param *dpfm_api_input_reader.OrdersDetailParams, *parameters.Parameters, oiRes *responses.OrdersRes, sID string) *models.Plant {
	return &models.Plant{
		General: models.PlantGeneral{
			BusinessPartner: *(*oiRes.Message.Item)[0].DeliverToParty,
			Plant:           *(*oiRes.Message.Item)[0].DeliverToPlant,
		},
		Accepter: []string{
			"General",
		},
		RuntimeSessionID: sID,
	}
}
