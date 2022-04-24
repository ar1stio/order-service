package validation

import (
	"order-service/exception"
	"order-service/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidationLoginBuyer(request model.LoginBuyerReq) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func ValidationRegisterBuyer(request model.CreateBuyer) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.AlamatPengiriman, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func ValidationUpdateBuyer(request model.UpdateBuyer) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.AlamatPengiriman, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}