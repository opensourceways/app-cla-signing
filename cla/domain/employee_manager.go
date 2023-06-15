package domain

import "github.com/opensourceways/app-cla-signing/cla/domain/dp"

type Manager struct {
	Name    dp.Name
	Email   dp.EmailAddr
	Account dp.Account
}

func (m *Manager) IsSame(m1 *Manager) bool {
	return m.Email.EmailAddr() == m1.Email.EmailAddr() || m.Account.Account() == m1.Account.Account()
}

func (m *Manager) isMe(a dp.Account) bool {
	return m.Account.Account() == a.Account()
}
