package dao

// Summary summariesテーブルのDAO
type Summary struct {
	ID      int64  `db:"id"`
	BookID  int64  `db:"book_id"`
	Content string `db:"content"`
}
