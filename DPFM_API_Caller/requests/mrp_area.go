package requests

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
	StandardDeliveryDuration	              *float32 `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit              *string  `json:"StandardDeliveryDurationUnit"`
	CreationDate                  			  string   `json:"CreationDate"`
	LastChangeDate                			  string   `json:"LastChangeDate"`
	IsMarkedForDeletion                       *bool    `json:"IsMarkedForDeletion"`
}
