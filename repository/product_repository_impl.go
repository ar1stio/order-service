package repository

import (
	"database/sql"
	"order-service/config"
	"order-service/model"
)

type productRepoImpl struct {
	sqlDb sql.DB
}

func NewProductRepository(sqlDb *sql.DB) ProductRepository {
	return &productRepoImpl{
		sqlDb: *sqlDb,
	}
}

func (repo *productRepoImpl) Create(req model.CreateProduct)(err error){
	ctx, cancel := config.NewMySqlContext()
	defer cancel()
	
	query := "INSERT INTO product(product_name,description,price,seller_id,created_at) VALUES(?, ?, ?, ?, ?)"
	stmt, err := repo.sqlDb.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, req.ProductName, req.Description, req.Price, req.SellerId, req.CreatedAt)
	return err
}

func (repo *productRepoImpl) Update(req model.UpdateProduct)(err error){
	ctx, cancel := config.NewMySqlContext()
	defer cancel()

	query := "UPDATE product SET product_name = ? ,description = ? ,price = ? ,seller_id = ? ,updated_at = ? WHERE id = ? "
	stmt, err := repo.sqlDb.PrepareContext(ctx, query)
	if err != nil {
		return err
	}	

	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, req.ProductName, req.Description, req.Price, req.SellerId, req.UpdatedAt, req.Id)
	
	return err
}

func (repo *productRepoImpl) FindSingleProduct(id int)(res model.ShowSingleProduct, err error){
	ctx, cancel := config.NewMySqlContext()
	defer cancel()

	sqlStatement := "SELECT product_name,description,price,seller_id,created_at,update_at FROM product WHERE id= ?"
	row := repo.sqlDb.QueryRowContext(ctx, sqlStatement, id)
	err = row.Scan(&res.ProductName, &res.Description, &res.Price, &res.SellerId, &res.CreatedAt, &res.CreatedAt)
	
	return res, err
}

func (repo *productRepoImpl) FindAllProduct(filter string) (res []model.ShowProduct, err error) {
	ctx, cancel := config.NewMySqlContext()
	defer cancel()
	sqlStatement := "SELECT id,product_name, seller.name as seller_name, description, price, seller_id from product inner join seller on seller_id=id "+ filter 
	record, err := repo.sqlDb.QueryContext(ctx, sqlStatement)
	if err != nil {
		return nil, err
	}
	for record.Next() {
		var Id int
		var ProductName string
		var SellerName string
		var Description string
		var Price int
		var SellerId int

		err = record.Scan(&Id,&ProductName,&SellerName,&Description,&Price,&SellerId)
		if err != nil {
			return nil, err
		}

		res = append(res, model.ShowProduct{
			Id : Id,
			ProductName : ProductName,
			SellerName : SellerName,
			Description : Description,
			Price : Price,
			SellerId : SellerId,
		})
	}

	return res, err
}


