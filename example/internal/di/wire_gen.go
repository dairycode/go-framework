package di

import (
	"go-framework/example/internal/server/http"
	"go-framework/example/internal/service"
)

// InitApp todo 后续改成 wire 自动生成
func InitApp() (*App, func(), error) {
	serviceService, cf1, err := service.New()
	if err != nil {
		cf1()
	}
	engine, err := http.New(serviceService)
	if err != nil {
		cf1()
	}
	app, cf2, err := NewApp(serviceService, engine)
	if err != nil {
		cf2()
		cf1()
	}
	return app, func() {
		cf2()
		cf1()
	}, nil
}
