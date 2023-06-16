package domain

import (
	"strings"

	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
)

type Corporation struct {
	Name               dp.CorpName
	EmailDomains       []string
	PrimaryEmailDomain string
}

func (c *Corporation) isMyEmail(email dp.EmailAddr) bool {
	domain := email.Domain()

	for _, v := range c.EmailDomains {
		if v == domain {
			return true
		}
	}

	return false
}

func (c *Corporation) addEmailDomain(ed string) error {
	for _, v := range c.EmailDomains {
		if v == ed {
			return errorEmailDomainExists
		}
	}

	if err := c.isValidEmailDomain(ed); err != nil {
		return err
	}

	c.EmailDomains = append(c.EmailDomains, ed)

	return nil
}

func (c *Corporation) isValidEmailDomain(ed string) error {
	e1 := strings.Split(c.PrimaryEmailDomain, ".")
	e2 := strings.Split(ed, ".")

	n1 := len(e1) - 1
	j := len(e2) - 1
	i := n1
	for ; i >= 0; i-- {
		if j < 0 {
			break
		}

		if e1[i] != e2[j] {
			break
		}

		j--
	}

	if i < 0 || n1-i >= config.MinNumOfSameEmailDomainParts {
		return nil
	}

	return errorUnmatchedEmailDomain
}
