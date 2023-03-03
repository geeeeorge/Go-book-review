package dao

// Tag tagsテーブルのDAO
type Tag struct {
	ID     int    `db:"id"`
	UserID int    `db:"user_id"`
	Name   string `db:"name"`
}

// TagBook tag_booksテーブルのDAO
type TagBook struct {
	BookID int `db:"book_id"`
	TagID  int `db:"tag_id"`
}
