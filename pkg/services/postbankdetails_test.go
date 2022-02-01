package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/suprageeks/tech-task/pkg/mocks"
	"github.com/suprageeks/tech-task/pkg/models"
)

func Test_PostBankDetails(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	dbMock := mocks.NewMockDbIface(controller)
	dbMock.EXPECT().DbExecuteScalar(gomock.Any(), gomock.Any()).AnyTimes().Return(nil, nil)
	dbMock.EXPECT().DbClose().AnyTimes()
	service := DatabaseService{
		DatabaseService: dbMock,
	}

	bankDetails := &models.Bank{
		BankName:   "test",
		IFSCCode:   "ifsc",
		BranchName: "branch",
	}

	uuid, err := service.PostBankDetails(bankDetails)
	assert.Nil(t, err)
	assert.NotNil(t, uuid)

}
