// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"crypto-finance/src/domain/usecase"
	"crypto-finance/src/infra"
	"crypto-finance/src/infra/http/impl"
	"crypto-finance/src/presentation/controller"
	"crypto-finance/src/presentation/route"
)

// Injectors from wire.go:

func InitializeApp() error {
	httpRepository := impl.NewHttpRouter()
	firebaseRepository := infra.NewFirebaseRepository()
	authUseCase := usecase.NewAuthUseCase(firebaseRepository)
	authController := controller.NewAuthController(authUseCase)
	statusController := controller.NewStatusController()
	error2 := route.NewRoutes(httpRepository, authController, statusController)
	return error2
}
