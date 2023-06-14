package emailclient

import "github.com/opensourceways/app-cla-signing/cla/domain/emaildelivery"

type emailDeliveryBuilder interface {
	Build(linkId string) (emaildelivery.EmailDelivery, error)
}

func NewEmailClient(builder emailDeliveryBuilder) *EmailClient {
	return &EmailClient{
		builder: builder,
	}
}

// EmailClient
type EmailClient struct {
	builder emailDeliveryBuilder
}

func (cli *EmailClient) Builder(linkId string) emailDeliveryBuilderAdapter {
	return emailDeliveryBuilderAdapter{builder: cli.builder, linkId: linkId}
}

// emailDeliveryBuilderAdapter
type emailDeliveryBuilderAdapter struct {
	builder emailDeliveryBuilder
	linkId  string
}

func (b emailDeliveryBuilderAdapter) Build() (emaildelivery.EmailDelivery, error) {
	return b.builder.Build(b.linkId)
}
