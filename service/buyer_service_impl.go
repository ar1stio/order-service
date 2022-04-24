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

type buyerServiceImpl struct {
	BuyerRepository repository.BuyerRepository
}

func NewBuyerService(BuyerRepository *repository.BuyerRepository) BuyerService {
	return &buyerServiceImpl{
		BuyerRepository: *BuyerRepository,
	}
}

func (service *buyerServiceImpl) Register(req model.CreateBuyer) (err error) {
	validation.ValidationRegisterBuyer(req)
	currentTime := time.Now()

	h := sha256.New()
	h.Write([]byte(req.Password))
	sha := h.Sum(nil) // "sha" is uint8 type, encoded in base16

	shaStr := hex.EncodeToString(sha) // String representation
	req.Password = string(shaStr)

	req.CreatedAt = currentTime.Format("2006-01-02 15:04")

	err = service.BuyerRepository.Create(req)
	return err
}

func (service *buyerServiceImpl) Login(req model.LoginBuyerReq) (res model.LoginBuyerRes) {
	validation.ValidationLoginBuyer(req)

	hpass := sha256.New()
	hpass.Write([]byte(req.Password))
	shapass := hpass.Sum(nil) // "sha" is uint8 type, encoded in base16

	shaStrpass := hex.EncodeToString(shapass) // String representation

	req.Password = string(shaStrpass)

	res, err := service.BuyerRepository.Login(req)
	if err != nil {
		exception.PanicIfNeeded(exception.AuthorizedError{Status: "UNAUTHORIZED", Message: "Tidak dapat menghubungkan ke database"})
	}
	if res.Email == "" {
		exception.PanicIfNeeded(exception.AuthorizedError{Status: "UNAUTHORIZED", Message: "Password dan Email tidak ditemukan"})
	}
	// currentTime := time.Now()

	// h := sha256.New()
	// h.Write([]byte(currentTime.Format("2006.01.02 15:04:05")))
	// sha := h.Sum(nil) // "sha" is uint8 type, encoded in base16

	// shaStr := hex.EncodeToString(sha) // String representation

	// Token := string(shaStr)
	// // req.BuyerToken = Token
	// tokendata := model.GetUpdateTokenRequest{
	// 	BuyerEmail: req.BuyerEmail,
	// 	BuyerToken: Token,
	// }

	// service.BuyerRepository.UpdateToken(tokendata)
	return res
}

func (service *buyerServiceImpl) UpdateBuyer(req model.UpdateBuyer, AdminToken string) (err error) {
	// validation.ValidationIdBuyer(req)
	//testmncportaladmin sha256
	if AdminToken != "cc22cb9fe343b911f74d2cde1e1d9a8ebfd3b4785decc436e67190b8132aaf1d" {
		exception.PanicIfNeeded(exception.AuthorizedError{Status: "UNAUTHORIZED", Message: "Token admin tidak sesuai"})
	}

	currentTime := time.Now()

	req.UpdatedAt = currentTime.Format("2006-01-02 15:04:05")
	h := sha256.New()
	h.Write([]byte(req.Password))
	sha := h.Sum(nil) // "sha" is uint8 type, encoded in base16

	shaStr := hex.EncodeToString(sha) // String representation
	req.Password = string(shaStr)

	err = service.BuyerRepository.Update(req)
	return err
}
