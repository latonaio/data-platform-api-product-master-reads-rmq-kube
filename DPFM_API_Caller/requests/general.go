package requests

type General struct {
	Product                       string   `json:"Product"`
	ProductType                   *string  `json:"ProductType"`
	BaseUnit                      *string  `json:"BaseUnit"`
	ValidityStartDate             *string  `json:"ValidityStartDate"`
	ProductGroup                  *string  `json:"ProductGroup"`
	Division                      *string  `json:"Division"`
	GrossWeight                   *float32 `json:"GrossWeight"`
	WeightUnit                    *string  `json:"WeightUnit"`
	SizeOrDimensionText           *string  `json:"SizeOrDimensionText"`
	IndustryStandardName          *string  `json:"IndustryStandardName"`
	ProductStandardID             *string  `json:"ProductStandardID"`
	CreationDate                  *string  `json:"CreationDate"`
	LastChangeDate                *string  `json:"LastChangeDate"`
	NetWeight                     *float32 `json:"NetWeight"`
	CountryOfOrigin               *string  `json:"CountryOfOrigin"`
	ItemCategory                  *string  `json:"ItemCategory"`
	ProductAccountAssignmentGroup *string  `json:"ProductAccountAssignmentGroup"`
	IsMarkedForDeletion           *bool    `json:"IsMarkedForDeletion"`
}
