package dao

// Summary summariesテーブルのDAO
type Summary struct {
	ID      int    `db:"id"`
	BookID  int    `db:"book_id"`
	Content string `db:"content"`
}
