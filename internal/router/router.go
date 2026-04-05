package router

import (
	"github.com/DilanHera/mockTP/internal/app"
	pgzinv "github.com/DilanHera/mockTP/internal/services/pgzinv"

	"github.com/go-chi/chi/v5"
)

func SetupRouter(app *app.App) chi.Router {
	pgzinvHandler := pgzinv.NewPgzinvHandler(app)
	router := chi.NewRouter()

	router.Post("/pgzinv/service-provisioning", pgzinvHandler.ServiceProvisioningHandler)

	return router
}
