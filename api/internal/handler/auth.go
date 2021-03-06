package handler

import (
	"encoding/json"
	"net/http"

	"github.com/badis/hackathon/internal/model"
)

type loginInput struct {
	Email    string
	Password string
}

func (h *handler) login(w http.ResponseWriter, r *http.Request) {
	var in loginInput

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := h.Login(r.Context(), in.Email, in.Password)

	if err == model.ErrInvalidEmail {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err == model.ErrPatientNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	respond(w, out, http.StatusOK)
}
