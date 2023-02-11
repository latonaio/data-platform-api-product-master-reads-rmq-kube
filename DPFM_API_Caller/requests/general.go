package requests

type General struct {
	Product                       string   `json:"Product"`
	ProductType                   *string  `json:"ProductType"`
	BaseUnit                      *string  `json:"BaseUnit"`
	ValidityStartDate             *string  `json:"ValidityStartDate"`
	ProductGroup                  *string  `json:"ProductGroup"`
	GrossWeight                   *float32 `json:"GrossWeight"`
	NetWeight                     *float32 `json:"NetWeight"`
	WeightUnit                    *string  `json:"WeightUnit"`
	InternalCapacityQuantity      *float32 `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit  *string  `json:"InternalCapacityQuantityUnit"`
	SizeOrDimensionText           *string  `json:"SizeOrDimensionText"`
	ProductStandardID             *string  `json:"ProductStandardID"`
	IndustryStandardName          *string  `json:"IndustryStandardName"`
	ItemCategory                  *string  `json:"ItemCategory"`
	BarcodeType                   *string  `json:"BarcodeType"`
	CountryOfOrigin               *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage       *string  `json:"CountryOfOriginLanguage"`
	ProductAccountAssignmentGroup *string  `json:"ProductAccountAssignmentGroup"`
	CreationDate                  *string  `json:"CreationDate"`
	LastChangeDate                *string  `json:"LastChangeDate"`
	IsMarkedForDeletion           *bool    `json:"IsMarkedForDeletion"`
}
