package requests

type Tax struct {
	Product                  string  `json:"Product"`
	BusinessPartner          int     `json:"BusinessPartner"`
	Country                  *string `json:"Country"`
	TaxCategory              *string `json:"TaxCategory"`
	ProductTaxClassification *string `json:"ProductTaxClassification"`
}
