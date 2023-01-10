package http

import (
	"github.com/gin-gonic/gin"
	"go-framework/example/internal/service"
	"net/http"
)

var svc *service.Service

func New(s *service.Service) (engine *gin.Engine, err error) {
	svc = s
	engine = gin.Default()
	initRouter(engine)
	err = engine.Run()
	return
}

func initRouter(e *gin.Engine) {
	// 健康检查
	e.GET("/ping", ping)
}

func ping(ctx *gin.Context) {
	if err := svc.Ping(ctx); err != nil {
		// todo 缺少日志
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
