package ordersdetaillist

import (
	"context"
	dpfm_api_input_reader "data-platform-orders-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-orders-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-orders-cache-manager-rmq-kube/api_requests/creator/ordersdetaillist"
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

type OrdersDetailListCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewOrdersDetailListCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *OrdersDetailListCtrl {
	return &OrdersDetailListCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *OrdersDetailListCtrl) OrdersDetailList(msg rabbitmq.RabbitmqMessage) error {
	start := time.Now()
	params := extractOrderDetailListParam(msg)
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

	oRes, err := c.ordersRequest(&params.Params, sID, reqKey)
	if err != nil {
		return err
	}

	c.pushOrdersDetailHeader(params, oRes, reqKey, &cacheResult)
	c.pushOrdersDetailList(params, oRes, reqKey, &cacheResult)

	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
	return nil
}

func (c *OrdersDetailListCtrl) ordersRequest(
	params *dpfm_api_input_reader.OrdersDetailListParams,
	sID string,
	reqKey string,
) (*apiresponses.OrdersRes, error) {
	oiReq := ordersdetaillist.CreateOrdersItemsReq(params, sID)
	res, err := c.request("data-platform-api-orders-reads-queue", oiReq, sID, reqKey, "Orders")
	if err != nil {
		return nil, xerrors.Errorf("orders cache set error: %w", err)
	}
	oiRes, err := apiresponses.CreateOrdersRes(res)
	if err != nil {
		return nil, xerrors.Errorf("orders response parse error: %w", err)
	}
	return oiRes, nil
}

func (c *OrdersDetailListCtrl) productRequest(
	params *dpfm_api_input_reader.OrdersDetailListParams,
	oRes *apiresponses.OrdersRes,
	sID string,
	reqKey string,
) (*apiresponses.ProductMasterRes, error) {
	oiReq := ordersdetaillist.CreateProductRequest(params, oRes, sID)
	res, err := c.request("data-platform-api-product-master-reads-queue", oiReq, sID, reqKey, "Orders")
	if err != nil {
		return nil, xerrors.Errorf("orders cache set error: %w", err)
	}
	pmRes, err := apiresponses.CreateProductMasterRes(res)
	if err != nil {
		return nil, xerrors.Errorf("product master response parse error: %w", err)
	}

	return pmRes, nil
}

func (c *OrdersDetailListCtrl) request(queue string, req interface{}, sID string, url, api string) (rabbitmq.RabbitmqMessage, error) {
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

func extractOrderDetailListParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.OrdersDetailList {
	data := dpfm_api_input_reader.ReadOrdersDetailList(msg)
	return data
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

func (c *OrdersDetailListCtrl) pushOrdersDetailHeader(
	params *dpfm_api_input_reader.OrdersDetailList,
	oRes *apiresponses.OrdersRes,
	url string, setFlag *RedisCacheApiName,
) error {
	api := "OrdersDetailHeader"
	key := fmt.Sprintf("%s/orders/list/User=%s/HeaderCompleteDeliveryIsDefined=%v/HeaderDeliveryStatus=NP/HeaderDeliveryBlockStatus=%v/HeaderIsCanceled=%v/HeaderIsDeleted=%v/Orders",
		params.UIKeyGeneralUserID, *params.Params.User, params.Params.HeaderCompleteDeliveryIsDefined, params.Params.HeaderDeliveryBlockStatus, params.Params.HeaderIsCanceled, params.Params.HeaderIsDeleted)
	m, err := c.cache.GetMap(c.ctx, key)
	if err != nil {
		return err
	}
	s := m["Orders"].([]interface{})

	idx := 0
	for i, v := range s {
		m := v.(map[string]interface{})
		if int(m["OrderID"].(float64)) == *params.Params.OrderID {
			idx = i
			break
		}
	}

	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": key, "index": idx}
	return nil

}

func (c *OrdersDetailListCtrl) pushOrdersDetailList(
	params *dpfm_api_input_reader.OrdersDetailList,
	oRes *apiresponses.OrdersRes,
	url string, setFlag *RedisCacheApiName,
) error {
	api := "OrdersDetailList"
	header := dpfm_api_output_formatter.OrdersDetailHeader{}
	details := make([]dpfm_api_output_formatter.OrdersItemSummary, 0, len(*oRes.Message.Item))

	key := fmt.Sprintf("%s/orders/list/User=%s/HeaderCompleteDeliveryIsDefined=%v/HeaderDeliveryStatus=NP/HeaderDeliveryBlockStatus=%v/HeaderIsCanceled=%v/HeaderIsDeleted=%v/Orders",
		params.UIKeyGeneralUserID, *params.Params.User, params.Params.HeaderCompleteDeliveryIsDefined, params.Params.HeaderDeliveryBlockStatus, params.Params.HeaderIsCanceled, params.Params.HeaderIsDeleted)
	m, err := c.cache.GetMap(c.ctx, key)
	if err != nil {
		return err
	}
	s := m["Orders"].([]interface{})

	idx := 0
	for i, v := range s {
		m := v.(map[string]interface{})
		if int(m["OrderID"].(float64)) == *params.Params.OrderID {
			idx = i
			break
		}
	}
	header = dpfm_api_output_formatter.OrdersDetailHeader{
		Index: idx,
		Key:   key,
	}

	pricingMapper := map[int]int{}
	for _, pricing := range *oRes.Message.ItemPricingElement {
		if pricing.ConditionType != nil && *pricing.ConditionType == "MWST" {
			continue
		}
		pricingMapper[pricing.OrderItem] += int(*pricing.ConditionRateValue)
	}

	for _, item := range *oRes.Message.Item {
		oqdu := ""
		if item.OrderQuantityInBaseUnit != nil {
			oqdu = fmt.Sprintf("%v", *item.OrderQuantityInDeliveryUnit)
		}
		netAmount := ""
		if item.NetAmount != nil {
			netAmount = fmt.Sprintf("%v", *item.NetAmount)
		}
		condRateVal := fmt.Sprintf("%v", pricingMapper[item.OrderItem])

		details = append(details,
			dpfm_api_output_formatter.OrdersItemSummary{
				OrderItem:                   item.OrderItem,
				Product:                     *item.Product,
				OrderItemTextByBuyer:        *item.OrderItemTextByBuyer,
				OrderItemTextBySeller:       *item.OrderItemTextBySeller,
				OrderQuantityInDeliveryUnit: oqdu,
				DeliveryUnit:                *item.DeliveryUnit,
				ConditionRateValue:          condRateVal,
				RequestedDeliveryDate:       *item.RequestedDeliveryDate,
				NetAmount:                   netAmount,
				IsCanceled:                  *item.ItemIsCancelled,
			})
	}
	redisKey := strings.Join([]string{url, api}, "/")
	data := dpfm_api_output_formatter.OrdersDetailList{
		Header:  header,
		Details: details,
	}
	b, _ := json.Marshal(data)
	err = c.cache.Set(c.ctx, redisKey, b, 1*time.Hour)
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
