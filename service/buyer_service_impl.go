package service

import (
	"crypto/sha256"
	"encoding/hex"
	"order-service/exception"
	"order-service/model"
	"order-service/repository"
	"order-service/validation"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

	claims := jwt.MapClaims{
		"sub":  res.Id,
		"name": res.Name,
		"role": "buyer",
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		exception.PanicIfNeeded("Gagal membuat token")
	}

	res.Token = t

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

func (service *buyerServiceImpl) UpdateBuyer(req model.UpdateBuyer) (err error) {
	// validation.ValidationIdBuyer(req)

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
