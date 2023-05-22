package requests

import (
	dpfm_api_input_reader "data-platform-orders-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-orders-cache-manager-rmq-kube/requests/models"
	"data-platform-orders-cache-manager-rmq-kube/responses"
)

func ReadBusinessPartnerGeneralOrdersList(param *dpfm_api_input_reader.OrdersListParams, oRes *responses.OrdersRes, sID string) *models.BusinessPartner {
	bpIDs := make([]int, 0)
	dupCheck := make(map[int]struct{})
	for _, v := range *oRes.Message.Header {
		id := *v.Seller
		if _, ok := dupCheck[id]; ok {
			continue
		}
		dupCheck[id] = struct{}{}
		bpIDs = append(bpIDs, id)
	}
	for _, v := range *oRes.Message.Header {
		id := *v.Buyer
		if _, ok := dupCheck[id]; ok {
			continue
		}
		dupCheck[id] = struct{}{}
		bpIDs = append(bpIDs, id)
	}

	return &models.BusinessPartner{
		Generals: models.BPGenerals{
			BusinessPartners: bpIDs,
		},
		Accepter: []string{
			"Generals",
		},
		RuntimeSessionID: sID,
	}
}
