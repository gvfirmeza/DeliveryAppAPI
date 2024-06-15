package loja

import (
    "fmt"
    "net/http"
    "time"
)

var StoreOpen bool

func HandleOpenClose(w http.ResponseWriter, r *http.Request) {
    tm := time.Now()
    timestamp := tm.Format("02/01/2006 15:04:05")
    
    switch r.URL.Path {
    case "/abrir":
        StoreOpen = true
        fmt.Fprintf(w, "Loja aberta.")
        fmt.Printf("%s - Loja aberta.\n", timestamp)
    case "/fechar":
        StoreOpen = false
        fmt.Fprintf(w, "Loja fechada.")
        fmt.Printf("%s - Loja fechada.\n", timestamp)
    }
}