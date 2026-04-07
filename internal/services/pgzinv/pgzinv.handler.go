package pgzinv

import (
	"encoding/json"
	"net/http"

	"github.com/DilanHera/mockTP/internal/app"
	serviceprovisioningmodel "github.com/DilanHera/mockTP/internal/services/pgzinv/model"
)

type PgzinvHandler struct {
	app    *app.App
	pgzinv Pgzinv
}

func NewPgzinvHandler(app *app.App) *PgzinvHandler {
	pgzinv := NewPgzinv(app)
	return &PgzinvHandler{
		app:    app,
		pgzinv: pgzinv,
	}
}

func (h *PgzinvHandler) ServiceProvisioningHandler(w http.ResponseWriter, r *http.Request) {
	request := &serviceprovisioningmodel.ServiceProvisioningRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(request.ResourceItemList) > 0 && request.ResourceItemList[0].ResourceName == "" {
		http.Error(w, "ResourceName is required", http.StatusBadRequest)
		return
	}

	resourceItem, err := json.Marshal(request.ResourceItemList[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payload := &serviceprovisioningmodel.ServiceProvisioningPayload{
		ResourceName:  request.ResourceItemList[0].ResourceName,
		RequestHeader: request.RequestHeader,
		Payload:       resourceItem,
	}
	response, err := h.pgzinv.ServiceProvisioning(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}
