package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/vladisvrau/FamilyTree/lib/database"
	"github.com/vladisvrau/FamilyTree/lib/util"
	"github.com/vladisvrau/FamilyTree/pkg/entity"
)

type kinshipHandler struct {
	*handler
}

func NewKinship(h *handler) *kinshipHandler {
	return &kinshipHandler{h}
}

func (h *kinshipHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	case http.MethodPost:
		h.post(w, r)
	case http.MethodDelete:
		h.delete(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *kinshipHandler) get(w http.ResponseWriter, r *http.Request) {
	child := r.URL.Query().Get("person")
	if child != "" {
		result := entity.KinshipCollection{}
		childId, err := strconv.Atoi(child)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err = h.services.Person.GetByID(r.Context(), childId)
		if err != nil {
			h.handleServiceError(w, err)
			return
		}

		kinships, err := h.services.Kinship.GetParents(r.Context(), childId)
		if err != nil && err != database.ErrEntityNotFound {
			h.handleServiceError(w, err)
			return
		}

		if kinships != nil {
			result = append(result, kinships...)
		}

		kinships, err = h.services.Kinship.GetChildren(r.Context(), childId)
		if err != nil && err != database.ErrEntityNotFound {
			h.handleServiceError(w, err)
			return
		}

		if kinships != nil {
			result = append(result, kinships...)
		}

		util.WriteJson(w, http.StatusOK, result)
		return
	}

}

func (h *kinshipHandler) handleServiceError(w http.ResponseWriter, err error) {
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

func (h *kinshipHandler) post(w http.ResponseWriter, r *http.Request) {
	p := entity.Kinship{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.services.Kinship.Create(r.Context(), &p)
	if err != nil {
		h.handleServiceError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *kinshipHandler) delete(w http.ResponseWriter, r *http.Request) {
	p := entity.Kinship{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	h.handler.services.Kinship.Delete(r.Context(), &p)
	w.WriteHeader(http.StatusOK)
}
