package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	General            *[]General            `json:"General"`
	BusinessPartner    *[]BusinessPartner    `json:"BusinessPartner"`
	Allergen           *[]Allergen           `json:"Allergen"`
	NutritionalInfo    *[]NutritionalInfo    `json:"NutritionalInfo"`
	Calories           *[]Calories           `json:"Calories"`
	BPPlant            *[]BPPlant            `json:"BPPlant"`
	StorageLocation    *[]StorageLocation    `json:"StorageLocation"`
	MRPArea            *[]MRPArea            `json:"MRPArea"`
	WorkScheduling     *[]WorkScheduling     `json:"WorkScheduling"`
	Accounting         *[]Accounting         `json:"Accounting"`
	ProductDescription *[]ProductDescription `json:"ProductDescription"`
	ProductDescByBP    *[]ProductDescByBP    `json:"ProductDescByBP"`
	Tax                *[]Tax                `json:"Tax"`
}

type General struct {
	Product                       string   `json:"Product"`
	ProductType                   *string  `json:"ProductType"`
	BaseUnit                      *string  `json:"BaseUnit"`
	ValidityStartDate             *string  `json:"ValidityStartDate"`
	ProductGroup                  *string  `json:"ProductGroup"`
	GrossWeight                   *float32 `json:"GrossWeight"`
	NetWeight                     *float32 `json:"NetWeight"`
	WeightUnit                    *string  `json:"WeightUnit"`
	InternalCapacityQuantity      *float32 `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit  *string  `json:"InternalCapacityQuantityUnit"`
	SizeOrDimensionText           *string  `json:"SizeOrDimensionText"`
	ProductStandardID             *string  `json:"ProductStandardID"`
	IndustryStandardName          *string  `json:"IndustryStandardName"`
	ItemCategory                  *string  `json:"ItemCategory"`
	BarcodeType                   *string  `json:"BarcodeType"`
	CountryOfOrigin               *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage       *string  `json:"CountryOfOriginLanguage"`
	ProductAccountAssignmentGroup *string  `json:"ProductAccountAssignmentGroup"`
	CreationDate                  *string  `json:"CreationDate"`
	LastChangeDate                *string  `json:"LastChangeDate"`
	IsMarkedForDeletion           *bool    `json:"IsMarkedForDeletion"`
}

type BusinessPartner struct {
	Product                string  `json:"Product"`
	BusinessPartner        int     `json:"BusinessPartner"`
	ValidityEndDate        string  `json:"ValidityEndDate"`
	ValidityStartDate      string  `json:"ValidityStartDate"`
	BusinessPartnerProduct *string `json:"BusinessPartnerProduct"`
	IsMarkedForDeletion    *bool   `json:"IsMarkedForDeletion"`
}

type Allergen struct {
	Product             string `json:"Product"`
	BusinessPartner     int    `json:"BusinessPartner"`
	Allergen            string `json:"Allergen"`
	AllergenIsContained *bool  `json:"AllergenIsContained"`
}

type NutritionalInfo struct {
	Product             string  `json:"Product"`
	BusinessPartner     int     `json:"BusinessPartner"`
	Nutrient            int     `json:"Nutrient"`
	NutrientContent     *int    `json:"NutrientContent"`
	NutrientContentUnit *string `json:"NutrientContentUnit"`
}

type Calories struct {
	Product            string  `json:"Product"`
	BusinessPartner    int     `json:"BusinessPartner"`
	CaloryUnitQuantity int     `json:"CaloryUnitQuantity"`
	Calories           *int    `json:"Calories"`
	CaloryUnit         *string `json:"CaloryUnit"`
}

type BPPlant struct {
	Product                                   string   `json:"Product"`
	BusinessPartner                           int      `json:"BusinessPartner"`
	Plant                                     string   `json:"Plant"`
	AvailabilityCheckType                     *string  `json:"AvailabilityCheckType"`
	MRPType                                   *string  `json:"MRPType"`
	MRPController                             *string  `json:"MRPController"`
	ReorderThresholdQuantity                  *float32 `json:"ReorderThresholdQuantity"`
	PlanningTimeFence                         *int     `json:"PlanningTimeFence"`
	MRPPlanningCalendar                       *string  `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32 `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *int     `json:"SafetyDuration"`
	MaximumStockQuantityInBaseUnit            *float32 `json:"MaximumStockQuantityInBaseUnit"`
	MinimumDeliveryQuantityInBaseUnit         *float32 `json:"MinimumDeliveryQuantityInBaseUnit"`
	MinimumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MinimumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32 `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32 `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32 `json:"MaximumDeliveryQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool    `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDurationInDays            *int     `json:"StandardDeliveryDurationInDays"`
	IsBatchManagementRequired                 *bool    `json:"IsBatchManagementRequired"`
	BatchManagementPolicy                     *string  `json:"BatchManagementPolicy"`
	InventoryUnit                             *string  `json:"InventoryUnit"`
	ProfitCenter                              *string  `json:"ProfitCenter"`
	IsMarkedForDeletion                       *bool    `json:"IsMarkedForDeletion"`
}

type StorageLocation struct {
	Product              string  `json:"Product"`
	BusinessPartner      int     `json:"BusinessPartner"`
	Plant                string  `json:"Plant"`
	StorageLocation      string  `json:"StorageLocation"`
	CreationDate         *string `json:"CreationDate"`
	InventoryBlockStatus *bool   `json:"InventoryBlockStatus"`
	IsMarkedForDeletion  *bool   `json:"IsMarkedForDeletion"`
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
	PDTCompIsMarkedForBackflush   *bool   `json:"PDTCompIsMarkedForBackflush"`
	ProductionSchedulingProfile   *string `json:"ProductionSchedulingProfile"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
}

type Accounting struct {
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

type ProductDescription struct {
	Product            string  `json:"Product"`
	Language           string  `json:"Language"`
	ProductDescription *string `json:"ProductDescription"`
}

type ProductDescByBP struct {
	Product            string  `json:"Product"`
	BusinessPartner    int     `json:"BusinessPartner"`
	Language           string  `json:"Language"`
	ProductDescription *string `json:"ProductDescription"`
}

type Tax struct {
	Product                  string  `json:"Product"`
	Country                  string  `json:"Country"`
	ProductTaxCategory       string  `json:"ProductTaxCategory"`
	ProductTaxClassification *string `json:"ProductTaxClassification"`
}
