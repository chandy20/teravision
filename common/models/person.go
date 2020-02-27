package models

import (
	"errors"
	"time"
)

//Person model to representate an person
type Person struct {
	ID        int64      `json:"id"`
	FullName  string     `json:"full_name"`
	DNI       string     `json:"dni"`
	BirthDate time.Time `json:"birth_date"`
}

//validate function to validate a model of person
func (p *Person) Validate() error {
	if p.FullName == "" {
		return errors.New("full_name_cannot_be_empty")
	}
	if p.DNI == "" {
		return errors.New("dni_cannot_be_empty")
	}
	if p.BirthDate == nil {
		return errors.New("birt_date_cannot_be_nil")
	}
	return nil
}