package repository

import (
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
	commonRepo "github.com/opensourceways/app-cla-signing/common/domain/repository"
)

var errorNotFound = dp.NewNotFoundDomainError("not_found")

func TryToConvertToNotFound(err error) error {
	if commonRepo.IsErrorResourceNotFound(err) {
		return errorNotFound
	}

	return err
}
