package requests

import (
	models "data-platform-orders-cache-manager-rmq-kube/requests/models"
	"data-platform-orders-cache-manager-rmq-kube/responses"
	"data-platform-orders-cache-manager-rmq-kube/parameters"
)

func ReadProductStockTo(oiRes *parameters.Parameters, *responses.OrdersRes, accountingRes *responses.ProductMasterRes, sID string) *models.ProductStock {
	return &models.ProductStock{
		ProductStock: models.ProductStock{
			Product:         *(*oiRes.Message.Item)[0].Product,
			BusinessPartner: *(*oiRes.Message.Item)[0].DeliverToParty,
			Plant:           *(*oiRes.Message.Item)[0].DeliverToPlant,
			ProductStockAvailability: models.ProductStockAvailability{
				ProductStockAvailabilityDate: *accountingRes.Message.Accounting.PriceLastChangeDate,
			},
		},
		Accepter: []string{
			"ProductStock",
			"ProductStockAvailability",
		},
		RuntimeSessionID: sID,
	}
}
