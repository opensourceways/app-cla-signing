package controller

import (
	"github.com/opensourceways/app-cla-signing/cla/app"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
)

type reqToCreateCodeForSigning struct {
	Email string `json:"email" binding:"required"`
}

func (req *reqToCreateCodeForSigning) toCmd(linkId string) (cmd app.CmdToCreateCodeForSigning, err error) {
	if cmd.EmailAddr, err = dp.NewEmailAddr(req.Email); err != nil {
		return
	}

	cmd.LinkId = linkId

	return
}
