package app

import (
	"github.com/opensourceways/app-cla-signing/cla/domain/repository"
	"github.com/opensourceways/app-cla-signing/cla/domain/userservice"
)

func NewEmployeeManagerService(
	repo repository.CorpSigning,
	userService userservice.UserService,
) EmployeeManagerService {
	return &employeeManagerService{
		repo:        repo,
		userService: userService,
	}
}

type EmployeeManagerService interface {
	Add(cmd *CmdToAddEmployeeManager) (err error)
	Remove(cmd *CmdToRemoveEmployeeManager) (err error)
}

type employeeManagerService struct {
	repo        repository.CorpSigning
	userService userservice.UserService
}

func (s *employeeManagerService) Add(cmd *CmdToAddEmployeeManager) (err error) {
	cs, err := s.repo.Find(cmd.CorpSigningId)
	if err != nil {
		return
	}

	if err = cs.AddManagers(cmd.Managers); err != nil {
		return
	}

	if err = s.userService.Add(cmd.CorpSigningId, cmd.Managers); err != nil {
		return
	}

	if err = s.repo.SaveManagers(&cs); err != nil {
		s.userService.Remove(cmd.toManagerAccounts())
	}

	// TODO send email

	return
}

func (s *employeeManagerService) Remove(cmd *CmdToRemoveEmployeeManager) (err error) {
	cs, err := s.repo.Find(cmd.CorpSigningId)
	if err != nil {
		return
	}

	if err = cs.RemoveManagers(cmd.Managers); err != nil {
		return
	}

	if err = s.repo.SaveManagers(&cs); err != nil {
		s.userService.Remove(cmd.Managers)
	}

	// TODO send email

	return
}
