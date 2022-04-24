package repository

import (
	"database/sql"
	"order-service/config"
	"order-service/model"
)

type orderRepoImpl struct {
	sqlDb sql.DB
}

func NewOrderRepository(sqlDb *sql.DB) OrderRepository {
	return &orderRepoImpl{
		sqlDb: *sqlDb,
	}
}

func (repo *orderRepoImpl) Create(req model.CreateOrder) (err error) {
	ctx, cancel := config.NewMySqlContext()
	defer cancel()

	query := "INSERT INTO orders(buyer_id,seller_id,buyer_name,seller_name,delivery_source_address,delivery_destination_address,items,quantity,price,total_price,created_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := repo.sqlDb.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, req.BuyerId, req.SellerId, req.BuyerName, req.SellerName, req.DeliverySourceAddress, req.DeliveryDestinationAddress, req.Items, req.Quantity, req.Price, req.TotalPrice, req.CreatedAt)
	return err
}

func (repo *orderRepoImpl) Update(req model.UpdateOrder) (err error) {
	ctx, cancel := config.NewMySqlContext()
	defer cancel()

	query := "UPDATE orders SET buyer_id = ? ,seller_id = ? ,buyer_name = ? ,seller_name = ? ,delivery_source_address = ? ,delivery_destination_address = ? ,items = ? ,quantity = ? ,price = ? ,total_price = ? ,updated_at = ? WHERE order_uuid = ? "
	stmt, err := repo.sqlDb.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, req.BuyerId, req.SellerId, req.BuyerName, req.SellerName, req.DeliverySourceAddress, req.DeliveryDestinationAddress, req.Items, req.Quantity, req.Price, req.TotalPrice, req.CreatedAt, req.Id)

	return err
}

func (repo *orderRepoImpl) Delivered(req model.Deliveredreq) (err error) {
	ctx, cancel := config.NewMySqlContext()
	defer cancel()

	query := "UPDATE orders SET status = 1  WHERE id = ? "
	stmt, err := repo.sqlDb.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, req.Id)

	return err
}

func (repo *orderRepoImpl) ShowOrder(filter string) (res []model.ShowOrder, err error) {
	ctx, cancel := config.NewMySqlContext()
	defer cancel()
	sqlStatement := "SELECT id,buyer_id,seller_id,buyer_name,seller_name,delivery_source_address,delivery_destination_address,items,quantity,status,price,total_price,created_at from orders " + filter
	record, err := repo.sqlDb.QueryContext(ctx, sqlStatement)
	if err != nil {
		return nil, err
	}
	for record.Next() {
		var Id int
		var BuyerId int
		var SellerId int
		var BuyerName string
		var SellerName string
		var DeliverySourceAddress string
		var DeliveryDestinationAddress string
		var Items string
		var Quantity int
		var Status int
		var Price int
		var TotalPrice int
		var CreatedAt string

		err = record.Scan(&Id, &BuyerId, &SellerId, &BuyerName, &SellerName, &DeliverySourceAddress, &DeliveryDestinationAddress, &Items, &Quantity, &Status, &Price, &TotalPrice, &CreatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, model.ShowOrder{
			Id:                         Id,
			BuyerId:                    BuyerId,
			SellerId:                   SellerId,
			BuyerName:                  BuyerName,
			SellerName:                 SellerName,
			DeliverySourceAddress:      DeliverySourceAddress,
			DeliveryDestinationAddress: DeliveryDestinationAddress,
			Status:                     Status,
			Items:                      Items,
			Quantity:                   Quantity,
			Price:                      Price,
			TotalPrice:                 TotalPrice,
			CreatedAt:                  CreatedAt,
		})
	}

	return res, err
}
