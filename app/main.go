package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Response struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_server_projeto_korp_requests_total",
		Help: "Total de requisicoes recebidas.",
	},
	[]string{"endpoint"},
)

func init() {
	prometheus.MustRegister(totalRequests)
}

func projetoKorp(w http.ResponseWriter, r *http.Request) {

	totalRequests.WithLabelValues("/projeto-korp").Inc()

	response := Response{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func health(w http.ResponseWriter, r *http.Request) {

	totalRequests.WithLabelValues("/health").Inc()

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}

func main() {

	http.HandleFunc("/projeto-korp", projetoKorp)
	http.HandleFunc("/health", health)

	http.Handle("/metrics", promhttp.Handler())

	println("Servidor iniciado na porta 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
