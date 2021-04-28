package model

// Disease Model.
type Disease struct {
	ID          int64  `json:"id"`
	OMIMID      string `json:"omim_id"`
	LinkFBGroup string `json:"link_fcbk_group"`
	CreatedAt   string `json:"created_at"`
}
