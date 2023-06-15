package domain

import "github.com/opensourceways/app-cla-signing/cla/domain/dp"

type Corporation struct {
	Name         dp.CorpName
	EmailDomains []string
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
