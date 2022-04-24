package service

import (
	"order-service/model"
	"order-service/repository"
	"time"
)

type orderServiceImpl struct {
	OrderRepository repository.OrderRepository
}

func NewOrderService(OrderRepository *repository.OrderRepository) OrderService {
	return &orderServiceImpl{
		OrderRepository: *OrderRepository,
	}
}

func (service *orderServiceImpl) Register(req model.CreateOrder) (err error) {
	// validation.ValidationRegisterOrder(req)
	currentTime := time.Now()

	req.CreatedAt = currentTime.Format("2006-01-02 15:04")

	err = service.OrderRepository.Create(req)
	return err
}

func (service *orderServiceImpl) Delrivered(req model.Deliveredreq) (err error) {
	// validation.ValidationIdOrder(req)

	err = service.OrderRepository.Delivered(req)
	return err
}

func (service *orderServiceImpl) FindsOrderProduct(req model.GetFindOrder) (res []model.ShowOrder) {
	filter := ""
	dataorder, _ := service.OrderRepository.ShowOrder(filter)
	res = dataorder

	return res
}

func (service *orderServiceImpl) FindsOrderList(req model.GetFindOrder) (res []model.ShowOrder) {
	filter := ""
	dataorder, _ := service.OrderRepository.ShowOrder(filter)
	res = dataorder

	return res
}
