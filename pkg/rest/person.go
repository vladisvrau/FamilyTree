package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/vladisvrau/FamilyTree/lib/database"
	"github.com/vladisvrau/FamilyTree/lib/util"
	"github.com/vladisvrau/FamilyTree/pkg/entity"
)

type personHandler struct {
	*handler
}

func NewPersonHanlder(h *handler) *personHandler {
	return &personHandler{h}
}

func (h *personHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	case http.MethodPost:
		h.post(w, r)
	case http.MethodDelete:
		h.delete(w, r)
	case http.MethodPut:
		h.put(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *personHandler) get(w http.ResponseWriter, r *http.Request) {

	queryField := ""
	if r.URL.Query().Has("id") {
		queryField = "id"
	} else if r.URL.Query().Has("name") {
		queryField = "name"
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	field := r.URL.Query().Get(queryField)

	var people []*entity.Person
	var err error
	if queryField == "id" {
		id, err := strconv.Atoi(field)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		person, err := h.services.Person.GetByID(r.Context(), id)

		if err != nil {
			h.handleServiceError(w, err)
			return
		}
		people = append(people, person)
	} else {
		people, err = h.services.Person.GetByName(r.Context(), field)
		if err != nil {
			h.handleServiceError(w, err)
			return
		}
	}

	util.WriteJson(w, http.StatusOK, people)
}

func (h *personHandler) post(w http.ResponseWriter, r *http.Request) {
	p := entity.Person{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	person, err := h.services.Person.Save(r.Context(), &p)
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	util.WriteJson(w, http.StatusOK, person)
}

func (h *personHandler) delete(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("id") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	field := r.URL.Query().Get("id")

	id, err := strconv.Atoi(field)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.services.Person.Delete(r.Context(), id)
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *personHandler) put(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("id") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	field := r.URL.Query().Get("id")
	id, err := strconv.Atoi(field)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p := entity.Person{}
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p.ID = id

	result, err := h.services.Person.Edit(r.Context(), &p)

	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	util.WriteJson(w, http.StatusOK, result)
}

func (h *personHandler) handleServiceError(w http.ResponseWriter, err error) {
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
