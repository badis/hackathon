package service

import (
	"database/sql"

	"github.com/hako/branca"
)

// Service contains all the core logic. It backs a REST API.
type Service struct {
	db    *sql.DB
	codec *branca.Branca
}

// New service implementation.
func New(db *sql.DB, codec *branca.Branca) *Service {
	return &Service{
		db:    db,
		codec: codec,
	}
}
