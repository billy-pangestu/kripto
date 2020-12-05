package request

//UserAuthCheckPhoneHandler ...
type UserAuthCheckPhoneHandler struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
}
