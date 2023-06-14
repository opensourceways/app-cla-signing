package repository

import "github.com/opensourceways/app-cla-signing/cla/domain"

type VerificationCode interface {
	Add(*domain.VerificationCode) error
	Find(*domain.VerificationCodeKey) (domain.VerificationCode, error)
}
