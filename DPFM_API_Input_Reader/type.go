package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string    `json:"connection_key"`
	Result            bool      `json:"result"`
	RedisKey          string    `json:"redis_key"`
	Filepath          string    `json:"filepath"`
	APIStatusCode     int       `json:"api_status_code"`
	RuntimeSessionID  string    `json:"runtime_session_id"`
	BusinessPartnerID *int      `json:"business_partner"`
	ServiceLabel      string    `json:"service_label"`
	APIType           string    `json:"api_type"`
	General           General   `json:"ProductMaster"`
	Generals          []General `json:"ProductMasters"`
	APISchema         string    `json:"api_schema"`
	Accepter          []string  `json:"accepter"`
	Deleted           bool      `json:"deleted"`
}

type General struct {
	Product                       string               `json:"Product"`
	ProductType                   *string              `json:"ProductType"`
	BaseUnit                      *string              `json:"BaseUnit"`
	ValidityStartDate             *string              `json:"ValidityStartDate"`
	ValidityEndDate               *string              `json:"ValidityEndDate"`
	ItemCategory                  *string              `json:"ItemCategory"`
	ProductGroup                  *string              `json:"ProductGroup"`
	GrossWeight                   *float32             `json:"GrossWeight"`
	NetWeight                     *float32             `json:"NetWeight"`
	WeightUnit                    *string              `json:"WeightUnit"`
	InternalCapacityQuantity      *float32             `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit  *string              `json:"InternalCapacityQuantityUnit"`
	SizeOrDimensionText           *string              `json:"SizeOrDimensionText"`
	ProductStandardID             *string              `json:"ProductStandardID"`
	IndustryStandardName          *string              `json:"IndustryStandardName"`
	CountryOfOrigin               *string              `json:"CountryOfOrigin"`
	CountryOfOriginLanguage       *string              `json:"CountryOfOriginLanguage"`
	LocalRegionOfOrigin           *string              `json:"LocalRegionOfOrigin"`
	LocalSubRegionOfOrigin        *string              `json:"LocalSubRegionOfOrigin"`
	BarcodeType                   *string              `json:"BarcodeType"`
	ProductAccountAssignmentGroup *string              `json:"ProductAccountAssignmentGroup"`
	CreationDate                  *string              `json:"CreationDate"`
	LastChangeDate                *string              `json:"LastChangeDate"`
	IsMarkedForDeletion           *bool                `json:"IsMarkedForDeletion"`
	ProductDescription            []ProductDescription `json:"ProductDescription"`
	BusinessPartner               []BusinessPartner    `json:"BusinessPartner"`
	GeneralDoc                    []GeneralDoc         `json:"GeneralDoc"`
	Tax                           []Tax                `json:"Tax"`
}

type BusinessPartner struct {
	Product                string            `json:"Product"`
	BusinessPartner        int               `json:"BusinessPartner"`
	ValidityStartDate      string            `json:"ValidityStartDate"`
	ValidityEndDate        string            `json:"ValidityEndDate"`
	BusinessPartnerProduct *string           `json:"BusinessPartnerProduct"`
	CreationDate           *string           `json:"CreationDate"`
	LastChangeDate         *string           `json:"LastChangeDate"`
	IsMarkedForDeletion    *bool             `json:"IsMarkedForDeletion"`
	BPPlant                []BPPlant         `json:"BPPlant"`
	ProductDescByBP        []ProductDescByBP `json:"ProductDescByBP"`
	Allergen               []Allergen        `json:"Allergen"`
	Calories               []Calories        `json:"Calories"`
	NutritionalInfo        []NutritionalInfo `json:"NutritionalInfo"`
}

type ProductDescription struct {
	Product             string            `json:"Product"`
	Language            string            `json:"Language"`
	ProductDescription  *string           `json:"ProductDescription"`
	CreationDate        *string           `json:"CreationDate"`
	LastChangeDate      *string           `json:"LastChangeDate"`
	IsMarkedForDeletion *bool             `json:"IsMarkedForDeletion"`
	ProductDescByBP     []ProductDescByBP `json:"ProductDescByBP"`
}

type ProductDescByBP struct {
	Product             string  `json:"Product"`
	BusinessPartner     int     `json:"BusinessPartner"`
	Language            string  `json:"Language"`
	ProductDescription  *string `json:"ProductDescription"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type BPPlant struct {
	Product                                   string            `json:"Product"`
	BusinessPartner                           int               `json:"BusinessPartner"`
	Plant                                     string            `json:"Plant"`
	MRPType                                   *string           `json:"MRPType"`
	MRPController                             *string           `json:"MRPController"`
	ReorderThresholdQuantityInBaseUnit        *float32          `json:"ReorderThresholdQuantityInBaseUnit"`
	PlanningTimeFenceInDays                   *int              `json:"PlanningTimeFenceInDays"`
	MRPPlanningCalendar                       *string           `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32          `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *float32          `json:"SafetyDuration"`
	SafetyDurationUnit                        *string           `json:"SafetyDurationUnit"`
	MaximumStockQuantityInBaseUnit            *float32          `json:"MaximumStockQuantityInBaseUnit"`
	MinimumDeliveryQuantityInBaseUnit         *float32          `json:"MinimumDeliveryQuantityInBaseUnit"`
	MinimumDeliveryLotSizeQuantityInBaseUnit  *float32          `json:"MinimumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryQuantityInBaseUnit        *float32          `json:"StandardDeliveryQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32          `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32          `json:"MaximumDeliveryQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32          `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32          `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool             `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDuration                  *float32          `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit              *string           `json:"StandardDeliveryDurationUnit"`
	IsBatchManagementRequired                 *bool             `json:"IsBatchManagementRequired"`
	BatchManagementPolicy                     *string           `json:"BatchManagementPolicy"`
	ProfitCenter                              *string           `json:"ProfitCenter"`
	CreationDate                              *string           `json:"CreationDate"`
	LastChangeDate                            *string           `json:"LastChangeDate"`
	IsMarkedForDeletion                       *bool             `json:"IsMarkedForDeletion"`
	MRPArea                                   []MRPArea         `json:"MRPArea"`
	Production                                []Production      `json:"Production"`
	Quality                                   []Quality         `json:"Quality"`
	Accounting                                []Accounting      `json:"Accounting"`
	StorageLocation                           []StorageLocation `json:"StorageLocation"`
	BPPlantDoc                                []BPPlantDoc      `json:"BPPlantDoc"`
}

type MRPArea struct {
	Product                                   string   `json:"Product"`
	BusinessPartner                           int      `json:"BusinessPartner"`
	Plant                                     string   `json:"Plant"`
	MRPArea                                   string   `json:"MRPArea"`
	MRPType                                   *string  `json:"MRPType"`
	MRPController                             *string  `json:"MRPController"`
	StorageLocationForMRP                     *string  `json:"StorageLocationForMRP"`
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
	CreationDate                              *string  `json:"CreationDate"`
	LastChangeDate                            *string  `json:"LastChangeDate"`
	IsMarkedForDeletion                       *bool    `json:"IsMarkedForDeletion"`
}

type Production struct {
	Product                                           string   `json:"Product"`
	BusinessPartner                                   int      `json:"BusinessPartner"`
	Plant                                             string   `json:"Plant"`
	ProductionStorageLocation                         *string  `json:"ProductionStorageLocation"`
	ProductProcessingDuration                         *float32 `json:"ProductProcessingDuration"`
	ProductProductionQuantityUnit                     *string  `json:"ProductProductionQuantityUnit"`
	MinimumProductionQuantityInBaseUnit               *float32 `json:"MinimumProductionQuantityInBaseUnit"`
	MinimumProductionLotSizeQuantityInBaseUnit        *float32 `json:"MinimumProductionLotSizeQuantityInBaseUnit"`
	StandardProductionQuantityInBaseUnit              *float32 `json:"StandardProductionQuantityInBaseUnit"`
	StandardProductionLotSizeQuantityInBaseUnit       *float32 `json:"StandardProductionLotSizeQuantityInBaseUnit"`
	MaximumProductionQuantityInBaseUnit               *float32 `json:"MaximumProductionQuantityInBaseUnit"`
	MaximumProductionLotSizeQuantityInBaseUnit        *float32 `json:"MaximumProductionLotSizeQuantityInBaseUnit"`
	ProductionLotSizeRoundingQuantityInBaseUnit       *float32 `json:"ProductionLotSizeRoundingQuantityInBaseUnit"`
	MinimumProductionQuantityInProductionUnit         *float32 `json:"MinimumProductionQuantityInProductionUnit"`
	MinimumProductionLotSizeQuantityInProductionUnit  *float32 `json:"MinimumProductionLotSizeQuantityInProductionUnit"`
	StandardProductionQuantityInProductionUnit        *float32 `json:"StandardProductionQuantityInProductionUnit"`
	StandardProductionLotSizeQuantityInProductionUnit *float32 `json:"StandardProductionLotSizeQuantityInProductionUnit"`
	MaximumProductionLotSizeQuantityInProductionUnit  *float32 `json:"MaximumProductionLotSizeQuantityInProductionUnit"`
	MaximumProductionQuantityInProductionUnit         *float32 `json:"MaximumProductionQuantityInProductionUnit"`
	ProductionLotSizeRoundingQuantityInProductionUnit *float32 `json:"ProductionLotSizeRoundingQuantityInProductionUnit"`
	ProductionLotSizeIsFixed                          *bool    `json:"ProductionLotSizeIsFixed"`
	ProductIsBatchManagedInProductionPlant            *bool    `json:"ProductIsBatchManagedInProductionPlant"`
	ProductIsMarkedForBackflush                       *bool    `json:"ProductIsMarkedForBackflush"`
	ProductionSchedulingProfile                       *string  `json:"ProductionSchedulingProfile"`
	CreationDate                                      *string  `json:"CreationDate"`
	LastChangeDate                                    *string  `json:"LastChangeDate"`
	IsMarkedForDeletion                               *bool    `json:"IsMarkedForDeletion"`
}

type Quality struct {
	Product                        string  `json:"Product"`
	BusinessPartner                int     `json:"BusinessPartner"`
	Plant                          string  `json:"Plant"`
	MaximumStoragePeriod           *string `json:"MaximumStoragePeriod"`
	QualityMgmtCtrlKey             *string `json:"QualityMgmtCtrlKey"`
	MatlQualityAuthorizationGroup  *string `json:"MatlQualityAuthorizationGroup"`
	HasPostToInspectionStock       *bool   `json:"HasPostToInspectionStock"`
	InspLotDocumentationIsRequired *bool   `json:"InspLotDocumentationIsRequired"`
	SuplrQualityManagementSystem   *string `json:"SuplrQualityManagementSystem"`
	RecrrgInspIntervalTimeInDays   *int    `json:"RecrrgInspIntervalTimeInDays"`
	ProductQualityCertificateType  *string `json:"ProductQualityCertificateType"`
	CreationDate                   *string `json:"CreationDate"`
	LastChangeDate                 *string `json:"LastChangeDate"`
	IsMarkedForDeletion            *bool   `json:"IsMarkedForDeletion"`
}

type StorageLocation struct {
	Product             string       `json:"Product"`
	BusinessPartner     int          `json:"BusinessPartner"`
	Plant               string       `json:"Plant"`
	StorageLocation     string       `json:"StorageLocation"`
	BlockStatus         *bool        `json:"BlockStatus"`
	CreationDate        *string      `json:"CreationDate"`
	LastChangeDate      *string      `json:"LastChangeDate"`
	IsMarkedForDeletion *bool        `json:"IsMarkedForDeletion"`
	StorageBin          []StorageBin `json:"StorageBin"`
}

type StorageBin struct {
	Product             string  `json:"Product"`
	BusinessPartner     int     `json:"BusinessPartner"`
	Plant               string  `json:"Plant"`
	StorageLocation     string  `json:"StorageLocation"`
	StorageBin          string  `json:"StorageBin"`
	BlockStatus         *bool   `json:"BlockStatus"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type GeneralDoc struct {
	Product                  string  `json:"Product"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	FileExtension            *string `json:"FileExtension"`
	FileName                 *string `json:"FileName"`
	FilePath                 *string `json:"FilePath"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}

type Accounting struct {
	Product             string   `json:"Product"`
	BusinessPartner     int      `json:"BusinessPartner"`
	Plant               string   `json:"Plant"`
	ValuationClass      *string  `json:"ValuationClass"`
	CostingPolicy       *string  `json:"CostingPolicy"`
	PriceUnitQty        *int     `json:"PriceUnitQty"`
	StandardPrice       *float32 `json:"StandardPrice"`
	MovingAveragePrice  *float32 `json:"MovingAveragePrice"`
	CreationDate        *string  `json:"CreationDate"`
	LastChangeDate      *string  `json:"LastChangeDate"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}

type Tax struct {
	Product                  string  `json:"Product"`
	Country                  string  `json:"Country"`
	ProductTaxCategory       string  `json:"ProductTaxCategory"`
	ProductTaxClassification *string `json:"ProductTaxClassification"`
	CreationDate             *string `json:"CreationDate"`
	LastChangeDate           *string `json:"LastChangeDate"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
}

type Allergen struct {
	Product             string  `json:"Product"`
	BusinessPartner     int     `json:"BusinessPartner"`
	Allergen            string  `json:"Allergen"`
	AllergenIsContained *bool   `json:"AllergenIsContained"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type Calories struct {
	Product             string   `json:"Product"`
	BusinessPartner     int      `json:"BusinessPartner"`
	Calories            *float32 `json:"Calories"`
	CaloryUnit          *string  `json:"CaloryUnit"`
	CaloryUnitQuantity  *int     `json:"CaloryUnitQuantity"`
	CreationDate        *string  `json:"CreationDate"`
	LastChangeDate      *string  `json:"LastChangeDate"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}

type NutritionalInfo struct {
	Product             string   `json:"Product"`
	BusinessPartner     int      `json:"BusinessPartner"`
	Nutrient            string   `json:"Nutrient"`
	NutrientContent     *float32 `json:"NutrientContent"`
	NutrientContentUnit *string  `json:"NutrientContentUnit"`
	CreationDate        *string  `json:"CreationDate"`
	LastChangeDate      *string  `json:"LastChangeDate"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}

type BPPlantDoc struct {
	Product                  string  `json:"Product"`
	BusinessPartner          int     `json:"BusinessPartner"`
	Plant                    string  `json:"Plant"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	FileExtension            *string `json:"FileExtension"`
	FileName                 *string `json:"FileName"`
	FilePath                 *string `json:"FilePath"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}
