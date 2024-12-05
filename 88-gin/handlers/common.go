package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	// ctx.Writer.Write([]byte("pong"))
	// ctx.Writer.WriteHeader(http.StatusOK)
	//ctx.JSON(200, gin.H{"message": "pong"})
	ctx.String(http.StatusOK, "pong")
}
func Health(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}
