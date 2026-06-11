Projeto Korp

Descrição
- Serviço HTTP em Go

Endpoints
- /projeto-korp
- /health
- /metrics

Tecnologias
- Golang
- Prometheus
- Grafana
- Nginx
- Podman/Docker

Arquitetura

Cliente
   |
 Nginx
   |
 Aplicação Go
   |
 Prometheus
   |
 Grafana

Como executar

podman build ...
podman run ...

Monitoramento

Prometheus:
http://localhost:9091

Grafana:
http://localhost:3000
