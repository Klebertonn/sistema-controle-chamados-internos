# sistema-controle-chamados-internos

## Objetivo

Sistema web tem o objetivo de administrar os chamados internos de TI. Ele será responsável por 
- Cadastrar chamado
- Consulta de chamados
- Atualizar chamdos
- Atribuição de responsaveis
 Controle de prioriade e Status
 Tornando dinâmico a distribuição e a análise das solicitações.

## Tecnologias
### Backend
-Gin
-GO
## Banco de Dados
-SQL Server
### Frontend/ Web
-HTML

-CSS

-JS

-Bootstrap5

## Funcionalidades
-Login

-Abertura de chamados

-Atribuiçao automática 

-Seleção Manual

-Atualização de Status

-Consultas de CHamados


## Arquitetura

O projeto foi desenvolvido utiliznado uma arquiterura de camadas

```text
 Handler
    |
 Service
    |
 Respository
    |
 SQL Server     
```
### Handler
- Responavel requisiçôes HTTP

- Validar entrada de dados

- Retornar respostas HTTP

### Service
- Regras de negóscio

- Validações

- Fluxo da aplicação

### Repository
- Responavel comunicação com o banco

- Queries SQL

- Persistencia dos dados



## Estrutura
```text
sistema-controle-chamados
 |
 |__cmd
 |    |__server
 |         |__main
 |
 |__internal
 |      |__hadlers
 |      |__service
 |      |__repositories
 |      |__models
 |
 |__web
 |    |__templates
 |
 |__database
 |
 |_docs
 
```
## Motivo da Escolha da Arquitetura
A arquitetura em camadas foi escolhida para a separação das responsabilidades, facilitando a manutenção e a adoção dos testes unitários, com baixa dependência entre os componentes e na escalabilidade futura. Além de tudo, essa arquitetura permite alterar o banco de dados, regras de negócio ou interface sem impactar toda a aplicação.
    
## Endpoints
   ### Criar Chamado
   Rota:

 ```http
    POST/chamados
```
Body:
```json
   "titulo": "Erro na PC",
  "descricao": "Impressora nao imprime",
  "prioridade": "ALTA",
  "solicitante": "Analista",
  "responsavelId": 1
```
Resposta:
```json
   {
    "message":"Chamado criado com sucesso"
   }
```
cURL Postman
```json
    postman request POST 'http://localhost:8080/chamados' \
  --header 'Content-Type: application/json' \
  --body '{
  "titulo": "Erro na PC",
  "descricao": "Impressora nao imprime",
  "prioridade": "ALTA",
  "solicitante": "Analista",
  "responsavelId": 1
   }'
```
  ### Lista Chamados
   Rota:

 ```http
    GET/chamados
```
Resposta:
```json
   {  
     "ordemID":1,
     "titulo": "Erro na PC",
     "descricao": "Impressora nao imprime",
     "prioridade": "ALTA",
     "solicitante": "Analista",
     "responsavelId": 1
   }

   {  
     "ordemID":2,
     "titulo": "Erro na PC",
     "descricao": "Impressora nao imprime",
     "prioridade": "ALTA",
     "solicitante": "Analista",
     "responsavelId": 1
   }
```
cURL Postman
```json
    postman request 'http://localhost:8080/chamados'
```

 ### Buscar Chamados por ID
   Rota:

```http
    GET/chamados/:id
```
```http
    GET/chamados/1
```
Resposta:
```json
   {  
     "ordemID":1,
     "titulo": "Erro na PC",
     "descricao": "Impressora nao imprime",
     "prioridade": "ALTA",
     "solicitante": "Analista",
     "responsavelId": 1
   }
```
cURL Postman
```json
   postman request 'http://localhost:8080/chamados/1'
```

 ### Atualizar Chamado
   Rota:
```http
    PUT/chamados/:id
```
```http
    GET/chamados/1
```
Body:
```json
   {  
     "ordemID":1,
     "titulo": "Erro na PC",
     "descricao": "Impressora nao imprime",
     "prioridade": "ALTA",
     "solicitante": "Analista",
     "responsavelId": 1
   }
```
Resposta:
```json
   {
    "message":"Chamado atualizado com sucesso"
   }
```

cURL Postman
```json
  postman request PUT 'http://localhost:8080/chamados/1' \
  --header 'Content-Type: application/json' \
  --body '{
  "titulo": "Internet voltou",
  "descricao": "Normalizado",
  "status": "Resolvido",
  "prioridade": "Baixa",
  "responsavelId": 1
}'
```

## Testes
O projeto possui testes unitários na camada da Service
- Criação das chamadas
- Atualização de chamados
- Consulta por ID
- Listagem de chamados
Executar 

```Bash
  go test ./... -v
  go test -v
```

## Como executar o projeto
Clonar 
```Bash
  git clone https://github.com/Klebertonn/sistema-controle-chamados-internos.git
```

Dependências
```Bash
   go mod tidy
```
 Configura Banco
  Atualizar string de conexão em 
    cmd/server/main.go

 Executar
 ```Bash
    go run cmd/server/main.go
``` 


## Referências
 GO
 ```Link
    https://go.dev/tour/welcome/1
 ```

 Arquitetura em Camadas
  ```link
    https://dev.to/yuripeixinho/arquitetura-em-camadas-layered-architecture-a68
  ````
  Banco SQL Server




  ```Link
      https://learn.microsoft.com/pt-br/ssms/sql-server-management-studio-ssms
  ```
