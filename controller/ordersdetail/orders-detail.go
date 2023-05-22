package ordersdetail

import (
	"context"
	dpfm_api_input_reader "data-platform-orders-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-orders-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-orders-cache-manager-rmq-kube/cache"
	"data-platform-orders-cache-manager-rmq-kube/requests/creator/ordersdetail"
	apiresponses "data-platform-orders-cache-manager-rmq-kube/responses"
	rmqsessioncontroller "data-platform-orders-cache-manager-rmq-kube/rmq_session_controller"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type OrdersDetailCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewOrdersDetailCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *OrdersDetailCtrl {
	return &OrdersDetailCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *OrdersDetailCtrl) OrdersDetail(msg rabbitmq.RabbitmqMessage) error {
	start := time.Now()
	params := extractOrderDetailParam(msg)
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

	pmRes, err := c.productRequest(params, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}
	pgRes, err := c.productGroupRequest(params, pmRes, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}
	oiRes, err := c.orderItemRequest(params, pmRes, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}
	bpRes, err := c.businessPartnerRequest(params, pmRes, oiRes, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}
	pmdRes, err := c.productMasterDocRequest(params, pmRes, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}

	stockFromRes, err := c.productStockFromRequest(params, oiRes, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}
	stockToRes, err := c.productStockToRequest(params, oiRes, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}
	fromPlantRes, err := c.deliverFromPlantRequest(params, pmRes, oiRes, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}
	toPlantRes, err := c.deliverToPlantRequest(params, pmRes, oiRes, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}
	tagRes, err := c.tagRequest(params, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}
	c.pushOrdersDetail(
		params, pmRes, pmdRes, pgRes, bpRes, fromPlantRes, toPlantRes, stockFromRes, stockToRes, oiRes, tagRes, sID, reqKey, &cacheResult,
	)

	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
	return nil
}

func (c *OrdersDetailCtrl) tagRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductTagRes, error) {
	ptReq := ordersdetail.CreateProductTagReq(params, sID)
	res, err := c.request("data-platform-api-product-tag-reads-queue", ptReq, sID, reqKey, "ProductTag", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("product cache set error: %w", err)
	}
	ptRes, err := apiresponses.CreateProductTagRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Product master response parse error: %w", err)
	}
	return ptRes, nil
}

func (c *OrdersDetailCtrl) productRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductMasterRes, error) {
	pmReq := ordersdetail.CreateProductRequest(params, sID)
	res, err := c.request("data-platform-api-product-master-reads-queue", pmReq, sID, reqKey, "Product", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("product cache set error: %w", err)
	}
	pmRes, err := apiresponses.CreateProductMasterRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Product master response parse error: %w", err)
	}
	return pmRes, nil
}

func (c *OrdersDetailCtrl) productGroupRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductGroupRes, error) {
	pgReq := ordersdetail.CreateProductGroupReq(params, pmRes, sID)
	res, err := c.request("data-platform-api-product-group-reads-queue", pgReq, sID, reqKey, "ProductGroup", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("ProductGroup cache set error: %w", err)
	}
	pgRes, err := apiresponses.CreateProductGroupRes(res)
	if err != nil {
		return nil, xerrors.Errorf("ProductGroup response parse error: %w", err)
	}
	return pgRes, nil
}

func (c *OrdersDetailCtrl) bpPlantRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductMasterRes, error) {
	bpPlantReq := ordersdetail.CreateBPPlantReq(params, pmRes, sID)
	res, err := c.request("data-platform-api-product-master-reads-queue", bpPlantReq, sID, reqKey, "BPPlant", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("BPPlant cache set error: %w", err)
	}
	bpPlantRes, err := apiresponses.CreateProductMasterRes(res)
	if err != nil {
		return nil, xerrors.Errorf("BPPlant response parse error: %w", err)
	}
	return bpPlantRes, nil
}

func (c *OrdersDetailCtrl) businessPartnerRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	oiRes *apiresponses.OrdersRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.BusinessPartnerRes, error) {
	bpGeneralReq := ordersdetail.CreateBusinessPartnerReq(params, pmRes, oiRes, sID)
	res, err := c.request("data-platform-api-business-partner-reads-general-queue", bpGeneralReq, sID, reqKey, "BusinessPartnerGeneral", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("BusinessPartnerGeneral cache set error: %w", err)
	}
	bpRes, err := apiresponses.CreateBusinessPartnerRes(res)
	if err != nil {
		return nil, xerrors.Errorf("BusinessPartnerGeneral response parse error: %w", err)
	}
	return bpRes, nil
}

func (c *OrdersDetailCtrl) productMasterDocRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductMasterDocRes, error) {
	pmDocReq := ordersdetail.CreateProductMasterDocReq(params, pmRes, sID)
	res, err := c.request("data-platform-api-product-master-doc-reads-queue", pmDocReq, sID, reqKey, "ProductMasterDoc", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("ProductMasterDoc cache set error: %w", err)
	}
	pmdRes, err := apiresponses.CreateProductMasterDocRes(res)
	if err != nil {
		return nil, xerrors.Errorf("ProductMasterDoc response parse error: %w", err)
	}
	return pmdRes, nil
}

func (c *OrdersDetailCtrl) productStockFromRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	oiRes *apiresponses.OrdersRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductStockRes, error) {
	psReq := ordersdetail.CreateProductStockFromReq(oiRes, sID)
	res, err := c.request("data-platform-api-product-stock-reads-queue", psReq, sID, reqKey, "ProductStockFrom", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("ProductStock cache set error: %w", err)
	}
	psRes, err := apiresponses.CreateProductStockRes(res)
	if err != nil {
		return nil, xerrors.Errorf("ProductStock response parse error: %w", err)
	}
	return psRes, nil
}
func (c *OrdersDetailCtrl) productStockToRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	oiRes *apiresponses.OrdersRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductStockRes, error) {
	psReq := ordersdetail.CreateProductStockToReq(oiRes, sID)
	res, err := c.request("data-platform-api-product-stock-reads-queue", psReq, sID, reqKey, "ProductStockTo", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("ProductStock cache set error: %w", err)
	}
	psRes, err := apiresponses.CreateProductStockRes(res)
	if err != nil {
		return nil, xerrors.Errorf("ProductStock response parse error: %w", err)
	}
	return psRes, nil
}

func (c *OrdersDetailCtrl) orderItemRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.OrdersRes, error) {
	oiReq := ordersdetail.CreateOrdersItemReq(params, 1, sID)
	res, err := c.request("data-platform-api-orders-reads-queue", oiReq, sID, reqKey, "OrderItem", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("OrderItem cache set error: %w", err)
	}
	oiRes, err := apiresponses.CreateOrdersRes(res)
	if err != nil {
		return nil, xerrors.Errorf("OrderItem response parse error: %w", err)
	}
	return oiRes, nil
}

func (c *OrdersDetailCtrl) deliverFromPlantRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	oiRes *apiresponses.OrdersRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.PlantRes, error) {
	dfpReq := ordersdetail.CreateDeliverFromPlantReq(params, oiRes, sID)
	res, err := c.request("data-platform-api-plant-reads-queue", dfpReq, sID, reqKey, "DeliverFromPlant", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("DeliverFromPlant cache set error: %w", err)
	}
	dfpRes, err := apiresponses.CreatePlantRes(res)
	if err != nil {
		return nil, xerrors.Errorf("DeliverFromPlant response parse error: %w", err)
	}
	return dfpRes, nil
}

func (c *OrdersDetailCtrl) deliverToPlantRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	oiRes *apiresponses.OrdersRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.PlantRes, error) {
	dtpReq := ordersdetail.CreateDeliverToPlantReq(params, oiRes, sID)
	res, err := c.request("data-platform-api-plant-reads-queue", dtpReq, sID, reqKey, "DeliverToPlant", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("DeliverToPlant cache set error: %w", err)
	}
	dtpRes, err := apiresponses.CreatePlantRes(res)
	if err != nil {
		return nil, xerrors.Errorf("DeliverToPlant response parse error: %w", err)
	}
	return dtpRes, nil
}

func (c *OrdersDetailCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func extractOrderDetailParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.OrdersDetailParams {
	data := dpfm_api_input_reader.ReadOrdersDetail(msg)
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

func (c *OrdersDetailCtrl) pushOrdersDetail(
	params *dpfm_api_input_reader.OrdersDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	pmdRes *apiresponses.ProductMasterDocRes,
	pgRes *apiresponses.ProductGroupRes,
	bpRes *apiresponses.BusinessPartnerRes,
	accountingRes *apiresponses.ProductMasterRes,
	fromPlantRes *apiresponses.PlantRes,
	toPlantRes *apiresponses.PlantRes,
	stockFromRes *apiresponses.ProductStockRes,
	stockToRes *apiresponses.ProductStockRes,
	oiRes *apiresponses.OrdersRes,
	ptRes *apiresponses.ProductTagRes,
	sID string,
	url string,
	setFlag *RedisCacheApiName,
) error {
	api := "OrdersDetail"
	pInfo := c.CreateProductInfo(pgRes, accountingRes, oiRes, bpRes)
	tags := c.CreateProductTag(ptRes)
	data := dpfm_api_output_formatter.OrdersDetail{
		ProductName: *pmRes.Message.ProductDescriptionByBusinessPartner.ProductDescription,
		ProductCode: pmRes.Message.ProductDescriptionByBusinessPartner.Product,
		ProductInfo: pInfo,
		ProductTag:  tags,
		Images: dpfm_api_output_formatter.Images{
			Product: dpfm_api_output_formatter.ProductImage{
				BusinessPartnerID: pmRes.Message.BusinessPartner.BusinessPartner,
				DocID:             pmdRes.Message.DocID,
				FileExtension:     pmdRes.Message.FileExtension,
			},
			Barcord: dpfm_api_output_formatter.BarcordImage{
				ID:          *pmRes.Message.General.ProductStandardID,
				Barcode:     "",
				BarcodeType: *pmRes.Message.General.BarcodeType,
			},
		},
		Stock: dpfm_api_output_formatter.Stock{
			ProductStock:    int(*stockToRes.Message.ProductStock.ProductStock),
			StorageLocation: *toPlantRes.Message.General.PlantFullName,
		},
		AvailabilityStock: dpfm_api_output_formatter.Stock{
			ProductStock:    int(*stockFromRes.Message.ProductStock.ProductStock),
			StorageLocation: *fromPlantRes.Message.General.PlantFullName,
		},
		OrderQuantityInDelivery: dpfm_api_output_formatter.OrderQuantityInDelivery{
			Quantity: int(*(*oiRes.Message.Item)[0].OrderQuantityInDeliveryUnit),
			Unit:     *(*oiRes.Message.Item)[0].DeliveryUnit,
		},
		OrderQuantityInBase: dpfm_api_output_formatter.OrderQuantityInBase{
			Quantity: int(*(*oiRes.Message.Item)[0].OrderQuantityInBaseUnit),
			Unit:     *(*oiRes.Message.Item)[0].BaseUnit,
		},
		ConfirmedOrderQuantityByPDTAvailCheck: dpfm_api_output_formatter.ConfirmedOrderQuantityByPDTAvailCheck{
			Quantity: int(*(*oiRes.Message.Item)[0].ConfirmedOrderQuantityInBaseUnit),
			Unit:     *(*oiRes.Message.Item)[0].BaseUnit,
		},
	}

	redisKey := strings.Join([]string{url, api}, "/")

	b, _ := json.Marshal(data)
	err := c.cache.Set(c.ctx, redisKey, b, 1*time.Hour)
	if err != nil {
		return nil
	}

	(*setFlag)["redisCacheApiName"][api] = map[string]string{"keyName": redisKey}
	return nil
}

func (c *OrdersDetailCtrl) CreateProductTag(ptRes *apiresponses.ProductTagRes) *[]dpfm_api_output_formatter.ProductTag {
	if ptRes == nil || ptRes.Message.ProductTag == nil {
		return &[]dpfm_api_output_formatter.ProductTag{}
	}
	tags := make([]dpfm_api_output_formatter.ProductTag, 0, len(*ptRes.Message.ProductTag))
	for _, v := range *ptRes.Message.ProductTag {
		tags = append(tags, dpfm_api_output_formatter.ProductTag{
			Key:      v.Key,
			DocCount: v.DocCount,
		},
		)
	}
	return &tags
}

func recovery(l *logger.Logger) {
	if e := recover(); e != nil {
		l.Error("%+v", e)
		return
	}
}
func (c *OrdersDetailCtrl) CreateProductInfo(
	pgRes *apiresponses.ProductGroupRes,
	aRes *apiresponses.ProductMasterRes,
	oRes *apiresponses.OrdersRes,
	bpRes *apiresponses.BusinessPartnerRes,
) (d []dpfm_api_output_formatter.ProductInfo) {
	defer recovery(c.log)
	d = make([]dpfm_api_output_formatter.ProductInfo, 0, 4)
	d = append(d, c.CreateProductInfoProductGroup(pgRes))
	// d = append(d, c.CreateProductInfoPriceUnitQty(aRes))
	d = append(d, c.CreateInternalCapacityQuantity(oRes))
	d = append(d, c.CreateProductInfoPrice(oRes, aRes))
	d = append(d, c.CreateProductInfoBPName(bpRes))
	d = append(d, c.CreateProductInfoAllergy(pgRes))
	return d
}

func (c *OrdersDetailCtrl) CreateProductInfoProductGroup(pgRes *apiresponses.ProductGroupRes) dpfm_api_output_formatter.ProductInfo {
	return dpfm_api_output_formatter.ProductInfo{
		KeyName: "ProductGroupName",
		Key:     "商品分類",
		Value:   *pgRes.Message.ProductGroupText.ProductGroupName,
	}
}

func (c *OrdersDetailCtrl) CreateProductInfoPriceUnitQty(
	aRes *apiresponses.ProductMasterRes,
) dpfm_api_output_formatter.ProductInfo {
	return dpfm_api_output_formatter.ProductInfo{
		KeyName: "PriceUnitQty",
		Key:     "価格単位",
		Value:   aRes.Message.Accounting.PriceUnitQty,
	}
}

func (c *OrdersDetailCtrl) CreateProductInfoBPName(
	bpRes *apiresponses.BusinessPartnerRes,
) dpfm_api_output_formatter.ProductInfo {
	return dpfm_api_output_formatter.ProductInfo{
		KeyName: "BusinessPartnerName",
		Key:     "製造者",
		Value:   bpRes.Message.General.BusinessPartnerFullName,
	}
}

func (c *OrdersDetailCtrl) CreateProductInfoPrice(
	oRes *apiresponses.OrdersRes,
	aRes *apiresponses.ProductMasterRes,
) dpfm_api_output_formatter.ProductInfo {
	netAmount := 0
	grossAmount := 0

	for _, v := range *oRes.Message.ItemPricingElement {
		switch *v.ConditionType {
		case "PR00":
			netAmount = int(*v.ConditionRateValue)
		case "MWST":
			grossAmount = int(*v.ConditionRateValue)
		}
	}

	return dpfm_api_output_formatter.ProductInfo{
		KeyName: "Price",
		Key:     "価格",
		Value: map[string]interface{}{
			"PriceWithoutTax": 0,
			"NetAmount":       netAmount,
			"PriceUnitQty":    aRes.Message.Accounting.PriceUnitQty,
			"GrossAmount":     netAmount + grossAmount,
		},
	}
}

func (c *OrdersDetailCtrl) CreateInternalCapacityQuantity(
	oRes *apiresponses.OrdersRes,
) dpfm_api_output_formatter.ProductInfo {

	return dpfm_api_output_formatter.ProductInfo{
		KeyName: "InternalCapacityQuantity",
		Key:     "内容量",
		Value: map[string]interface{}{
			"InternalCapacityQuantity":     (*oRes.Message.Item)[0].InternalCapacityQuantity,
			"InternalCapacityQuantityUnit": (*oRes.Message.Item)[0].InternalCapacityQuantityUnit,
		},
	}
}

func (c *OrdersDetailCtrl) CreateProductInfoAllergy(pgRes *apiresponses.ProductGroupRes) dpfm_api_output_formatter.ProductInfo {
	return dpfm_api_output_formatter.ProductInfo{
		KeyName: "Allergy",
		Key:     "アレルゲン",
		Value:   nil,
	}
}
