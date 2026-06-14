# http-server-projeto-korp

## Descrição

Serviço HTTP desenvolvido em Golang para o desafio técnico da Korp.

A aplicação disponibiliza endpoints para consulta da aplicação, health check e métricas Prometheus, executando em containers com Nginx como reverse proxy e monitoramento através de Prometheus e Grafana.

---

## Tecnologias Utilizadas

* Golang
* Nginx
* Prometheus
* Grafana
* Docker / Podman
* Docker Compose / Podman Compose
* Ansible

---

## Funcionalidades

* Endpoint principal `/projeto-korp`
* Endpoint de saúde `/health`
* Exposição de métricas Prometheus `/metrics`
* Reverse Proxy com Nginx
* Monitoramento com Prometheus
* Dashboard Grafana
* Provisionamento automatizado com Ansible

---

## Estrutura do Projeto

```text
.
├── app
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── nginx
│   └── conf.d
│       └── http-server-projeto-korp.conf
├── prometheus
│   └── prometheus.yml
├── grafana
│   ├── provisioning
│   │   ├── datasources
│   │   └── dashboards
│   └── dashboards
├── ansible
│   ├── inventory
│   ├── site.yml
│   └── roles
├── docs
│   └── evidencias
├── docker-compose.yml
└── README.md
```

---

## Arquitetura

```text
Cliente
   │
   ▼
Nginx (8081)
   │
   ▼
Aplicação Go (8080)
   │
   ├── /projeto-korp
   ├── /health
   └── /metrics
            │
            ▼
      Prometheus
            │
            ▼
         Grafana
```

---

## Endpoints

| Endpoint      | Método | Descrição                        |
| ------------- | ------ | -------------------------------- |
| /projeto-korp | GET    | Retorna informações da aplicação |
| /health       | GET    | Endpoint de verificação de saúde |
| /metrics      | GET    | Métricas Prometheus              |

---

## Exemplo de Retorno

### GET /projeto-korp

```json
{
  "nome": "Projeto Korp",
  "horario": "2026-06-14T13:53:43Z"
}
```

### GET /health

```json
{
  "status": "ok"
}
```

---

## Como Executar

### Clonar o repositório

```bash
git clone https://github.com/rafaelpedrosorosa/http-server-projeto-korp.git
cd http-server-projeto-korp
```

### Subir os containers

Docker:

```bash
docker compose up -d
```

Podman:

```bash
podman-compose up -d
```

---

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

### Consulta Prometheus

```bash
curl -s 'http://localhost:9091/api/v1/query?query=up'
```

---

## Monitoramento

### Prometheus

```text
http://localhost:9091
```

### Grafana

```text
http://localhost:3000
```

Credenciais padrão:

```text
Usuário: admin
Senha: admin
```

---

## Dashboard Grafana

O dashboard disponibiliza informações sobre:

* Status da aplicação
* Uso de memória
* Uso de CPU
* Heap da aplicação
* Quantidade de goroutines
* Total de requisições
* Requisições por segundo
* Distribuição de requisições por endpoint

O Grafana é provisionado automaticamente através dos arquivos:

```text
grafana/provisioning/datasources/datasources.yml
grafana/provisioning/dashboards/dashboards.yml
grafana/dashboards/http-server-projeto-korp-dashboard.json
```

Não é necessária importação manual após a execução do ambiente.

---

## Automação com Ansible

O ambiente pode ser provisionado automaticamente utilizando Ansible.

O playbook realiza:

* Instalação do Docker/Podman
* Criação da rede de containers
* Build da aplicação
* Deploy dos containers
* Configuração do Nginx
* Configuração do Prometheus
* Provisionamento do Grafana
* Validação da aplicação através de requisição HTTP

Execução:

```bash
ansible-playbook -i inventory site.yml --ask-become-pass
```

---

## Evidências

### Endpoint /projeto-korp

![Endpoint Projeto Korp](docs/evidencias/app_endpoint_projeto_korp.png)

---

### Endpoint /health

![Health Check](docs/evidencias/app_healthcheck.png)

---

### Métricas Prometheus

![Prometheus Query](docs/evidencias/prometheus_query_up.png)

---

### Targets Prometheus

![Prometheus Targets](docs/evidencias/prometheus_targets_up.png)

---

### Dashboard Grafana

![Grafana Dashboard](docs/evidencias/grafana_dashboard.png)

---

### Deploy Automatizado com Ansible

![Ansible Deploy](docs/evidencias/ansible_deploy_sucesso.png)

---

### Validação da Aplicação via Ansible

![Ansible Validation](docs/evidencias/ansible_validacao_aplicacao.png)

---

## Validação do Ambiente

Após a execução do playbook Ansible ou do Docker Compose, os seguintes componentes devem estar disponíveis:

| Serviço      | URL                                |
| ------------ | ---------------------------------- |
| Aplicação    | http://localhost:8081/projeto-korp |
| Health Check | http://localhost:8081/health       |
| Prometheus   | http://localhost:9091              |
| Grafana      | http://localhost:3000              |

---

## Autor

Rafael Rosa

Desafio Técnico Korp - Infraestrutura, Containers, Monitoramento e Automação com Ansible.

