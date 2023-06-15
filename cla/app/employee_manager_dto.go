package app

import (
	"github.com/opensourceways/app-cla-signing/cla/domain"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
)

type CmdToAddEmployeeManager struct {
	CorpSigningId string
	Managers      []domain.Manager
}

func (cmd *CmdToAddEmployeeManager) toManagerAccounts() []dp.Account {
	r := make([]dp.Account, len(cmd.Managers))

	for i := range cmd.Managers {
		r[i] = cmd.Managers[i].Account
	}

	return r
}

type CmdToRemoveEmployeeManager struct {
	CorpSigningId string
	Managers      []dp.Account
}

type EmployeeManagerDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
