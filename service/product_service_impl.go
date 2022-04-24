package service

import (
	"order-service/model"
	"order-service/repository"
	"time"
)

type productServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(ProductRepository *repository.ProductRepository) ProductService {
	return &productServiceImpl{
		ProductRepository: *ProductRepository,
	}
}

func (service *productServiceImpl) Register(req model.CreateProduct) (err error) {
	// validation.ValidationRegisterProduct(req)
	currentTime := time.Now()

	req.CreatedAt = currentTime.Format("2006-01-02 15:04")

	err = service.ProductRepository.Create(req)
	return err
}

func (service *productServiceImpl) UpdateProduct(req model.UpdateProduct, AdminToken string) (err error) {
	// validation.ValidationIdProduct(req)

	currentTime := time.Now()

	req.UpdatedAt = currentTime.Format("2006-01-02 15:04:05")

	err = service.ProductRepository.Update(req)
	return err
}

func (service *productServiceImpl) FindSingleProduct(id int) (res model.ShowSingleProduct) {
	// validation.Validate(req)

	datasingleproduct, _ := service.ProductRepository.FindSingleProduct(id)
	res = datasingleproduct

	return res
}

func (service *productServiceImpl) FindAllProduct(req model.GetAllProductReq) (res []model.ShowProduct) {
	filter := ""
	datallproduct, _ := service.ProductRepository.FindAllProduct(filter)
	res = datallproduct

	return res
}
