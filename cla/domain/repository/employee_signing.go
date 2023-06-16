package repository

import (
	"github.com/opensourceways/app-cla-signing/cla/domain"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
)

type EmployeeSigning interface {
	SaveEnabled(*domain.EmployeeSigning) error
	Remove(domain.SigningIndex) error
	Find(domain.SigningIndex) (domain.EmployeeSigning, error)
	FindAllOfCorp(domain.SigningIndex, dp.Language) ([]domain.EmployeeSigning, error)
}
