package service

import (
	"order-service/model"
)

type SellerService interface {
	Register(req model.CreateSeller) (err error)
	Login(req model.LoginSellerReq) (res model.LoginSellerRes)
	UpdateSeller(req model.UpdateSeller, AdminToken string) (err error)
}
