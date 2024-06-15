package processamento

import (
	"api/handlers/loja"
	"fmt"
	"io"
	"net/http"
	"time"
)

func LojaAberta() {
	idPedido := 1
	tm := time.Now()
	timestamp := tm.Format("02/01/2006 15:04:05")
	for {
		if loja.StoreOpen {
			// chamando rota para expedir pedidos "/pedido/{id}"
			url := fmt.Sprintf("http://localhost:8080/pedido/%d", idPedido)

			resp, err := http.Post(url, "application/json", nil)
			if err != nil {
				fmt.Printf("Erro ao enviar pedido: %s %s\n", err, timestamp)
				break
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Erro ao ler resposta: %s %s\n", err, timestamp)
				break
			}

			if resp.StatusCode == http.StatusOK {
				fmt.Println("Pedido", idPedido, "expedido com sucesso.", timestamp)
				idPedido++
			} else {
				fmt.Println("Erro ao expedir pedido", idPedido, ":", string(body), timestamp)
			}
		} else {
			fmt.Printf("%s - Loja Fechada.\n", timestamp)
		}
		time.Sleep(30 * time.Second)
	}
}
