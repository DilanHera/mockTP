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
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Modes:")
		fmt.Fprintln(os.Stderr, "  -mode server   Run only the HTTP mock server")
		fmt.Fprintln(os.Stderr, "  -mode tui      Run only the TUI")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Examples:")
		fmt.Fprintf(os.Stderr, "  %s -mode server\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -mode tui\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Flags:")
		flag.PrintDefaults()
	}

	showVersion := flag.Bool("version", false, "show version")
	mode := flag.String("mode", "", "run mode: server or tui")
	flag.Parse()

	if *showVersion {
		fmt.Println(getVersion())
		os.Exit(0)
	}

	if *mode == "" {
		fmt.Fprintln(os.Stderr, "missing required flag: -mode")
		fmt.Fprintln(os.Stderr, "valid modes are: server, tui")
		flag.Usage()
		os.Exit(1)
	}

	switch *mode {
	case "server":
		runServer()
	case "tui":
		runTUI()
	default:
		fmt.Fprintln(os.Stderr, "invalid mode:", *mode)
		fmt.Fprintln(os.Stderr, "valid modes are: server, tui")
		flag.Usage()
		os.Exit(1)
	}
}

func runServer() {
	app := app.NewApp()

	r := router.SetupRouter(app)

	fmt.Fprintln(os.Stderr, "HTTP server listening on :3001")
	if err := http.ListenAndServe(":3001", r); err != nil {
		fmt.Fprintln(os.Stderr, "HTTP server error:", err)
		os.Exit(1)
	}
}

func runTUI() {
	app := app.NewApp()

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
