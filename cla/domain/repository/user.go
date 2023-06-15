package repository

import (
	"github.com/opensourceways/app-cla-signing/cla/domain"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
)

type User interface {
	Add(*domain.User) error
	Remove(dp.Account) error
}
