package di

import (
	"github.com/gin-gonic/gin"
	"go-framework/example/internal/service"
)

type App struct {
	svc  *service.Service
	http *gin.Engine
}

func NewApp(svc *service.Service, h *gin.Engine) (app *App, closeFunc func(), err error) {
	app = &App{
		svc:  svc,
		http: h,
	}
	closeFunc = func() {
		// todo
	}
	return
}
