package models

//Account ...
type Account struct {
	AccountID         int    `json:"accountid"`
	AccountHolderName string `json:"accountholdername"`
	BankName          string `json:"bankname"`
	BranchName        string `json:"branchname"`
	IdentityID        string `json:"identityid"`
	FirstName         string `json:"firstname"`
	LastName          string `json:"lastname"`
	Address           string `json:"address"`
}
