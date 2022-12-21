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
	ConnectionKey     string   `json:"connection_key"`
	Result            bool     `json:"result"`
	RedisKey          string   `json:"redis_key"`
	Filepath          string   `json:"filepath"`
	APIStatusCode     int      `json:"api_status_code"`
	RuntimeSessionID  string   `json:"runtime_session_id"`
	BusinessPartnerID *int     `json:"business_partner"`
	ServiceLabel      string   `json:"service_label"`
	APIType           string   `json:"api_type"`
	General           General  `json:"Product"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	Deleted           bool     `json:"deleted"`
}

type General struct {
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
	BusinessPartner                     BusinessPartner                     `json:"BusinessPartner"`
	ProductDescription                  ProductDescription                  `json:"ProductDescription"`
	ProductDescriptionByBusinessPartner ProductDescriptionByBusinessPartner `json:"ProductDescriptionByBusinessPartner"`
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
	IsMarkedForDeletion    *bool   `json:"IsMarkedForDeletion"`
	BusinessPartnerProduct *string `json:"BusinessPartnerProduct"`
	BPPlant                BPPlant `json:"BPPlant"`
	Sales                  Sales   `json:"Sales"`
	Tax                    Tax     `json:"Tax"`
}

type BPPlant struct {
	Product                                   string          `json:"Product"`
	BusinessPartner                           int             `json:"BusinessPartner"`
	Plant                                     string          `json:"Plant"`
	AvailabilityCheckType                     *string         `json:"AvailabilityCheckType"`
	ProfitCenter                              *string         `json:"ProfitCenter"`
	MRPType                                   *string         `json:"MRPType"`
	MRPController                             *string         `json:"MRPController"`
	ReorderThresholdQuantity                  *float32        `json:"ReorderThresholdQuantity"`
	PlanningTimeFence                         *int            `json:"PlanningTimeFence"`
	MRPPlanningCalendar                       *string         `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32        `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *int            `json:"SafetyDuration"`
	MaximumStockQuantityInBaseUnit            *float32        `json:"MaximumStockQuantityInBaseUnit"`
	MinumumDeliveryQuantityInBaseUnit         *float32        `json:"MinumumDeliveryQuantityInBaseUnit"`
	MinumumDeliveryLotSizeQuantityInBaseUnit  *float32        `json:"MinumumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32        `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32        `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32        `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32        `json:"MaximumDeliveryQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool           `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDurationInDays            *int            `json:"StandardDeliveryDurationInDays"`
	IsBatchManagementRequired                 *bool           `json:"IsBatchManagementRequired"`
	BatchManagementPolicy                     *string         `json:"BatchManagementPolicy"`
	InventoryUnit                             *string         `json:"InventoryUnit"`
	IsMarkedForDeletion                       *bool           `json:"IsMarkedForDeletion"`
	BPPlantPDF                                BPPlantPDF      `json:"BPPlantPDF"`
	StorageLocation                           StorageLocation `json:"StorageLocation"`
	Procurement                               Procurement     `json:"Procurement"`
	MRPArea                                   MRPArea         `json:"MRPArea"`
	WorkScheduling                            WorkScheduling  `json:"WorkScheduling"`
	Accounting                                Accounting      `json:"Accounting"`
}

type BPPlantPDF struct {
	DocType      string `json:"DocType"`
	DocVersionID *int   `json:"DocVersionID"`
	DocID        string `json:"DocID"`
	FileName     string `json:"FileName"`
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
