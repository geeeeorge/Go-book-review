package dao

// Tag tagsテーブルのDAO
type Tag struct {
	ID     int64  `db:"id"`
	UserID int64  `db:"user_id"`
	Name   string `db:"name"`
}

// TagBook tag_booksテーブルのDAO
type TagBook struct {
	BookID int64 `db:"book_id"`
	TagID  int64 `db:"tag_id"`
}
