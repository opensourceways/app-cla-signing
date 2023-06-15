package signingcodeemail

import (
	"github.com/opensourceways/app-cla-signing/cla/domain"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
	"github.com/opensourceways/app-cla-signing/cla/domain/emaildelivery"
)

func NewSigningCodeEmail(builder emailBodyBuilder) SigningCodeEmail {
	return &signingCodeEmail{builder: builder}
}

// EmailBodyParameter
type EmailBodyParameter struct {
	Code      string
	EmailAddr dp.EmailAddr
	Community domain.Community
}

// emailBodyBuilder
type emailBodyBuilder interface {
	Build(*EmailBodyParameter) (emaildelivery.EmailBody, error)
}

// SigningCodeEmail
type SigningCodeEmail interface {
	Builder(p *EmailBodyParameter) signingCodeEmailAdapter
}

// signingCodeEmail
type signingCodeEmail struct {
	builder emailBodyBuilder
}

func (e *signingCodeEmail) Builder(p *EmailBodyParameter) signingCodeEmailAdapter {
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
