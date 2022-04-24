package service

import (
	"order-service/model"
)

type OrderService interface {
	Register(req model.CreateUser) (err error)
	Login(req model.GetLoginRequest) (res model.GetFindTokenUserResponse)
	AuthenticationToken(req model.GetUserFindRequest) (res bool)
	RegisterMember(req model.CreateUserCompany) (err error)

	CreateUser(req model.CreateUser, AdminToken string) (err error)
	UpdateUser(req model.UpdateUser, AdminToken string) (err error)
	FindTokenUser(req model.GetLoginRequest, AdminToken string) (res model.GetFindTokenUserResponse)
	ActivateUser(req model.UpdateActivateRequest, AdminToken string) (err error)
	NonActivateUser(req model.NonActivateUser, AdminToken string) (err error)
	FinsUserCompany(companyid int) (res []model.GetUserCompanyRespon)
}
