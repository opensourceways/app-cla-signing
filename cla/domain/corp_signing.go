package domain

import "github.com/opensourceways/app-cla-signing/cla/domain/dp"

type Representative struct {
	Name  dp.Name
	Email dp.EmailAddr
	Title dp.Title
}

type Link struct {
	Id       string
	CLAId    string
	Language dp.Language
}

type CorpSigning struct {
	Id             string
	PDF            string
	Date           int64
	Link           Link
	Admin          Manager
	Managers       []Manager
	Corporation    Corporation
	Representative Representative

	Version int
}

func (cs *CorpSigning) EmailDomains() []string {
	return cs.Corporation.EmailDomains
}

func (cs *CorpSigning) AddEmailDomain(email dp.EmailAddr) error {
	return cs.Corporation.addEmailDomain(email.Domain())
}

func (cs *CorpSigning) AddManagers(managers []Manager) error {
	if len(cs.Managers)+len(managers) > config.MaxNumOfEmployeeManager {
		return errorTooManyEmployeeManagers
	}

	for i := range managers {
		item := &managers[i]

		if !cs.IsSameCorp(item.Email) {
			return errorNotSameCorp
		}

		if cs.hasManager(item) {
			return errorEmployeeManagerExists
		}

		if cs.Admin.IsSame(item) {
			return errorErrAdminAsManager
		}
	}

	return nil
}

func (cs *CorpSigning) RemoveManagers(managers []dp.Account) error {
	toRemove := make(map[int]bool)

	for i := range managers {
		j, exists := cs.includeManager(managers[i])
		if !exists {
			return errorEmployeeManagerNotExists
		}
		toRemove[j] = true
	}

	if n := len(cs.Managers) - len(toRemove); n <= 0 {
		cs.Managers = nil
	} else {
		m := make([]Manager, 0, n)
		for i := range cs.Managers {
			if !toRemove[i] {
				m = append(m, cs.Managers[i])
			}
		}
		cs.Managers = m
	}

	return nil
}

func (cs *CorpSigning) IsSameCorp(email dp.EmailAddr) bool {
	return cs.Corporation.isMyEmail(email)
}

func (cs *CorpSigning) hasManager(m *Manager) bool {
	for j := range cs.Managers {
		if cs.Managers[j].IsSame(m) {
			return true
		}
	}

	return false
}

func (cs *CorpSigning) includeManager(m dp.Account) (int, bool) {
	for j := range cs.Managers {
		if cs.Managers[j].isMe(m) {
			return j, true
		}
	}

	return 0, false
}
