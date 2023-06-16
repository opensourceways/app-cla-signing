package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/opensourceways/app-cla-signing/cla/app"
	commonctl "github.com/opensourceways/app-cla-signing/common/controller"
)

func AddRouteForEmployeeSigningController(
	r *gin.RouterGroup,
	s app.EmployeeSigningService,
) {
	ctl := employeeSigningController{
		s: s,
	}

	r.GET("/v1/employee-signing", ctl.ListEmployeeSigning)
	r.PUT("/v1/employee-signing/:signing_id", ctl.UpdateEmployeeSigning)
	r.DELETE("/v1/employee-signing/:signing_id", ctl.RemoveEmployeeSigning)
}

type employeeSigningController struct {
	s app.EmployeeSigningService
}

// RemoveEmployeeSigning
// @Description remove employee signing
// @Tags   EmployeeSigning
// @Accept json
// @Param  signing_id  path  string  true  "employee signing id"
// @Success 204 {object} commonctl.ResponseData
// @Failure 400 {object} commonctl.ResponseData
// @router /{signing_id} [put]
func (ctl employeeSigningController) RemoveEmployeeSigning(ctx *gin.Context) {
	// TODO mising index
	cmd := app.CmdToRemoveEmployeeSigning{
		EmployeeSigningId: ctx.Param("signing_id"),
	}

	if err := ctl.s.Remove(cmd); err != nil {
		commonctl.SendFailedResp(ctx, err)
	} else {
		commonctl.SendRespOfDelete(ctx)
	}
}

// UpdateEmployeeSigning
// @Description enable/unable employee signing
// @Tags   EmployeeSigning
// @Accept json
// @Param  signing_id  path  string                      true  "employee signing id"
// @Param  param       body  reqToUpdateEmployeeSigning  true  "body of updating employee signing"
// @Success 202 {object} commonctl.ResponseData
// @Failure 400 {object} commonctl.ResponseData
// @router /{signing_id} [put]
func (ctl employeeSigningController) UpdateEmployeeSigning(ctx *gin.Context) {
	var req reqToUpdateEmployeeSigning
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		commonctl.SendBadRequestBody(ctx, err)

		return
	}

	// TODO mising index
	cmd := app.CmdToUpdateEmployeeSigning{}
	cmd.EmployeeSigningId = ctx.Param("signing_id")
	req.toCmd(&cmd)

	if err := ctl.s.Update(cmd); err != nil {
		commonctl.SendFailedResp(ctx, err)
	} else {
		commonctl.SendRespOfPut(ctx)
	}
}

// ListEmployeeSigning
// @Description list employee signing
// @Tags   EmployeeSigning
// @Accept json
// @Param  cla_language  query  string  false  "query by language"
// @Success 200 {object} app.EmployeeSigningDTO
// @Failure 400 {object} commonctl.ResponseData
// @router / [get]
func (ctl employeeSigningController) ListEmployeeSigning(ctx *gin.Context) {
	var req reqToListEmployeeSigning
	if err := ctx.ShouldBindQuery(&req); err != nil {
		commonctl.SendBadRequestParam(ctx, err)

		return
	}

	// TODO missing index
	cmd := app.CmdToListEmployeeSigning{}
	if err := req.toCmd(&cmd); err != nil {
		commonctl.SendBadRequestParam(ctx, err)

		return
	}

	if v, err := ctl.s.List(cmd); err != nil {
		commonctl.SendFailedResp(ctx, err)
	} else {
		commonctl.SendRespOfGet(ctx, v)
	}
}
