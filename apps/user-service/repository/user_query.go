package repository

var userQuery = struct {
	selectUser string
	insertUser string
}{
	selectUser: "SELECT * FROM users",
	insertUser: "INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3)",
}
