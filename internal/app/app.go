package app

import (
	internal "github.com/DilanHera/mockTP/internal"
)

type App struct {
	Helper internal.Helper
}

func NewApp() *App {
	return &App{
		Helper: internal.NewHelper(),
	}
}
