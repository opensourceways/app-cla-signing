package app

import "github.com/opensourceways/app-cla-signing/cla/domain/dp"

var (
	errorSamePassword        = dp.NewDomainError("same_password")
	errorInvalidPassword     = dp.NewDomainError("invalid_password")
	errorWrongUserOrPassword = dp.NewDomainError("wrong_user_or_password")
)
