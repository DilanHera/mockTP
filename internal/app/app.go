package app

import (
	internal "github.com/DilanHera/mockTP/internal"
	"github.com/DilanHera/mockTP/internal/services"
	"github.com/DilanHera/mockTP/internal/store"
)

type App struct {
	Helper       internal.Helper
	ApiInfoStore store.ApiInfoStore
	Service      services.Service
}

func NewApp() *App {
	db, err := store.Open()
	apiInfoStore := store.NewApiInfoStore(db)
	if err != nil {
		panic(err)
	}
	return &App{
		Helper:       internal.NewHelper(),
		ApiInfoStore: *apiInfoStore,
		Service:      services.NewService(apiInfoStore),
	}
}
