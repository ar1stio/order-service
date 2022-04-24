package repository

import "order-service/model"

type SellerRepository interface {
	Login(req model.LoginSellerReq)(res model.LoginSellerRes, err error)
	
	Create(req model.CreateSeller)(err error)
	Update(req model.UpdateSeller)(err error)
}