package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/suprageeks/tech-task/pkg/models"
	"github.com/suprageeks/tech-task/pkg/services"
)

//BankAndAccountHandler is the class implementation for CompositeIface Interface
type BankAndAccountHandler struct {
	DatabaseService services.ISQLService
}

//NewBankAndAccountsHandlerInstance instantiates the struct
func NewBankAndAccountsHandlerInstance() CompositeIface {
	return &BankAndAccountHandler{
		DatabaseService: services.NewDatabaseServicesInstance(),
	}
}

//---------------------------------Bank related endpoints-------------------------------------------------

//CreateBank creates a bank in DB
func (handler *BankAndAccountHandler) CreateBank(w http.ResponseWriter, r *http.Request) {
	bankDetails := models.Bank{}
	bodyBytes, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		responseController(w, http.StatusInternalServerError, readErr)
		return
	}

	strBufferValue := string(bodyBytes)
	err := json.Unmarshal([]byte(strBufferValue), &bankDetails)
	if err != nil {
		responseController(w, http.StatusInternalServerError, err)
		return
	}

	errorMessages := []string{}
	postRequestBodyInitialValidation(bankDetails, &errorMessages)
	if len(errorMessages) > 0 {
		responseController(w, http.StatusBadRequest, errorMessages)
		return
	}

	uniqueID, err := handler.DatabaseService.PostBankDetails(&bankDetails)
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseController(w, http.StatusOK, map[string]string{
		"BankUUID": *uniqueID,
	})
}

//GetAllBanks gets all banks from DB
func (handler *BankAndAccountHandler) GetAllBanks(w http.ResponseWriter, r *http.Request) {
	lstBanks, err := handler.DatabaseService.ListAllBanks()
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, "Error occured while fetch the userdetails")
		return
	}

	responseController(w, http.StatusOK, lstBanks)
}

//DeleteBank deletes a bank from DB
func (handler *BankAndAccountHandler) DeleteBank(w http.ResponseWriter, r *http.Request) {}

//GetBankDetails gets the bank details
func (handler *BankAndAccountHandler) GetBankDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bankID := params["id"]

	bankDetails, err := handler.DatabaseService.GetBankDetails(bankID)
	if bankDetails.UUID == "" {
		responseController(w, http.StatusNotFound, "Bank Not Found")
		return
	}

	if err != nil {
		responseController(w, http.StatusInternalServerError, "Error occured while fetching the bank details")
		return
	}

	responseController(w, http.StatusOK, bankDetails)
}

//UpdateBankDetails updates the bank details
func (handler *BankAndAccountHandler) UpdateBankDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bankID := params["id"]

	bankDetails, err := handler.DatabaseService.GetBankDetails(bankID)
	if err != nil {
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	if bankDetails.UUID == "" {
		responseController(w, http.StatusBadRequest, "Bank Not Found")
		return
	}

	requestedBankDetails := models.Bank{}
	bodyBytes, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Println("reading from buffer")
		err := errors.New("error reading data from the response " + readErr.Error())
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	strBufferValue := string(bodyBytes)
	err = json.Unmarshal([]byte(strBufferValue), &requestedBankDetails)
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	errorMessages := []string{}
	patchCallInitialValidation(requestedBankDetails, &errorMessages)
	if len(errorMessages) > 0 {
		responseController(w, http.StatusBadRequest, errorMessages)
		return
	}

	err = handler.DatabaseService.PatchBankDetails(bankID, requestedBankDetails)
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseController(w, http.StatusNoContent, "Updated Sucessfully")
}

//RemoveBank deletes the bank record
func (handler *BankAndAccountHandler) RemoveBank(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bankID := params["id"]

	err := handler.DatabaseService.DeleteBank(bankID)
	if err != nil {
		responseController(w, http.StatusInternalServerError, err.Error())
	}

	responseController(w, http.StatusNoContent, "Successfully Deleted")
}

func postRequestBodyInitialValidation(bankDetails models.Bank, errorMessages *[]string) {
	if strings.TrimSpace(bankDetails.BankName) == "" {
		errorMessage := "Attribute Missing: Name in the request body"
		*errorMessages = append(*errorMessages, errorMessage)
	}
}

func responseController(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func patchCallInitialValidation(bankDetails models.Bank, errorMessages *[]string) {
	if bankDetails.IFSCCode == "" {
		errorMessage := "Invalid Attribute: IFSCCode in the request body"
		*errorMessages = append(*errorMessages, errorMessage)
	}
}
