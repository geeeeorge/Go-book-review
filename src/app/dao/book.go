package dao

// Book booksテーブルのDAO
type Book struct {
	ID        int    `db:"id"`
	Title     string `db:"title"`
	Image     []byte `db:"image"`
	AmazonURL string `db:"amazon_url"`
	Status    string `db:"status"`
}

// BookSummary summariesテーブルとbooksテーブルのDAO
type BookSummary struct {
	ID             int    `db:"id"`
	Title          string `db:"title"`
	Image          []byte `db:"image"`
	AmazonURL      string `db:"amazon_url"`
	Status         string `db:"status"`
	SummaryID      string `db:"summary_id"`
	SummaryContent string `db:"summary_content"`
}

// BookTag tagsテーブルとbooksテーブルのDAO
type BookTag struct {
	ID        int    `db:"id"`
	Title     string `db:"title"`
	Image     []byte `db:"image"`
	AmazonURL string `db:"amazon_url"`
	Status    string `db:"status"`
	TagID     string `db:"tag_id"`
	TagName   string `db:"tag_name"`
}

// UserBook user_booksテーブルのDAO
type UserBook struct {
	UserID string `db:"user_id"`
	BookID string `db:"book_id"`
}
