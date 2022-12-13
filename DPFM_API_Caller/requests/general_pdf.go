package requests

type GeneralPDF struct {
	Product      string `json:"Product"`
	DocType      string `json:"DocType"`
	DocVersionID *int   `json:"DocVersionID"`
	DocID        string `json:"DocID"`
	FileName     string `json:"FileName"`
}
