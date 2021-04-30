package service

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/badis/hackathon/internal/model"
)

const (
	// TokenLifeSpan time until token expires.
	TokenLifeSpan = 14 * 24 * time.Hour
)

type LoginOutput struct {
	Token       string    `json:"token"`
	ExpiresAt   time.Time `json:"expiresAt"`
	AuthPatient struct {
		ID          int64  `json:"id"`
		Firstname   string `json:"firstname"`
		Lastname    string `json:"lastname"`
		Email       string `json:"email"`
		DiseaseID   int64  `json:"diseaseId"`
		DiseaseName string `json:"diseaseName"`
		CreatedAt   string `json:"createdAt"`
	} `json:"authPatient"`
}

func (s *Service) Login(ctx context.Context, email string) (LoginOutput, error) {
	var out LoginOutput

	email = strings.TrimSpace(email)
	if !model.RegExpEmail.MatchString(email) {
		return out, model.ErrInvalidEmail
	}

	query := "SELECT p.id as id, p.firstname, p.lastname, p.email, p.disease_id, d.name as disease_name, p.created_at " +
		"FROM patients as p INNER JOIN diseases as d ON  p.disease_id = d.id AND p.email = $1"

	err := s.db.QueryRowContext(ctx, query, email).Scan(
		&out.AuthPatient.ID,
		&out.AuthPatient.Firstname,
		&out.AuthPatient.Lastname,
		&out.AuthPatient.Email,
		&out.AuthPatient.DiseaseID,
		&out.AuthPatient.DiseaseName,
		&out.AuthPatient.CreatedAt)

	if err == sql.ErrNoRows {
		return out, model.ErrPatientNotFound
	}

	if err != nil {
		return out, fmt.Errorf("could not query select user: %v", err)
	}

	out.Token, err = s.codec.EncodeToString(strconv.FormatInt(out.AuthPatient.ID, 10))
	if err != nil {
		return out, fmt.Errorf("could not create user token: %v", err)
	}

	out.ExpiresAt = time.Now().Add(TokenLifeSpan)

	return out, nil

}
