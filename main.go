package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/DilanHera/mockTP/internal/app"
	"github.com/DilanHera/mockTP/internal/router"
	"github.com/DilanHera/mockTP/internal/tui"
)

func main() {
	showVersion := flag.Bool("version", false, "show version")
	flag.Parse()
	if *showVersion {
		fmt.Println(getVersion())
		os.Exit(0)
	}
	app := app.NewApp()
	r := router.SetupRouter(app)

	fmt.Fprintln(os.Stderr, "HTTP server listening on :3001")

	go func() {
		if err := http.ListenAndServe(":3001", r); err != nil {
			fmt.Fprintln(os.Stderr, "HTTP server error:", err)
			os.Exit(1)
		}
	}()

	if err := tui.Run(app); err != nil {
		fmt.Fprintln(os.Stderr, "TUI error:", err)
		os.Exit(1)
	}
}

func getVersion() string {
    info, ok := debug.ReadBuildInfo()
    if !ok {
        return "unknown"
    }
    return info.Main.Version
}
