package controller

import (
	"errors"

	"github.com/opensourceways/app-cla-signing/cla/app"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
)

// reqToLogin
type reqToLogin struct {
	User     string `json:"user"      binding:"required"`
	LinkId   string `json:"link_id"   binding:"required"`
	Password string `json:"password"  binding:"required"`
}

func (req *reqToLogin) toCmd() (cmd app.CmdToLogin, err error) {
	if cmd.Account, err = dp.NewAccount(req.User); err != nil {
		if cmd.Email, err = dp.NewEmailAddr(req.User); err != nil {
			err = errors.New("invalid username")

			return
		}
	}

	if cmd.Password, err = dp.NewPassword(req.Password); err != nil {
		return
	}

	cmd.LinkId = req.LinkId

	return
}

type loginResp struct {
	//models.OrgRepo

	Role             string `json:"role"`
	Token            string `json:"token"`
	InitialPWChanged bool   `json:"initial_pw_changed"`
}

// reqToChangePassword
type reqToChangePassword struct {
	OldOne string `json:"old_password"  binding:"required"`
	NewOne string `json:"new_password"  binding:"required"`
}

func (req *reqToChangePassword) toCmd(a dp.Account) (cmd app.CmdToChangePassword, err error) {
	if cmd.OldOne, err = dp.NewPassword(req.OldOne); err != nil {
		return
	}

	if cmd.NewOne, err = dp.NewPassword(req.NewOne); err != nil {
		return
	}

	cmd.Account = a

	return
}

// tokenPayload
type tokenPayload struct {
	Role          string `json:"role"`
	LinkId        string `json:"link_id"`
	Account       string `json:"account"`
	CheckTime     int64  `json:"check_time"`
	CorpSigningId string `json:"corp_signing_id"`
}
