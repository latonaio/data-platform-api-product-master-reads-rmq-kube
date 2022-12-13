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
	General                             *General                             `json:"General"`
	ProductDescriptionByBusinessPartner *ProductDescriptionByBusinessPartner `json:"ProductDescriptionByBusinessPartner"`
	BusinessPartner                     *BusinessPartner                     `json:"BusinessPartner"`
	BPPlant                             *BPPlant                             `json:"BPPlant"`
	Tax                                 *Tax                                 `json:"Tax"`
	Accounting                          *Accounting                          `json:"Accounting"`
	MRPArea                             *MRPArea                             `json:"MRPArea"`
	Procurement                         *Procurement                         `json:"Procurement"`
	ProductDescription                  *ProductDescription                  `json:"ProductDescription"`
	Sales                               *Sales                               `json:"Sales"`
	StorageLocation                     *StorageLocation                     `json:"StorageLocation"`
	WorkScheduling                      *WorkScheduling                      `json:"WorkScheduling"`
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
	Product                       string   `json:"Product"`
	ProductType                   *string  `json:"ProductType"`
	BaseUnit                      *string  `json:"BaseUnit"`
	ValidityStartDate             *string  `json:"ValidityStartDate"`
	ProductGroup                  *string  `json:"ProductGroup"`
	Division                      *string  `json:"Division"`
	GrossWeight                   *float32 `json:"GrossWeight"`
	WeightUnit                    *string  `json:"WeightUnit"`
	SizeOrDimensionText           *string  `json:"SizeOrDimensionText"`
	IndustryStandardName          *string  `json:"IndustryStandardName"`
	ProductStandardID             *string  `json:"ProductStandardID"`
	CreationDate                  *string  `json:"CreationDate"`
	LastChangeDate                *string  `json:"LastChangeDate"`
	NetWeight                     *float32 `json:"NetWeight"`
	CountryOfOrigin               *string  `json:"CountryOfOrigin"`
	ItemCategory                  *string  `json:"ItemCategory"`
	ProductAccountAssignmentGroup *string  `json:"ProductAccountAssignmentGroup"`
	IsMarkedForDeletion           *bool    `json:"IsMarkedForDeletion"`
}
type GeneralPDF struct {
	DocType      string `json:"DocType"`
	DocVersionID *int   `json:"DocVersionID"`
	DocID        string `json:"DocID"`
	FileName     string `json:"FileName"`
}
type BusinessPartner struct {
	Product                string  `json:"Product"`
	BusinessPartner        int     `json:"BusinessPartner"`
	ValidityEndDate        *string `json:"ValidityEndDate"`
	ValidityStartDate      *string `json:"ValidityStartDate"`
	BusinessPartnerProduct *string `json:"BusinessPartnerProduct"`
	IsMarkedForDeletion    *bool   `json:"IsMarkedForDeletion"`
}

type BPPlant struct {
	Product                   string  `json:"Product"`
	BusinessPartner           int     `json:"BusinessPartner"`
	Plant                     string  `json:"Plant"`
	Issuable                  *bool   `json:"Issuable"`
	Receivable                *bool   `json:"Receivable"`
	IssuingStorageLocation    *string `json:"IssuingStorageLocation"`
	ReceivingStorageLocation  *string `json:"ReceivingStorageLocation"`
	AvailabilityCheckType     *string `json:"AvailabilityCheckType"`
	ProfitCenter              *string `json:"ProfitCenter"`
	MRPType                   *string `json:"MRPType"`
	MRPResponsible            *string `json:"MRPResponsible"`
	MinimumLotSizeQuantity    *string `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity    *string `json:"MaximumLotSizeQuantity"`
	FixedLotSizeQuantity      *string `json:"FixedLotSizeQuantity"`
	IsBatchManagementRequired *bool   `json:"IsBatchManagementRequired"`
	ProcurementType           *string `json:"ProcurementType"`
	InventoryUnit             *string `json:"InventoryUnit"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}

type BPPlantPDF struct {
	Product         string `json:"Product"`
	BusinessPartner *int   `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	DocType         string `json:"DocType"`
	DocVersionID    *int   `json:"DocVersionID"`
	DocID           string `json:"DocID"`
	FileName        string `json:"FileName"`
}

type StorageLocation struct {
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
	Product                       string   `json:"Product"`
	BusinessPartner               int      `json:"BusinessPartner"`
	Plant                         string   `json:"Plant"`
	MRPArea                       *string  `json:"MRPArea"`
	MRPType                       *string  `json:"MRPType"`
	MRPResponsible                *string  `json:"MRPResponsible"`
	MRPGroup                      *string  `json:"MRPGroup"`
	ReorderThresholdQuantity      *float32 `json:"ReorderThresholdQuantity"`
	PlanningTimeFence             *int     `json:"PlanningTimeFence"`
	LotSizeRoundingQuantity       *float32 `json:"LotSizeRoundingQuantity"`
	MinimumLotSizeQuantity        *float32 `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity        *float32 `json:"MaximumLotSizeQuantity"`
	MaximumStockQuantity          *float32 `json:"MaximumStockQuantity"`
	DfltStorageLocationExtProcmt  *string  `json:"DfltStorageLocationExtProcmt"`
	MRPPlanningCalendar           *string  `json:"MRPPlanningCalendar"`
	SafetyStockQuantity           *float32 `json:"SafetyStockQuantity"`
	SafetyDuration                *int     `json:"SafetyDuration"`
	FixedLotSizeQuantity          *float32 `json:"FixedLotSizeQuantity"`
	PlannedDeliveryDurationInDays *int     `json:"PlannedDeliveryDurationInDays"`
	StorageLocation               *string  `json:"StorageLocation"`
	IsMarkedForDeletion           *bool    `json:"IsMarkedForDeletion"`
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
