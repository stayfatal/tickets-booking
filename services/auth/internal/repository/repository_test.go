package repository

import (
	"booking/libs/entities"
	"booking/services/auth/internal/testhelpers"
	"fmt"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	tx, err := testhelpers.PreparePostgres(t)
	if err != nil {
		t.Fatal(err)
	}

	repo := New(tx)

	expected := entities.User{
		Name:     "test",
		Email:    fmt.Sprintf("test%s@gmail.com", uuid.New().String()),
		Password: "123",
	}

	id, err := repo.CreateUser(expected)
	if err != nil {
		t.Fatal(err)
	}

	got := entities.User{}

	err = tx.Get(&got, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected.Email, got.Email)
}

func TestGetUserByEmail(t *testing.T) {
	tx, err := testhelpers.PreparePostgres(t)
	if err != nil {
		t.Fatal(err)
	}

	repo := New(tx)

	expected := entities.User{
		Name:     "test",
		Email:    fmt.Sprintf("test%s@gmail.com", uuid.New().String()),
		Password: "123",
	}

	_, err = tx.Exec("INSERT INTO users (name,email,password) VALUES ($1,$2,$3)", expected.Name, expected.Email, expected.Password)
	if err != nil {
		t.Fatal(err)
	}

	got, err := repo.GetUserByEmail(expected)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected.Email, got.Email)
}
