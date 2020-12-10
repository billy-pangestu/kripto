package request

// SecretGetNotes ....
type SecretGetNotes struct {
	UserID   string `json:"user_id" validate:"required"`
	Notes    string `json:"notes" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// SecretInsertNotes ....
type SecretInsertNotes struct {
	UserID   string `json:"user_id" validate:"required"`
	Subject  string `json:"subject" validate:"required"`
	Notes    string `json:"notes" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// SecretUpdateNote ...
type SecretUpdateNote struct {
	Subject  string `json:"subject" validate:"required"`
	Notes    string `json:"notes" validate:"required"`
	Password string `json:"password" validate:"required"`
}
