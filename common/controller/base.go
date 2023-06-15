package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendBadRequestBody(ctx *gin.Context, err error) {
	ctx.JSON(
		http.StatusBadRequest,
		newResponseCodeMsg(errorBadRequestBody, err.Error()),
	)
}

func SendBadRequestParam(ctx *gin.Context, err error) {
	ctx.JSON(
		http.StatusBadRequest,
		newResponseCodeMsg(errorBadRequestParam, err.Error()),
	)
}

func SendRespOfCreate(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, newResponseCodeMsg("", "success"))
}

func SendRespOfPut(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, newResponseCodeMsg("", "success"))
}

func SendRespOfDelete(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, newResponseCodeMsg("", "success"))
}

func SendRespOfGet(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, newResponseData(data))
}

func SendRespOfPost(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, newResponseData(data))
}

// errorWithCode
type errorWithCode interface {
	Code() string
}

// errorOfNotFound
type errorOfNotFound interface {
	NotFound()
}

// SendFailedResp
func SendFailedResp(ctx *gin.Context, err error) {
	code := ""
	if v, ok := err.(errorWithCode); ok {
		code = v.Code()
	}

	if code == "" {
		ctx.JSON(
			http.StatusInternalServerError,
			newResponseCodeMsg(errorSystemError, err.Error()),
		)

		return
	}

	status := http.StatusBadRequest
	if _, ok := err.(errorOfNotFound); ok {
		status = http.StatusNotFound
	}

	ctx.JSON(status, newResponseCodeMsg(code, err.Error()))
}
