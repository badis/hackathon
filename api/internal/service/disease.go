package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/badis/hackathon/internal/model"
)

// InsertDisease insert a disease in database.
func (s *Service) InsertDisease(ctx context.Context, disease_name, disease_omimid string) (disease_id int64, err error) {

	query := "INSERT INTO diseases (name, omim_id) VALUES ($1, $2) RETURNING id"

	err = s.db.QueryRowContext(ctx, query, disease_name, disease_omimid).Scan(&disease_id)

	unique := model.IsUniqueViolation(err)

	if unique && strings.Contains(err.Error(), "disease") {
		return 0, model.ErrDiseaseInserted
	}

	if err != nil {
		return 0, fmt.Errorf("could not insert disease: %v", err)
	}

	return
}
