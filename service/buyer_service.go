package service

import (
	"order-service/model"
)

type BuyerService interface {
	Register(req model.CreateBuyer) (err error)
	UpdateBuyer(req model.UpdateBuyer) (err error)
	Login(req model.LoginBuyerReq) (res model.LoginBuyerRes)
}
