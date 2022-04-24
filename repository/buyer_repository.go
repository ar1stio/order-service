package repository

import "order-service/model"

type BuyerRepository interface {
	Login(req model.LoginBuyerReq)(res model.LoginBuyerRes, err error)
	
	Create(req model.CreateBuyer)(err error)
	Update(req model.UpdateBuyer)(err error)
}