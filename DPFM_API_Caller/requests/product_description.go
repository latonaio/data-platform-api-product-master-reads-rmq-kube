package requests

type ProductDescription struct {
	Product             string  `json:"Product"`
	Language            string  `json:"Language"`
	ProductDescription  string  `json:"ProductDescription"`
	CreationDate        string  `json:"CreationDate"`
	LastChangeDate      string  `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}
