package manager

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Infra interface {
	SqlDb() *sqlx.DB
}

type infra struct {
	db *sqlx.DB
}

func (i infra) SqlDb() *sqlx.DB {
	return i.db
}

func NewInfra(datasourcename string) Infra {
	conn, err := sqlx.Connect("mysql", datasourcename)
	if err != nil {
		panic(err)
	}
	conn.SetMaxOpenConns(10)
	return infra{
		db: conn,
	}
}
