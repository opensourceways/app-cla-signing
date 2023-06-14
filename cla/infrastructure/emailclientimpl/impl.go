package emailclientimpl

import "github.com/opensourceways/app-cla-signing/cla/domain/emaildelivery"

func NewEmailClientImpl() *emailClientImpl {
	return &emailClientImpl{}
}

type emailClientImpl struct{}

func (impl *emailClientImpl) Build(linkId string) (emaildelivery.EmailDelivery, error) {
	return nil, nil
}
