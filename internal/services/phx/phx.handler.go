package phx

import (
	"encoding/json"
	"net/http"

	"github.com/DilanHera/mockTP/internal/app"
)

type phxHandler struct {
	app *app.App
	phx Phx
}

func NewPhxHandler(app *app.App) *phxHandler {
	phx := NewPhx(app)
	return &phxHandler{
		app: app,
		phx: phx,
	}
}

func (h *phxHandler) RequestESIMHandler(w http.ResponseWriter, r *http.Request) {
	request := &RequestESIMRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.phx.RequestESIM(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func (h *phxHandler) NewRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	request := &NewRegistrationRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.phx.NewRegistration(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func (h *phxHandler) EncryptLibHandler(w http.ResponseWriter, r *http.Request) {
	request := &EncryptLibRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.phx.EncryptLib(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func (h *phxHandler) CheckPersoHandler(w http.ResponseWriter, r *http.Request) {
	request := &CheckPersoRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.phx.CheckPerso(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func (h *phxHandler) ProductProvisioningHandler(w http.ResponseWriter, r *http.Request) {
	request := &ProductProvisioningRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.phx.ProductProvisioning(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}
