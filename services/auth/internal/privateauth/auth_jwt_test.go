package privateauth

import (
	"testing"
	"ticketsbooking/libs/entities"
	"ticketsbooking/libs/publicauth"

	"github.com/stretchr/testify/assert"
)

var expected = entities.User{
	Id:    4,
	Email: "test@testmail.com",
}

func TestCreatingAndValidatingToken(t *testing.T) {
	token, err := CreateToken(expected)
	if err != nil {
		t.Error(err)
	}

	claims, err := publicauth.ValidateToken(token)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, claims)

	assert.Equal(t, expected.Id, claims.Id)
	assert.Equal(t, expected.Email, claims.Email)
}
