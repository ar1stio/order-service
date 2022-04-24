package service

import (
	"crypto/sha256"
	"encoding/hex"
	"order-service/exception"
	"order-service/model"
	"order-service/repository"
	"order-service/validation"
	"time"
)

type sellerServiceImpl struct {
	SellerRepository repository.SellerRepository
}

func NewSellerService(SellerRepository *repository.SellerRepository) SellerService {
	return &sellerServiceImpl{
		SellerRepository: *SellerRepository,
	}
}

func (service *sellerServiceImpl) Register(req model.CreateSeller) (err error) {
	validation.ValidationRegisterSeller(req)
	currentTime := time.Now()

	h := sha256.New()
	h.Write([]byte(req.Password))
	sha := h.Sum(nil) // "sha" is uint8 type, encoded in base16

	shaStr := hex.EncodeToString(sha) // String representation
	req.Password = string(shaStr)

	req.CreatedAt = currentTime.Format("2006-01-02 15:04")

	err = service.SellerRepository.Create(req)
	return err
}

func (service *sellerServiceImpl) Login(req model.LoginSellerReq) (res model.LoginSellerRes) {
	validation.ValidationLoginSeller(req)

	hpass := sha256.New()
	hpass.Write([]byte(req.Password))
	shapass := hpass.Sum(nil) // "sha" is uint8 type, encoded in base16

	shaStrpass := hex.EncodeToString(shapass) // String representation

	req.Password = string(shaStrpass)

	res, err := service.SellerRepository.Login(req)
	if err != nil {
		exception.PanicIfNeeded(exception.AuthorizedError{Status: "UNAUTHORIZED", Message: "Tidak dapat menghubungkan ke database"})
	}
	if res.Email == "" {
		exception.PanicIfNeeded(exception.AuthorizedError{Status: "UNAUTHORIZED", Message: "Password dan Email tidak ditemukan"})
	}

	return res
}

func (service *sellerServiceImpl) UpdateSeller(req model.UpdateSeller, AdminToken string) (err error) {
	// validation.ValidationIdSeller(req)

	currentTime := time.Now()

	req.UpdatedAt = currentTime.Format("2006-01-02 15:04:05")
	h := sha256.New()
	h.Write([]byte(req.Password))
	sha := h.Sum(nil) // "sha" is uint8 type, encoded in base16

	shaStr := hex.EncodeToString(sha) // String representation
	req.Password = string(shaStr)

	err = service.SellerRepository.Update(req)
	return err
}
