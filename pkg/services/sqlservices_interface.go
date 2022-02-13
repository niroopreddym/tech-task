package services

import "github.com/suprageeks/tech-task/pkg/models"

//ISQLService provides the data interface
type ISQLService interface {
	PostBankDetails(bankDetails *models.Bank) (*string, error)
	ListAllBanks(pageNumber int) ([]*models.Bank, error)
	GetBankDetails(id string) (*models.Bank, error)
	PatchBankDetails(id string, bankDetails models.Bank) error
	DeleteBank(id string) error

	PostAccountDetails(bankDetails *models.Account) (*string, error)
	GetAccountDetails(id string) (*models.Account, error)
}
