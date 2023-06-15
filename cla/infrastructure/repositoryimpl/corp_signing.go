package repositoryimpl

import (
	"github.com/opensourceways/app-cla-signing/cla/domain"
)

func NewCorpSigning(dao dao) *corpSigning {
	return &corpSigning{
		dao: dao,
	}
}

type corpSigning struct {
	dao dao
}

func (impl *corpSigning) Find(string) (domain.CorpSigning, error) {
	return domain.CorpSigning{}, nil
}

func (impl *corpSigning) SaveManagers(*domain.CorpSigning) error {
	return nil
}
