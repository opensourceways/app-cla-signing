package domain

import "github.com/opensourceways/app-cla-signing/cla/domain/dp"

type User struct {
	Email          dp.EmailAddr
	Account        dp.Account
	Password       string
	CorpSigningId  string
	PasswordChaged bool
}

func (u *User) ChangePassword(p string) {
	u.Password = p
	u.PasswordChaged = true
}
