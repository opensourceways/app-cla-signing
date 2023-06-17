package app

import (
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
	"github.com/opensourceways/app-cla-signing/cla/domain/userservice"
)

// CmdToLogin
type CmdToLogin struct {
	LinkId   string
	Email    dp.EmailAddr
	Account  dp.Account
	Password dp.Password
}

func (cmd *CmdToLogin) checkPassword(us userservice.UserService) error {
	if !us.IsValidPassword(cmd.Password) {
		return errorWrongUserOrPassword
	}

	return nil
}

// UserLoginDTO
type UserLoginDTO struct {
	//models.OrgRepo
	Role             string
	Account          string
	CorpSigningId    string
	InitialPWChanged bool
}

// CmdToChangePassword
type CmdToChangePassword struct {
	Account dp.Account
	OldOne  dp.Password
	NewOne  dp.Password
}

func (cmd *CmdToChangePassword) checkPassword(us userservice.UserService) error {
	if cmd.OldOne.Password() == cmd.NewOne.Password() {
		return errorSamePassword
	}

	if !us.IsValidPassword(cmd.OldOne) || !us.IsValidPassword(cmd.NewOne) {
		return errorInvalidPassword
	}

	return nil
}
