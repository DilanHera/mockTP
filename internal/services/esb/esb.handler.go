package esb

import (
	"encoding/json"
	"net/http"

	"github.com/DilanHera/mockTP/internal/app"
)

type ESBHandler struct {
	app *app.App
	esb ESB
}

func NewESBHandler(app *app.App) *ESBHandler {
	esbSvc := NewESB(app)
	return &ESBHandler{
		app: app,
		esb: esbSvc,
	}
}

func (h *ESBHandler) OauthTokenHandler(w http.ResponseWriter, r *http.Request) {
	request := &OauthTokenRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.esb.OauthToken(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(response.HttpStatusCode)
	json.NewEncoder(w).Encode(response)
}

func (h *ESBHandler) CreateFreightOrderHandler(w http.ResponseWriter, r *http.Request) {
	request := &CreateFreightOrderRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.esb.CreateFreightOrder(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func (h *ESBHandler) DOCreationHandler(w http.ResponseWriter, r *http.Request) {
	request := &DOCreationRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.esb.DOCreation(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func (h *ESBHandler) LegoupdateOrderStatusHandler(w http.ResponseWriter, r *http.Request) {
	request := &LegoupdateOrderStatusRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.esb.LegoupdateOrderStatus(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func (h *ESBHandler) PersosimHandler(w http.ResponseWriter, r *http.Request) {
	request := &PersosimRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.esb.Persosim(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func (h *ESBHandler) SerialNumberExpirationDateHandler(w http.ResponseWriter, r *http.Request) {
	request := &SerialNumberExpirationDateRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.esb.SerialNumberExpirationDate(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}
