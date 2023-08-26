package requests

type Quality struct {
	Product                        string  `json:"Product"`
	BusinessPartner                int     `json:"BusinessPartner"`
	Plant                          string  `json:"Plant"`
	QualityMgmtCtrlKey             *string `json:"QualityMgmtCtrlKey"`
	ProductSpecification		   *string `json:"ProductSpecification"`
	MaximumStoragePeriodInDays     *int	   `json:"MaximumStoragePeriodInDays"`
	RecrrgInspIntervalTimeInDays   *int    `json:"RecrrgInspIntervalTimeInDays"`
	ProductQualityCertificateType  *string `json:"ProductQualityCertificateType"`
	CreationDate                   string  `json:"CreationDate"`
	LastChangeDate                 string  `json:"LastChangeDate"`
	IsMarkedForDeletion            *bool   `json:"IsMarkedForDeletion"`
}
