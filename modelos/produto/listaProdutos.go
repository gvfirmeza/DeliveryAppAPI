package produtos

import "errors"

type ListaProdutos struct {
	Produtos []*Produto
}

func (lp *ListaProdutos) BuscarProdutoByID(id int) (*Produto, error) {
	for _, p := range lp.Produtos {
		if p.ID == id {
			return p, nil
		}
	}
	return nil, errors.New("produto não encontrado")
}