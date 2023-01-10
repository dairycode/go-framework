//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"go-framework/example/internal/server/http"
	"go-framework/example/internal/service"
)

//go:generate wire gen
func InitApp() (*App, func(), error) {
	panic(wire.Build(
		// todo dao
		// service
		service.Provider,
		// server
		http.New,
		// app
		NewApp,
	))
}
