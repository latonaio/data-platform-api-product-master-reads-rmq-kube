package requests

type BusinessPartner struct {
	Product                string  `json:"Product"`
	BusinessPartner        int     `json:"BusinessPartner"`
	ValidityStartDate      string  `json:"ValidityStartDate"`
	ValidityEndDate        string  `json:"ValidityEndDate"`
	BusinessPartnerProduct *string `json:"BusinessPartnerProduct"`
    CreationDate           string  `json:"CreationDate"`
    LastChangeDate         string  `json:"LastChangeDate"`
	IsMarkedForDeletion    *bool   `json:"IsMarkedForDeletion"`
}
