package app

import (
	"github.com/opensourceways/app-cla-signing/cla/domain"
	"github.com/opensourceways/app-cla-signing/cla/domain/repository"
	"github.com/opensourceways/app-cla-signing/utils"
)

func NewEmployeeSigningService(
	repo repository.EmployeeSigning,
) EmployeeSigningService {
	return &employeeSigningService{
		repo: repo,
	}
}

type EmployeeSigningService interface {
	List(cmd CmdToListEmployeeSigning) ([]EmployeeSigningDTO, error)
	Update(cmd CmdToUpdateEmployeeSigning) error
	Remove(cmd CmdToRemoveEmployeeSigning) error
}

type employeeSigningService struct {
	repo repository.EmployeeSigning
}

func (s *employeeSigningService) List(cmd CmdToListEmployeeSigning) ([]EmployeeSigningDTO, error) {
	v, err := s.repo.FindAllOfCorp(cmd.CorpSigningIndex, cmd.Lang)
	if err != nil || len(v) == 0 {
		return nil, err
	}

	r := make([]EmployeeSigningDTO, len(v))
	for i := range v {
		r[i] = s.toEmployeeSigningDTO(&v[i])
	}

	return r, nil
}

func (s *employeeSigningService) Update(cmd CmdToUpdateEmployeeSigning) error {
	v, err := s.repo.Find(cmd.employeeSigningIndex())
	if err != nil {
		err = repository.TryToConvertToNotFound(err)

		return err
	}

	if !v.ChangeEnabled(cmd.Enabled) {
		return nil
	}

	if err := s.repo.SaveEnabled(&v); err != nil {
		return err
	}

	// send email

	return nil
}

func (s *employeeSigningService) Remove(cmd CmdToRemoveEmployeeSigning) error {
	index := cmd.employeeSigningIndex()

	v, err := s.repo.Find(index)
	if err != nil {
		err = repository.TryToConvertToNotFound(err)

		return err
	}

	if err := v.Remove(); err != nil {
		return err
	}

	if err := s.repo.Remove(index); err != nil {
		return err
	}

	// send email
	// maybe it need save the deleted the employee signing

	return nil
}

func (s *employeeSigningService) toEmployeeSigningDTO(v *domain.EmployeeSigning) EmployeeSigningDTO {
	dto := IndividualSigningDTO{
		ID:    v.Id,
		Name:  v.Name.Name(),
		Date:  utils.Date(v.Date),
		Email: v.Email.EmailAddr(),
	}
	return EmployeeSigningDTO{
		IndividualSigningDTO: dto,
		Enabled:              v.Enabled,
	}
}
