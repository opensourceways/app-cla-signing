package emaildelivery

import "github.com/opensourceways/app-cla-signing/cla/domain/dp"

type Email struct {
	To []dp.EmailAddr
	EmailBody
}

type EmailBody struct {
	Subject    string
	Content    string
	Attachment string
	MIME       string
}

// EmailBuilder
type EmailBuilder interface {
	Build() (Email, error)
}

// EmailDelivery
type EmailDelivery interface {
	Send(*Email) error
}

// EmailDeliveryBuilder
type EmailDeliveryBuilder interface {
	Build() (EmailDelivery, error)
}

// EmailDeliveryService
type EmailDeliveryService interface {
	Deliver(EmailBuilder, EmailDeliveryBuilder) error
}
