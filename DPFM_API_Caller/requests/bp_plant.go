package requests

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
