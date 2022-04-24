package usecase

import (
	"mohamadelabror.com/posapi/model"
	"mohamadelabror.com/posapi/repository"
)

type GetCashierDetailUseCase interface {
	GetDetail(cashierId string) (*model.Cashier, error)
}

type getCashierDetailUseCase struct {
	cashierRepo repository.CashierRepo
}

func (g *getCashierDetailUseCase) GetDetail(cashierId string) (*model.Cashier, error) {
	return g.cashierRepo.GetCashierDetail(cashierId)
}

func NewGetCashierUseCase(cashierRepo repository.CashierRepo) GetCashierDetailUseCase {
	return &getCashierDetailUseCase{
		cashierRepo: cashierRepo,
	}
}
