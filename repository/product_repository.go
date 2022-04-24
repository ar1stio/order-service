package repository

import "order-service/model"

type ProductRepository interface {
	Create(req model.CreateProduct)(err error)
	Update(req model.UpdateProduct)(err error)

	FindSingleProduct(id int)(res model.ShowSingleProduct, err error)
	FindAllProduct(filter string) (res []model.ShowProduct, err error)
}