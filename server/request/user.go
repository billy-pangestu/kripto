package request

//UserRegisterRequest ...
type UserRegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"username" validate:"required"`
	Pin         string `json:"pin" validate:"required"`
}

//UserVerifyOTP ...
type UserVerifyOTP struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	OTP         string `json:"otp" validate:"required"`
}

//UserSetPin ...
type UserSetPin struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Pin         string `json:"pin" validate:"required"`
}

//UserSetPinWOToken ...
type UserSetPinWOToken struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	OldPin      string `json:"old_pin" validate:"required"`
	NewPin      string `json:"new_pin" validate:"required"`
}

//UserChangePinRequest ...
type UserChangePinRequest struct {
	OldPin string `json:"old_pin" validate:"required"`
	NewPin string `json:"new_pin" validate:"required"`
}

//UserVerifyPinRequest ...
type UserVerifyPinRequest struct {
	Pin string `json:"pin" validate:"required"`
}

//UserLoginRequest ...
type UserLoginRequest struct {
	PhoneNumber string `json:"username" validate:"required"`
	Pin         string `json:"pin"`
}
