package types

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	PostgresConnection *pgxpool.Pool
}

type Member struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Active    bool      `json:"active"`
	JoinedAt  time.Time `json:"joined_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

type Book struct {
	Id      int      `json:"id"`
	Title   string   `json:"title"`
	Year    int      `json:"year"`
	Genre   string   `json:"genre"`
	Authors []Author `json:"authors"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Instance struct {
	ID        int  `json:"id"`
	BookId    int  `json:"book_id"`
	Available bool `json:"available"`
}

type Loan struct {
	ID         int       `json:"id"`
	InstanceID int       `json:"instance_id"`
	MemberID   int       `json:"member_id"`
	IssuedAt   time.Time `json:"issued_at"`
	ReturnedAt time.Time `json:"returned_at"`
	Pending    bool      `json:"pending"`
}
