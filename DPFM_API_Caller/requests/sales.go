package requests

type Sales struct {
	Product             string `json:"Product"`
	BusinessPartner     *int   `json:"BusinessPartner"`
	Sellable            *bool  `json:"Sellable"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
