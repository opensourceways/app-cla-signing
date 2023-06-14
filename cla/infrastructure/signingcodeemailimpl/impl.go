package signingcodeemailimpl

import (
	"github.com/opensourceways/app-cla-signing/cla/domain/emaildelivery"
	"github.com/opensourceways/app-cla-signing/cla/domain/signingcodeemail"
)

func NewSigningCodeEmailImpl() *signingCodeEmailImpl {
	return &signingCodeEmailImpl{}
}

type signingCodeEmailImpl struct{}

func (impl *signingCodeEmailImpl) Build(*signingcodeemail.EmailBodyParameter) (
	emaildelivery.EmailBody, error,
) {
	return emaildelivery.EmailBody{}, nil
}
