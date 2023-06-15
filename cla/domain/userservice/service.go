package userservice

import (
	"github.com/sirupsen/logrus"

	"github.com/opensourceways/app-cla-signing/cla/domain"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
	"github.com/opensourceways/app-cla-signing/cla/domain/encryption"
	"github.com/opensourceways/app-cla-signing/cla/domain/repository"
	"github.com/opensourceways/app-cla-signing/cla/domain/userpassword"
	commonRepo "github.com/opensourceways/app-cla-signing/common/domain/repository"
)

var errorUserExists = dp.NewDomainError("user_exists")

func NewUserService(
	repo repository.User,
	encrypt encryption.Encryption,
	password userpassword.UserPassword,
) UserService {
	return &userService{
		repo:     repo,
		encrypt:  encrypt,
		password: password,
	}
}

type UserService interface {
	Add(csId string, managers []domain.Manager) (err error)
	Remove(accounts []dp.Account)
}

type userService struct {
	repo     repository.User
	encrypt  encryption.Encryption
	password userpassword.UserPassword
}

func (s *userService) Add(csId string, managers []domain.Manager) (err error) {
	r := make([]dp.Account, 0, len(managers))

	for i := range managers {
		item := &managers[i]

		if err = s.add(csId, item); err != nil {
			if commonRepo.IsErrorDuplicateCreating(err) {
				err = errorUserExists
			}

			break
		}

		r = append(r, item.Account)
	}

	if err != nil && len(r) > 0 {
		s.Remove(r)
	}

	return
}

func (s *userService) Remove(accounts []dp.Account) {
	for _, v := range accounts {
		if err := s.repo.Remove(v); err != nil {
			logrus.Errorf(
				"remove user failed, user: %s, err: %s",
				v.Account(), err.Error(),
			)
		}
	}
}

func (s *userService) add(csId string, manager *domain.Manager) error {
	p, err := s.password.New()
	if err != nil {
		return err
	}

	v, err := s.encrypt.Ecrypt(p)
	if err != nil {
		return err
	}

	return s.repo.Add(&domain.User{
		Email:         manager.Email,
		Account:       manager.Account,
		Password:      v,
		CorpSigningId: csId,
	})
}
