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
