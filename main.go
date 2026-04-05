package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/DilanHera/mockTP/internal/app"
	"github.com/DilanHera/mockTP/internal/router"
)

func main() {
	app := app.NewApp()
	router := router.SetupRouter(app)
	fmt.Println("Server is running on port 3000")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
