package service

import (
	"order-service/model"
)

type ProductService interface {
	Register(req model.CreateProduct) (err error)
	UpdateProduct(req model.UpdateProduct, AdminToken string) (err error)

	FindSingleProduct(id int) (res model.ShowSingleProduct)
	FindAllProduct(req model.GetAllProductReq) (res []model.ShowProduct)
}
