package produtos

import "errors"

type Produto struct {
	ID        int     `json:"id"`
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Valor     float64 `json:"valor"`
}

type Nodo struct {
	Produto *Produto
	Esq     *Nodo
	Dir     *Nodo
}

type BSTProdutos struct {
	Raiz *Nodo
}

func (bst *BSTProdutos) AdicionarProduto(produto *Produto) {
	nodo := &Nodo{Produto: produto}
	if bst.Raiz == nil {
		bst.Raiz = nodo
	} else {
		bst.Raiz.adicionar(nodo)
	}
}

func (n *Nodo) adicionar(novoNodo *Nodo) {
	if novoNodo.Produto.Nome < n.Produto.Nome {
		if n.Esq == nil {
			n.Esq = novoNodo
		} else {
			n.Esq.adicionar(novoNodo)
		}
	} else {
		if n.Dir == nil {
			n.Dir = novoNodo
		} else {
			n.Dir.adicionar(novoNodo)
		}
	}
}

func (bst *BSTProdutos) BuscarProdutoByID(id int) (*Produto, error) {
	return bst.Raiz.buscarByID(id)
}

func (n *Nodo) buscarByID(id int) (*Produto, error) {
	if n == nil {
		return nil, errors.New("produto não encontrado")
	}
	if n.Produto.ID == id {
		return n.Produto, nil
	}
	if produto, err := n.Esq.buscarByID(id); err == nil {
		return produto, nil
	}
	return n.Dir.buscarByID(id)
}

func (bst *BSTProdutos) RemoverProduto(id int) {
	bst.Raiz = bst.Raiz.remover(id)
}

func (n *Nodo) remover(id int) *Nodo {
	if n == nil {
		return nil
	}
	if id < n.Produto.ID {
		n.Esq = n.Esq.remover(id)
		return n
	}
	if id > n.Produto.ID {
		n.Dir = n.Dir.remover(id)
		return n
	}
	if n.Esq == nil && n.Dir == nil {
		return nil
	}
	if n.Esq == nil {
		return n.Dir
	}
	if n.Dir == nil {
		return n.Esq
	}

	menorDireita := n.Dir.encontrarMin()
	n.Produto = menorDireita.Produto
	n.Dir = n.Dir.remover(menorDireita.Produto.ID)
	return n
}

func (n *Nodo) encontrarMin() *Nodo {
	if n.Esq == nil {
		return n
	}
	return n.Esq.encontrarMin()
}

func (bst *BSTProdutos) ListarProdutos() []*Produto {
	var produtos []*Produto
	bst.Raiz.inOrder(&produtos)
	return produtos
}

func (n *Nodo) inOrder(produtos *[]*Produto) {
	if n == nil {
		return
	}
	n.Esq.inOrder(produtos)
	*produtos = append(*produtos, n.Produto)
	n.Dir.inOrder(produtos)
}