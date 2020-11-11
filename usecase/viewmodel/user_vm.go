package viewmodel

// UserVM ...
type UserVM struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	PhoneValidAt string `json:"phoneValidAt"`
	Pin          string `json:"pin"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
	DeletedAt    string `json:"deletedAt"`
}

//UserStoreVM ...
type UserStoreVM struct {
	PhoneNumber string `json:"phone_number"`
	Pin         string `json:"password"`
	Name        string `json:"name"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// UserRegisterVM ...
type UserRegisterVM struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	OTP         string `json:"otp"`
}

//UserLoginVM ...
type UserLoginVM struct {
	Token       string `json:"token"`
	ExpiredDate string `json:"expired_date"`
}
