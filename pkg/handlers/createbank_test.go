package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/suprageeks/tech-task/pkg/mocks"
	"github.com/suprageeks/tech-task/pkg/models"
)

func Test_CreateBank_ReturnsSuccess(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	dbMock := mocks.NewMockISQLService(controller)
	uuidValue := uuid.New().String()
	dbMock.EXPECT().PostBankDetails(gomock.Any()).AnyTimes().Return(&uuidValue, nil)
	handler := BankAndAccountHandler{
		DatabaseService: dbMock,
	}

	bankData := models.Bank{
		BankName:   "test",
		IFSCCode:   "ifsc",
		BranchName: "branch",
	}
	byteArr, _ := json.Marshal(bankData)
	ioReader := io.NopCloser(bytes.NewReader(byteArr))

	req := httptest.NewRequest(http.MethodPost, "/bank", ioReader)
	w := httptest.NewRecorder()
	handler.CreateBank(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	assert.NotEmpty(t, data)
}
