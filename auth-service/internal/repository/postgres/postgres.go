package postgres

import (
	"database/sql"
	"os"

	"github.com/DurkaVerder/Scalable-E-Commerce-Platform/auth-service/internal/models"
	elk "github.com/DurkaVerder/elk-send-logs/elk"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func ConnectToDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		elk.Log.SendMsg(
			elk.LogMessage{
				Level:   'E',
				Message: "Failed to connect to database",
				Fields: map[string]interface{}{
					"method": "ConnectToDB()",
					"action": "connecting to database",
					"error":  err.Error(),
				},
			})
		panic(err)
	}
	return db
}

func (p *Postgres) CreateUser(user models.User) error {
	if _, err := p.db.Exec(createUserQuery, user.Username, user.Email, user.Password); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetUser(email string) (models.User, error) {
	var user models.User
	user.Email = email
	if err := p.db.QueryRow(getUserQuery, email).Scan(&user.ID, &user.Username, &user.Password); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (p *Postgres) Close() {
	p.db.Close()
}
