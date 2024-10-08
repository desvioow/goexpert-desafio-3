# Go Expert Desafio 3
Desafio Go Expert sobre clean arch.

## Descrição do Desafio
Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
  Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

## Subindo a Aplicação
Para subir a aplicação é preciso primeiro iniciar os containers definidos no docker-compose.yaml. Este irá iniciar os containeres
mySQL, rabbitMQ e já criar e popular a tabela orders no container mySQL com alguns dados dummy.
```
docker-compose up
```

Após subidos os containeres via docker-compose, iniciar a aplicação principal. Esta irá subir 3 servidores locais:
1) rest na porta 8000
2) gRPC na porta 50051
3) graphQL na porta 8080
```
go run main.go
```

### Testando APIs
#### Criação de novos pedidos
##### rest - POST /order
Usar algum cliente REST ou IDE com suporte para arquivos .http
e executar a request que se encontra no arquivo apis.http
![image](https://github.com/user-attachments/assets/5aa2a34c-21e5-4425-987d-1610d74fb046)


O resultado esperado é http 200 e a criação de um novo pedido via rest na tabela orders conforme os parâmetros escolhidos.

##### gRPC - CreateOrder
Usar algum cliente gRPC como qualquer um listado aqui [gRPC clients](https://github.com/grpc-ecosystem/awesome-grpc?tab=readme-ov-file#tools)
e testar o serviço CreateOrder definido no arquivo order.proto.
![image](https://github.com/user-attachments/assets/1f881672-f800-4ea0-9cf3-e73293c26f76)


O resultado esperado é sucesso na request e a criação de um novo pedido via gRPC na tabela orders conforme os parâmetros escolhidos.

##### graphQL - CreateOrder
Acessar o graphQL playground usando um browser no endpoint http://localhost:8080/.
Executar a mutation CreateOrder conforme o exemplo a seguir:
```
mutation CreateOrder {
    createOrder(input: {id: "graphQL_order", Price: 50.0, Tax: 25.0}) {
        id
        Price
        Tax
        FinalPrice
    }
}
```
![image](https://github.com/user-attachments/assets/0e9a2482-6446-4554-b086-848b42060c8e)


O resultado esperado é sucesso na request e a criação de um novo pedido via graphQL na tabela orders conforme os parâmetros escolhidos.
(imagem_database)
#### Listar pedidos criados
##### rest - GET /order
Usar algum cliente REST ou IDE com suporte para arquivos .http
e executar a request que se encontra no arquivo apis.http
![image](https://github.com/user-attachments/assets/337e55cc-3ee6-4228-b9e1-1c402de8c873)

O resultado esperado é http 200 e a listagem de todos pedidos criados na tabela orders.

##### gRPC - ListOrders
Usar algum cliente gRPC como qualquer um listado aqui [gRPC clients](https://github.com/grpc-ecosystem/awesome-grpc?tab=readme-ov-file#tools)
e testar o serviço ListOrders definido no arquivo order.proto.
![image](https://github.com/user-attachments/assets/155c4dae-ca52-4166-9188-b4e601ed1958)


O resultado esperado é sucesso na request e a listagem de todos pedidos criados na tabela orders.

##### graphQL - ListOrders
Acessar o graphQL playground usando um browser no endpoint http://localhost:8080/.
Executar a mutation CreateOrder conforme o exemplo a seguir:
```
query ListAllOrders {
  orders {
    id
    Price
    Tax
    FinalPrice
  }
}
```
![image](https://github.com/user-attachments/assets/790358d2-9a61-494c-98d7-be4b9a86d1b2)


O resultado esperado é sucesso na request e a listagem de todos pedidos criados na tabela orders.

### Database
#### Antes
![image](https://github.com/user-attachments/assets/5f5ec3c1-b159-4e90-b4d3-62bfcc9712d3)

#### Após criações de novos pedidos via Rest, gRPC e GraphQL
![image](https://github.com/user-attachments/assets/fb34def7-c987-4be6-ba46-69874d6c0e5b)
