package repository

import "order-service/model"

type OrderRepository interface {
	Create(req model.CreateOrder)(err error)
	Update(req model.UpdateOrder)(err error)
	ShowOrder(filter string) (res []model.ShowOrder, err error)

	Delivered(req model.Deliveredreq)(err error)
}