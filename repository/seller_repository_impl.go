package repository

import (
	"database/sql"
	"order-service/config"
	"order-service/model"
)

type sellerRepoImpl struct {
	sqlDb sql.DB
}

func NewSellerRepository(sqlDb *sql.DB) SellerRepository {
	return &sellerRepoImpl{
		sqlDb: *sqlDb,
	}
}

func (repo *sellerRepoImpl) Create(req model.CreateSeller) (err error) {
	ctx, cancel := config.NewMySqlContext()
	defer cancel()

	query := "INSERT INTO seller(email,name,password,alamat_pickup,created_at) VALUES(?, ?, ?, ?, ?)"
	stmt, err := repo.sqlDb.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, req.Email, req.Name, req.Password, req.AlamatPickup, req.CreatedAt)
	return err
}

func (repo *sellerRepoImpl) Update(req model.UpdateSeller) (err error) {
	ctx, cancel := config.NewMySqlContext()
	defer cancel()

	query := "UPDATE seller SET email = ? ,name = ? ,password = ? ,alamat_pickup = ? ,updated_at = ? WHERE id = ? "
	stmt, err := repo.sqlDb.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, req.Email, req.Name, req.Password, req.AlamatPickup, req.UpdatedAt, req.Id)

	return err
}

func (repo *sellerRepoImpl) Login(req model.LoginSellerReq) (res model.LoginSellerRes, err error) {
	ctx, cancel := config.NewMySqlContext()
	defer cancel()
	sqlStatement := "SELECT id,name,password,email,created_at FROM seller WHERE email= '" + req.Email + "' and password = '" + req.Password + "'"
	row := repo.sqlDb.QueryRowContext(ctx, sqlStatement)
	err = row.Scan(&res.Id, &res.Name, &res.Password, &res.Email, &res.CreatedAt)

	return res, err
}
