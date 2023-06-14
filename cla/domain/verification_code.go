package domain

import (
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
	"github.com/opensourceways/app-cla-signing/utils"
)

func NewVerificationCode(code string, Purpose dp.Purpose) VerificationCode {
	return VerificationCode{
		Expiry: utils.Now() + config.VerificationCodeExpiry,
		VerificationCodeKey: VerificationCodeKey{
			Code:    code,
			Purpose: Purpose,
		},
	}
}

type VerificationCodeKey struct {
	Code    string
	Purpose dp.Purpose
}

type VerificationCode struct {
	VerificationCodeKey

	Expiry int64
}

func (vc *VerificationCode) IsExpired() bool {
	return vc.Expiry < utils.Now()
}
