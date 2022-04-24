package repository

import (
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"mohamadelabror.com/posapi/model"
)

type CashierRepo interface {
	GetListCashier(limit, skip string) (*[]model.Cashier, error)
	GetCashierDetail(cashierId string) (*model.Cashier, error)
	CreateCashier(cashierName, cashierPass string) (*model.Cashier, error)
	UpdateCashier(cashierId, cashierName, passcode string) error
	DeleteCashier(cashierId string) error
}

type CashierRepoImpl struct {
	cashierDb *sqlx.DB
}

func (c *CashierRepoImpl) GetListCashier(limit, skip string) (*[]model.Cashier, error) {
	var cashiers []model.Cashier
	intLimit, _ := strconv.Atoi(limit)
	intSkip, _ := strconv.Atoi(skip)
	// err := c.cashierDb.Select(&cashiers, fmt.Sprintf("SELECT cashierId, name FROM cashiers LIMIT %d OFFSET %d", intLimit, intSkip))
	err := c.cashierDb.Select(&cashiers, "SELECT cashierId, name FROM cashiers WHERE isDelete = 0 LIMIT ? OFFSET ?", intLimit, intSkip)
	if err != nil {
		return nil, err
	}
	return &cashiers, nil
}

func (c *CashierRepoImpl) GetCashierDetail(cashierId string) (*model.Cashier, error) {
	var cashier model.Cashier
	id, _ := strconv.Atoi(cashierId)
	err := c.cashierDb.Get(&cashier, "SELECT cashierId, name FROM cashiers WHERE cashierId = ?", id)
	if err != nil {
		return nil, err
	}
	return &cashier, nil
}

func (c *CashierRepoImpl) CreateCashier(cashierName, cashierPass string) (*model.Cashier, error) {
	var cashier model.Cashier
	time := time.Now().Format("2006-01-02 15:04:05")
	_, err := c.cashierDb.Exec("INSERT INTO cashiers(name, passcode) VALUES(?, ?)", cashierName, cashierPass)
	if err != nil {
		return nil, err
	}
	err = c.cashierDb.Get(&cashier, "SELECT * FROM cashiers WHERE createdAt = ?", time)
	if err != nil {
		return nil, err
	}
	return &cashier, nil
}

func (c *CashierRepoImpl) UpdateCashier(cashierId, cashierName, passcode string) error {
	_, err := c.cashierDb.Exec("UPDATE cashiers SET name = ?, passcode = ? WHERE cashierId = ?", cashierName, passcode, cashierId)
	if err != nil {
		return err
	}
	return nil
}

func (c *CashierRepoImpl) DeleteCashier(cashierId string) error {
	id, _ := strconv.Atoi(cashierId)
	_, err := c.cashierDb.Exec("UPDATE cashiers SET isDelete = 1 WHERE cashierId = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func NewCashierRepo(cashierDb *sqlx.DB) CashierRepo {
	cashierRepo := CashierRepoImpl{
		cashierDb: cashierDb,
	}
	return &cashierRepo
}
