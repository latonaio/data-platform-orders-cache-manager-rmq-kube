package requests

import (
	models "data-platform-orders-cache-manager-rmq-kube/requests/models"
	"data-platform-orders-cache-manager-rmq-kube/responses"
	"data-platform-orders-cache-manager-rmq-kube/parameters"
)

func ReadProductStockFrom(oiRes *parameters.Parameters, *responses.OrdersRes, accountingRes *responses.ProductMasterRes, sID string) *models.ProductStock {
	item := (*oiRes.Message.Item)[0]
	return &models.ProductStock{
		ProductStock: models.ProductStock{
			Product:         *item.Product,
			BusinessPartner: *item.DeliverFromParty,
			Plant:           *item.DeliverFromPlant,
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
