package app

import (
	"github.com/opensourceways/app-cla-signing/cla/domain"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
)

type CmdToAddEmailDomain struct {
	EmailAddr        dp.EmailAddr
	CorpSigningId    string
	VerificationCode string
}

func (cmd *CmdToAddEmailDomain) toVerificationCodeKey() domain.VerificationCodeKey {
	return domain.VerificationCodeKey{
		Code:    cmd.VerificationCode,
		Purpose: dp.NewPurposeOfAddingEmailDomain(cmd.EmailAddr),
	}
}
