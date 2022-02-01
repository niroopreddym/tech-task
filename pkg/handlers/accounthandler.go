package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/suprageeks/tech-task/pkg/models"
)

func postAccountRequestBodyInitialValidation(accountDetails models.Account, errorMessages *[]string) {
	if strings.TrimSpace(accountDetails.BankName) == "" {
		errorMessage := "Attribute Missing: Name in the request body"
		*errorMessages = append(*errorMessages, errorMessage)
	}
}

//---------------------------------Account related endpoints-------------------------------------------------

//CreateAccount creates a account in bank
func (handler *BankAndAccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	accountDetails := models.Account{}
	log.Println("request:", r)
	bodyBytes, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		responseController(w, http.StatusInternalServerError, readErr)
		return
	}

	strBufferValue := string(bodyBytes)
	err := json.Unmarshal([]byte(strBufferValue), &accountDetails)
	if err != nil {
		responseController(w, http.StatusInternalServerError, err)
		return
	}

	errorMessages := []string{}
	postAccountRequestBodyInitialValidation(accountDetails, &errorMessages)
	if len(errorMessages) > 0 {
		responseController(w, http.StatusBadRequest, errorMessages)
		return
	}

	uniqueID, err := handler.DatabaseService.PostAccountDetails(&accountDetails)
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseController(w, http.StatusOK, map[string]string{
		"IdentityID": *uniqueID,
	})
}

//GetAccountDetails gets account details
func (handler *BankAndAccountHandler) GetAccountDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountID := params["id"]

	accDetails, err := handler.DatabaseService.GetAccountDetails(accountID)
	if accDetails.IdentityID == "" {
		responseController(w, http.StatusNotFound, "Bank Not Found")
		return
	}

	if err != nil {
		responseController(w, http.StatusInternalServerError, "Error occured while fetching the bank details")
		return
	}

	responseController(w, http.StatusOK, accDetails)
}
