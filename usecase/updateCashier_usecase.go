package usecase

import "mohamadelabror.com/posapi/repository"

type UpdateCashierUseCase interface {
	UpdateCashier(cashierId, name, passcode string) error
}

type updateCashierUseCase struct {
	cashierRepo repository.CashierRepo
}

func (u *updateCashierUseCase) UpdateCashier(cashierId, name, passcode string) error {
	return u.cashierRepo.UpdateCashier(cashierId, name, passcode)
}

func NewUpdateCashierUseCase(cashierRepo repository.CashierRepo) UpdateCashierUseCase {
	return &updateCashierUseCase{
		cashierRepo: cashierRepo,
	}
}
