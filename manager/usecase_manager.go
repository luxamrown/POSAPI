package manager

import "mohamadelabror.com/posapi/usecase"

type UseCaseManager interface {
	ShowAllCashierUseCase() usecase.ShowAllCashierUseCase
	GetCashierDetail() usecase.GetCashierDetailUseCase
	CreateCashier() usecase.CreateCashierUseCase
	UpdateCashier() usecase.UpdateCashierUseCase
	DeleteCashier() usecase.DeleteCashierUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) ShowAllCashierUseCase() usecase.ShowAllCashierUseCase {
	return usecase.NewShowAllUseCase(u.repo.CashierRepo())
}

func (u *useCaseManager) GetCashierDetail() usecase.GetCashierDetailUseCase {
	return usecase.NewGetCashierUseCase(u.repo.CashierRepo())
}

func (u *useCaseManager) CreateCashier() usecase.CreateCashierUseCase {
	return usecase.NewCreateCashierUseCase(u.repo.CashierRepo())
}

func (u *useCaseManager) UpdateCashier() usecase.UpdateCashierUseCase {
	return usecase.NewUpdateCashierUseCase(u.repo.CashierRepo())
}

func (u *useCaseManager) DeleteCashier() usecase.DeleteCashierUseCase {
	return usecase.NewDeleteCashierUseCase(u.repo.CashierRepo())
}

func NewUseCaseManager(manager RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: manager,
	}
}
