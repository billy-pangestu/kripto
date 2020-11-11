package request

// OTPRequest ...
type OTPRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
}
