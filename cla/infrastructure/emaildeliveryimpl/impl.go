package emaildeliveryimpl

import "github.com/opensourceways/app-cla-signing/cla/domain/emaildelivery"

func NewEmailDeliveryImpl() *emailDeliveryImpl {
	return &emailDeliveryImpl{}
}

type emailDeliveryImpl struct{}

func (impl *emailDeliveryImpl) Deliver(emaildelivery.EmailBuilder, emaildelivery.EmailDeliveryBuilder) error {
	return nil
}
