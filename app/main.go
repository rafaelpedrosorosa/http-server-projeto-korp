
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
                Help: "Total de requisicoes recebidas pelo servico.",
        },
        []string{"endpoint", "method", "status"},
)

func init() {
        prometheus.MustRegister(totalRequests)
}

func projetoKorp(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
                totalRequests.WithLabelValues("/projeto-korp", r.Method, "405").Inc()
                http.Error(w, "Metodo nao permitido", http.StatusMethodNotAllowed)
                return
        }

        response := Response{
                Nome:    "Projeto Korp",
                Horario: time.Now().UTC().Format(time.RFC3339),
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)

        totalRequests.WithLabelValues("/projeto-korp", r.Method, "200").Inc()
        json.NewEncoder(w).Encode(response)
	err := json.NewEncoder(w).Encode(response)
if err != nil {
	http.Error(w, "Erro ao gerar resposta", http.StatusInternalServerError)
	return
}
}

func health(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
                totalRequests.WithLabelValues("/health", r.Method, "405").Inc()
                http.Error(w, "Metodo nao permitido", http.StatusMethodNotAllowed)
                return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)

        totalRequests.WithLabelValues("/health", r.Method, "200").Inc()
        w.Write([]byte(`{"status":"ok"}`))
}

func main() {
        http.HandleFunc("/projeto-korp", projetoKorp)
        http.HandleFunc("/health", health)
        http.Handle("/metrics", promhttp.Handler())

        println("http-server-projeto-korp iniciado na porta 8080")

        err := http.ListenAndServe(":8080", nil)
        if err != nil {
                panic(err)
        }
}

