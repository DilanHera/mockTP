package router

import (
	"github.com/DilanHera/mockTP/internal/app"
	"github.com/DilanHera/mockTP/internal/services/dt"
	"github.com/DilanHera/mockTP/internal/services/eos"
	"github.com/DilanHera/mockTP/internal/services/esb"
	"github.com/DilanHera/mockTP/internal/services/ids"
	"github.com/DilanHera/mockTP/internal/services/im"
	"github.com/DilanHera/mockTP/internal/services/mychannel"
	"github.com/DilanHera/mockTP/internal/services/pgzinv"
	"github.com/DilanHera/mockTP/internal/services/phx"
	"github.com/DilanHera/mockTP/internal/services/smis"

	"github.com/go-chi/chi/v5"
)

func SetupRouter(app *app.App) chi.Router {
	pgzinvHandler := pgzinv.NewPgzinvHandler(app)
	phxHandler := phx.NewPhxHandler(app)
	dtHandler := dt.NewDTHandler(app)
	imHandler := im.NewIMHandler(app)
	esbHandler := esb.NewESBHandler(app)
	mychannel := mychannel.NewMyChannelHandler(app)
	eos := eos.NewEosHandler(app)
	smis := smis.NewSmisHandler(app)
	ids := ids.NewIdsHandler(app)
	router := chi.NewRouter()

	router.Post("/api/v2/PGZInventory/synchronous/ServiceProvisioning", pgzinvHandler.ServiceProvisioningHandler)
	router.Post("/api/v1/product/requestESIM", phxHandler.RequestESIMHandler)
	router.Post("/api/v1/broker/createOrder/newRegistration/msisdn/{msisdn}.json", phxHandler.NewRegistrationHandler)
	router.Post("/api/v1/product/appConfig", phxHandler.EncryptLibHandler)
	router.Post("/api/v1/product/checkPerso", phxHandler.CheckPersoHandler)
	router.Post("/api/v1/fulFillment/productProvisioning", phxHandler.ProductProvisioningHandler)
	router.Post("/DTWS/api/stock/v1/queryStockImeiMyStore", dtHandler.QueryStockImeiMyStoreHandler)
	router.Post("/DTWS/api/after-sale/v1/picking-document", dtHandler.PickingDocumentHandler)
	router.Get("/DTWS/api/reprint-form/v1/print-receipt", dtHandler.ReprintReceiptFormHandler)
	router.Post("/DTWS/api/after-sale/v1/update-sim-serial-perso", dtHandler.UpdateSimSerialPersoHandler)
	router.Post("/DTWS/api/sale/v1/list-orderno-by-dono", dtHandler.ListOrderNoByDonoHandler)

	router.Post("/dtauth/api/auth/authenticate", dtHandler.AuthenticateHandler)

	router.Post("/dtcore-saleout/api/stock/v1/queryStockImeiMyStore", dtHandler.QueryStockImeiMyStoreHandler)
	router.Post("/dtcore-saleout/api/after-sale/v1/picking-document", dtHandler.PickingDocumentHandler)
	router.Get("/dtcore-saleout/api/reprint-form/v1/print-receipt", dtHandler.ReprintReceiptFormHandler)
	router.Post("/dtcore-saleout/api/saleout/v1/query-print", dtHandler.QueryPrintHandler)

	router.Post("/dtreport-replicas/api/sale/v1/list-orderno-by-dono", dtHandler.ListOrderNoByDonoHandler)

	router.Post("/order/sendGoodsReceive", dtHandler.ListOrderNoByDonoHandler)

	router.Post("/prweb/PRRestService/AISNIMWorkRequestDelivery/Services/SendSimSerialNo", imHandler.SendSimSerialNoHandler)

	router.Post("/auth/v3.2/oauth/token", esbHandler.OauthTokenHandler)
	router.Post("/sap-px/v1/FreightOrder/0170/CreateFreightOrder", esbHandler.CreateFreightOrderHandler)
	router.Post("/sap-px/v1/DeliveryOrder/0145/DOCreation", esbHandler.DOCreationHandler)
	router.Post("/lego-be-updateorderstatus/action/updateOrderStatus", esbHandler.LegoupdateOrderStatusHandler)
	router.Post("/gomo-px/api/warehouse/sim/persosim", esbHandler.PersosimHandler)
	router.Post("/sap-px/v1/BatchMaster/0025/SerialNumberExpirationDate", esbHandler.SerialNumberExpirationDateHandler)

	router.Post("/api/receive/sim-serial-no", mychannel.SimSerialNoHandler)

	router.Post("/admin/api/v1.0/warehouse/sim/persosim", eos.UpdateSimSerialNoHandler)

	router.Post("/CPSM/RestAPIsService/api/UpdateDataReceiveOptimus", smis.UpdateSerialHandler)

	router.Post("/oauth2/userinfo", ids.UserInfoHandler)
	return router
}
