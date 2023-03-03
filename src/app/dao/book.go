package dao

// Book booksテーブルのDAO
type Book struct {
	ID        int64  `db:"id"`
	UserID    int64  `db:"user_id"`
	Title     string `db:"title"`
	Image     []byte `db:"image"`
	AmazonURL string `db:"amazon_url"`
	Status    string `db:"status"`
}

// BookSummaryTag summariesテーブルとTagテーブルとbooksテーブルのDAO
type BookSummaryTag struct {
	ID             int64  `db:"id"`
	UserID         int64  `db:"user_id"`
	Title          string `db:"title"`
	Image          []byte `db:"image"`
	AmazonURL      string `db:"amazon_url"`
	Status         string `db:"status"`
	SummaryID      int64  `db:"summary_id"`
	SummaryContent string `db:"summary_content"`
	TagID          int64  `db:"tag_id"`
	TagName        string `db:"tag_name"`
}
