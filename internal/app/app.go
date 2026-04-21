package app

import (
	internal "github.com/DilanHera/mockTP/internal"
	"github.com/DilanHera/mockTP/internal/store"
)

type App struct {
	Helper       internal.Helper
	AppInfoStore store.ApiInfoStore
}

func NewApp() *App {
	db, err := store.Open()
	if err != nil {
		panic(err)
	}
	return &App{
		Helper:       internal.NewHelper(),
		AppInfoStore: *store.NewApiInfoStore(db),
	}
}
