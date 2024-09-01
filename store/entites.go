package store

type Items struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Qtd   int64   `json:"qtd"`
	Price float32 `json:"price"`
}

type listItems struct {
	Items []Items `json:"items"`
}
