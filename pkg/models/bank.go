package models

//Bank ...
type Bank struct {
	BankID     int    `json:"bankid"`
	UUID       string `json:"uuid"`
	BankName   string `json:"bankname"`
	IFSCCode   string `json:"ifsccode"`
	BranchName string `json:"branchname"`
}
