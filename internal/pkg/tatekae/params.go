package tatekae

type TatekaeParams struct {
	Members      []string      `json:"members"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}
