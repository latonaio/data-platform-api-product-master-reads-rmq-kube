package requests

type Calories struct {
	Product            	string   `json:"Product"`
	BusinessPartner    	int      `json:"BusinessPartner"`
	Calories           	*float32 `json:"Calories"`
	CaloryUnit         	*string  `json:"CaloryUnit"`
	CaloryUnitQuantity 	*int     `json:"CaloryUnitQuantity"`
	CreationDate       	string   `json:"CreationDate"`
	LastChangeDate     	string   `json:"LastChangeDate"`
	IsMarkedForDeletion	*bool    `json:"IsMarkedForDeletion"`
}
