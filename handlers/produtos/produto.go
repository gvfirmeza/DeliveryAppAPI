package produtos

import (
	"api/modelos/metricas"
	"api/modelos/produto"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

var varID int

func AdicionarProduto(w http.ResponseWriter, r *http.Request) {
	var novoProduto produtos.Produto
	err := json.NewDecoder(r.Body).Decode(&novoProduto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// att metricas
	metricas.MetricasSistema.TotalProdutos++

	varID++
	novoProduto.ID = varID

	produtos.LProdutos.Produtos = append(produtos.LProdutos.Produtos, &novoProduto)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Novo produto adicionado com ID: %d", novoProduto.ID)
}

func RemoverProduto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	produtoID := params["id"]

	id, err := strconv.Atoi(produtoID)
	if err != nil {
		http.Error(w, "ID do produto inválido", http.StatusBadRequest)
		return
	}

	for i, produto := range produtos.LProdutos.Produtos {
		if produto.ID == id {
			produtos.LProdutos.Produtos = append(produtos.LProdutos.Produtos[:i], produtos.LProdutos.Produtos[i+1:]...)

			fmt.Fprintf(w, "Produto com ID %d removido com sucesso", id)

			// att metricas
			metricas.MetricasSistema.TotalProdutos--
			
			return
		}
	}

	http.Error(w, "Produto não encontrado", http.StatusNotFound)
}

func BuscarProduto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	produtoID := params["id"]

	id, err := strconv.Atoi(produtoID)
	if err != nil {
		http.Error(w, "ID do produto inválido", http.StatusBadRequest)
		return
	}

	for _, produto := range produtos.LProdutos.Produtos {
		if produto.ID == id {
			json.NewEncoder(w).Encode(produto)
			return
		}
	}

	http.Error(w, "Produto não encontrado", http.StatusNotFound)
}

func ListarProdutos(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(produtos.LProdutos.Produtos)
}