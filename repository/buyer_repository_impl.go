package repository

import (
	"database/sql"
	"order-service/config"
	"order-service/model"
)

type buyerRepoImpl struct {
	sqlDb sql.DB
}

func NewBuyerRepository(sqlDb *sql.DB) BuyerRepository {
	return &buyerRepoImpl{
		sqlDb: *sqlDb,
	}
}

func (repo *buyerRepoImpl) Create(req model.CreateBuyer)(err error){
	ctx, cancel := config.NewMySqlContext()
	defer cancel()
	
	query := "INSERT INTO buyer(email,name,password,alamat_pengiriman,created_at) VALUES(?, ?, ?, ?, ?)"
	stmt, err := repo.sqlDb.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, req.Email, req.Name, req.Password, req.AlamatPengiriman, req.CreatedAt)
	return err
}

func (repo *buyerRepoImpl) Update(req model.UpdateBuyer)(err error){
	ctx, cancel := config.NewMySqlContext()
	defer cancel()

	query := "UPDATE buyer SET email = ? ,name = ? ,password = ? ,updated_at = ? WHERE id = ? "
	stmt, err := repo.sqlDb.PrepareContext(ctx, query)
	if err != nil {
		return err
	}	

	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, req.Email, req.Name, req.Password, req.AlamatPengiriman, req.UpdatedAt, req.Id)
	
	return err
}

func (repo *buyerRepoImpl) Login(req model.LoginBuyerReq)(res model.LoginBuyerRes, err error){
	ctx, cancel := config.NewMySqlContext()
	defer cancel()
	sqlStatement := "SELECT email,name,alamat_pengiriman,created_at FROM buyer WHERE email= '"+ req.Email + "' and password = '"+ req.Password +"'"
	row := repo.sqlDb.QueryRowContext(ctx, sqlStatement)
	err = row.Scan(&res.Email, &res.Name, &res.AlamatPengiriman, &res.CreatedAt)
	
	return res, err
}


