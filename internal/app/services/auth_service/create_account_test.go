package authservice_test

import (
	"testing"

	"github.com/arthur-juan/voting-golang-rabbitmq/config"
	authservice "github.com/arthur-juan/voting-golang-rabbitmq/internal/app/services/auth_service"
	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type Sut struct {
	suite.Suite
}

func (sut *Sut) SetupTest() {

	config.Init()
}

func (sut *Sut) Test_ShouldCreateAccount(t *testing.T) {
	input := types.CreateAccountInput{
		FirstName:       "User",
		LastName:        "Teste",
		Email:           "user@email.com",
		Password:        "root",
		ConfirmPassword: "root",
	}

	svc := authservice.NewAuthService()

	result, err := svc.CreateAccount(&input)
	t.Logf(result)
	t.Logf(err.Error())

	assert.Nil(t, err)
	assert.NotNil(t, result)

}
