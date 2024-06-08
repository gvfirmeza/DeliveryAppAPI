package loja

import (
    "fmt"
    "net/http"
)

var StoreOpen bool

func HandleOpenClose(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
    case "/abrir":
        StoreOpen = true
        fmt.Fprintf(w, "Loja aberta.")
    case "/fechar":
        StoreOpen = false
        fmt.Fprintf(w, "Loja fechada.")
    }
}