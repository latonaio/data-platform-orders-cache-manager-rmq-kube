package requests

import (
	dpfm_api_input_reader "data-platform-orders-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-orders-cache-manager-rmq-kube/requests/models"
)

func ReadProductMasterGeneral(param *dpfm_api_input_reader.OrdersDetailParams, sID string) *models.ProductMasterGeneral {
	return &models.ProductMasterGeneral{
		BusinessPartnerID: &param.BusinessPartner,
		General: models.PMGeneral{
			BusinessPartner: models.PMBusinessPartner{
				BusinessPartner: param.BusinessPartner,
			},
			Product: param.Product,
			ProductDescriptionByBusinessPartner: models.ProductDescriptionByBusinessPartner{
				BusinessPartner: param.BusinessPartner,
				Language:        param.Language,
			},
		},
		Accepter: []string{
			"General",
			"ProductDescriptionByBusinessPartner",
			"BusinessPartner",
		},
		RuntimeSessionID: sID,
	}
}
