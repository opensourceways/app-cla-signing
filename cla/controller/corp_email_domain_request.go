package controller

import (
	"github.com/opensourceways/app-cla-signing/cla/app"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
)

type reqToAddCorpEmailDomain struct {
	Email string `json:"sub_email"          binding:"required"`
	Code  string `json:"verification_code"  binding:"required"`
}

func (req *reqToAddCorpEmailDomain) toCmd(CorpSigningId string) (cmd app.CmdToAddEmailDomain, err error) {
	if cmd.EmailAddr, err = dp.NewEmailAddr(req.Email); err != nil {
		return
	}

	cmd.CorpSigningId = CorpSigningId
	cmd.VerificationCode = req.Code

	return
}
