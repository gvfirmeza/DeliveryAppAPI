package metricas

type Metricas struct {
	TotalProdutos    int
	PedidosCompletos int
	PedidosAtivos    int
	LucroTotal       float64
}

var MetricasSistema Metricas