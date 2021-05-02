package model

import (
	"encoding/base64"
	"fmt"
	"regexp"

	"github.com/jackc/pgx"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
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

// HashPassword hashes the clear-text password and encodes it as base64,
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10 /*cost*/)
	if err != nil {
		return "", err
	}

	// Encode the entire thing as base64 and return
	hashBase64 := base64.StdEncoding.EncodeToString(hashedBytes)

	return hashBase64, nil
}

// ComparePassword hashes the test password and then compares
// the two hashes.
func ComparePassword(hashBase64, testPassword string) bool {

	// Decode the real hashed and salted password so we can
	// split out the salt
	hashBytes, err := base64.StdEncoding.DecodeString(hashBase64)
	if err != nil {
		fmt.Println("Error, we were given invalid base64 string", err)
		return false
	}

	err = bcrypt.CompareHashAndPassword(hashBytes, []byte(testPassword))
	return err == nil
}
