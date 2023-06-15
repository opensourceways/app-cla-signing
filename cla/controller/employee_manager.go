package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/opensourceways/app-cla-signing/cla/app"
	commonctl "github.com/opensourceways/app-cla-signing/common/controller"
)

func AddRouteForEmployeeManagerController(
	r *gin.RouterGroup,
	s app.EmployeeManagerService,
) {
	ctl := employeeManagerController{
		s: s,
	}

	r.POST("/v1/employee-manager", ctl.AddEmployeeManager)
	r.DELETE("/v1/employee-manager", ctl.RemoveEmployeeManager)
}

type employeeManagerController struct {
	s app.EmployeeManagerService
}

// AddEmployeeManager
// @Description add employee manager
// @Tags   EmployeeManager
// @Accept json
// @Param  param  body  reqToAddEmployeeManager  true  "body of adding employee manager"
// @Success 201 {object} commonctl.ResponseData
// @Failure 400 {object} commonctl.ResponseData
// @router / [post]
func (ctl employeeManagerController) AddEmployeeManager(ctx *gin.Context) {
	var req reqToAddEmployeeManager
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		commonctl.SendBadRequestBody(ctx, err)

		return
	}

	// TODO missing corp signing id
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

// RemoveEmployeeManager
// @Description remove employee manager
// @Tags   EmployeeManager
// @Accept json
// @Param  param  body  reqToRemoveEmployeeManager  true  "body of adding employee manager"
// @Success 204 {object} commonctl.ResponseData
// @Failure 400 {object} commonctl.ResponseData
// @router / [delete]
func (ctl employeeManagerController) RemoveEmployeeManager(ctx *gin.Context) {
	var req reqToRemoveEmployeeManager
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		commonctl.SendBadRequestBody(ctx, err)

		return
	}

	// TODO missing corp signing id
	cmd, err := req.toCmd("")
	if err != nil {
		commonctl.SendBadRequestParam(ctx, err)

		return
	}

	if err := ctl.s.Remove(&cmd); err != nil {
		commonctl.SendFailedResp(ctx, err)
	} else {
		commonctl.SendRespOfDelete(ctx)
	}
}
