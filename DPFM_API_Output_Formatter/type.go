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
	ProductDescription *[]ProductDescription `json:"ProductDescription"`
	BusinessPartner    *[]BusinessPartner    `json:"BusinessPartner"`
	ProductDescByBP    *[]ProductDescByBP    `json:"ProductDescByBP"`
	BPPlant            *[]BPPlant            `json:"BPPlant"`
	MRPArea            *[]MRPArea            `json:"MRPArea"`
	StorageLocation    *[]StorageLocation    `json:"StorageLocation"`
	Production         *[]Production         `json:"Production"`
	Quality            *[]Quality            `json:"Quality"`
	StorageBin         *[]StorageBin         `json:"StorageBin"`
	Accounting         *[]Accounting         `json:"Accounting"`
	Allergen           *[]Allergen           `json:"Allergen"`
	Calories           *[]Calories           `json:"Calories"`
	NutritionalInfo    *[]NutritionalInfo    `json:"NutritionalInfo"`
	Tax                *[]Tax                `json:"Tax"`
}

type General struct {
	Product                       string   `json:"Product"`
	ProductType                   string   `json:"ProductType"`
	BaseUnit                      string   `json:"BaseUnit"`
	ValidityStartDate             string   `json:"ValidityStartDate"`
	ValidityEndDate               string   `json:"ValidityEndDate"`
	ItemCategory                  string   `json:"ItemCategory"`
	ProductGroup                  *string  `json:"ProductGroup"`
	GrossWeight                   *float32 `json:"GrossWeight"`
	NetWeight                     *float32 `json:"NetWeight"`
	WeightUnit                    *string  `json:"WeightUnit"`
	InternalCapacityQuantity      *float32 `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit  *string  `json:"InternalCapacityQuantityUnit"`
	SizeOrDimensionText           *string  `json:"SizeOrDimensionText"`
	ProductStandardID             *string  `json:"ProductStandardID"`
	IndustryStandardName          *string  `json:"IndustryStandardName"`
	MarkingOfMaterial             *string  `json:"MarkingOfMaterial"`
	CountryOfOrigin               *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage       *string  `json:"CountryOfOriginLanguage"`
	LocalRegionOfOrigin           *string  `json:"LocalRegionOfOrigin"`
	LocalSubRegionOfOrigin        *string  `json:"LocalSubRegionOfOrigin"`
	BarcodeType                   *string  `json:"BarcodeType"`
	ProductAccountAssignmentGroup *string  `json:"ProductAccountAssignmentGroup"`
	CertificateAuthorityChain     *string  `json:"CertificateAuthorityChain"`
	UsageControlChain        	  *string  `json:"UsageControlChain"`
	CreationDate                  string   `json:"CreationDate"`
	LastChangeDate                string   `json:"LastChangeDate"`
	IsMarkedForDeletion           *bool    `json:"IsMarkedForDeletion"`
}

type BusinessPartner struct {
	Product                string  `json:"Product"`
	BusinessPartner        int     `json:"BusinessPartner"`
	ValidityStartDate      string  `json:"ValidityStartDate"`
	ValidityEndDate        string  `json:"ValidityEndDate"`
	BusinessPartnerProduct *string `json:"BusinessPartnerProduct"`
	CreationDate           string  `json:"CreationDate"`
	LastChangeDate         string  `json:"LastChangeDate"`
	IsMarkedForDeletion    *bool   `json:"IsMarkedForDeletion"`
}

type BPPlant struct {
	Product                                   string   `json:"Product"`
	BusinessPartner                           int      `json:"BusinessPartner"`
	Plant                                     string   `json:"Plant"`
	MRPType                                   string   `json:"MRPType"`
	MRPController                             *string  `json:"MRPController"`
	ReorderThresholdQuantityInBaseUnit        *float32 `json:"ReorderThresholdQuantityInBaseUnit"`
	PlanningTimeFenceInDays                   *int     `json:"PlanningTimeFenceInDays"`
	MRPPlanningCalendar                       *string  `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32 `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *float32 `json:"SafetyDuration"`
	SafetyDurationUnit                        *string  `json:"SafetyDurationUnit"`
	MaximumStockQuantityInBaseUnit            *float32 `json:"MaximumStockQuantityInBaseUnit"`
	MinimumDeliveryQuantityInBaseUnit         *float32 `json:"MinimumDeliveryQuantityInBaseUnit"`
	MinimumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MinimumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryQuantityInBaseUnit        *float32 `json:"StandardDeliveryQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32 `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32 `json:"MaximumDeliveryQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32 `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool    `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDuration                  *float32 `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit              *string  `json:"StandardDeliveryDurationUnit"`
	IsBatchManagementRequired                 *bool    `json:"IsBatchManagementRequired"`
	BatchManagementPolicy                     *string  `json:"BatchManagementPolicy"`
	ProfitCenter                              *string  `json:"ProfitCenter"`
	CreationDate                              string   `json:"CreationDate"`
	LastChangeDate                            string   `json:"LastChangeDate"`
	IsMarkedForDeletion                       *bool    `json:"IsMarkedForDeletion"`
}

type MRPArea struct {
	Product                                   string   `json:"Product"`
	BusinessPartner                           int      `json:"BusinessPartner"`
	Plant                                     string   `json:"Plant"`
	MRPArea                                   string   `json:"MRPArea"`
	MRPType                                   string   `json:"MRPType"`
	MRPController                             string   `json:"MRPController"`
	StorageLocationForMRP                     string   `json:"StorageLocationForMRP"`
	ReorderThresholdQuantityInBaseUnit        *float32 `json:"ReorderThresholdQuantityInBaseUnit"`
	PlanningTimeFenceInDays                   *int     `json:"PlanningTimeFenceInDays"`
	MRPPlanningCalendar                       *string  `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32 `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *float32 `json:"SafetyDuration"`
	SafetyDurationUnit                        *string  `json:"SafetyDurationUnit"`
	MaximumStockQuantityInBaseUnit            *float32 `json:"MaximumStockQuantityInBaseUnit"`
	MinumumDeliveryQuantityInBaseUnit         *float32 `json:"MinumumDeliveryQuantityInBaseUnit"`
	MinumumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MinumumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryQuantityInBaseUnit        *float32 `json:"StandardDeliveryQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32 `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32 `json:"MaximumDeliveryQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32 `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool    `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDuration                  *float32 `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit              *string  `json:"StandardDeliveryDurationUnit"`
	CreationDate                              string   `json:"CreationDate"`
	LastChangeDate                            string   `json:"LastChangeDate"`
	IsMarkedForDeletion                       *bool    `json:"IsMarkedForDeletion"`
}

type ProductDescription struct {
	Product             string `json:"Product"`
	Language            string `json:"Language"`
	ProductDescription  string `json:"ProductDescription"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type ProductDescByBP struct {
	Product             string `json:"Product"`
	BusinessPartner     int    `json:"BusinessPartner"`
	Language            string `json:"Language"`
	ProductDescription  string `json:"ProductDescription"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type Production struct {
	Product                                           string   `json:"Product"`
	BusinessPartner                                   int      `json:"BusinessPartner"`
	Plant                                             string   `json:"Plant"`
	ProductionStorageLocation                         string   `json:"ProductionStorageLocation"`
	ProductionDuration                         		  float32  `json:"ProductionDuration"`
	ProductionDurationUnit                     		  string   `json:"ProductionDurationUnit"`
	ProductionQuantityUnit                     		  string   `json:"ProductionQuantityUnit"`
	MinimumProductionQuantityInBaseUnit               float32  `json:"MinimumProductionQuantityInBaseUnit"`
	MinimumProductionLotSizeQuantityInBaseUnit        float32  `json:"MinimumProductionLotSizeQuantityInBaseUnit"`
	StandardProductionQuantityInBaseUnit              float32  `json:"StandardProductionQuantityInBaseUnit"`
	StandardProductionLotSizeQuantityInBaseUnit       float32  `json:"StandardProductionLotSizeQuantityInBaseUnit"`
	MaximumProductionQuantityInBaseUnit               float32  `json:"MaximumProductionQuantityInBaseUnit"`
	MaximumProductionLotSizeQuantityInBaseUnit        float32  `json:"MaximumProductionLotSizeQuantityInBaseUnit"`
	ProductionLotSizeRoundingQuantityInBaseUnit       *float32 `json:"ProductionLotSizeRoundingQuantityInBaseUnit"`
	MinimumProductionQuantityInProductionUnit         float32  `json:"MinimumProductionQuantityInProductionUnit"`
	MinimumProductionLotSizeQuantityInProductionUnit  float32  `json:"MinimumProductionLotSizeQuantityInProductionUnit"`
	StandardProductionQuantityInProductionUnit        float32  `json:"StandardProductionQuantityInProductionUnit"`
	StandardProductionLotSizeQuantityInProductionUnit float32  `json:"StandardProductionLotSizeQuantityInProductionUnit"`
	MaximumProductionQuantityInProductionUnit         float32  `json:"MaximumProductionQuantityInProductionUnit"`
	MaximumProductionLotSizeQuantityInProductionUnit  float32  `json:"MaximumProductionLotSizeQuantityInProductionUnit"`
	ProductionLotSizeRoundingQuantityInProductionUnit *float32 `json:"ProductionLotSizeRoundingQuantityInProductionUnit"`
	ProductionLotSizeIsFixed                          *bool    `json:"ProductionLotSizeIsFixed"`
	ProductIsBatchManagedInProductionPlant            *bool    `json:"ProductIsBatchManagedInProductionPlant"`
	ProductIsMarkedForBackflush                       *bool    `json:"ProductIsMarkedForBackflush"`
	ProductionSchedulingProfile                       *string  `json:"ProductionSchedulingProfile"`
	CreationDate                                      string   `json:"CreationDate"`
	LastChangeDate                                    string   `json:"LastChangeDate"`
	IsMarkedForDeletion                               *bool    `json:"IsMarkedForDeletion"`
}

type Quality struct {
	Product                        string  `json:"Product"`
	BusinessPartner                int     `json:"BusinessPartner"`
	Plant                          string  `json:"Plant"`
	QualityMgmtCtrlKey             *string `json:"QualityMgmtCtrlKey"`
	ProductSpecification		   *string `json:"ProductSpecification"`
	MaximumStoragePeriodInDays     *int	   `json:"MaximumStoragePeriodInDays"`
	RecrrgInspIntervalTimeInDays   *int    `json:"RecrrgInspIntervalTimeInDays"`
	ProductQualityCertificateType  *string `json:"ProductQualityCertificateType"`
	CreationDate                   string  `json:"CreationDate"`
	LastChangeDate                 string  `json:"LastChangeDate"`
	IsMarkedForDeletion            *bool   `json:"IsMarkedForDeletion"`
}

type StorageLocation struct {
	Product             string `json:"Product"`
	BusinessPartner     int    `json:"BusinessPartner"`
	Plant               string `json:"Plant"`
	StorageLocation     string `json:"StorageLocation"`
	BlockStatus         *bool  `json:"BlockStatus"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type StorageBin struct {
	Product             string `json:"Product"`
	BusinessPartner     int    `json:"BusinessPartner"`
	Plant               string `json:"Plant"`
	StorageLocation     string `json:"StorageLocation"`
	StorageBin          string `json:"StorageBin"`
	BlockStatus         *bool  `json:"BlockStatus"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type Tax struct {
	Product                  string  `json:"Product"`
	Country                  string  `json:"Country"`
	ProductTaxCategory       string  `json:"ProductTaxCategory"`
	ProductTaxClassification *string `json:"ProductTaxClassification"`
	CreationDate             string  `json:"CreationDate"`
	LastChangeDate           string  `json:"LastChangeDate"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
}

type Accounting struct {
	Product             string   `json:"Product"`
	BusinessPartner     int      `json:"BusinessPartner"`
	Plant               string   `json:"Plant"`
	ValuationClass      string   `json:"ValuationClass"`
	CostingPolicy       *string  `json:"CostingPolicy"`
	PriceUnitQty        *int     `json:"PriceUnitQty"`
	StandardPrice       *float32 `json:"StandardPrice"`
	MovingAveragePrice  *float32 `json:"MovingAveragePrice"`
	CreationDate        string   `json:"CreationDate"`
	LastChangeDate      string   `json:"LastChangeDate"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}

type Allergen struct {
	Product             string  `json:"Product"`
	BusinessPartner     int     `json:"BusinessPartner"`
	Allergen            string  `json:"Allergen"`
	AllergenIsContained *bool   `json:"AllergenIsContained"`
	CreationDate        string  `json:"CreationDate"`
	LastChangeDate      string  `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type Calories struct {
	Product             string   `json:"Product"`
	BusinessPartner     int      `json:"BusinessPartner"`
	Calories            *float32 `json:"Calories"`
	CaloryUnit          *string  `json:"CaloryUnit"`
	CaloryUnitQuantity  *int     `json:"CaloryUnitQuantity"`
	CreationDate        string   `json:"CreationDate"`
	LastChangeDate      string   `json:"LastChangeDate"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}

type NutritionalInfo struct {
	Product             string   `json:"Product"`
	BusinessPartner     int      `json:"BusinessPartner"`
	Nutrient            string   `json:"Nutrient"`
	NutrientContent     *float32 `json:"NutrientContent"`
	NutrientContentUnit *string  `json:"NutrientContentUnit"`
	CreationDate        string   `json:"CreationDate"`
	LastChangeDate      string   `json:"LastChangeDate"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}
