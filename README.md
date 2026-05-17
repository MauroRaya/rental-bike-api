## 📋 Pré-requisitos
* [Git](https://git-scm.com/)
* [Docker](https://docs.docker.com/get-started/docker-overview/)
* [Docker compose](https://docs.docker.com/compose/)
* [Goose](https://github.com/pressly/goose)
* [sqlc](https://github.com/sqlc-dev/sqlc)

## ⚙️ Como executar o projeto

### 1. Clone o repositório.
```bash
git clone https://github.com/MauroRaya/rental-bike-api
```

### 2. Mude para o diretório do projeto.
```bash
cd rental-bike-api
```

### 3. Configure o arquivo `.env` na raiz do projeto, conforme o arquivo `.env.example`.
```bash
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=postgres
POSTGRES_PORT=5432

GOOSE_DRIVER=postgres
GOOSE_DBSTRING=host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable
GOOSE_MIGRATION_DIR=./migrations

PORT=8080
```

### 4. Suba o banco de dados localmente.
```bash
docker compose up -d
```

> [!WARNING]  
> Não esqueça de executar `docker compose down -v` após finalizar a aplicação.

### 5. Execute as migrations.
```bash
goose up
```

### 6. Gere os arquivos do sqlc.
```bash
sqlc generate
```

### 7. Inicie a aplicação.
```bash
go run ./cmd/server/main.go
```