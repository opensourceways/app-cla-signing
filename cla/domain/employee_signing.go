package domain

import "github.com/opensourceways/app-cla-signing/cla/domain/dp"

type SigningIndex struct {
	LinkId    string
	SigningId string
}

type IndividualSigning struct {
	Id    string
	Date  int64
	Link  Link
	Name  dp.Name
	Email dp.EmailAddr
}

type EmployeeSigning struct {
	IndividualSigning

	Enabled       bool
	CorpSigningId string
}

func (es *EmployeeSigning) ChangeEnabled(b bool) bool {
	if es.Enabled == b {
		return false
	}

	es.Enabled = b

	return true
}

func (es *EmployeeSigning) Remove() error {
	if es.Enabled {
		return errorRemoveEnabledEmployee
	}

	return nil
}
