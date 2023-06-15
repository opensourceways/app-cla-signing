package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/opensourceways/app-cla-signing/cla/app"
	commonctl "github.com/opensourceways/app-cla-signing/common/controller"
)

func AddRouteForVerificationCodeController(
	r *gin.RouterGroup,
	s app.SigningCodeService,
	es app.EmailDomainCodeService,
) {
	ctl := verificationCodeController{
		signingCodeService:     s,
		emailDomainCodeService: es,
	}

	r.POST("/v1/verification-code/:link_id", ctl.NewCodeForSigning)
	r.POST("/v1/verification-code", ctl.NewCodeForAddingEmailDomain)
}

type verificationCodeController struct {
	signingCodeService     app.SigningCodeService
	emailDomainCodeService app.EmailDomainCodeService
}

// NewCodeForSigning
// @Description apply a new verification code for signing
// @Tags   VerificationCode
// @Accept json
// @Param  link_id  path  string                     true  "link id"
// @Param  param    body  reqToCreateCodeForSigning  true  "body of applying a new verification code"
// @Success 201 {object} commonctl.ResponseData
// @Failure 400 {object} commonctl.ResponseData
// @router /{link_id} [post]
func (ctl verificationCodeController) NewCodeForSigning(ctx *gin.Context) {
	var req reqToCreateCodeForSigning
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		commonctl.SendBadRequestBody(ctx, err)

		return
	}

	cmd, err := req.toCmd(ctx.Param("link_id"))
	if err != nil {
		commonctl.SendBadRequestParam(ctx, err)

		return
	}

	if err := ctl.signingCodeService.Create(&cmd); err != nil {
		commonctl.SendFailedResp(ctx, err)
	} else {
		commonctl.SendRespOfCreate(ctx)
	}
}

// NewCodeForAddingEmailDomain
// @Description apply a new verification code for adding email domain
// @Tags   VerificationCode
// @Accept json
// @Success 201 {object} commonctl.ResponseData
// @Failure 400 {object} commonctl.ResponseData
// @router / [post]
func (ctl verificationCodeController) NewCodeForAddingEmailDomain(ctx *gin.Context) {
	cmd := app.CmdToCreateCodeForEmailDomain{} // TODO build the cmd

	if err := ctl.emailDomainCodeService.Create(&cmd); err != nil {
		commonctl.SendFailedResp(ctx, err)
	} else {
		commonctl.SendRespOfCreate(ctx)
	}
}
