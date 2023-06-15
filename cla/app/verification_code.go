package app

import (
	"github.com/opensourceways/app-cla-signing/cla/domain"
	"github.com/opensourceways/app-cla-signing/cla/domain/corpemaildomainemail"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
	"github.com/opensourceways/app-cla-signing/cla/domain/emailclient"
	"github.com/opensourceways/app-cla-signing/cla/domain/emaildelivery"
	"github.com/opensourceways/app-cla-signing/cla/domain/randomcode"
	"github.com/opensourceways/app-cla-signing/cla/domain/repository"
	"github.com/opensourceways/app-cla-signing/cla/domain/signingcodeemail"
)

type SigningCodeService interface {
	Create(cmd *CmdToCreateCodeForSigning) (err error)
}

func NewSigningCodeService(
	cli emailclient.EmailClient,
	repo repository.VerificationCode,
	builder signingcodeemail.SigningCodeEmail,
	service emaildelivery.EmailDeliveryService,
	randomCode randomcode.RandomCode,
) SigningCodeService {
	return &signingCodeService{
		cli:        cli,
		repo:       repo,
		builder:    builder,
		service:    service,
		randomCode: randomCode,
	}
}

// signingCodeService
type signingCodeService struct {
	cli        emailclient.EmailClient
	repo       repository.VerificationCode
	builder    signingcodeemail.SigningCodeEmail
	service    emaildelivery.EmailDeliveryService
	randomCode randomcode.RandomCode
}

func (s *signingCodeService) Create(cmd *CmdToCreateCodeForSigning) (err error) {
	code, err := s.randomCode.New()
	if err != nil {
		return
	}

	vc := domain.NewVerificationCode(code, dp.NewPurposeOfSigning(
		cmd.LinkId, cmd.EmailAddr,
	))
	if err = s.repo.Add(&vc); err != nil {
		return
	}

	err = s.service.Deliver(
		s.builder.Builder(&signingcodeemail.EmailBodyParameter{
			Code:      code,
			EmailAddr: cmd.EmailAddr,
			// Community:  // TODO
		}),
		s.cli.Builder(cmd.LinkId),
	)

	return
}

type EmailDomainCodeService interface {
	Create(cmd *CmdToCreateCodeForEmailDomain) (err error)
}

func NewEmailDomainCodeService(
	cli emailclient.EmailClient,
	repo repository.VerificationCode,
	builder corpemaildomainemail.CorpEmailDomainEmail,
	service emaildelivery.EmailDeliveryService,
	randomCode randomcode.RandomCode,
) EmailDomainCodeService {
	return &emailDomainCodeService{
		cli:        cli,
		repo:       repo,
		builder:    builder,
		service:    service,
		randomCode: randomCode,
	}
}

// emailDomainCodeService
type emailDomainCodeService struct {
	cli        emailclient.EmailClient
	repo       repository.VerificationCode
	builder    corpemaildomainemail.CorpEmailDomainEmail
	service    emaildelivery.EmailDeliveryService
	randomCode randomcode.RandomCode
}

func (s *emailDomainCodeService) Create(cmd *CmdToCreateCodeForEmailDomain) (err error) {
	// get corp name and email addr by signing id

	code, err := s.randomCode.New()
	if err != nil {
		return
	}

	vc := domain.NewVerificationCode(
		code, dp.NewPurposeOfAddingEmailDomain(nil), // TODO
	)
	if err = s.repo.Add(&vc); err != nil {
		return
	}

	err = s.service.Deliver(
		s.builder.Builder(&corpemaildomainemail.EmailBodyParameter{
			Code:      code,
			CorpName:  nil, //TODO
			EmailAddr: nil, // TODO
			//Community: cmd.Community,
		}),
		s.cli.Builder(cmd.LinkId),
	)

	return
}
