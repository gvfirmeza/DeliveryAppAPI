package pedidos

import (
	"api/modelos/metricas"
	"api/modelos/pedido"
	"api/modelos/produto"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var filaPedidos pedido.FilaPedidos
var varID int

func IncluirPedido(w http.ResponseWriter, r *http.Request) {
	var novoPedido pedido.Pedido
	err := json.NewDecoder(r.Body).Decode(&novoPedido)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(novoPedido.Produtos) == 0 {
		http.Error(w, "Nenhum produto fornecido no pedido", http.StatusBadRequest)
		return
	}

	produtosValidos := []*produtos.Produto{}

	novoPedido.ValorTotal = 0

	for _, produtoID := range novoPedido.Produtos {
		produto, err := produtos.LProdutos.BuscarProdutoByID(produtoID.ID)
		if err != nil {
			http.Error(w, "Produto não encontrado", http.StatusNotFound)
			return
		}

		produtosValidos = append(produtosValidos, produto)
		novoPedido.ValorTotal += produto.Valor
	}

	// taxa de entrega
	if novoPedido.Delivery {
		novoPedido.ValorTotal += 7
	}

	novoPedido.Produtos = produtosValidos

	varID++
	novoPedido.ID = varID

	filaPedidos.Pedidos = append(filaPedidos.Pedidos, &novoPedido)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Novo pedido incluído com ID: %d", novoPedido.ID)

	// att metricas
	metricas.MetricasSistema.PedidosAtivos += 1
}

func ExibirPedidosAbertos(w http.ResponseWriter, r *http.Request) {
	var pedidosAbertos []*pedido.Pedido

	pedidosAbertos = append(pedidosAbertos, filaPedidos.Pedidos...)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(pedidosAbertos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Essa função não deveria ser um handler, afinal de contas isso não está aberto na API...
func ExpedirPedido(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pedidoID := params["id"]

	id, err := strconv.Atoi(pedidoID)
	if err != nil {
		http.Error(w, "ID do pedido inválido", http.StatusBadRequest)
		return
	}

	for i, pedido := range filaPedidos.Pedidos {
		if pedido.ID == id {
			filaPedidos.Pedidos = append(filaPedidos.Pedidos[:i], filaPedidos.Pedidos[i+1:]...)

			fmt.Fprintf(w, "Pedido com ID %d expedido com sucesso", id)

			// att metricas
			metricas.MetricasSistema.PedidosCompletos++
			metricas.MetricasSistema.PedidosAtivos--
			metricas.MetricasSistema.LucroTotal += pedido.ValorTotal

			return
		}
	}

	http.Error(w, "Pedido não encontrado ou já expedido", http.StatusNotFound)
}
