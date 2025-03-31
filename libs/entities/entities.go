package entities

import (
	"time"
)

type User struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

type Ticket struct {
	Id     int       `db:"id"`
	Name   string    `db:"name"`
	Price  int       `db:"price"`
	Amount int       `db:"amount"`
	Date   time.Time `db:"date"`
}
