package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/opensourceways/app-cla-signing/cla/app"
	commonctl "github.com/opensourceways/app-cla-signing/common/controller"
)

func AddRouteForCorpEmailDomainController(
	r *gin.RouterGroup,
	s app.CorpEmailDomainService,
) {
	ctl := corpEmailDomainController{
		s: s,
	}

	r.POST("/v1/corporation-email-domain", ctl.AddCorpEmailDomain)
	r.GET("/v1/corporation-email-domain", ctl.ListCorpEmailDomain)
}

type corpEmailDomainController struct {
	s app.CorpEmailDomainService
}

// AddCorpEmailDomain
// @Description add corp email domain
// @Tags   CorpEmailDomain
// @Accept json
// @Param  param    body  reqToAddCorpEmailDomain  true  "body of adding corp email domain"
// @Success 201 {object} commonctl.ResponseData
// @Failure 400 {object} commonctl.ResponseData
// @router / [post]
func (ctl corpEmailDomainController) AddCorpEmailDomain(ctx *gin.Context) {
	var req reqToAddCorpEmailDomain
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		commonctl.SendBadRequestBody(ctx, err)

		return
	}

	// TODO
	cmd, err := req.toCmd("")
	if err != nil {
		commonctl.SendBadRequestParam(ctx, err)

		return
	}

	if err := ctl.s.Add(&cmd); err != nil {
		commonctl.SendFailedResp(ctx, err)
	} else {
		commonctl.SendRespOfCreate(ctx)
	}
}

// ListCorpEmailDomain
// @Description list corp email domains
// @Tags   CorpEmailDomain
// @Accept json
// @Success 201 {object} commonctl.ResponseData
// @Failure 400 {object} commonctl.ResponseData
// @router / [get]
func (ctl corpEmailDomainController) ListCorpEmailDomain(ctx *gin.Context) {
	if v, err := ctl.s.List(""); err != nil {
		commonctl.SendFailedResp(ctx, err)
	} else {
		commonctl.SendRespOfGet(ctx, v)
	}
}
