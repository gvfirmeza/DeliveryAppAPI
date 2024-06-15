# API McRonalds

## Como Utilizar

Clone o repositório para sua máquina.
Certifique-se de ter o Go instalado.
No diretório do projeto, execute o comando go run main.go para iniciar o servidor.
O servidor será iniciado em http://localhost:8080.
Utilize os comandos descritos abaixo para interagir com a API.

## Rotas

| Rota            | Método | Funcionalidade                       |
|-----------------|--------|--------------------------------------|
| /abrir          | POST   | Abre a loja                          |
| /fechar         | POST   | Fecha a loja                         |
| /produto        | POST   | Adiciona um novo produto             |
| /produto/{id}   | DELETE | Remove um produto pelo ID            |
| /produto/{id}   | GET    | Busca um produto pelo ID             |
| /produtos       | GET    | Lista todos os produtos disponíveis  |
| /pedido         | POST   | Inclui um novo pedido                |
| /pedidos        | GET    | Exibe todos os pedidos abertos       |
| /pedido/{id}    | POST   | Expede um pedido pelo ID             |
| /metricas       | GET    | Obtém as métricas do sistema         |


## Exemplos para Rotas:

### Adicionar Produtos - POST /produto

JSON para ser enviado:
```json
{
  "nome": "Produto Teste",
  "descricao": "Este é um produto de exemplo",
  "valor": 99.99
}
```
Resposta Esperada:
```
Novo produto adicionado com ID: 1
```

### Listar todos Produtos - GET /produto

Resposta Esperada:
```json
{
    "ID": 1,
    "nome": "Produto Teste",
    "descricao": "Este é um produto de exemplo",
    "valor": 99.99
}
```

### Buscar Produtos por ID - GET /produto/{id}

Exemplo de envio:
```
http://localhost:8080/produto/1
```
Resposta Esperada:
```json
{
    "ID": 1,
    "nome": "Produto Teste",
    "descricao": "Este é um produto de exemplo",
    "valor": 99.99
}
```

### Remover Produtos - DELETE /produto/{id}

Exemplo de envio:
```
http://localhost:8080/produto/1
```
Resposta Esperada:
```
Produto com ID 1 removido com sucesso
```

### Adicionar Pedidos - POST /pedido

JSON para ser enviado:
```json
{
  "delivery": false,
  "produtos": [
    {
        "id" : 1
    }
  ]
}
```
Resposta Esperada:
```
Novo pedido incluído com ID: {id}
```

### Listar Pedidos - GET /pedidos

Para ordenar usando Bubble Sort: /pedidos?ordenacao=bubblesort
Para ordenar usando Quick Sort: /pedidos?ordenacao=quicksort
Para a lista sem ordenação (ordem original pelo ID): /pedidos

## Integrantes

- Gabriel Firmamento
- Guilherme Firmeza
- Paulo Henrique
- Jorge Felipe Magarão
- Cesar Viana
