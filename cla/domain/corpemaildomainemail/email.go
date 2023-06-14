package corpemaildomainemail

import (
	"github.com/opensourceways/app-cla-signing/cla/domain"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
	"github.com/opensourceways/app-cla-signing/cla/domain/emaildelivery"
)

type EmailBodyParameter struct {
	Code      string
	CorpName  dp.CorpName
	EmailAddr dp.EmailAddr
	Community domain.Community
}

type emailBodyBuilder interface {
	Build(*EmailBodyParameter) (emaildelivery.EmailBody, error)
}

func NewCorpEmailDomainEmail(builder emailBodyBuilder) *CorpEmailDomainEmail {
	return &CorpEmailDomainEmail{builder: builder}
}

// CorpEmailDomainEmail
type CorpEmailDomainEmail struct {
	builder emailBodyBuilder
}

func (e *CorpEmailDomainEmail) Builder(p *EmailBodyParameter) corpEmailDomainEmailAdapter {
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
