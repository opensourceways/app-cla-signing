package repositoryimpl

import (
	"github.com/opensourceways/app-cla-signing/cla/domain"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
)

func NewUser(dao dao) *user {
	return &user{
		dao: dao,
	}
}

type user struct {
	dao dao
}

func (impl *user) Add(*domain.User) error {
	return nil
}

func (impl *user) Remove(dp.Account) error {
	return nil
}
