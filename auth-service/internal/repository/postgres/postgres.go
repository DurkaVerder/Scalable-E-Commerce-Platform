package postgres

import (
	"auth-service/internal/models"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) CreateUser(user models.User) error {
	if _, err := p.db.Exec(createUserQuery, user.Username, user.Email, user.Password); err != nil {
		log.Printf("Failed to create user: %v", err)
		return err
	}

	return nil
}

func (p *Postgres) GetUser(email string) (models.User, error) {
	var user models.User
	user.Email = email
	if err := p.db.QueryRow(getUserQuery, email).Scan(&user.ID, &user.Username, &user.Password); err != nil {
		log.Printf("Failed to get user: %v", err)
		return models.User{}, err
	}

	return user, nil
}
