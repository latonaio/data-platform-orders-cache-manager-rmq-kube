package models

type ProductMasterReq struct {
	ConnectionKey     string    `json:"connection_key"`
	Result            bool      `json:"result"`
	RedisKey          string    `json:"redis_key"`
	Filepath          string    `json:"filepath"`
	APIStatusCode     int       `json:"api_status_code"`
	RuntimeSessionID  string    `json:"runtime_session_id"`
	BusinessPartnerID *int      `json:"business_partner"`
	ServiceLabel      string    `json:"service_label"`
	APIType           string    `json:"api_type"`
	General           PMGeneral `json:"Product"`
	APISchema         string    `json:"api_schema"`
	Accepter          []string  `json:"accepter"`
	Deleted           bool      `json:"deleted"`
}

type PMGeneral struct {
	Product                             string                              `json:"Product"`
	ProductType                         *string                             `json:"ProductType"`
	BaseUnit                            *string                             `json:"BaseUnit"`
	ValidityStartDate                   *string                             `json:"ValidityStartDate"`
	ProductGroup                        *string                             `json:"ProductGroup"`
	Division                            *string                             `json:"Division"`
	GrossWeight                         *float32                            `json:"GrossWeight"`
	WeightUnit                          *string                             `json:"WeightUnit"`
	SizeOrDimensionText                 *string                             `json:"SizeOrDimensionText"`
	IndustryStandardName                *string                             `json:"IndustryStandardName"`
	ProductStandardID                   *string                             `json:"ProductStandardID"`
	CreationDate                        *string                             `json:"CreationDate"`
	LastChangeDate                      *string                             `json:"LastChangeDate"`
	NetWeight                           *float32                            `json:"NetWeight"`
	CountryOfOrigin                     *string                             `json:"CountryOfOrigin"`
	ItemCategory                        *string                             `json:"ItemCategory"`
	ProductAccountAssignmentGroup       *string                             `json:"ProductAccountAssignmentGroup"`
	IsMarkedForDeletion                 *bool                               `json:"IsMarkedForDeletion"`
	GeneralPDF                          GeneralPDF                          `json:"GeneralPDF"`
	BusinessPartner                     PMBusinessPartner                   `json:"BusinessPartner"`
	ProductDescription                  ProductDescription                  `json:"ProductDescription"`
	ProductDescriptionByBusinessPartner ProductDescriptionByBusinessPartner `json:"ProductDescriptionByBusinessPartner"`
}

type GeneralPDF struct {
	DocType      string `json:"DocType"`
	DocVersionID *int   `json:"DocVersionID"`
	DocID        string `json:"DocID"`
	FileName     string `json:"FileName"`
}

type PMBusinessPartner struct {
	Product                string  `json:"Product"`
	BusinessPartner        int     `json:"BusinessPartner"`
	ValidityEndDate        *string `json:"ValidityEndDate"`
	ValidityStartDate      *string `json:"ValidityStartDate"`
	IsMarkedForDeletion    *bool   `json:"IsMarkedForDeletion"`
	BusinessPartnerProduct *string `json:"BusinessPartnerProduct"`
	BPPlant                BPPlant `json:"BPPlant"`
	Sales                  Sales   `json:"Sales"`
	Tax                    Tax     `json:"Tax"`
}

type BPPlant struct {
	Product                                   string            `json:"Product"`
	BusinessPartner                           int               `json:"BusinessPartner"`
	Plant                                     string            `json:"Plant"`
	AvailabilityCheckType                     *string           `json:"AvailabilityCheckType"`
	ProfitCenter                              *string           `json:"ProfitCenter"`
	MRPType                                   *string           `json:"MRPType"`
	MRPController                             *string           `json:"MRPController"`
	ReorderThresholdQuantity                  *float32          `json:"ReorderThresholdQuantity"`
	PlanningTimeFence                         *int              `json:"PlanningTimeFence"`
	MRPPlanningCalendar                       *string           `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32          `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *int              `json:"SafetyDuration"`
	MaximumStockQuantityInBaseUnit            *float32          `json:"MaximumStockQuantityInBaseUnit"`
	MinumumDeliveryQuantityInBaseUnit         *float32          `json:"MinumumDeliveryQuantityInBaseUnit"`
	MinumumDeliveryLotSizeQuantityInBaseUnit  *float32          `json:"MinumumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32          `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32          `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32          `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32          `json:"MaximumDeliveryQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool             `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDurationInDays            *int              `json:"StandardDeliveryDurationInDays"`
	IsBatchManagementRequired                 *bool             `json:"IsBatchManagementRequired"`
	BatchManagementPolicy                     *string           `json:"BatchManagementPolicy"`
	InventoryUnit                             *string           `json:"InventoryUnit"`
	IsMarkedForDeletion                       *bool             `json:"IsMarkedForDeletion"`
	BPPlantPDF                                BPPlantPDF        `json:"BPPlantPDF"`
	StorageLocation                           PMStorageLocation `json:"StorageLocation"`
	Procurement                               Procurement       `json:"Procurement"`
	MRPArea                                   MRPArea           `json:"MRPArea"`
	WorkScheduling                            WorkScheduling    `json:"WorkScheduling"`
	Accounting                                PMAccounting      `json:"Accounting"`
}

type BPPlantPDF struct {
	DocType      string `json:"DocType"`
	DocVersionID *int   `json:"DocVersionID"`
	DocID        string `json:"DocID"`
	FileName     string `json:"FileName"`
}

type PMStorageLocation struct {
	Product              string  `json:"Product"`
	BusinessPartner      int     `json:"BusinessPartner"`
	Plant                string  `json:"Plant"`
	StorageLocation      *string `json:"StorageLocation"`
	CreationDate         *string `json:"CreationDate"`
	InventoryBlockStatus *bool   `json:"InventoryBlockStatus"`
	IsMarkedForDeletion  *bool   `json:"IsMarkedForDeletion"`
}

type Procurement struct {
	Product                     string `json:"Product"`
	BusinessPartner             int    `json:"BusinessPartner"`
	Plant                       string `json:"Plant"`
	Buyable                     *bool  `json:"Buyable"`
	IsAutoPurOrdCreationAllowed *bool  `json:"IsAutoPurOrdCreationAllowed"`
	IsSourceListRequired        *bool  `json:"IsSourceListRequired"`
	IsMarkedForDeletion         *bool  `json:"IsMarkedForDeletion"`
}

type MRPArea struct {
	Product                                   string   `json:"Product"`
	BusinessPartner                           int      `json:"BusinessPartner"`
	Plant                                     string   `json:"Plant"`
	MRPArea                                   string   `json:"MRPArea"`
	StorageLocationForMRP                     *string  `json:"StorageLocationForMRP"`
	MRPType                                   *string  `json:"MRPType"`
	MRPController                             *string  `json:"MRPController"`
	ReorderThresholdQuantity                  *float32 `json:"ReorderThresholdQuantity"`
	PlanningTimeFence                         *int     `json:"PlanningTimeFence"`
	MRPPlanningCalendar                       *string  `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32 `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *int     `json:"SafetyDuration"`
	MaximumStockQuantityInBaseUnit            *float32 `json:"MaximumStockQuantityInBaseUnit"`
	MinumumDeliveryQuantityInBaseUnit         *float32 `json:"MinumumDeliveryQuantityInBaseUnit"`
	MinumumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MinumumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32 `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32 `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32 `json:"MaximumDeliveryQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool    `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDurationInDays            *int     `json:"StandardDeliveryDurationInDays"`
	IsMarkedForDeletion                       *bool    `json:"IsMarkedForDeletion"`
}

type WorkScheduling struct {
	Product                       string  `json:"Product"`
	BusinessPartner               int     `json:"BusinessPartner"`
	Plant                         string  `json:"Plant"`
	ProductionInvtryManagedLoc    *string `json:"ProductionInvtryManagedLoc"`
	ProductProcessingTime         *int    `json:"ProductProcessingTime"`
	ProductionSupervisor          *string `json:"ProductionSupervisor"`
	ProductProductionQuantityUnit *string `json:"ProductProductionQuantityUnit"`
	ProdnOrderIsBatchRequired     *bool   `json:"ProdnOrderIsBatchRequired"`
	MatlCompIsMarkedForBackflush  *bool   `json:"MatlCompIsMarkedForBackflush"`
	ProductionSchedulingProfile   *string `json:"ProductionSchedulingProfile"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
}
type PMAccounting struct {
	Product             string   `json:"Product"`
	BusinessPartner     int      `json:"BusinessPartner"`
	Plant               string   `json:"Plant"`
	ValuationClass      *string  `json:"ValuationClass"`
	CostingPolicy       *string  `json:"CostingPolicy"`
	PriceUnitQty        *string  `json:"PriceUnitQty"`
	StandardPrice       *float32 `json:"StandardPrice"`
	MovingAveragePrice  *float32 `json:"MovingAveragePrice"`
	PriceLastChangeDate *string  `json:"PriceLastChangeDate"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}

type Sales struct {
	Product             string `json:"Product"`
	BusinessPartner     int    `json:"BusinessPartner"`
	Sellable            *bool  `json:"Sellable"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type Tax struct {
	Product                  string  `json:"Product"`
	BusinessPartner          int     `json:"BusinessPartner"`
	Country                  *string `json:"Country"`
	TaxCategory              *string `json:"TaxCategory"`
	ProductTaxClassification *string `json:"ProductTaxClassification"`
}

type ProductDescription struct {
	Product            string  `json:"Product"`
	Language           string  `json:"Language"`
	ProductDescription *string `json:"ProductDescription"`
}

type ProductDescriptionByBusinessPartner struct {
	Product            string  `json:"Product"`
	BusinessPartner    int     `json:"BusinessPartner"`
	Language           string  `json:"Language"`
	ProductDescription *string `json:"ProductDescription"`
}
