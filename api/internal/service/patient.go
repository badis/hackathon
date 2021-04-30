package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/badis/hackathon/internal/model"
)

// RegisterPatient create a patient in database.
func (s *Service) RegisterPatient(ctx context.Context, firstname, lastname, email, disease_name, disease_omimid string, disease_id int64) error {

	firstname = strings.TrimSpace(firstname)
	if !model.RegExpName.MatchString(firstname) {
		return model.ErrInvalidFirstname
	}

	lastname = strings.TrimSpace(lastname)
	if !model.RegExpName.MatchString(lastname) {
		return model.ErrInvalidLastname
	}

	email = strings.TrimSpace(email)
	if !model.RegExpEmail.MatchString(email) {
		return model.ErrInvalidEmail
	}

	var err error

	if disease_id == 0 {
		disease_id, err = s.InsertDisease(ctx, disease_name, disease_omimid)

		if err != nil {
			return err
		}
	}

	query := "INSERT INTO patients(firstname, lastname, email, disease_id) VALUES ($1, $2, $3, $4)"
	_, err = s.db.ExecContext(ctx, query, firstname, lastname, email, disease_id)

	unique := model.IsUniqueViolation(err)

	if unique && strings.Contains(err.Error(), "email") {
		return model.ErrEmailTaken
	}

	if err != nil {
		return fmt.Errorf("could not insert patient: %v", err)
	}

	return err
}
