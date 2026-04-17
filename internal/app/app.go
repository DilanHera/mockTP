package app

import (
	internal "github.com/DilanHera/mockTP/internal"
)

var responseState string = "SUCCESS"

type App struct {
	Helper        internal.Helper
	ResponseState string
}

func NewApp() *App {
	return &App{
		Helper:        internal.NewHelper(),
		ResponseState: responseState,
	}
}
