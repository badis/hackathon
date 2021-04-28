package handler

import (
	"encoding/json"
	"net/http"

	"github.com/badis/hackathon/internal/model"
)

type registerPatientInput struct {
	Firstname, Lastname, Email string
}

func (h *handler) registerPatient(w http.ResponseWriter, r *http.Request) {
	var in registerPatientInput

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.RegisterPatient(r.Context(), in.Firstname, in.Lastname, in.Email)

	if err == model.ErrInvalidEmail || err == model.ErrInvalidFirstname || err == model.ErrInvalidLastname {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err == model.ErrEmailTaken {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}