package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"order-service/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderController_Create(t *testing.T) {
	// memberaddressRepository.DeleteAll()
	createAddressRequest := model.CreateOrder{
		Uuid:       "",
		Name:       "Aristio3",
		Password:   "admin123",
		Email:      "theeo_chau@yahoo.com",
		Handphone:  "085523865720",
		Token:      "",
		Status:     0,
		CreatedAt:  "",
		UpdatedAt:  "",
		VerifiedAt: "",
	}
	requestBody, _ := json.Marshal(createAddressRequest)

	request := httptest.NewRequest("POST", "/order-service/buyer/create-order", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("x-auth-token", "cc22cb9fe343b911f74d2cde1e1d9a8ebfd3b4785decc436e67190b8132aaf1d")

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	assert.Equal(t, "create user successfull", webResponse.Data)
}