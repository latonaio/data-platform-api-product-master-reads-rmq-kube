package requests

type WorkScheduling struct {
	Product                       string   `json:"Product"`
	BusinessPartner               int      `json:"BusinessPartner"`
	Plant                         string   `json:"Plant"`
	ProductionInvtryManagedLoc    *string  `json:"ProductionInvtryManagedLoc"`
	ProductProcessingTime         *int     `json:"ProductProcessingTime"`
	ProductionSupervisor          *string  `json:"ProductionSupervisor"`
	ProductProductionQuantityUnit *string  `json:"ProductProductionQuantityUnit"`
	ProdnOrderIsBatchRequired     *bool    `json:"ProdnOrderIsBatchRequired"`
	PDTCompIsMarkedForBackflush   *bool    `json:"PDTCompIsMarkedForBackflush"`
	ProductionSchedulingProfile   *string  `json:"ProductionSchedulingProfile"`
	MinimumLotSizeQuantity        *float32 `json:"MinimumLotSizeQuantity"`
	StandardLotSizeQuantity       *float32 `json:"StandardLotSizeQuantity"`
	LotSizeRoundingQuantity       *float32 `json:"LotSizeRoundingQuantity"`
	MaximumLotSizeQuantity        *float32 `json:"MaximumLotSizeQuantity"`
	LotSizeIsFixed                *bool    `json:"LotSizeIsFixed"`
	IsMarkedForDeletion           *bool    `json:"IsMarkedForDeletion"`
}
