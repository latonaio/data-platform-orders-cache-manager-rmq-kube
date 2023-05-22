package orderslist

import (
	"context"
	dpfm_api_input_reader "data-platform-orders-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-orders-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-orders-cache-manager-rmq-kube/api_requests/creator/orderslist"
	"data-platform-orders-cache-manager-rmq-kube/cache"
	apiresponses "data-platform-orders-cache-manager-rmq-kube/responses"
	rmqsessioncontroller "data-platform-orders-cache-manager-rmq-kube/rmq_session_controller"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

// CreateOrdersItemsReq
type OrdersListCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewOrdersListCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *OrdersListCtrl {
	return &OrdersListCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *OrdersListCtrl) OrdersList(msg rabbitmq.RabbitmqMessage) error {
	start := time.Now()
	params := extractOrdersListParam(msg)
	reqKey, err := getRequestKey(msg.Data())
	if err != nil {
		return xerrors.Errorf("reqKey error: %w", err)
	}
	sID, err := getSessionID(msg.Data())
	if err != nil {
		return xerrors.Errorf("session ID error: %w", err)
	}
	cacheResult := RedisCacheApiName{
		"redisCacheApiName": map[string]interface{}{},
	}
	defer func() {
		b, _ := json.Marshal(cacheResult)
		err = c.cache.Set(c.ctx, reqKey, b, 1*time.Hour)
		if err != nil {
			c.log.Error("cache set error: %w", err)
		}
	}()

	oRes, err := c.ordersRequest(params, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}
	bpRes, err := c.businessPartnerRequest(params, oRes, sID, reqKey, &cacheResult)
	if err != nil {
		// return err
		return nil
	}

	c.fin(params, oRes, bpRes, reqKey, "Orders", &cacheResult)
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
	return nil
}

func (c *OrdersListCtrl) ordersRequest(
	params *dpfm_api_input_reader.OrdersListParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.OrdersRes, error) {
	oiReq := orderslist.CreateOrdersReq(params, sID)
	res, err := c.request("data-platform-api-orders-reads-queue", oiReq, sID, reqKey, "Orders", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("orders cache set error: %w", err)
	}
	oiRes, err := apiresponses.CreateOrdersRes(res)
	if err != nil {
		return nil, xerrors.Errorf("orders response parse error: %w", err)
	}
	return oiRes, nil
}
func (c *OrdersListCtrl) businessPartnerRequest(
	params *dpfm_api_input_reader.OrdersListParams,
	oRes *apiresponses.OrdersRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.BusinessPartnerRes, error) {
	bpReq := orderslist.CreateBusinessPartnerReq(params, oRes, sID)
	res, err := c.request("data-platform-api-business-partner-reads-general-queue", bpReq, sID, reqKey, "BusinessPartner", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("business partner cache set error: %w", err)
	}
	bpRes, err := apiresponses.CreateBusinessPartnerRes(res)
	if err != nil {
		return nil, xerrors.Errorf("business partner response parse error: %w", err)
	}
	return bpRes, nil
}

func (c *OrdersListCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
	resFunc := c.rmq.SessionRequest(queue, req, sID)
	res := resFunc()
	if res == nil {
		return nil, xerrors.Errorf("receive nil response")
	}
	// redisKey := strings.Join([]string{url, api}, "/")
	// err := c.cache.Set(c.ctx, redisKey, res.Raw(), 1*time.Hour)
	// if err != nil {
	// 	return nil, xerrors.Errorf("cache set error: %w", err)
	// }
	// (*setFlag)["redisCacheApiName"][api] = map[string]string{"keyName": redisKey}
	return res, nil
}

func extractOrdersListParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.OrdersListParams {
	data := dpfm_api_input_reader.ReadOrdersList(msg)
	return &data.Params
}

func getSessionID(req interface{}) (string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(b, &m)
	if err != nil {
		return "", err
	}
	rawSID, ok := m["runtime_session_id"]
	if !ok {
		return "", xerrors.Errorf("runtime_session_id not included")
	}

	return fmt.Sprintf("%v", rawSID), nil
}

func getRequestKey(req interface{}) (string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(b, &m)
	if err != nil {
		return "", err
	}
	rawReqID, ok := m["ui_key_function_url"]
	if !ok {
		return "", xerrors.Errorf("keyName not included")
	}

	return fmt.Sprintf("%v", rawReqID), nil
}

type RedisCacheApiName map[string]map[string]interface{}

func (c *OrdersListCtrl) fin(
	params *dpfm_api_input_reader.OrdersListParams,
	oRes *apiresponses.OrdersRes,
	bpRes *apiresponses.BusinessPartnerRes,
	url, api string, setFlag *RedisCacheApiName,
) error {
	type bpIDRel struct {
		orderID        int
		sellerID       int
		buyerID        int
		deliveryStatus string
	}
	data := dpfm_api_output_formatter.OrdersList{
		Orders: make([]dpfm_api_output_formatter.Orders, 0),
	}
	bpMapper := map[int]apiresponses.BPGeneral{}
	for _, v := range *bpRes.Message.Generals {
		bpMapper[v.BusinessPartner] = v
	}
	ordersInfo := map[int]apiresponses.Header{}
	for _, v := range *oRes.Message.Header {
		if *v.HeaderDeliveryStatus == "CL" {
			continue
		}
		ordersInfo[v.OrderID] = v
	}
	infos := orderDesc(ordersInfo)

	for _, info := range infos {
		if params.User == "Buyer" {
			if *info.Buyer != *params.Buyer {
				continue
			}
		} else if params.User == "Seller" {
			if *info.Seller != *params.Seller {
				continue
			}
		}
		buyerName := ""
		sellerName := ""
		buyer, ok := bpMapper[*info.Buyer]
		if ok {
			buyerName = *buyer.BusinessPartnerFullName
		}
		seller, ok := bpMapper[*info.Seller]
		if ok {
			sellerName = *seller.BusinessPartnerFullName
		}

		data.Orders = append(data.Orders,
			dpfm_api_output_formatter.Orders{
				OrderID:             info.OrderID,
				SellerName:          sellerName,
				BuyerName:           buyerName,
				DeliveryStatus:      info.HeaderDeliveryStatus,
				OrderDate:           info.OrderDate,
				PaymentTerms:        info.PaymentTerms,
				PaymentMethod:       info.PaymentMethod,
				TransactionCurrency: info.TransactionCurrency,
				OrderType:           info.OrderType,
			},
		)
	}

	redisKey := strings.Join([]string{url, api}, "/")
	// redisKey := strings.Join([]string{url, api, params.User}, "/")
	b, _ := json.Marshal(data)
	err := c.cache.Set(c.ctx, redisKey, b, 1*time.Hour)
	if err != nil {
		return nil
	}
	(*setFlag)["redisCacheApiName"][api] = map[string]string{"keyName": redisKey}
	return nil
}

func orderAsc[T any](d map[int]T) []T {
	ids := make([]int, 0, len(d))
	for i := range d {
		ids = append(ids, i)
	}
	sort.Ints(ids)
	sli := make([]T, 0, len(d))
	for _, i := range ids {
		sli = append(sli, d[i])
	}
	return sli
}

func orderDesc[T any](d map[int]T) []T {
	ids := make([]int, 0, len(d))
	for i := range d {
		ids = append(ids, i)
	}
	sort.Ints(ids)
	sli := make([]T, 0, len(d))
	for i := len(ids) - 1; i >= 0; i-- {
		sli = append(sli, d[ids[i]])
	}
	return sli
}
