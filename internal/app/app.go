package app

import (
	internal "github.com/DilanHera/mockTP/internal"
	"github.com/DilanHera/mockTP/internal/store"
)

type App struct {
	Helper          internal.Helper
	CustomRespStore store.CustomResponseStore
}

func NewApp() *App {
	db, err := store.Open()
	if err != nil {
		panic(err)
	}
	return &App{
		Helper:          internal.NewHelper(),
		CustomRespStore: *store.NewCustomResponseStore(db),
	}
}
