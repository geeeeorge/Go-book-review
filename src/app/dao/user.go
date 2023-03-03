package dao

// User usersテーブルのDAO
type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
}
