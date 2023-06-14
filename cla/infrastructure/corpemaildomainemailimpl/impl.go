package corpemaildomainemailimpl

import (
	"github.com/opensourceways/app-cla-signing/cla/domain/corpemaildomainemail"
	"github.com/opensourceways/app-cla-signing/cla/domain/emaildelivery"
)

func NewCorpEmailDomainEmailImpl() *corpEmailDomainEmailImpl {
	return &corpEmailDomainEmailImpl{}
}

type corpEmailDomainEmailImpl struct{}

func (impl *corpEmailDomainEmailImpl) Build(*corpemaildomainemail.EmailBodyParameter) (
	emaildelivery.EmailBody, error,
) {
	return emaildelivery.EmailBody{}, nil
}
