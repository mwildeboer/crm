package model

import "time"
import "github.com/jmoiron/sqlx/types"

type Contact struct {
	ID           string         `json:"id" db:"id"`
	PrimaryEmail string         `json:"primary_email" db:"primary_email"`
	Attributes   types.JSONText `json:"attributes" db:"attributes"`
	CreatedAt    time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at" db:"updated_at"`
}

type ContactField struct {
	Name  string
	Type  string
	Value interface{}
}

type ContactCollection []Contact

func (c *ContactCollection) Fields() ([]ContactField, error) {
	return []Field{}
}
