# Password vault Go
Desenvolvido em 2024, o password vault tem por finalidade ser um serviço de armazenamento
de logins, senhas e palavras chave.

## Instalação
### Pré-requisitos
- Go 1.23.5 or higher
- Docker & docker compose

### Passo a passo
1. Clone e acesse o diretório do repositório
    > git clone git@github.com:ropehapi/password-vault-go.git
    
    > cd password-vault-go
2. Copie o arquivo .env.example e configure as variáveis de ambiente
    > cp .env.example .env
3. Suba o conteiner do banco de dados
    > docker compose up -d
4. Execute a aplicação
    > make run

*Note que essa aplicação possui um middleware de autenticação que passa por um IDP externo.
Para ignorar essa validação, você pode comentar a linha que chama o middleware `jwt.ValidateToken` no arquivo
`main.go`.*

## Endpoints
Todos os consumos HTTP dessa aplicação podem ser encontrados dentro dos arquivos `account.http` e 
`account-codes.http` dentro do diretório `/api`.
### Pares login/senha
#### Criar login/senha
- **Endpoint**: `/account`
- **Método**: POST
- **Descrição**: Cria um novo par login/senha.
- **Corpo da requisição**:
```json
{
  "name":"Facebook", 
  "login": "joao.silva@example.com",
  "password": "Senhaforte123"
}
```
- **Resposta**: `201 Created` Detalhes do login/senha.
```json
{
   "message": "Conta criada com sucesso",
   "data": {
      "id": "1",
      "name": "Facebook",
      "login": "joao.silva@example.com",
      "password": "f793f87174464478d3fa0c241520d068e858b3758065730c353avsf625c6d9fcb45d2f6ec21049",
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": "0001-01-01T00:00:00Z"
   }
}
```

#### Alterar login/senha
- **Endpoint**: `/account/:id`
- **Método**: PUT
- **Descrição**: Altera um par login/senha.
- **Corpo da requisição**:
```json
{
  "name":"Instagram", 
  "login": "joao.silva@blank.com",
  "password": "senhaFraca321"
}
```
- **Resposta**: `200 ok` Detalhes do login/senha.
```json
{
   "message": "Conta atualizada com sucesso",
   "data": {
      "id": "1",
      "name": "Instagram",
      "login": "joao.silva@blank.com",
      "password": "f793f87174464478d3fa0c241520d068e858b3758065730c353avsf625c6d9fcb45d2f6ec21049",
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": "0001-01-01T00:00:00Z"
   }
}
```

#### Deletar login/senha
- **Endpoint**: `/account/:id`
- **Método**: DELETE
- **Descrição**: Deleta um par login/senha.
- **Resposta**: `200 ok` Detalhes do login/senha.
```json
{
   "message": "Conta deletada com sucesso",
   "data": null
}
```

#### Listar todos os login/senha
- **Endpoint**: `/account`
- **Método**: GET
- **Descrição**: Lista todos os pares login/senha.
- **Resposta**: `200 ok` Detalhes dos login/senha.
```json
{
   "message": "Contas encontradas com sucesso",
   "data": [
      {
         "id": "1",
         "name": "Instagram",
         "login": "joao.silva@blank.com",
         "password": "SenhaForte123",
         "created_at": "0001-01-01T00:00:00Z",
         "updated_at": "0001-01-01T00:00:00Z"
      },
      {
         "id": "2",
         "name": "Facebook",
         "login": "joao.silva@example.com",
         "password": "senhaFraca321",
         "created_at": "0001-01-01T00:00:00Z",
         "updated_at": "0001-01-01T00:00:00Z"
      }
   ]
}
```

#### Listar todos os login/senha por nome
- **Endpoint**: `/account/:name`
- **Método**: GET
- **Descrição**: Lista todos os pares login/senha com aquele nome.
- **Resposta**: `200 ok` Detalhes dos login/senha.
```json
{
   "message": "Contas encontradas com sucesso",
   "data": [
      {
         "id": "1",
         "name": "Instagram",
         "login": "joao.silva@blank.com",
         "password": "SenhaForte123",
         "created_at": "0001-01-01T00:00:00Z",
         "updated_at": "0001-01-01T00:00:00Z"
      },
      {
         "id": "2",
         "name": "Instagram",
         "login": "joaozinho.silvao@example.com",
         "password": "senhaFraca321",
         "created_at": "0001-01-01T00:00:00Z",
         "updated_at": "0001-01-01T00:00:00Z"
      }
   ]
}
```

### Palavras chave
#### Criar conjunto de palavras chave
- **Endpoint**: `/account-codes`
- **Método**: POST
- **Descrição**: Cria um novo conjunto de palavras chave.
- **Corpo da requisição**:
```json
{
  "name":"Trust wallet", 
  "codes": "arroz, feijao"
}
```
- **Resposta**: `201 Created` Detalhes do conjunto de palavras chave.
```json
{
   "message": "Conta criada com sucesso",
   "data": {
      "id": "1",
      "name": "Trust wallet",
      "codes": "9d94cf01e025f4221f75834c7561510ad762461a3ef08bfbe24c11ce442c38de6a81d2913e611af1626fb297ef",
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": "0001-01-01T00:00:00Z"
   }
}
```

#### Alterar conjunto de palavras chave
- **Endpoint**: `/account-codes/:id`
- **Método**: PUT
- **Descrição**: Altera um conjunto de palavras chave.
- **Corpo da requisição**:
```json
{
  "name":"metamask", 
  "codes": "strogonoff, salame"
}
```
- **Resposta**: `200 ok` Detalhes do conjunto de palavras chave.
```json
{
   "message": "Conta atualizada com sucesso",
   "data": {
      "id": "1",
      "name": "metamask",
      "codes": "9d94cf01e025f4221f75834c7561510ad762461a3ef08bfbe24c11ce442c38de6a81d2913e611af1626fb297ef",
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": "0001-01-01T00:00:00Z"
   }
}
```

#### Deletar conjunto de palavras chave
- **Endpoint**: `/account-codes/:id`
- **Método**: DELETE
- **Descrição**: Deleta um conjunto de palavras chave.
- **Resposta**: `200 ok` Detalhes do conjunto de palavras chave.
```json
{
   "message": "Conta deletada com sucesso",
   "data": null
}
```

#### Listar todos os conjunto de palavras chave
- **Endpoint**: `/account-codes`
- **Método**: GET
- **Descrição**: Lista todos os pares conjunto de palavras chave.
- **Resposta**: `200 ok` Detalhes dos conjunto de palavras chave.
```json
{
   "message": "Contas encontradas com sucesso",
   "data": [
      {
         "id": "1",
         "name": "Coinbase",
         "codes": "arroz, feijao",
         "created_at": "0001-01-01T00:00:00Z",
         "updated_at": "0001-01-01T00:00:00Z"
      },
      {
         "id": "2",
         "name": "Metamask",
         "codes": "strogonoff, salame",
         "created_at": "0001-01-01T00:00:00Z",
         "updated_at": "0001-01-01T00:00:00Z"
      }
   ]
}
```

#### Listar todos os conjunto de palavras chave por nome
- **Endpoint**: `/account-codes/:name`
- **Método**: GET
- **Descrição**: Lista todos os pares conjunto de palavras chave com aquele nome.
- **Resposta**: `200 ok` Detalhes dos conjunto de palavras chave.
```json
{
   "message": "Contas encontradas com sucesso",
   "data": [
      {
         "id": "1",
         "name": "Metamask",
         "codes": "arroz, feijao",
         "created_at": "0001-01-01T00:00:00Z",
         "updated_at": "0001-01-01T00:00:00Z"
      },
      {
         "id": "2",
         "name": "Metamask",
         "codes": "strogonoff, salame",
         "created_at": "0001-01-01T00:00:00Z",
         "updated_at": "0001-01-01T00:00:00Z"
      }
   ]
}
```

## Funcionalidades
- [x] Armazena seus logins e senhas de forma criptografada
- [x] Armazena palavras chaves, como keys de recuperação, seed phrases