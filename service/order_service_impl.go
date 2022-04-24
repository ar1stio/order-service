package service

import (
	"crypto/sha256"
	"encoding/hex"
	"order-service/exception"
	"order-service/model"
	"order-service/repository"
	"order-service/validation"
	"time"

	"github.com/google/uuid"
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
	validation.ValidationRegisterOrder(req)
	currentTime := time.Now()

	req.OrderUuid = uuid.New().String()
	h := sha256.New()
	h.Write([]byte(req.OrderPassword))
	sha := h.Sum(nil) // "sha" is uint8 type, encoded in base16

	shaStr := hex.EncodeToString(sha) // String representation
	req.OrderPassword = string(shaStr)

	req.OrderCreatedAt = currentTime.Format("2006-01-02 15:04")

	err = service.OrderRepository.Create(req)
	return err
}

func (service *orderServiceImpl) RegisterMember(req model.CreateOrderCompany) (err error) {
	currentTime := time.Now()

	req.MemberCreatedAt = currentTime.Format("2006-01-02 15:04")

	err = service.OrderRepository.CreateMember(req)
	return err
}

func (service *orderServiceImpl) Login(req model.GetLoginRequest) (res model.GetFindTokenOrderResponse) {
	validation.ValidationLoginOrder(req)

	hpass := sha256.New()
	hpass.Write([]byte(req.OrderPassword))
	shapass := hpass.Sum(nil) // "sha" is uint8 type, encoded in base16

	shaStrpass := hex.EncodeToString(shapass) // String representation

	req.OrderPassword = string(shaStrpass)

	res, err := service.OrderRepository.Login(req)
	if err != nil {
		exception.PanicIfNeeded(exception.AuthorizedError{Status: "UNAUTHORIZED", Message: "Tidak dapat menghubungkan ke database"})
	}
	if res.OrderUuid == "" {
		exception.PanicIfNeeded(exception.AuthorizedError{Status: "UNAUTHORIZED", Message: "Password dan Email tidak ditemukan"})
	}
	currentTime := time.Now()

	h := sha256.New()
	h.Write([]byte(currentTime.Format("2006.01.02 15:04:05")))
	sha := h.Sum(nil) // "sha" is uint8 type, encoded in base16

	shaStr := hex.EncodeToString(sha) // String representation

	Token := string(shaStr)
	// req.OrderToken = Token
	tokendata := model.GetUpdateTokenRequest{
		OrderEmail: req.OrderEmail,
		OrderToken: Token,
	}

	service.OrderRepository.UpdateToken(tokendata)
	return res
}

func (service *orderServiceImpl) AuthenticationToken(req model.GetOrderFindRequest) (res bool) {
	// validation.Validate(req)
	dataTokenValid, err := service.OrderRepository.FindValidToken(req)

	if err != nil {
		exception.PanicIfNeeded(exception.AuthorizedError{Status: "UNAUTHORIZED", Message: "Tidak dapat menghubungkan ke database"})
	}

	if dataTokenValid.Total == "0" {
		exception.PanicIfNeeded(exception.AuthorizedError{Status: "UNAUTHORIZED", Message: "Token anda tidak ditemukan"})
	}

	res = true

	return res
}

func (service *orderServiceImpl) ActivateOrder(req model.UpdateActivateRequest, AdminToken string) (err error) {
	validation.ValidationActivatetOrder(req)
	//testmncportaladmin sha256
	if AdminToken != "cc22cb9fe343b911f74d2cde1e1d9a8ebfd3b4785decc436e67190b8132aaf1d" {
		exception.PanicIfNeeded(exception.AuthorizedError{Status: "UNAUTHORIZED", Message: "Token admin tidak sesuai"})
	}

	currentTime := time.Now()

	req.OrderVerifiedAt = currentTime.Format("2006-01-02 15:04")
	req.OrderStatus = 1

	err = service.OrderRepository.Activate(req)
	return err
}

func (service *orderServiceImpl) NonActivateOrder(req model.NonActivateOrder, AdminToken string) (err error) {
	validation.ValidationNonActivateOrder(req)
	//testmncportaladmin sha256
	if AdminToken != "cc22cb9fe343b911f74d2cde1e1d9a8ebfd3b4785decc436e67190b8132aaf1d" {
		exception.PanicIfNeeded(exception.AuthorizedError{Status: "UNAUTHORIZED", Message: "Token admin tidak sesuai"})
	}
	req.OrderStatus = 0

	err = service.OrderRepository.NonActivate(req)
	return err
}

func (service *orderServiceImpl) CreateOrder(req model.CreateOrder, AdminToken string) (err error) {
	validation.ValidationCreateOrder(req)
	//testmncportaladmin sha256
	if AdminToken != "cc22cb9fe343b911f74d2cde1e1d9a8ebfd3b4785decc436e67190b8132aaf1d" {
		exception.PanicIfNeeded(exception.AuthorizedError{Status: "UNAUTHORIZED", Message: "Token admin tidak sesuai"})
	}

	currentTime := time.Now()

	req.OrderUuid = uuid.New().String()
	req.OrderCreatedAt = currentTime.Format("2006-01-02 15:04")

	err = service.OrderRepository.Create(req)
	return err
}

func (service *orderServiceImpl) UpdateOrder(req model.UpdateOrder, AdminToken string) (err error) {
	// validation.ValidationIdOrder(req)
	//testmncportaladmin sha256
	if AdminToken != "cc22cb9fe343b911f74d2cde1e1d9a8ebfd3b4785decc436e67190b8132aaf1d" {
		exception.PanicIfNeeded(exception.AuthorizedError{Status: "UNAUTHORIZED", Message: "Token admin tidak sesuai"})
	}

	currentTime := time.Now()

	req.OrderUpdatedAt = currentTime.Format("2006-01-02 15:04:05")
	h := sha256.New()
	h.Write([]byte(req.OrderPassword))
	sha := h.Sum(nil) // "sha" is uint8 type, encoded in base16

	shaStr := hex.EncodeToString(sha) // String representation
	req.OrderPassword = string(shaStr)

	err = service.OrderRepository.Update(req)
	return err
}

func (service *orderServiceImpl) Delrivered(req model.DeliveredreqUpdateOrder, AdminToken string) (err error) {
	// validation.ValidationIdOrder(req)
	//testmncportaladmin sha256
	if AdminToken != "cc22cb9fe343b911f74d2cde1e1d9a8ebfd3b4785decc436e67190b8132aaf1d" {
		exception.PanicIfNeeded(exception.AuthorizedError{Status: "UNAUTHORIZED", Message: "Token admin tidak sesuai"})
	}

	currentTime := time.Now()

	req.OrderUpdatedAt = currentTime.Format("2006-01-02 15:04:05")
	h := sha256.New()
	h.Write([]byte(req.OrderPassword))
	sha := h.Sum(nil) // "sha" is uint8 type, encoded in base16

	shaStr := hex.EncodeToString(sha) // String representation
	req.OrderPassword = string(shaStr)

	err = service.OrderRepository.Update(req)
	return err
}

func (service *orderServiceImpl) FindTokenOrder(req model.GetLoginRequest, AdminToken string) (res model.GetFindTokenOrderResponse) {
	// validation.Validate(req)
	//testmncportaladmin sha256
	if AdminToken != "cc22cb9fe343b911f74d2cde1e1d9a8ebfd3b4785decc436e67190b8132aaf1d" {
		exception.PanicIfNeeded(exception.AuthorizedError{Status: "UNAUTHORIZED", Message: "Token admin tidak sesuai"})
	}

	dataorder, _ := service.OrderRepository.FindOrderToken(req)
	res = dataorder

	return res
}

func (service *orderServiceImpl) FinsOrderCompany(companyid int) (res []model.GetOrderCompanyRespon) {
	dataorder, _ := service.OrderRepository.ShowOrder(companyid)
	res = dataorder

	return res
}
