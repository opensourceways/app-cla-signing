package dp

import (
	"errors"
	"regexp"

	"github.com/opensourceways/app-cla-signing/utils"
)

var reEmailAddr = regexp.MustCompile(`^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*(\.[a-zA-Z]{2,6})$`)

func NewEmailAddr(v string) (EmailAddr, error) {
	err := errors.New("invalid email address")

	if utils.StrLen(v) > config.MaxLengthOfEmail {
		return nil, err
	}

	if v == "" || !reEmailAddr.MatchString(v) {
		return nil, err
	}

	return emailAddr(v), nil
}

// EmailAddr
type EmailAddr interface {
	EmailAddr() string
	Domain() string
}

type emailAddr string

func (r emailAddr) EmailAddr() string {
	return string(r)
}

func (r emailAddr) Domain() string {
	return utils.EmailSuffix(string(r))
}
