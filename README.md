# cadUserInGo

## Desafio - Criando uma API REST

### Introdução

Faaala Dev,

Este projeto é parte de um desafio prático para reforçar os conceitos aprendidos nos módulos da trilha Go. O objetivo é construir uma API RESTful que realiza operações CRUD in-memory, aplicando boas práticas de desenvolvimento e organização de código.

Este desafio é opcional e não possui correção oficial, mas é uma excelente oportunidade para praticar e consolidar conhecimentos. Lembre-se: **tenha calma** e **acredite no seu processo**. 💜

---

## Progresso dos Endpoints

| Método | URL               | Descrição                                                                 | Status       |
|--------|--------------------|---------------------------------------------------------------------------|--------------|
| POST   | /api/users         | Cria um usuário usando as informações enviadas no corpo da requisição.   | ✔ Implementado |
| GET    | /api/users         | Retorna um array de usuários.                                            | ✔ Implementado |
| GET    | /api/users/:id     | Retorna o objeto do usuário com o id especificado.                       | ✔ Implementado |
| DELETE | /api/users/:id     | Remove o usuário com o id especificado e retorna o usuário deletado.     | ✔ Implementado |
| PUT    | /api/users/:id     | Atualiza o usuário com o id especificado usando dados do corpo da requisição. Retorna o usuário modificado. | ✔ Implementado |

---

## Esquema do Usuário

Cada recurso de usuário segue o seguinte esquema:

```json
{
  "id": "",                     // UUID, obrigatório
  "first_name": "Jane Doe",     // String, obrigatório (len >= 2 && len <= 20)
  "last_name": "Jane Doe",      // String, obrigatório (len >= 2 && len <= 20)
  "biography": "Tendo diversão" // String, obrigatório (len >= 20 && len <= 450)
}
```

## Banco de Dados em Memória

Como ainda não cobrimos como lidar com bancos de dados em Go, este projeto utiliza um "banco de dados" em memória, implementado como um hash map, onde o `id` é a chave. As funções implementadas incluem:

- **FindAll**: Retorna a lista de usuários (ou array vazio).
- **FindById**: Retorna o usuário com o id especificado (ou `null` se o id não existir).
- **Insert**: Adiciona um novo usuário e retorna o usuário recém-criado.
- **Update**: Atualiza um usuário existente e retorna o usuário atualizado (retorna erro caso o id não exista).
- **Delete**: Remove um usuário e retorna o usuário deletado.

---

## Especificações dos Endpoints

### POST /api/users

- **Regras**:
    - Se faltar alguma propriedade obrigatória (`first_name`, `last_name` ou `biography`):
        - Retorna `400 Bad Request` com a mensagem: `{"message": "Please provide FirstName LastName and bio for the user"}`.
    - Se os dados forem válidos:
        - Salva o usuário no banco de dados.
        - Retorna `201 Created` com o usuário criado.
    - Em caso de erro ao salvar:
        - Retorna `500 Internal Server Error` com a mensagem: `{"message": "There was an error while saving the user to the database"}`.

### GET /api/users

- **Regras**:
    - Retorna todos os usuários cadastrados.
    - Em caso de erro ao recuperar os dados:
        - Retorna `500 Internal Server Error` com a mensagem: `{"message": "The users information could not be retrieved"}`.

### GET /api/users/:id

- **Regras**:
    - Se o usuário com o id especificado não for encontrado:
        - Retorna `404 Not Found` com a mensagem: `{"message": "The user with the specified ID does not exist"}`.
    - Em caso de erro ao recuperar os dados:
        - Retorna `500 Internal Server Error` com a mensagem: `{"message": "The user information could not be retrieved"}`.

### DELETE /api/users/:id

- **Regras**:
    - Se o usuário com o id especificado não for encontrado:
        - Retorna `404 Not Found` com a mensagem: `{"message": "The user with the specified ID does not exist"}`.
    - Em caso de erro ao remover o usuário:
        - Retorna `500 Internal Server Error` com a mensagem: `{"message": "The user could not be removed"}`.

### PUT /api/users/:id

- **Regras**:
    - Se o usuário com o id especificado não for encontrado:
        - Retorna `404 Not Found` com a mensagem: `{"message": "The user with the specified ID does not exist"}`.
    - Se faltar alguma propriedade obrigatória (`first_name`, `last_name` ou `biography`):
        - Retorna `400 Bad Request` com a mensagem: `{"message": "Please provide name and bio for the user"}`.
    - Em caso de erro ao atualizar os dados:
        - Retorna `500 Internal Server Error` com a mensagem: `{"message": "The user information could not be modified"}`.
    - Se os dados forem válidos:
        - Atualiza o usuário no banco de dados.
        - Retorna `200 OK` com o usuário atualizado.

---

## Como Testar

1. Clone o repositório:
     git clone https://github.com/seu-usuario/cadUserInGo.git
2. Inicie o servidor:
     go run main.go

3. Use ferramentas como Postman, Insomnia ou cURL para testar os endpoints.

---
## Conclusão

Este projeto é uma excelente oportunidade para praticar o desenvolvimento de APIs RESTful em Go. À medida que os endpoints forem implementados, o progresso será atualizado na tabela acima. 🚀 ```
