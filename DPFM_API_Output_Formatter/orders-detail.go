package dpfm_api_output_formatter

type OrdersDetail struct {
	ProductName                           string                                `json:"productName"`
	ProductCode                           string                                `json:"productCode"`
	ProductInfo                           []ProductInfo                         `json:"productInfo"`
	ProductTag                            *[]ProductTag                         `json:"ProductTag"`
	Images                                Images                                `json:"images"`
	Stock                                 Stock                                 `json:"stock"`
	AvailabilityStock                     Stock                                 `json:"availabilityStock"`
	OrderQuantityInDelivery               OrderQuantityInDelivery               `json:"orderQuantityInDelivery"`
	OrderQuantityInBase                   OrderQuantityInBase                   `json:"orderQuantityInBase"`
	ConfirmedOrderQuantityByPDTAvailCheck ConfirmedOrderQuantityByPDTAvailCheck `json:"confirmedOrderQuantityByPDTAvailCheck"`
}

type ProductTagReads struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Deleted       string `json:"deleted"`
}

type ProductTag struct {
	Key      string `json:"key"`
	DocCount int    `json:"doc_count"`
}

type OrderQuantityInDelivery struct {
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

type OrderQuantityInBase struct {
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}
type ConfirmedOrderQuantityByPDTAvailCheck struct {
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

type ProductInfo struct {
	KeyName string      `json:"keyName"`
	Key     string      `json:"key"`
	Value   interface{} `json:"value,omitempty"`
}
type Images struct {
	Product ProductImage `json:"product"`
	Barcord BarcordImage `json:"barcode"`
}

type ProductImage struct {
	BusinessPartnerID int    `json:"BusinessPartnerID"`
	DocID             string `json:"DocID"`
	FileExtension     string `json:"FileExtension"`
}

type BarcordImage struct {
	ID          string `json:"Id"`
	Barcode     string `json:"barcode"`
	BarcodeType string `json:"barcodeType"`
}
type Stock struct {
	ProductStock    int    `json:"productStock"`
	StorageLocation string `json:"storageLocation"`
}
type AvailabilityStock struct {
	AvailableProductStock    int    `json:"availableProductStock"`
	AvailableStorageLocation string `json:"availableStorageLocation"`
}
