package dp

import (
	"errors"
	"fmt"

	"github.com/opensourceways/app-cla-signing/utils"
)

type Purpose interface {
	Purpose() string
}

type purpose string

func (v purpose) Purpose() string {
	return string(v)
}

func NewPurpose(v string) (Purpose, error) {
	if v == "" {
		return nil, errors.New("invalid purpose")
	}

	return purpose(v), nil
}

func NewPurposeOfSigning(linkId string, email EmailAddr) Purpose {
	return purpose(utils.GenMD5([]byte(
		fmt.Sprintf("sign %s, %s", linkId, email.EmailAddr()),
	)))
}

func NewPurposeOfAddingEmailDomain(email EmailAddr) Purpose {
	return purpose(utils.GenMD5([]byte(
		fmt.Sprintf("add email domain: %s", email.EmailAddr()),
	)))
}
