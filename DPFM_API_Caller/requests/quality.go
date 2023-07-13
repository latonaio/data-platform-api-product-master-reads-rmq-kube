package requests

type Quality struct {
	Product                        string  `json:"Product"`
	BusinessPartner                int     `json:"BusinessPartner"`
	Plant                          string  `json:"Plant"`
	MaximumStoragePeriod           *string `json:"MaximumStoragePeriod"`
	QualityMgmtCtrlKey             *string `json:"QualityMgmtCtrlKey"`
	MatlQualityAuthorizationGroup  *string `json:"MatlQualityAuthorizationGroup"`
	HasPostToInspectionStock       *bool   `json:"HasPostToInspectionStock"`
	InspLotDocumentationIsRequired *bool   `json:"InspLotDocumentationIsRequired"`
	SuplrQualityManagementSystem   *string `json:"SuplrQualityManagementSystem"`
	RecrrgInspIntervalTimeInDays   *int    `json:"RecrrgInspIntervalTimeInDays"`
	ProductQualityCertificateType  *string `json:"ProductQualityCertificateType"`
	CreationDate                   *string `json:"CreationDate"`
	LastChangeDate                 *string `json:"LastChangeDate"`
	IsMarkedForDeletion            *bool   `json:"IsMarkedForDeletion"`
}
