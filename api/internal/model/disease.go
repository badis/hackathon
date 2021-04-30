package model

import "errors"

var (
	// ErrDiseaseInserted denotes a disease already inserted.
	ErrDiseaseInserted = errors.New("disease inserted")
)

// Disease Model.
type Disease struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	OMIMID      string `json:"omim_id"`
	LinkFBGroup string `json:"link_fcbk_group"`
	CreatedAt   string `json:"created_at"`
}
