package controller

import (
	"errors"

	"github.com/opensourceways/app-cla-signing/cla/app"
	"github.com/opensourceways/app-cla-signing/cla/domain"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
)

// reqToAddEmployeeManager
type reqToAddEmployeeManager struct {
	Managers []employeeManager `json:"managers" binding:"required"`
}

func (req *reqToAddEmployeeManager) toCmd(corpSigningId string) (
	cmd app.CmdToAddEmployeeManager, err error,
) {
	idMap := make(map[string]bool)
	emailMap := make(map[string]bool)
	for i := range req.Managers {
		item := &req.Managers[i]

		// idMap
		if idMap[item.ID] {
			err = errors.New("duplicate manager id")

			return
		}
		idMap[item.ID] = true

		// emailMap
		if emailMap[item.Email] {
			err = errors.New("duplicate manager email")

			return
		}
		emailMap[item.Email] = true
	}

	ms := make([]domain.Manager, len(req.Managers))
	for i := range req.Managers {
		if ms[i], err = req.Managers[i].toManager(); err != nil {
			return
		}
	}

	cmd.Managers = ms
	cmd.CorpSigningId = corpSigningId

	return
}

// employeeManager
type employeeManager struct {
	ID    string `json:"id"      binding:"required"`
	Email string `json:"email"   binding:"required"`
	Name  string `json:"name"    binding:"required"`
}

func (req *employeeManager) toManager() (m domain.Manager, err error) {
	if m.Name, err = dp.NewName(req.Name); err != nil {
		return
	}

	if m.Email, err = dp.NewEmailAddr(req.Email); err != nil {
		return
	}

	// must be after the email
	m.Account, err = dp.NewManagerAccount(req.ID, m.Email)

	return
}

// reqToRemoveEmployeeManager
type reqToRemoveEmployeeManager struct {
	Managers []employeeManagerAccount `json:"managers"  binding:"required"`
}

func (req *reqToRemoveEmployeeManager) toCmd(corpSigningId string) (
	cmd app.CmdToRemoveEmployeeManager, err error,
) {
	ms := make([]dp.Account, len(req.Managers))
	for i := range req.Managers {
		if ms[i], err = req.Managers[i].toAccount(); err != nil {
			return
		}
	}

	cmd.Managers = ms
	cmd.CorpSigningId = corpSigningId

	return
}

// employeeManagerAccount
type employeeManagerAccount struct {
	ID string `json:"id"  binding:"required"`
}

func (req *employeeManagerAccount) toAccount() (dp.Account, error) {
	return dp.NewAccount(req.ID)
}
