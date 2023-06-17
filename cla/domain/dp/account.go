package dp

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/opensourceways/app-cla-signing/utils"
)

var reAccount = regexp.MustCompile("^[a-zA-Z0-9_.-]+_[a-zA-Z0-9-]+(\\.[a-zA-Z0-9-]+)*\\.[a-zA-Z]{2,6}$")

func NewAccount(v string) (Account, error) {
	err := errors.New("invalid account")

	if utils.StrLen(v) > config.MaxLengthOfAccount {
		return nil, err
	}

	if v == "" || !reAccount.MatchString(v) {
		return nil, err
	}

	return account(v), nil
}

func NewManagerAccount(account string, email EmailAddr) (Account, error) {
	return NewAccount(fmt.Sprintf("%s_%s", account, email.Domain()))
}

// Account
type Account interface {
	Account() string
}

type account string

func (r account) Account() string {
	return string(r)
}
