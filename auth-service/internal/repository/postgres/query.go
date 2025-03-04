package postgres

const (
	createUserQuery = `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
	getUserQuery    = `SELECT id, username, password FROM users WHERE email = $1`
)
