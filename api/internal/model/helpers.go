package model

import (
	"regexp"

	"github.com/jackc/pgx"
	"github.com/lib/pq"
)

var (
	// RegExpEmail a regular expression for validating emails.
	RegExpEmail = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

	// RegExpName a regular expression for validating names.
	RegExpName = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]{0,17}$`)
)

// IsUniqueViolation check if a unique constraint has been violated in
// database.
func IsUniqueViolation(err error) bool {
	pgerr, ok := err.(pgx.PgError)
	return ok && pgerr.Code == "23505"
}

// IsForeignKeyViolation check if a foreign key constraint has been violated in
// database.
func IsForeignKeyViolation(err error) bool {
	pqerr, ok := err.(*pq.Error)
	return ok && pqerr.Code == "23503"
}
