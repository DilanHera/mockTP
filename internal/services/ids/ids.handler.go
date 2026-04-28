package ids

import (
	"encoding/json"
	"net/http"

	"github.com/DilanHera/mockTP/internal/app"
)

type idsHandler struct {
	app *app.App
	ids Ids
}

func NewIdsHandler(app *app.App) *idsHandler {
	return &idsHandler{app: app}
}

func (h *idsHandler) AuthenHandler(w http.ResponseWriter, r *http.Request) {
	request := &AuthenRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.ids.Authen(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func (h *idsHandler) UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	request := &UserInfoRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.ids.UserInfo(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}
