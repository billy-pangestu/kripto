package viewmodel

// SecretVM ...
type SecretVM struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Subject   string `json:"subject"`
	Notes     string `json:"notes"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

// SecretInsertVM ...
type SecretInsertVM struct {
	UserID    string `json:"user_id"`
	Subject   string `json:"subject"`
	Notes     string `json:"notes"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

// SecretUpdateVM ...
type SecretUpdateVM struct {
	Subject   string `json:"subject"`
	Notes     string `json:"notes"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
