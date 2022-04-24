package usecase

import (
	"mohamadelabror.com/posapi/model"
	"mohamadelabror.com/posapi/repository"
)

type CreateCashierUseCase interface {
	Insert(cashierName, cashierPass string) (*model.Cashier, error)
}

type createCashierUseCase struct {
	cashierRepo repository.CashierRepo
}

func (c *createCashierUseCase) Insert(cashierName, cashierPass string) (*model.Cashier, error) {
	return c.cashierRepo.CreateCashier(cashierName, cashierPass)
}

func NewCreateCashierUseCase(cashierRepo repository.CashierRepo) CreateCashierUseCase {
	return &createCashierUseCase{
		cashierRepo: cashierRepo,
	}
}
