package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"

	"github.com/DilanHera/mockTP/internal/app"
	"github.com/DilanHera/mockTP/internal/kafka"
	"github.com/DilanHera/mockTP/internal/router"
	"github.com/DilanHera/mockTP/internal/tui"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Modes:")
		fmt.Fprintln(os.Stderr, "  -mode server -port <port> -kafkabroker <kafka broker> -kafkauser <username> -kafkapass <password> Run only the HTTP mock server with optional port and kafka broker")
		fmt.Fprintln(os.Stderr, "  -mode tui      Run only the TUI")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Flags:")
		flag.PrintDefaults()
	}

	showVersion := flag.Bool("version", false, "show version")
	mode := flag.String("mode", "", "run mode: server or tui")
	port := flag.Int("port", 3001, "port to listen on")
	kafkaBroker := flag.String("kafkabroker", "192.168.2.106:9092", "kafka broker to connect to")
	kafkaUsername := flag.String("kafkauser", "", "username for kafka authentication")
	kafkaPassword := flag.String("kafkapass", "", "password for kafka authentication")

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
		kafkaConfig := kafka.KafkaAppConfig{
			Broker:   *kafkaBroker,
			Username: *kafkaUsername,
			Password: *kafkaPassword,
		}
		runServer(strconv.Itoa(*port), kafkaConfig)
	case "tui":
		runTUI()
	default:
		fmt.Fprintln(os.Stderr, "invalid mode:", *mode)
		fmt.Fprintln(os.Stderr, "valid modes are: server, tui")
		flag.Usage()
		os.Exit(1)
	}
}

func runServer(port string, kafkaConfig kafka.KafkaAppConfig) {
	app := app.NewApp(kafkaConfig)
	defer app.Kafka.CloseWriters()
	r := router.SetupRouter(app)

	fmt.Fprintln(os.Stderr, "HTTP server listening on :"+port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		fmt.Fprintln(os.Stderr, "HTTP server error:", err)
		os.Exit(1)
	}
}

func runTUI() {
	app := app.NewApp(kafka.KafkaAppConfig{})

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
