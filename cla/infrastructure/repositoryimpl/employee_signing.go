package repositoryimpl

import (
	"github.com/opensourceways/app-cla-signing/cla/domain"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
)

func NewEmployeeSigning(dao dao) *employeeSigning {
	return &employeeSigning{
		dao: dao,
	}
}

type employeeSigning struct {
	dao dao
}

func (impl *employeeSigning) SaveEnabled(*domain.EmployeeSigning) error {
	return nil
}

func (impl *employeeSigning) Remove(domain.SigningIndex) error {
	return nil
}

func (impl *employeeSigning) Find(domain.SigningIndex) (domain.EmployeeSigning, error) {
	return domain.EmployeeSigning{}, nil
}

func (impl *employeeSigning) FindAllOfCorp(domain.SigningIndex, dp.Language) ([]domain.EmployeeSigning, error) {
	return nil, nil
}
