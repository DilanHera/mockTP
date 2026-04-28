package smis

import (
	"encoding/json"
	"net/http"

	"github.com/DilanHera/mockTP/internal/app"
)

type SmisHandler struct {
	app  *app.App
	smis Smis
}

func NewSmisHandler(app *app.App) *SmisHandler {
	smis := NewSmis(app)
	return &SmisHandler{app: app, smis: smis}
}

func (h *SmisHandler) UpdateSerialHandler(w http.ResponseWriter, r *http.Request) {
	request := &UpdateSerialRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.smis.UpdateSerial(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}
