package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/DilanHera/mockTP/internal/app"
	"github.com/DilanHera/mockTP/internal/router"
	"github.com/DilanHera/mockTP/internal/tui"
)

func main() {
	app := app.NewApp()
	r := router.SetupRouter(app)

	fmt.Fprintln(os.Stderr, "HTTP server listening on :3000")

	go func() {
		if err := http.ListenAndServe(":3000", r); err != nil {
			fmt.Fprintln(os.Stderr, "HTTP server error:", err)
			os.Exit(1)
		}
	}()

	if err := tui.Run(app); err != nil {
		fmt.Fprintln(os.Stderr, "TUI error:", err)
		os.Exit(1)
	}
}
