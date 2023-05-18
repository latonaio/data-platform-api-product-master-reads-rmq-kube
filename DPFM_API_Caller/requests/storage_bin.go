package requests

type StorageBin struct {
	Product             string  `json:"Product"`
	BusinessPartner     int     `json:"BusinessPartner"`
	Plant               string  `json:"Plant"`
	StorageLocation     string  `json:"StorageLocation"`
	StorageBin          string  `json:"StorageBin"`
	CreationDate        *string `json:"CreationDate"`
	BlockStatus         *bool   `json:"BlockStatus"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}
