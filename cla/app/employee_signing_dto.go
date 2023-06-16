package app

import (
	"github.com/opensourceways/app-cla-signing/cla/domain"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
)

type IndividualSigningDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Date  string `json:"date"`
	Email string `json:"email"`
}

type EmployeeSigningDTO struct {
	IndividualSigningDTO

	Enabled bool `json:"enabled"`
}

// CmdToListEmployeeSigning
type CmdToListEmployeeSigning struct {
	Lang             dp.Language
	CorpSigningIndex domain.SigningIndex
}

// CmdToUpdateEmployeeSigning
type CmdToUpdateEmployeeSigning struct {
	CmdToRemoveEmployeeSigning

	Enabled bool
}

// CmdToRemoveEmployeeSigning
type CmdToRemoveEmployeeSigning struct {
	CorpSigningIndex  domain.SigningIndex
	EmployeeSigningId string
}

func (cmd *CmdToRemoveEmployeeSigning) employeeSigningIndex() domain.SigningIndex {
	return domain.SigningIndex{
		LinkId:    cmd.CorpSigningIndex.LinkId,
		SigningId: cmd.EmployeeSigningId,
	}
}
