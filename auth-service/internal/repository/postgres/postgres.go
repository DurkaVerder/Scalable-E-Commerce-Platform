package postgres

import (
	"database/sql"
	"os"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/auth-service/internal/models"
	elk "github.com/DurkaVerder/Scalable-E-Commerce-Platform/auth-service/pkg/logs"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		elk.Log.Error("Failed to connect to database", map[string]interface{}{
			"method": "ConnectToDB",
			"action": "connect_to_db",
			"error":  err,
		})
		return nil, err
	}
	return db, nil
}

func (p *Postgres) CreateUser(user models.User) error {
	if _, err := p.db.Exec(createUserQuery, user.Username, user.Email, user.Password); err != nil {
		elk.Log.Error("Failed to create user", map[string]interface{}{
			"method": "CreateUser",
			"action": "create_user",
			"error":  err,
		})
		return err
	}

	return nil
}

func (p *Postgres) GetUser(email string) (models.User, error) {
	var user models.User
	user.Email = email
	if err := p.db.QueryRow(getUserQuery, email).Scan(&user.ID, &user.Username, &user.Password); err != nil {
		elk.Log.Error("Failed to get user", map[string]interface{}{
			"method": "GetUser",
			"action": "get_user",
			"error":  err,
		})
		return models.User{}, err
	}

	return user, nil
}
