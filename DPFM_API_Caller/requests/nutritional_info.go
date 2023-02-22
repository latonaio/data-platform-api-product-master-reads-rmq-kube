package requests

type NutritionalInfo struct {
	Product             string  `json:"Product"`
	BusinessPartner     int     `json:"BusinessPartner"`
	Nutrient            int     `json:"Nutrient"`
	NutrientContent     *int    `json:"NutrientContent"`
	NutrientContentUnit *string `json:"NutrientContentUnit"`
}
