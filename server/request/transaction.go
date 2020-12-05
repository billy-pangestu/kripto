package request

// TransactionRequest ...
type TransactionRequest struct {
	MoneyIn  float64     `json:"money_in"`
	MoneyOut float64     `json:"money_out"`
	Notes    string      `json:"notes"`
	Tags     []DetailTag `json:"tags"`
}

// DetailTag ...
type DetailTag struct {
	Name string `json:"name"`
}
