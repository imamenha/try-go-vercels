//go:build wireinject
// +build wireinject

package di

import (
	"try-go-vercels/app/handlers"
	"try-go-vercels/app/repositories"
	"try-go-vercels/app/services"

	"github.com/google/wire"
)

func InitializeHandler() *handlers.Handler {
	wire.Build(
		repositories.NewRepository,
		services.NewService,
		handlers.NewHandler,
	)
	return &handlers.Handler{}
}
