package handler

import (
	"encoding/json"
	"net/http"

	"github.com/badis/hackathon/internal/model"
)

type registerPatientInput struct {
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	DiseaseName   string `json:"disease_name"`
	DiseaseOMIMID string `json:"disease_omimid"`
	DiseaseID     int64  `json:"disease_id"`
}

func (h *handler) registerPatient(w http.ResponseWriter, r *http.Request) {
	var in registerPatientInput

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.RegisterPatient(r.Context(), in.Firstname, in.Lastname, in.Email, in.Password, in.DiseaseName, in.DiseaseOMIMID, in.DiseaseID)

	if err == model.ErrInvalidEmail || err == model.ErrInvalidFirstname || err == model.ErrInvalidLastname {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err == model.ErrEmailTaken || err == model.ErrDiseaseInserted {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
