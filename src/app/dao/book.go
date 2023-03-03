package dao

// Book booksテーブルのDAO
type Book struct {
	ID        int    `db:"id"`
	UserID    int    `db:"user_id"`
	Title     string `db:"title"`
	Image     []byte `db:"image"`
	AmazonURL string `db:"amazon_url"`
	Status    string `db:"status"`
}

// BookSummary summariesテーブルとbooksテーブルのDAO
type BookSummary struct {
	ID             int    `db:"id"`
	UserID         int    `db:"user_id"`
	Title          string `db:"title"`
	Image          []byte `db:"image"`
	AmazonURL      string `db:"amazon_url"`
	Status         string `db:"status"`
	SummaryID      int    `db:"summary_id"`
	SummaryContent string `db:"summary_content"`
}

// BookTag tagsテーブルとbooksテーブルのDAO
type BookTag struct {
	ID        int    `db:"id"`
	UserID    int    `db:"user_id"`
	Title     string `db:"title"`
	Image     []byte `db:"image"`
	AmazonURL string `db:"amazon_url"`
	Status    string `db:"status"`
	TagID     int    `db:"tag_id"`
	TagName   string `db:"tag_name"`
}
