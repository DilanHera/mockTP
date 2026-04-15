package im

import (
	"encoding/json"
	"net/http"

	"github.com/DilanHera/mockTP/internal/app"
)

type IMHandler struct {
	app *app.App
	im  IM
}

func NewIMHandler(app *app.App) *IMHandler {
	im := NewIM(app)
	return &IMHandler{
		app: app,
		im:  im,
	}
}

func (h *IMHandler) SendSimSerialNoHandler(w http.ResponseWriter, r *http.Request) {
	request := &SendSimSerialNoRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.im.SendSimSerialNo(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}
