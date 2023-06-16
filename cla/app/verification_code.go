package app

import (
	"github.com/opensourceways/app-cla-signing/cla/domain/corpemaildomainemail"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
	"github.com/opensourceways/app-cla-signing/cla/domain/emailclient"
	"github.com/opensourceways/app-cla-signing/cla/domain/emaildelivery"
	"github.com/opensourceways/app-cla-signing/cla/domain/signingcodeemail"
	"github.com/opensourceways/app-cla-signing/cla/domain/vcservice"
)

type SigningCodeService interface {
	Create(cmd *CmdToCreateCodeForSigning) (err error)
}

func NewSigningCodeService(
	cli emailclient.EmailClient,
	builder signingcodeemail.SigningCodeEmail,
	delivery emaildelivery.EmailDeliveryService,
	vcService vcservice.VCService,
) SigningCodeService {
	return &signingCodeService{
		cli:       cli,
		builder:   builder,
		delivery:  delivery,
		vcService: vcService,
	}
}

// signingCodeService
type signingCodeService struct {
	cli       emailclient.EmailClient
	builder   signingcodeemail.SigningCodeEmail
	delivery  emaildelivery.EmailDeliveryService
	vcService vcservice.VCService
}

func (s *signingCodeService) Create(cmd *CmdToCreateCodeForSigning) (err error) {
	code, err := s.vcService.New(
		dp.NewPurposeOfSigning(cmd.LinkId, cmd.EmailAddr),
	)

	err = s.delivery.Deliver(
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
	builder corpemaildomainemail.CorpEmailDomainEmail,
	delivery emaildelivery.EmailDeliveryService,
	vcService vcservice.VCService,
) EmailDomainCodeService {
	return &emailDomainCodeService{
		cli:       cli,
		builder:   builder,
		delivery:  delivery,
		vcService: vcService,
	}
}

// emailDomainCodeService
type emailDomainCodeService struct {
	cli       emailclient.EmailClient
	builder   corpemaildomainemail.CorpEmailDomainEmail
	delivery  emaildelivery.EmailDeliveryService
	vcService vcservice.VCService
}

func (s *emailDomainCodeService) Create(cmd *CmdToCreateCodeForEmailDomain) (err error) {
	// get corp name and email addr by signing id

	code, err := s.vcService.New(
		dp.NewPurposeOfAddingEmailDomain(nil), // TODO
	)

	err = s.delivery.Deliver(
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
