package app

import (
	"github.com/opensourceways/app-cla-signing/cla/domain/repository"
	"github.com/opensourceways/app-cla-signing/cla/domain/vcservice"
)

func NewCorpEmailDomainService(
	repo repository.CorpSigning,
	vcService vcservice.VCService,
) CorpEmailDomainService {
	return &corpEmailDomainService{
		repo:      repo,
		vcService: vcService,
	}
}

type CorpEmailDomainService interface {
	Add(cmd *CmdToAddEmailDomain) error
	List(corpSigningId string) ([]string, error)
}

type corpEmailDomainService struct {
	repo      repository.CorpSigning
	vcService vcservice.VCService
}

func (s *corpEmailDomainService) Add(cmd *CmdToAddEmailDomain) error {
	cs, err := s.repo.Find(cmd.CorpSigningId)
	if err != nil {
		return err
	}

	key := cmd.toVerificationCodeKey()
	if err := s.vcService.Verify(&key); err != nil {
		return err
	}

	if err := cs.AddEmailDomain(cmd.EmailAddr); err != nil {
		return err
	}

	return s.repo.SaveEmailDomain(&cs)
}

func (s *corpEmailDomainService) List(corpSigningId string) ([]string, error) {
	cs, err := s.repo.Find(corpSigningId)
	if err != nil {
		return nil, err
	}

	return cs.Corporation.EmailDomains, nil
}
