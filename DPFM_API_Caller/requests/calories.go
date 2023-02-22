package requests

type Calories struct {
	Product            string  `json:"Product"`
	BusinessPartner    int     `json:"BusinessPartner"`
	CaloryUnitQuantity int     `json:"CaloryUnitQuantity"`
	Calories           *int    `json:"Calories"`
	CaloryUnit         *string `json:"CaloryUnit"`
}
