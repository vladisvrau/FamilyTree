package rest

import (
	"net/http"
	"time"

	"github.com/vladisvrau/FamilyTree/lib/util"
)

type healthHandler struct {
	*handler
}

func NewHealthHandler(h *handler) *healthHandler {
	return &healthHandler{h}
}

func (h *healthHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// test connections
	type Status struct {
		Time   string `json:"time"`
		Status string `json:"status"`
	}

	now := time.Now().UTC().Format(time.RFC3339)
	h.logger.Info("Health: OK")
	status := Status{
		Time:   now,
		Status: "OK",
	}

	util.WriteJson(w, http.StatusOK, status)
}
