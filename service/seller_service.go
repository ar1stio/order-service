package service

import (
	"order-service/model"
)

type SellerService interface {
	Register(req model.CreateSeller)(err error)
	Login(req model.GetLoginRequest)(res model.GetFindTokenSellerResponse)
	UpdateSeller(req model.UpdateSeller, AdminToken string)(err error)
	FinsSellerCompany(companyid int) (res []model.GetSellerCompanyRespon)
}
