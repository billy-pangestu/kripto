package viewmodel

// TagVM ...
type TagVM struct {
	ID            string `json:"id"`
	TransactionID string `json:"transaction_id"`
	UserID        string `json:"user_id"`
	Hashtag       string `json:"hashtag"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at"`
}

// DetailTagVM ...
type DetailTagVM struct {
	ID        string `json:"id"`
	Hashtag   string `json:"hashtag"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
