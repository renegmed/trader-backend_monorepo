package domain

type Order struct {
	ID        string  `json:"id"`
	Symbol    string  `json:"symbol"`
	Side      string  `json:"side"`
	Quantity  int     `json:"quantity"`
	StategyID string  `json:"strategy_id"`
	Price     float64 `json:"price"`
	Status    string  `json:"status"`
}
