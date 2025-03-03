package postgres

import (
	"auth-service/internal/models"
	"database/sql"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) CreateUser(models.User) error {
	return nil
}

func (p *Postgres) GetUser(email string) (models.User, error) {
	var user models.User

	return user, nil
}
