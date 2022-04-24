package config

import (
	"context"
	"database/sql"
	"order-service/exception"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySqlDatabase(configuration Config) *sql.DB {
	db, err := sql.Open("mysql", configuration.Get("USER")+":"+configuration.Get("PASS")+"@tcp(localhost:3306)/"+configuration.Get("DATABASE")+"?parseTime=true")
	exception.PanicIfNeeded(err)

	return db
}

func NewMySqlContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
