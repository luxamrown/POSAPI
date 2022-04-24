package usecase

import "mohamadelabror.com/posapi/repository"

type DeleteCashierUseCase interface {
	DeleteCashier(cashierId string) error
}

type deleteCashierUseCase struct {
	cashierRepo repository.CashierRepo
}

func (d *deleteCashierUseCase) DeleteCashier(cashierId string) error {
	return d.cashierRepo.DeleteCashier(cashierId)
}

func NewDeleteCashierUseCase(cashierRepo repository.CashierRepo) DeleteCashierUseCase {
	return &deleteCashierUseCase{
		cashierRepo: cashierRepo,
	}
}
