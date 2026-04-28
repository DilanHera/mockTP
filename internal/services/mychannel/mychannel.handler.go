package mychannel

import (
	"encoding/json"
	"net/http"

	"github.com/DilanHera/mockTP/internal/app"
)

type myChannelHandler struct {
	app       *app.App
	myChannel MyChannel
}

func NewMyChannelHandler(app *app.App) *myChannelHandler {
	myChannel := NewMyChannel(app)
	return &myChannelHandler{
		app:       app,
		myChannel: myChannel,
	}
}

func (h *myChannelHandler) SimSerialNoHandler(w http.ResponseWriter, r *http.Request) {
	request := &SimSerialNoRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := h.myChannel.SimSerialNo(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}
