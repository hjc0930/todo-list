package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-list/pkg/utils"
	"todo-list/service"
	"todo-list/types"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserRegisterReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetUserSrv()

			resp, err := l.UserRegister(ctx.Request.Context(), &req)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			utils.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}
