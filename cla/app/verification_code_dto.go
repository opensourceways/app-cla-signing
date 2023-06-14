package app

import "github.com/opensourceways/app-cla-signing/cla/domain/dp"

type CmdToCreateCodeForSigning struct {
	LinkId    string
	EmailAddr dp.EmailAddr
}

type CmdToCreateCodeForEmailDomain struct {
	LinkId    string
	SigningId string
}
