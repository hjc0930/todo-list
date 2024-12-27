package api

import (
	"net/http"
	"todo-list/pkg/utils"
	"todo-list/service"
	"todo-list/types"

	"github.com/gin-gonic/gin"
)

func CreateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CreateTaskReq

		if err := ctx.ShouldBind(&req); err == nil {
			l := service.GetTaskSrv()

			resp, err := l.CreateTask(ctx.Request.Context(), &req)
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
