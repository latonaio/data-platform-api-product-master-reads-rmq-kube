package requests

type StorageLocation struct {
	Product              string  `json:"Product"`
	BusinessPartner      int     `json:"BusinessPartner"`
	Plant                string  `json:"Plant"`
	StorageLocation      string  `json:"StorageLocation"`
	BlockStatus 		 *bool   `json:"BlockStatus"`
	CreationDate         string  `json:"CreationDate"`
	LastChangeDate       string  `json:"LastChangeDate"`
	IsMarkedForDeletion  *bool   `json:"IsMarkedForDeletion"`
}
