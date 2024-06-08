package processamento

import (
	"fmt"
	"io"
	"net/http"
	"time"
	"api/handlers/loja"
)

func LojaAberta() {
	idPedido := 1
	for {
		if loja.StoreOpen {
			// chamando rota para expedir pedidos "/pedido/{id}"
			url := fmt.Sprintf("http://localhost:8080/pedido/%d", idPedido)

			resp, err := http.Post(url, "application/json", nil)
			if err != nil {
				fmt.Println("Erro ao enviar pedido:", err)
				break
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Erro ao ler resposta:", err)
				break
			}

			if resp.StatusCode == http.StatusOK {
				fmt.Println("Pedido", idPedido, "expedido com sucesso.")
				idPedido++
			} else {
				fmt.Println("Erro ao expedir pedido", idPedido, ":", string(body))
			}
		} else {
			fmt.Println("Loja fechada")
		}
		time.Sleep(30 * time.Second)
	}
}
