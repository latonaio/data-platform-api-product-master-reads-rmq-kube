package requests

type ProductDescByBP struct {
	Product             string  `json:"Product"`
	BusinessPartner     int     `json:"BusinessPartner"`
	Language            string  `json:"Language"`
	ProductDescription  string  `json:"ProductDescription"`
	CreationDate        string  `json:"CreationDate"`
	LastChangeDate      string  `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}
