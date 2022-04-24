package manager

import "mohamadelabror.com/posapi/repository"

type RepoManager interface {
	CashierRepo() repository.CashierRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) CashierRepo() repository.CashierRepo {
	return repository.NewCashierRepo(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
