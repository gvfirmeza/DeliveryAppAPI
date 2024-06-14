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
	novoNodo := &Nodo{Produto: produto}
	if bst.Raiz == nil {
		bst.Raiz = novoNodo
	} else {
		adicionarNodo(bst.Raiz, novoNodo)
	}
}

func adicionarNodo(atual, novo *Nodo) {
	if novo.Produto.Nome < atual.Produto.Nome {
		if atual.Esq == nil {
			atual.Esq = novo
		} else {
			adicionarNodo(atual.Esq, novo)
		}
	} else {
		if atual.Dir == nil {
			atual.Dir = novo
		} else {
			adicionarNodo(atual.Dir, novo)
		}
	}
}

func (bst *BSTProdutos) BuscarProdutoByID(id int) (*Produto, error) {
	return buscarPorID(bst.Raiz, id)
}

func buscarPorID(nodo *Nodo, id int) (*Produto, error) {
	if nodo == nil {
		return nil, errors.New("produto não encontrado")
	}
	if nodo.Produto.ID == id {
		return nodo.Produto, nil
	}
	produto, err := buscarPorID(nodo.Esq, id)
	if err == nil {
		return produto, nil
	}
	return buscarPorID(nodo.Dir, id)
}

func (bst *BSTProdutos) RemoverProduto(id int) {
	bst.Raiz = removerNodo(bst.Raiz, id)
}

func removerNodo(nodo *Nodo, id int) *Nodo {
	if nodo == nil {
		return nil
	}
	if id < nodo.Produto.ID {
		nodo.Esq = removerNodo(nodo.Esq, id)
		return nodo
	}
	if id > nodo.Produto.ID {
		nodo.Dir = removerNodo(nodo.Dir, id)
		return nodo
	}
	if nodo.Esq == nil {
		return nodo.Dir
	}
	if nodo.Dir == nil {
		return nodo.Esq
	}
	sucessor := nodo.Dir
	for sucessor.Esq != nil {
		sucessor = sucessor.Esq
	}
	nodo.Produto = sucessor.Produto
	nodo.Dir = removerNodo(nodo.Dir, sucessor.Produto.ID)
	return nodo
}

func (bst *BSTProdutos) ListarProdutos() []*Produto {
	var produtos []*Produto
	inOrderTraversal(bst.Raiz, &produtos)
	return produtos
}

func inOrderTraversal(nodo *Nodo, produtos *[]*Produto) {
	if nodo != nil {
		inOrderTraversal(nodo.Esq, produtos)
		*produtos = append(*produtos, nodo.Produto)
		inOrderTraversal(nodo.Dir, produtos)
	}
}