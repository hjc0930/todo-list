package middleware

import (
	"todo-list/pkg/ctl"
	"todo-list/pkg/errorStatus"
	"todo-list/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := errorStatus.Success
		token := ctx.GetHeader("Authorization")

		if token == "" {
			code = errorStatus.Error
			ctx.JSON(errorStatus.InvalidParams, gin.H{
				"status": code,
				"msg":    errorStatus.GetMsg(code),
				"data":   "token is empty",
			})
			ctx.Abort()
			return
		}
		claims, err := utils.ParseToken(token)

		if err != nil {
			code = errorStatus.Error
			ctx.JSON(errorStatus.Error, gin.H{
				"status": code,
				"msg":    errorStatus.GetMsg(code),
				"data":   err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.Request = ctx.Request.WithContext(ctl.NewContext(ctx.Request.Context(), &ctl.UserInfo{
			Id:       claims.Id,
			UserName: claims.UserName,
		}))

		ctx.Next()
	}
}
