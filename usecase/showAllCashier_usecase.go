package usecase

import (
	"mohamadelabror.com/posapi/model"
	"mohamadelabror.com/posapi/repository"
)

type ShowAllCashierUseCase interface {
	ShowAll(limit, skip string) (*[]model.Cashier, error)
}

type showAllCashierUseCase struct {
	cashierRepo repository.CashierRepo
}

func (s *showAllCashierUseCase) ShowAll(limit, skip string) (*[]model.Cashier, error) {
	return s.cashierRepo.GetListCashier(limit, skip)
}

func NewShowAllUseCase(cashierRepo repository.CashierRepo) ShowAllCashierUseCase {
	return &showAllCashierUseCase{
		cashierRepo: cashierRepo,
	}
}
