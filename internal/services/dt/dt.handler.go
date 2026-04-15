package dt

import (
	"encoding/json"
	"net/http"

	"github.com/DilanHera/mockTP/internal/app"
)

type DTHandler struct {
	app *app.App
	dt  DT
}

func NewDTHandler(app *app.App) *DTHandler {
	dt := NewDT(app)
	return &DTHandler{
		app: app,
		dt:  dt,
	}
}

func (h *DTHandler) ListOrderNoByDonoHandler(w http.ResponseWriter, r *http.Request) {
	request := &ListOrderNoByDonoRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.dt.ListOrderNoByDono(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func (h *DTHandler) PickingDocumentHandler(w http.ResponseWriter, r *http.Request) {
	request := &PickingDocumentRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.dt.PickingDocument(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func (h *DTHandler) QueryPrintHandler(w http.ResponseWriter, r *http.Request) {
	request := &QueryPrintRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.dt.QueryPrint(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func (h *DTHandler) QueryStockImeiMyStoreHandler(w http.ResponseWriter, r *http.Request) {
	request := &QueryStockImeiMyStoreRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.dt.QueryStockImeiMyStore(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func (h *DTHandler) ReprintReceiptFormHandler(w http.ResponseWriter, r *http.Request) {
	response, err := h.dt.ReprintReceiptForm(&ReprintReceiptFormRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func (h *DTHandler) UpdateSimSerialPersoHandler(w http.ResponseWriter, r *http.Request) {
	request := &UpdateSimSerialPersoRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.dt.UpdateSimSerialPerso(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func (h *DTHandler) AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	request := &AuthenticateRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.dt.Authenticate(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}
