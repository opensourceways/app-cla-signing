package corpemaildomainemail

import (
	"github.com/opensourceways/app-cla-signing/cla/domain"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
	"github.com/opensourceways/app-cla-signing/cla/domain/emaildelivery"
)

func NewCorpEmailDomainEmail(builder emailBodyBuilder) CorpEmailDomainEmail {
	return &corpEmailDomainEmail{builder: builder}
}

// EmailBodyParameter
type EmailBodyParameter struct {
	Code      string
	CorpName  dp.CorpName
	EmailAddr dp.EmailAddr
	Community domain.Community
}

// emailBodyBuilder
type emailBodyBuilder interface {
	Build(*EmailBodyParameter) (emaildelivery.EmailBody, error)
}

// CorpEmailDomainEmail
type CorpEmailDomainEmail interface {
	Builder(p *EmailBodyParameter) corpEmailDomainEmailAdapter
}

// corpEmailDomainEmail
type corpEmailDomainEmail struct {
	builder emailBodyBuilder
}

func (e *corpEmailDomainEmail) Builder(p *EmailBodyParameter) corpEmailDomainEmailAdapter {
	return corpEmailDomainEmailAdapter{
		builder: e.builder,
		p:       *p,
	}
}

// corpEmailDomainEmailAdapter
type corpEmailDomainEmailAdapter struct {
	builder emailBodyBuilder
	p       EmailBodyParameter
}

func (adapter corpEmailDomainEmailAdapter) Build() (emaildelivery.Email, error) {
	body, err := adapter.builder.Build(&adapter.p)
	if err != nil {
		return emaildelivery.Email{}, err
	}

	return emaildelivery.Email{
		To:        []dp.EmailAddr{adapter.p.EmailAddr},
		EmailBody: body,
	}, nil
}
