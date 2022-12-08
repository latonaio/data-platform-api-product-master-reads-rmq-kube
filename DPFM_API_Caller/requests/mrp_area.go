package requests

type MRPArea struct {
	Product                       string   `json:"Product"`
	BusinessPartner               *int     `json:"BusinessPartner"`
	Plant                         string   `json:"Plant"`
	MRPArea                       string   `json:"MRPArea"`
	MRPType                       string   `json:"MRPType"`
	MRPResponsible                string   `json:"MRPResponsible"`
	MRPGroup                      string   `json:"MRPGroup"`
	ReorderThresholdQuantity      *float32 `json:"ReorderThresholdQuantity"`
	PlanningTimeFence             *int     `json:"PlanningTimeFence"`
	LotSizeRoundingQuantity       *float32 `json:"LotSizeRoundingQuantity"`
	MinimumLotSizeQuantity        *float32 `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity        *float32 `json:"MaximumLotSizeQuantity"`
	MaximumStockQuantity          *float32 `json:"MaximumStockQuantity"`
	DfltStorageLocationExtProcmt  string   `json:"DfltStorageLocationExtProcmt"`
	MRPPlanningCalendar           string   `json:"MRPPlanningCalendar"`
	SafetyStockQuantity           *float32 `json:"SafetyStockQuantity"`
	SafetyDuration                *int     `json:"SafetyDuration"`
	FixedLotSizeQuantity          *float32 `json:"FixedLotSizeQuantity"`
	PlannedDeliveryDurationInDays *int     `json:"PlannedDeliveryDurationInDays"`
	StorageLocation               string   `json:"StorageLocation"`
	IsMarkedForDeletion           *bool    `json:"IsMarkedForDeletion"`
}
