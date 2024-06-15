package pedidos

import (
	"api/modelos/metricas"
	"api/modelos/pedido"
	produtos "api/modelos/produto"
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

	// desconto de 10% para pedidos acima de 100
	if novoPedido.ValorTotal > 100 {
		novoPedido.ValorTotal *= 0.9
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
	ordenacao := r.URL.Query().Get("ordenacao")

	var pedidosAbertos []*pedido.Pedido
	pedidosAbertos = append(pedidosAbertos, filaPedidos.Pedidos...)

	switch ordenacao {
	case "bubblesort":
		bubbleSort(pedidosAbertos)
	case "quicksort":
		quickSort(pedidosAbertos, 0, len(pedidosAbertos)-1)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(pedidosAbertos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func bubbleSort(pedidos []*pedido.Pedido) {
	n := len(pedidos)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if pedidos[j].ValorTotal > pedidos[j+1].ValorTotal {
				pedidos[j], pedidos[j+1] = pedidos[j+1], pedidos[j]
			}
		}
	}
}

func quickSort(pedidos []*pedido.Pedido, low, high int) {
	if low < high {
		pi := partition(pedidos, low, high)

		quickSort(pedidos, low, pi-1)
		quickSort(pedidos, pi+1, high)
	}
}

func partition(pedidos []*pedido.Pedido, low, high int) int {
	pivot := pedidos[high].ValorTotal
	i := low - 1

	for j := low; j <= high-1; j++ {
		if pedidos[j].ValorTotal < pivot {
			i++
			pedidos[i], pedidos[j] = pedidos[j], pedidos[i]
		}
	}
	pedidos[i+1], pedidos[high] = pedidos[high], pedidos[i+1]
	return i + 1
}

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
