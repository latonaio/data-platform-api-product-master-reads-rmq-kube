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
	General            *General              `json:"General"`
	ProductDescByBP    *[]ProductDescByBP    `json:"ProductDescByBP"`
	BusinessPartner    *[]BusinessPartner    `json:"BusinessPartner"`
	BPPlant            *[]BPPlant            `json:"BPPlant"`
	Tax                *[]Tax                `json:"Tax"`
	Accounting         *[]Accounting         `json:"Accounting"`
	MRPArea            *[]MRPArea            `json:"MRPArea"`
	ProductDescription *[]ProductDescription `json:"ProductDescription"`
	StorageLocation    *[]StorageLocation    `json:"StorageLocation"`
	WorkScheduling     *[]WorkScheduling     `json:"WorkScheduling"`
}

type ProductMaster struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Deleted       string `json:"deleted"`
}

type General struct {
	Product                       string               `json:"Product"`
	ProductType                   *string              `json:"ProductType"`
	BaseUnit                      *string              `json:"BaseUnit"`
	ValidityStartDate             *string              `json:"ValidityStartDate"`
	ProductGroup                  *string              `json:"ProductGroup"`
	Division                      *string              `json:"Division"`
	GrossWeight                   *float32             `json:"GrossWeight"`
	WeightUnit                    *string              `json:"WeightUnit"`
	SizeOrDimensionText           *string              `json:"SizeOrDimensionText"`
	IndustryStandardName          *string              `json:"IndustryStandardName"`
	ProductStandardID             *string              `json:"ProductStandardID"`
	CreationDate                  *string              `json:"CreationDate"`
	LastChangeDate                *string              `json:"LastChangeDate"`
	NetWeight                     *float32             `json:"NetWeight"`
	CountryOfOrigin               *string              `json:"CountryOfOrigin"`
	CountryOfOriginLanguage       *string              `json:"CountryOfOriginLanguage"`
	ItemCategory                  *string              `json:"ItemCategory"`
	ProductAccountAssignmentGroup *string              `json:"ProductAccountAssignmentGroup"`
	IsMarkedForDeletion           *bool                `json:"IsMarkedForDeletion"`
	BusinessPartner               []BusinessPartner    `json:"BusinessPartner"`
	ProductDescription            []ProductDescription `json:"ProductDescription"`
	Tax                           []Tax                `json:"Tax"`
}

type BusinessPartner struct {
	Product                string            `json:"Product"`
	BusinessPartner        int               `json:"BusinessPartner"`
	ValidityEndDate        string            `json:"ValidityEndDate"`
	ValidityStartDate      string            `json:"ValidityStartDate"`
	BusinessPartnerProduct *string           `json:"BusinessPartnerProduct"`
	IsMarkedForDeletion    *bool             `json:"IsMarkedForDeletion"`
	BPPlant                []BPPlant         `json:"BPPlant"`
	ProductDescByBP        []ProductDescByBP `json:"ProductDescByBP"`
}

type BPPlant struct {
	Product                                   string            `json:"Product"`
	BusinessPartner                           int               `json:"BusinessPartner"`
	Plant                                     string            `json:"Plant"`
	AvailabilityCheckType                     *string           `json:"AvailabilityCheckType"`
	MRPType                                   *string           `json:"MRPType"`
	MRPController                             *string           `json:"MRPController"`
	ReorderThresholdQuantity                  *float32          `json:"ReorderThresholdQuantity"`
	PlanningTimeFence                         *int              `json:"PlanningTimeFence"`
	MRPPlanningCalender                       *string           `json:"MRPPlanningCalender"`
	SafetyStockQuantityInBaseUnit             *float32          `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *int              `json:"SafetyDuration"`
	MaximumStockQuantityInBaseUnit            *float32          `json:"MaximumStockQuantityInBaseUnit"`
	MinimumDeliveryQuantityInBaseUnit         *float32          `json:"MinimumDeliveryQuantityInBaseUnit"`
	MinimumDeliveryLotSizeQuantityInBaseUnit  *float32          `json:"MinimumDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32          `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32          `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32          `json:"MaximumDeliveryQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool             `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDurationInDays            *int              `json:"StandardDeliveryDurationInDays"`
	IsBatchManagementRequired                 *bool             `json:"IsBatchManagementRequired"`
	BatchManagementPolicy                     *string           `json:"BatchManagementPolicy"`
	InventoryUnit                             *string           `json:"InventoryUnit"`
	ProfitCenter                              *string           `json:"ProfitCenter"`
	IsMarkedForDeletion                       *bool             `json:"IsMarkedForDeletion"`
	StorageLocation                           []StorageLocation `json:"StorageLocation"`
	MRPArea                                   []MRPArea         `json:"MRPArea"`
	WorkScheduling                            []WorkScheduling  `json:"WorkScheduling"`
	Accounting                                []Accounting      `json:"Accounting"`
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
	MinimumDeliveryQuantityInBaseUnit         *float32 `json:"MinimumDeliveryQuantityInBaseUnit"`
	MinimumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MinimumDeliveryLotSizeQuantityInBaseUnit"`
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
