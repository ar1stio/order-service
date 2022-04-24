package validation

import (
	"order-service/exception"
	"order-service/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidationLoginSeller(request model.LoginSellerReq) {
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

func ValidationRegisterSeller(request model.CreateSeller) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.AlamatPickup, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func ValidationUpdateSeller(request model.UpdateSeller) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.AlamatPickup, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}