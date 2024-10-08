// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/userapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/baseadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter"
)

// Injectors from wire.go:

func initUserApp() *userapp.Application {
	adapter := baseadapter.New()
	svcoreadapterAdapter := svcoreadapter.New()
	application := userapp.New(adapter, svcoreadapterAdapter)
	return application
}
