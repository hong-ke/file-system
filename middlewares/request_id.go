package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

const RequestIDHeader = "X-Request-ID"

// 全局中间件 返回全局请求ID
func RequestID(ctx *gin.Context) {

	requestID := ctx.Request.Header.Get(RequestIDHeader)
	if requestID == "" {
		u4 := uuid.NewV4()
		requestID = u4.String()
	}
	ctx.Set(RequestIDHeader, requestID)

	ctx.Writer.Header().Set(RequestIDHeader, requestID)
	ctx.Next()

}
