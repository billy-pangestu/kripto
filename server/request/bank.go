package request

// BankRequest ...
type BankRequest struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}
