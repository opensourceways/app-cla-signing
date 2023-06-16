package repository

import "github.com/opensourceways/app-cla-signing/cla/domain"

type CorpSigning interface {
	Find(string) (domain.CorpSigning, error)
	SaveManagers(*domain.CorpSigning) error
	SaveEmailDomain(*domain.CorpSigning) error
}
