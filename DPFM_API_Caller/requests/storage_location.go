package requests

type StorageLocation struct {
	Product              string `json:"Product"`
	BusinessPartner      *int   `json:"BusinessPartner"`
	Plant                string `json:"Plant"`
	StorageLocation      string `json:"StorageLocation"`
	CreationDate         string `json:"CreationDate"`
	InventoryBlockStatus *bool  `json:"InventoryBlockStatus"`
	IsMarkedForDeletion  *bool  `json:"IsMarkedForDeletion"`
}
