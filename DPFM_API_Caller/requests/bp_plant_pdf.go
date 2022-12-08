package requests

type BPPlantPDF struct {
	Product         string `json:"Product"`
	BusinessPartner *int   `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	DocType         string `json:"DocType"`
	DocVersionID    *int   `json:"DocVersionID"`
	DocID           string `json:"DocID"`
	FileName        string `json:"FileName"`
}
