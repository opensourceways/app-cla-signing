package emailclient

import "github.com/opensourceways/app-cla-signing/cla/domain/emaildelivery"

func NewEmailClient(builder emailDeliveryBuilder) EmailClient {
	return &emailClient{
		builder: builder,
	}
}

// emailDeliveryBuilder
type emailDeliveryBuilder interface {
	Build(linkId string) (emaildelivery.EmailDelivery, error)
}

// EmailClient
type EmailClient interface {
	Builder(linkId string) emailDeliveryBuilderAdapter
}

// emailClient
type emailClient struct {
	builder emailDeliveryBuilder
}

func (cli *emailClient) Builder(linkId string) emailDeliveryBuilderAdapter {
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
