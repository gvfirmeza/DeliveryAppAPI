package produtos

type Produto struct {
	ID        int     // id é implementado automaticamente no handler
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Valor     float64 `json:"valor"`
}

var LProdutos ListaProdutos