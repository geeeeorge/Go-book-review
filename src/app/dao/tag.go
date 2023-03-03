package dao

// Tag tagsテーブルのDAO
type Tag struct {
	ID     int    `db:"id"`
	UserID string `db:"user_id"`
	Name   string `db:"name"`
}

// TagBook tag_booksテーブルのDAO
type TagBook struct {
	UserID string `db:"user_id"`
	TagID  string `db:"tag_id"`
}
