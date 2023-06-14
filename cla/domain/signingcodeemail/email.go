package signingcodeemail

import (
	"github.com/opensourceways/app-cla-signing/cla/domain"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
	"github.com/opensourceways/app-cla-signing/cla/domain/emaildelivery"
)

type EmailBodyParameter struct {
	Code      string
	EmailAddr dp.EmailAddr
	Community domain.Community
}

type emailBodyBuilder interface {
	Build(*EmailBodyParameter) (emaildelivery.EmailBody, error)
}

func NewSigningCodeEmail(builder emailBodyBuilder) *SigningCodeEmail {
	return &SigningCodeEmail{builder: builder}
}

// SigningCodeEmail
type SigningCodeEmail struct {
	builder emailBodyBuilder
}

func (e *SigningCodeEmail) Builder(p *EmailBodyParameter) signingCodeEmailAdapter {
	return signingCodeEmailAdapter{
		builder: e.builder,
		p:       *p,
	}
}

// signingCodeEmailAdapter
type signingCodeEmailAdapter struct {
	builder emailBodyBuilder
	p       EmailBodyParameter
}

func (adapter signingCodeEmailAdapter) Build() (emaildelivery.Email, error) {
	body, err := adapter.builder.Build(&adapter.p)
	if err != nil {
		return emaildelivery.Email{}, err
	}

	return emaildelivery.Email{
		To:        []dp.EmailAddr{adapter.p.EmailAddr},
		EmailBody: body,
	}, nil
}
