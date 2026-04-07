package router

import (
	"github.com/DilanHera/mockTP/internal/app"
	pgzinv "github.com/DilanHera/mockTP/internal/services/pgzinv"
	"github.com/DilanHera/mockTP/internal/services/phx"

	"github.com/go-chi/chi/v5"
)

func SetupRouter(app *app.App) chi.Router {
	pgzinvHandler := pgzinv.NewPgzinvHandler(app)
	phxHandler := phx.NewPhxHandler(app)
	router := chi.NewRouter()

	router.Post("/api/v2/PGZInventory/synchronous/ServiceProvisioning", pgzinvHandler.ServiceProvisioningHandler)
	router.Post("/api/v1/product/requestESIM", phxHandler.RequestESIMHandler)
	router.Post("/api/v1/broker/createOrder/newRegistration/msisdn/{msisdn}.json", phxHandler.NewRegistrationHandler)
	return router
}
