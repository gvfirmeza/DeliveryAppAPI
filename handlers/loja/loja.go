package loja

import (
    "fmt"
    "net/http"
    "strconv" // Importe strconv para converter o intervalo de string para int
    "time"
)

var StoreOpen bool
var ExpedicaoIntervalo int = 30

func HandleOpenClose(w http.ResponseWriter, r *http.Request) {
    tm := time.Now()
    timestamp := tm.Format("02/01/2006 15:04:05")

    switch r.URL.Path {
    case "/abrir":
        intervaloStr := r.URL.Query().Get("intervalo")
        if intervaloStr != "" {
            if intervalo, err := strconv.Atoi(intervaloStr); err == nil {
                ExpedicaoIntervalo = intervalo
            }
        }

        StoreOpen = true
        fmt.Fprintf(w, "Loja aberta. Intervalo de expedição atualizado para %d segundos.", ExpedicaoIntervalo)
        fmt.Printf("%s - Loja aberta. Intervalo de expedição atualizado para %d segundos.\n", timestamp, ExpedicaoIntervalo)
    case "/fechar":
        StoreOpen = false
        fmt.Fprintf(w, "Loja fechada.")
        fmt.Printf("%s - Loja fechada.\n", timestamp)
    }
}
