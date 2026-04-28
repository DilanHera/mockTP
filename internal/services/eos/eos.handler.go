package eos

import (
	"encoding/json"
	"net/http"

	"github.com/DilanHera/mockTP/internal/app"
)

type EosHandler struct {
	app *app.App
	Eos Eos
}

func NewEosHandler(app *app.App) *EosHandler {
	eos := NewEos(app)
	return &EosHandler{app: app, Eos: eos}
}

func (h *EosHandler) UpdateSimSerialNoHandler(w http.ResponseWriter, r *http.Request) {
	request := &UpdateSimSerialNoRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.Eos.UpdateSimSerialNo(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}
