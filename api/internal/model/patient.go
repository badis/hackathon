package model

import (
	"errors"
)

var (
	// ErrInvalidEmail denotes an invalid email address.
	ErrInvalidEmail = errors.New("invalid email")
	// ErrEmailTaken denotes an email already taken.
	ErrEmailTaken = errors.New("email taken")
	// ErrInvalidFirstname denotes an invalid firstname.
	ErrInvalidFirstname = errors.New("invalid firstname")
	// ErrInvalidLastname denotes an invalid lastname.
	ErrInvalidLastname = errors.New("invalid lastname")

	// ErrPatientNotFound used when patient doesn't exist in db.
	ErrPatientNotFound = errors.New("patient not found")
)

// Patient model.
type Patient struct {
	ID        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	DiseaseID int64  `json:"disease_id"`
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"`
}
