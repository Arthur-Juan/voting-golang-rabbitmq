package types_test

import (
	"strings"
	"testing"

	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
)

func TestNewUser(t *testing.T) {
	first_name := "teste"
	last_name := "test last name"
	email := "teste@email.com"
	password := "pass"

	_, err := types.NewUser(first_name, last_name, email, password)
	if err != nil {
		t.Errorf("ERROR: %s\n", err.Error())
	}
}

func TestInvalidParams(t *testing.T) {
	first_name := "teste"
	last_name := "test last name"
	email := "invalid email"
	password := "pass"

	_, err := types.NewUser(first_name, last_name, email, password)
	if err == nil {
		t.Errorf("expected invalid email")
	}

	first_name = ""
	last_name = "test last name"
	email = "valid@email"
	password = "pass"

	_, err = types.NewUser(first_name, last_name, email, password)
	if err == nil {
		t.Errorf("expected required first_name")
	}

	first_name = "test"
	last_name = ""
	email = "valid@email.com"
	password = "pass"

	_, err = types.NewUser(first_name, last_name, email, password)
	if err == nil {
		t.Errorf("expected required last_name")
	}

	first_name = "test"
	last_name = "test last"
	email = ""
	password = "pass"

	_, err = types.NewUser(first_name, last_name, email, password)
	if err == nil {
		t.Errorf("expected required email")
	}

	first_name = "test"
	last_name = "test last"
	email = "valid@email.com"
	password = ""

	_, err = types.NewUser(first_name, last_name, email, password)
	if err == nil {
		t.Errorf("expected required password")
	}

}
func FuzzTest(f *testing.F) {
	inputs := []struct {
		FirstName string
		LastName  string
		Email     string
		Password  string
	}{
		{"name", "last name", "valid@email.com", "password"},
	}

	expectedErrors := []string{
		"First name cannot be empty",
		"Last name cannot be empty",
		"email cannot be empty",
		"password cannot be empty",
		"invalid email",
	}

	for _, c := range inputs {
		f.Add(c.FirstName, c.LastName, c.Email, c.Password)
	}

	f.Fuzz(func(t *testing.T, a string, b string, c string, d string) {
		_, err := types.NewUser(a, b, c, d)
		t.Logf("creating user with inputs (%s, %s, %s, %s)", a, b, c, d)

		if err != nil {
			// Check if the error is in the list of expected errors
			found := false
			for _, expectedErr := range expectedErrors {
				if strings.Contains(err.Error(), expectedErr) {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("Unexpected error creating user with inputs (%s, %s, %s, %s): %v", a, b, c, d, err)
			}
		}
	})
}
