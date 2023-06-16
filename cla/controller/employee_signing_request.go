package controller

import (
	"github.com/opensourceways/app-cla-signing/cla/app"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
)

type reqToListEmployeeSigning struct {
	Language string `form:"cla_language"`
}

func (req *reqToListEmployeeSigning) toCmd(cmd *app.CmdToListEmployeeSigning) (err error) {
	if req.Language != "" {
		cmd.Lang, err = dp.NewLanguage(req.Language)
	}

	return
}

type reqToUpdateEmployeeSigning struct {
	Enabled bool `json:"enabled"`
}

func (req *reqToUpdateEmployeeSigning) toCmd(cmd *app.CmdToUpdateEmployeeSigning) {
	cmd.Enabled = req.Enabled
}
