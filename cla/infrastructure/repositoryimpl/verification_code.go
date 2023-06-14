package repositoryimpl

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/opensourceways/app-cla-signing/cla/domain"
	"github.com/opensourceways/app-cla-signing/common/domain/repository"
	"github.com/opensourceways/app-cla-signing/utils"
)

func NewVerificationCode(dao dao) *verificationCode {
	return &verificationCode{
		dao: dao,
	}
}

type verificationCode struct {
	dao dao
}

func (impl *verificationCode) Add(code *domain.VerificationCode) error {
	do := toVerificationCodeDO(code)
	body, err := do.toMap()
	if err != nil {
		return err
	}

	_, err = impl.dao.InsertDoc(body)

	_ = impl.dao.DeleteAll(bson.M{fieldExpiry: bson.M{"$lt": utils.Now()}})

	return err
}

func (impl *verificationCode) Find(key *domain.VerificationCodeKey) (domain.VerificationCode, error) {
	var do verificationCodeDO

	err := impl.dao.FindOneAndDelete(
		bson.M{
			fieldCode:    key.Code,
			fieldPurpose: key.Purpose,
		}, &do,
	)
	if err != nil {
		if impl.dao.IsDocNotExists(err) {
			err = repository.NewErrorResourceNotFound(err)
		}

		return domain.VerificationCode{}, err
	}

	return do.toVerificationCode()
}
