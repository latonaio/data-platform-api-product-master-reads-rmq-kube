package requests

type Accounting struct {
	Product             string   `json:"Product"`
	BusinessPartner     *int     `json:"BusinessPartner"`
	Plant               string   `json:"Plant"`
	ValuationClass      string   `json:"ValuationClass"`
	CostingPolicy       string   `json:"CostingPolicy"`
	PriceUnitQty        string   `json:"PriceUnitQty"`
	StandardPrice       *float32 `json:"StandardPrice"`
	MovingAveragePrice  *float32 `json:"MovingAveragePrice"`
	PriceLastChangeDate string   `json:"PriceLastChangeDate"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}
