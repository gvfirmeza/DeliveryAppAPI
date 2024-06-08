package metricas

import (
	"encoding/json"
	"net/http"
	"api/modelos/metricas"
)

func ObterMetricas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(metricas.MetricasSistema)
}
