package handlers

import "net/http"

//BankHandlerIface provides the interface for handling the bank related crud operations
type BankHandlerIface interface {
	CreateBank(w http.ResponseWriter, r *http.Request)
	GetAllBanks(w http.ResponseWriter, r *http.Request)
	RemoveBank(w http.ResponseWriter, r *http.Request)
	GetBankDetails(w http.ResponseWriter, r *http.Request)
	UpdateBankDetails(w http.ResponseWriter, r *http.Request)
}

//AccountHandlerIface provides the interface for handling the product related
type AccountHandlerIface interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
	GetAccountDetails(w http.ResponseWriter, r *http.Request)
}

//CompositeIface achieves interface composition
type CompositeIface interface {
	BankHandlerIface
	AccountHandlerIface
}
