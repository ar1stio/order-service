package service

import (
	"order-service/model"
	"order-service/repository"
	"strconv"
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
	id := strconv.Itoa(req.Id)
	sellerid := strconv.Itoa(req.SellerId)
	buyerid := strconv.Itoa(req.BuyerId)
	quantity := strconv.Itoa(req.Quantity)
	price := strconv.Itoa(req.Price)
	totalprice := strconv.Itoa(req.TotalPrice)
	filter := " where status = 0 and (id like '%" + id + "%' or buyer_name like '%" + req.BuyerName + "%' or seller_name like '%" + req.SellerName + "%' or items like '%" + req.Items + "%' or delivery_source_address like '%" + req.DeliverySourceAddress + "%' or delivery_destination_address like '%" + req.DeliveryDestinationAddress + "%' or buyer_id like '%" + buyerid + "%' or seller_id like '%" + sellerid + "%' or quantity like '%" + quantity + "%' or price like '%" + price + "%' or total_price like '%" + totalprice + "%' )"
	dataorder, _ := service.OrderRepository.ShowOrder(filter)
	res = dataorder

	if res == nil {
		res = []model.ShowOrder{}
	}

	return res
}

func (service *orderServiceImpl) FindsOrderList(req model.GetFindOrder) (res []model.ShowOrder) {
	id := strconv.Itoa(req.Id)
	sellerid := strconv.Itoa(req.SellerId)
	buyerid := strconv.Itoa(req.BuyerId)
	quantity := strconv.Itoa(req.Quantity)
	price := strconv.Itoa(req.Price)
	totalprice := strconv.Itoa(req.TotalPrice)
	filter := " where status = 1 and (id like '%" + id + "%' or buyer_name like '%" + req.BuyerName + "%' or seller_name like '%" + req.SellerName + "%' or items like '%" + req.Items + "%' or delivery_source_address like '%" + req.DeliverySourceAddress + "%' or delivery_destination_address like '%" + req.DeliveryDestinationAddress + "%' or buyer_id like '%" + buyerid + "%' or seller_id like '%" + sellerid + "%' or quantity like '%" + quantity + "%' or price like '%" + price + "%' or total_price like '%" + totalprice + "%' )"
	dataorder, _ := service.OrderRepository.ShowOrder(filter)
	res = dataorder

	if res == nil {
		res = []model.ShowOrder{}
	}

	return res
}
