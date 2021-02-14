package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zzfly256/IpService/src/error"
	"github.com/zzfly256/IpService/src/idl"
	"net/http"
)

func FormatResponse() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		// 格式化输出
		response := idl.StandardResponse{
			Code:    0,
			Message: "ok",
		}

		if err, exists := context.Get("error"); exists == true {
			response.Code = err.(error.StandardError).Code
			response.Message = err.(error.StandardError).Message
		}

		response.Data = context.Value("data")

		context.JSON(http.StatusOK, response)
	}
}
