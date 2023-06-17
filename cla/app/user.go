package app

import (
	"errors"

	"github.com/opensourceways/app-cla-signing/cla/domain"
	"github.com/opensourceways/app-cla-signing/cla/domain/repository"
	"github.com/opensourceways/app-cla-signing/cla/domain/userservice"
	commonRepo "github.com/opensourceways/app-cla-signing/common/domain/repository"
)

func NewUserService(
	repo repository.CorpSigning,
	us userservice.UserService,
) UserService {
	return &userService{
		repo: repo,
		us:   us,
	}
}

type UserService interface {
	Login(cmd *CmdToLogin) (dto UserLoginDTO, err error)
	ChangePassword(cmd *CmdToChangePassword) error
}

type userService struct {
	repo repository.CorpSigning
	us   userservice.UserService
}

func (s *userService) Login(cmd *CmdToLogin) (dto UserLoginDTO, err error) {
	if err = cmd.checkPassword(s.us); err != nil {
		return
	}

	u, err := s.find(cmd)
	if err != nil {
		if commonRepo.IsErrorResourceNotFound(err) {
			err = errorWrongUserOrPassword
		}

		return
	}

	cs, err := s.repo.Find(u.CorpSigningId)
	if err != nil {
		return
	}

	dto.Role = cs.GetRole(u.Account)
	if dto.Role == "" {
		err = errors.New("system error")

		return
	}

	dto.Account = u.Account.Account()
	dto.CorpSigningId = u.CorpSigningId
	dto.InitialPWChanged = u.PasswordChaged

	return
}

func (s *userService) ChangePassword(cmd *CmdToChangePassword) error {
	if err := cmd.checkPassword(s.us); err != nil {
		return err
	}

	u, err := s.us.FindByAccount(cmd.Account, cmd.OldOne)
	if err != nil {
		if commonRepo.IsErrorResourceNotFound(err) {
			err = errorInvalidPassword
		}

		return err
	}

	return s.us.ChangePassword(&u, cmd.NewOne)
}

func (s *userService) find(cmd *CmdToLogin) (domain.User, error) {
	if cmd.Account != nil {
		return s.us.FindByAccount(cmd.Account, cmd.Password)
	}

	return s.us.FindByEmail(cmd.Email, cmd.Password)
}
