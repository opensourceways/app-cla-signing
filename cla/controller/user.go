package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/opensourceways/app-cla-signing/cla/app"
	"github.com/opensourceways/app-cla-signing/cla/domain/token"
	commonctl "github.com/opensourceways/app-cla-signing/common/controller"
	"github.com/opensourceways/app-cla-signing/utils"
)

func AddRouteForUserController(
	r *gin.RouterGroup,
	s app.UserService,
	t token.Token,
) {
	ctl := userController{
		s: s,
		t: t,
	}

	r.POST("/v1/user", ctl.Login)
	r.PUT("/v1/user", ctl.ChangePassword)
}

type userController struct {
	s app.UserService
	t token.Token
}

// Login
// @Description login
// @Tags   User
// @Accept json
// @Param  param    body  reqToLogin  true  "body of login"
// @Success 201 {object} controller.loginResp
// @Failure 400 {object} commonctl.ResponseData
// @router / [post]
func (ctl userController) Login(ctx *gin.Context) {
	var req reqToLogin
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		commonctl.SendBadRequestBody(ctx, err)

		return
	}

	cmd, err := req.toCmd()
	if err != nil {
		commonctl.SendBadRequestBody(ctx, err)

		return
	}

	dto, err := ctl.s.Login(&cmd)
	if err != nil {
		commonctl.SendFailedResp(ctx, err)

		return
	}

	t, err := ctl.t.New(tokenPayload{
		Role:          dto.Role,
		LinkId:        req.LinkId,
		Account:       dto.Account,
		CheckTime:     utils.Now(),
		CorpSigningId: dto.CorpSigningId,
	})
	if err != nil {
		commonctl.SendFailedResp(ctx, err)

		return
	}

	resp := loginResp{
		Role:             dto.Role,
		Token:            t,
		InitialPWChanged: dto.InitialPWChanged,
	}

	commonctl.SendRespOfPost(ctx, resp)
}

// ChangePassword
// @Description change password
// @Tags   User
// @Accept json
// @Param  param    body  reqToChangePassword  true  "body of changing password"
// @Success 202 {object} commonctl.ResponseData
// @Failure 400 {object} commonctl.ResponseData
// @router / [put]
func (ctl userController) ChangePassword(ctx *gin.Context) {
	var req reqToChangePassword
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		commonctl.SendBadRequestBody(ctx, err)

		return
	}

	// TODO missing account
	cmd, err := req.toCmd(nil)
	if err != nil {
		commonctl.SendBadRequestParam(ctx, err)

		return
	}

	if err := ctl.s.ChangePassword(&cmd); err != nil {
		commonctl.SendFailedResp(ctx, err)
	} else {
		commonctl.SendRespOfUpdate(ctx)
	}
}
