# http-server-projeto-korp

## Descrição

Serviço HTTP desenvolvido em Golang para o desafio técnico da Korp.

A aplicação disponibiliza endpoints para consulta da aplicação, health check e métricas Prometheus, executando em containers com Nginx como reverse proxy e monitoramento através de Prometheus e Grafana.

## Tecnologias Utilizadas

* Golang
* Nginx
* Prometheus
* Grafana
* Podman / Docker
* Docker Compose

## Endpoints

| Endpoint      | Descrição                                        |
| ------------- | ------------------------------------------------ |
| /projeto-korp | Retorna informações da aplicação em formato JSON |
| /health       | Endpoint de verificação de saúde da aplicação    |
| /metrics      | Métricas Prometheus                              |

## Arquitetura

Cliente → Nginx → Aplicação Go

Prometheus → Coleta métricas da aplicação

Grafana → Visualização das métricas

## Como Executar

### Clonar o projeto

```bash
git clone https://github.com/rafaelpedrosorosa/http-server-projeto-korp.git
cd http-server-projeto-korp
```

### Subir os containers

```bash
docker compose up -d
```

ou

```bash
podman-compose up -d
```

## Testes

### Endpoint principal

```bash
curl http://localhost:8081/projeto-korp
```

### Health Check

```bash
curl http://localhost:8081/health
```

### Métricas

```bash
curl http://localhost:8081/metrics
```

## Monitoramento

Prometheus:

http://localhost:9091

Grafana:

http://localhost:3000

## Dashboard

O dashboard Grafana apresenta:

* Status da aplicação
* Uso de memória
* Uso de CPU
* Heap da aplicação
* Quantidade de goroutines
* Total de requisições
* Requisições por segundo
* Distribuição de requisições por endpoint

O dashboard exportado encontra-se no repositório através do arquivo JSON disponibilizado para importação.
