package rest

import (
	"net/http"
	"strconv"

	"github.com/vladisvrau/FamilyTree/lib/database"
	"github.com/vladisvrau/FamilyTree/lib/util"
)

type treeHandler struct {
	*handler
}

func NewTree(h *handler) *treeHandler {
	return &treeHandler{h}
}

func (h *treeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *treeHandler) get(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("id") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	child := r.URL.Query().Get("id")
	if child != "" {
		childId, err := strconv.Atoi(child)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		result, err := h.services.Tree.BuildFamilyTree(r.Context(), childId)
		if err != nil {
			h.handleServiceError(w, err)
			return
		}
		util.WriteJson(w, http.StatusOK, result)
		return
	}

}

func (h *treeHandler) handleServiceError(w http.ResponseWriter, err error) {
	switch err {
	case database.ErrEntityNotFound:
		w.WriteHeader(http.StatusNotFound)
	case database.ErrInvalidInsert:
		w.WriteHeader(http.StatusBadRequest)
	default:
		h.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
