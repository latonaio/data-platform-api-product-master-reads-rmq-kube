package requests

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
