package service

import (
	"order-service/model"
)

type OrderService interface {
	Register(req model.CreateOrder) (err error)
	Delrivered(req model.Deliveredreq) (err error)
	FindsOrderProduct(req model.GetFindOrder) (res []model.ShowOrder)
	FindsOrderList(req model.GetFindOrder) (res []model.ShowOrder)
}
