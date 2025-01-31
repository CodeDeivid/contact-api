# 📞 Contact API

Uma API RESTful para gerenciamento de contatos, construída com **Go**, **Gin** e **GORM**.

## 🚀 Tecnologias

- **Go 1.22**
- **Gin** (Framework Web)
- **GORM** (ORM para Go)
- **PostgreSQL** (Banco de Dados)
- **Docker & Docker Compose** (Ambiente de desenvolvimento)

---

## 📦 Instalação e Configuração

### 1️⃣ Clone o repositório:

```sh
git clone https://github.com/CodeDeivid/contact-api.git
cd contact-api
```

### 2️⃣ Configure as variáveis de ambiente:

Copie o arquivo `.env.example` para `.env` e ajuste as configurações:

```sh
cp .env.example .env
```

> **Exemplo de `.env`**:
> ```ini
> DB_HOST=localhost
> DB_USER=seu_usuario
> DB_PASSWORD=sua_senha
> DB_NAME=seu_banco
> DB_PORT=5432
> ```

### 3️⃣ Suba os serviços com Docker:

```sh
docker-compose up --build
```

A API estará disponível em **http://localhost:8080**.

---

## 🛠️ Endpoints da API

### 📌 Criar um contato

- **URL:** `/contacts`
- **Método:** `POST`
- **Corpo da requisição (`JSON`):**
  ```json
  {
    "name": "John Doe",
    "email": "john.doe@example.com",
    "phone": "1234567890"
  }
  ```
- **Respostas possíveis:**
  - ✅ `201 Created` – Contato criado com sucesso.
  - ❌ `400 Bad Request` – JSON inválido ou campos ausentes.
  - ❌ `409 Conflict` – E-mail ou telefone já cadastrado.

---

### 📌 Listar todos os contatos

- **URL:** `/contacts`
- **Método:** `GET`
- **Exemplo de resposta (`JSON`):**
  ```json
  [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "phone": "1234567890"
    }
  ]
  ```
- **Respostas possíveis:**
  - ✅ `200 OK` – Lista de contatos.

---

### 📌 Buscar contato por ID

- **URL:** `/contacts/:id`
- **Método:** `GET`
- **Exemplo de resposta (`JSON`):**
  ```json
  {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "phone": "1234567890"
  }
  ```
- **Respostas possíveis:**
  - ✅ `200 OK` – Contato encontrado.
  - ❌ `400 Bad Request` – ID inválido.
  - ❌ `404 Not Found` – Contato não encontrado.

---

### 📌 Atualizar um contato

- **URL:** `/contacts/:id`
- **Método:** `PUT`
- **Corpo da requisição (`JSON`):**
  ```json
  {
    "name": "Jane Doe",
    "email": "jane.doe@example.com",
    "phone": "0987654321"
  }
  ```
- **Respostas possíveis:**
  - ✅ `200 OK` – Contato atualizado com sucesso.
  - ❌ `400 Bad Request` – ID ou dados inválidos.
  - ❌ `404 Not Found` – Contato não encontrado.
  - ❌ `409 Conflict` – E-mail ou telefone já cadastrado.

---

### 📌 Deletar um contato

- **URL:** `/contacts/:id`
- **Método:** `DELETE`
- **Respostas possíveis:**
  - ✅ `204 No Content` – Contato deletado com sucesso.
  - ❌ `400 Bad Request` – ID inválido.
  - ❌ `404 Not Found` – Contato não encontrado.

---

## 🧪 Testes

Para rodar os testes automatizados, utilize:

```sh
go test ./test
```

---

## 🛠 Erros Comuns e Soluções

### ❌ Erro ao carregar o arquivo `.env`
> **Solução:** Certifique-se de que o arquivo `.env` existe e está configurado corretamente.

### ❌ Falha ao conectar ao banco de dados
> **Solução:** Verifique se o PostgreSQL está em execução e se as credenciais estão corretas.

### ❌ Formato JSON inválido
> **Solução:** Certifique-se de que a requisição contém um JSON válido e que todos os campos necessários estão preenchidos.