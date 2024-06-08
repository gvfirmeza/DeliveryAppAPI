package main

import (
	"fmt"
	"net/http"

	"api/handlers/loja"
	"api/handlers/metricas"
	"api/handlers/pedidos"
	"api/handlers/produtos"
	"api/processamento"

	"github.com/gorilla/mux"
)

func main() {
	// router
	r := mux.NewRouter()

	// Rotas para abrir e fechar a loja
	r.HandleFunc("/abrir", loja.HandleOpenClose).Methods("POST")
	r.HandleFunc("/fechar", loja.HandleOpenClose).Methods("POST")

	// Rotas para o pacote 'produtos'
	r.HandleFunc("/produto", produtos.AdicionarProduto).Methods("POST")
	r.HandleFunc("/produto/{id}", produtos.RemoverProduto).Methods("DELETE")
	r.HandleFunc("/produto/{id}", produtos.BuscarProduto).Methods("GET")
	r.HandleFunc("/produtos", produtos.ListarProdutos).Methods("GET")

	// Rotas para o pacote 'pedidos'
	r.HandleFunc("/pedido", pedidos.IncluirPedido).Methods("POST")
	r.HandleFunc("/pedidos", pedidos.ExibirPedidosAbertos).Methods("GET")
	r.HandleFunc("/pedido/{id}", pedidos.ExpedirPedido).Methods("POST")

	// Rota para as métricas
	r.HandleFunc("/metricas", metricas.ObterMetricas).Methods("GET")

	// Processando dados através de goroutine
	go processamento.LojaAberta()

	// Inicialização do servidor HTTP
	fmt.Println("Servidor iniciado em http://localhost:8080")
	http.ListenAndServe(":8080", r)
}