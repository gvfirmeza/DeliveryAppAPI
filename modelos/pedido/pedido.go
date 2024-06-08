package pedido

import produtos "api/modelos/produto"

type Pedido struct {
	ID         int                 // id é implementado automaticamente no handler
	Delivery   bool                `json:"delivery"`
	Produtos   []*produtos.Produto `json:"produtos"`
	ValorTotal float64             // calculado automaticamente no handler
}
