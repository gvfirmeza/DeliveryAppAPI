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

	metricas.MetricasSistema.TotalProdutos++

	varID++
	novoProduto.ID = varID

	produtos.LProdutos.AdicionarProduto(&novoProduto)

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

	produtos.LProdutos.RemoverProduto(id)

	metricas.MetricasSistema.TotalProdutos--

	fmt.Fprintf(w, "Produto com ID %d removido com sucesso", id)
}

func BuscarProduto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	produtoID := params["id"]

	id, err := strconv.Atoi(produtoID)
	if err != nil {
		http.Error(w, "ID do produto inválido", http.StatusBadRequest)
		return
	}

	produto, err := produtos.LProdutos.BuscarProdutoByID(id)
	if err != nil {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(produto)
}

func ListarProdutos(w http.ResponseWriter, r *http.Request) {
	produtosList := produtos.LProdutos.ListarProdutos()
	json.NewEncoder(w).Encode(produtosList)
}
