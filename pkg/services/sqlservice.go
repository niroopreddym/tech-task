package services

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/suprageeks/tech-task/pkg/database"
	"github.com/suprageeks/tech-task/pkg/models"
)

//DatabaseService is the class implementation for ProductServicesIface interface
type DatabaseService struct {
	DatabaseService database.DbIface
}

//NewDatabaseServicesInstance instantiates the struct
func NewDatabaseServicesInstance() *DatabaseService {
	return &DatabaseService{
		DatabaseService: database.DBNewHandler(),
	}
}

//PostBankDetails posts the bank data from DB
func (service *DatabaseService) PostBankDetails(bankDetails *models.Bank) (*string, error) {
	defer service.DatabaseService.DbClose()

	uuid := uuid.New().String()
	query := `INSERT INTO Bank (bank_uuid, bank_name, ifsc_code, branch_name) VALUES ($1, $2, $3, $4)`
	_, err := service.DatabaseService.DbExecuteScalar(query, uuid, bankDetails.BankName, bankDetails.IFSCCode, bankDetails.BranchName)
	if err != nil {
		log.Println(err)
		return nil, errors.New("internal server error")
	}

	return &uuid, nil
}

//ListAllBanks retrives all bank records from DB
func (service *DatabaseService) ListAllBanks() ([]*models.Bank, error) {
	defer service.DatabaseService.DbClose()
	query := "select * from Bank"
	tx, err := service.DatabaseService.TxBegin()
	rowsAffected, err := service.DatabaseService.TxQuery(tx, query)
	if err != nil {
		log.Println(err)
		return nil, errors.New("internal server error")
	}

	txResult := []*models.Bank{}
	for rowsAffected.Next() {
		var id int
		var bankUUID string
		var bankName string
		var ifscCode string
		var branchName string

		if err := rowsAffected.Scan(&id, &bankUUID, &bankName, &ifscCode, &branchName); err != nil {
			log.Println(err)
			log.Fatal(err)
		}

		bank := &models.Bank{
			BankID:     id,
			UUID:       bankUUID,
			BankName:   bankName,
			IFSCCode:   ifscCode,
			BranchName: branchName,
		}

		txResult = append(txResult, bank)
	}

	if err = service.DatabaseService.TxComplete(tx); err != nil {
		return nil, errors.New("internal server error")
	}

	return txResult, nil
}

//GetBankDetails fetches the bank details
func (service *DatabaseService) GetBankDetails(id string) (*models.Bank, error) {
	defer service.DatabaseService.DbClose()
	query := "select * from Bank where bank_id=" + id
	tx, err := service.DatabaseService.TxBegin()
	rowsAffected, err := service.DatabaseService.TxQuery(tx, query)
	if err != nil {
		log.Println(err)
		return nil, errors.New("internal server error")
	}

	txResult := &models.Bank{}
	for rowsAffected.Next() {
		var id int
		var bankUUID string
		var bankName string
		var ifscCode string
		var branchName string

		if err := rowsAffected.Scan(&id, &bankUUID, &bankName, &ifscCode, &branchName); err != nil {
			log.Println(err)
			log.Fatal(err)
		}

		txResult = &models.Bank{
			BankID:     id,
			UUID:       bankUUID,
			BankName:   bankName,
			IFSCCode:   ifscCode,
			BranchName: branchName,
		}

	}

	if err = service.DatabaseService.TxComplete(tx); err != nil {
		return nil, errors.New("internal server error")
	}

	return txResult, nil
}

//PatchBankDetails updates the bank details
func (service *DatabaseService) PatchBankDetails(id string, bankDetails models.Bank) error {
	defer service.DatabaseService.DbClose()
	query := fmt.Sprintf("update Bank set bank_name = '%v', branch_name='%v', ifsc_code='%v' where bank_id=%v", bankDetails.BankName, bankDetails.BranchName, bankDetails.IFSCCode, id)

	tx, err := service.DatabaseService.TxBegin()
	_, err = service.DatabaseService.TxExecuteStmt(tx, query)

	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}

	if err = service.DatabaseService.TxComplete(tx); err != nil {
		return errors.New("internal server error")
	}

	return nil
}

//DeleteBank deletes the bank record
func (service *DatabaseService) DeleteBank(id string) error {
	defer service.DatabaseService.DbClose()
	query := fmt.Sprintf("delete from Bank where bank_id=%v", id)

	tx, err := service.DatabaseService.TxBegin()
	_, err = service.DatabaseService.TxExecuteStmt(tx, query)

	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}

	if err = service.DatabaseService.TxComplete(tx); err != nil {
		return errors.New("internal server error")
	}

	return nil
}

//PostAccountDetails creates account data record
func (service *DatabaseService) PostAccountDetails(accountDetails *models.Account) (*string, error) {
	defer service.DatabaseService.DbClose()

	uuid := uuid.New().String()
	query := `INSERT INTO Account (bank_name, branch_name, account_holder_name, identity_id, first_name, last_name, acc_holder_addr) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := service.DatabaseService.DbExecuteScalar(query, accountDetails.BankName, accountDetails.BranchName,
		accountDetails.AccountHolderName, uuid, accountDetails.FirstName, accountDetails.LastName, accountDetails.Address)
	if err != nil {
		log.Println(err)
		return nil, errors.New("internal server error")
	}

	return &uuid, nil
}

//GetAccountDetails fetches acc details
func (service *DatabaseService) GetAccountDetails(id string) (*models.Account, error) {
	defer service.DatabaseService.DbClose()
	query := "select * from Account where account_id=" + id
	tx, err := service.DatabaseService.TxBegin()
	rowsAffected, err := service.DatabaseService.TxQuery(tx, query)
	if err != nil {
		log.Println(err)
		return nil, errors.New("internal server error")
	}

	txResult := &models.Account{}
	for rowsAffected.Next() {
		var id int
		var accHolderName string
		var bankName string
		var branchName string
		var identity string
		var firstName string
		var lastName string
		var address string

		if err := rowsAffected.Scan(&id, &bankName, &branchName, &accHolderName, &identity, &firstName, &lastName, &address); err != nil {
			log.Println(err)
			log.Fatal(err)
		}

		txResult = &models.Account{
			AccountID:         id,
			AccountHolderName: accHolderName,
			BankName:          bankName,
			BranchName:        branchName,
			IdentityID:        identity,
			FirstName:         firstName,
			LastName:          lastName,
			Address:           address,
		}
	}

	if err = service.DatabaseService.TxComplete(tx); err != nil {
		return nil, errors.New("internal server error")
	}

	return txResult, nil
}
