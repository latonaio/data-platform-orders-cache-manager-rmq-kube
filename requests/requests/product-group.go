package requests
import (
	dpfm_api_input_reader "data-platform-orders-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-orders-cache-manager-rmq-kube/requests/models"
	"data-platform-orders-cache-manager-rmq-kube/responses"
	"data-platform-orders-cache-manager-rmq-kube/parameters"
)

func ReadProductGroup(param *dpfm_api_input_reader.OrdersDetailParams, *parameters.Parameters, productRes *responses.ProductMasterRes, sID string) *models.ProductGroup {
	return &models.ProductGroup{
		ProductGroup: &models.ProductGroup{
			ProductGroup: *productRes.Message.General.ProductGroup,
			ProductGroupText: models.ProductGroupText{
				Language:         param.Language,
				ProductGroupName: productRes.Message.General.ProductGroup,
			},
		},
		Accepter: []string{
			"ProductGroup",
			"ProductGroupText",
		},
		RuntimeSessionID: sID,
	}
}
