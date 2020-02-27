package models

import (
	"testing"
	"time"
	"common/models"
)

func newPerson(id int64, fullName string, dni string, birtDate time.Time) *models.person {
	return &Person{
		ID:        id,
		FullName:  fullName,
		DNI:       dni,
		BirthDate: birtDate,
	}
}

func Test_validate(t *testing.T) {
	person := newPerson(1, "Alexander Correa", "123456789", nil)
	err := person.Validate()
	if err != nil {
		t.Errorf("Error during person creation %v", err)
		t.Fail()
	}
}
