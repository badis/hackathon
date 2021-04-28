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
	Token       string        `json:"token"`
	ExpiresAt   time.Time     `json:"expiresAt"`
	AuthPatient model.Patient `json:"authPatient"`
}

func (s *Service) Login(ctx context.Context, email string) (LoginOutput, error) {
	var out LoginOutput

	email = strings.TrimSpace(email)
	if !model.RegExpEmail.MatchString(email) {
		return out, model.ErrInvalidEmail
	}

	query := "SELECT id, firstname, lastname FROM patients WHERE email = $1"
	err := s.db.QueryRowContext(ctx, query, email).Scan(
		&out.AuthPatient.ID,
		&out.AuthPatient.Firstname,
		&out.AuthPatient.Lastname)

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
