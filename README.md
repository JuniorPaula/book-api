# Golang Book API

![GitHub](https://img.shields.io/badge/-golang-blue?logo=github&style=flat)
![MySQL](https://img.shields.io/badge/-mysql-orange?logo=mysql&style=flat)
![Docker](https://img.shields.io/badge/-docker-blue?logo=docker&style=flat)
![Git](https://img.shields.io/badge/-git-black?logo=git&style=flat)

O projeto "Book API" é uma API de cadastro de livros desenvolvida em Golang. A API permite realizar as principais operações do CRUD (Create, Read, Update, Delete) para gerenciar um catálogo de livros. O banco de dados utilizado é o MySQL, que está configurado em um container Docker usando Docker Compose.

## Rotas da API

A API possui as seguintes rotas para executar as operações do CRUD:

- **GET /book**: Retorna todos os livros cadastrados.
- **GET /book/{id}**: Retorna um livro específico com base no ID.
- **POST /book**: Cria um novo livro no catálogo.
- **PUT /book/{id}**: Atualiza as informações de um livro existente.
- **DELETE /book/{id}**: Remove um livro do catálogo.

## Dependências

- Golang: Certifique-se de ter o Golang instalado em sua máquina.
- Docker: Instale o Docker para criar e gerenciar os containers.
- Docker Compose: O Docker Compose é usado para orquestrar o container do MySQL.

Certifique-se de ter todas as dependências instaladas corretamente antes de executar o projeto.

## Executando o Projeto

1. Clone este repositório em sua máquina local.
2. Crie um arquivo `.env`com base no arquivo `.env.example`.
3. Navegue até o diretório raiz do projeto.
4. Execute o comando `docker-compose up -d` para iniciar o container do MySQL em segundo plano.
5. Execute o comando `go run cmd/main.go` para iniciar a API.
6. Acesse `http://localhost:8081/book` para interagir com a API por meio de uma ferramenta como o Postman ou um navegador.

## Personalização e Aprimoramento

Este projeto de API de cadastro de livros pode ser personalizado e aprimorado de várias maneiras. Você pode adicionar novos endpoints, implementar autenticação e autorização, adicionar validações de entrada, entre outros recursos.

Sinta-se à vontade para explorar e modificar o código-fonte de acordo com suas necessidades específicas.

## Contribuição

Contribuições são bem-vindas! Se você encontrar algum problema, tiver sugestões ou quiser contribuir de alguma forma para este projeto, sinta-se à vontade para enviar um pull request.

## Considerações Finais

O projeto "Book API" é uma API de cadastro de livros desenvolvida em Golang, com suporte ao MySQL usando Docker. Esperamos que este projeto seja útil e inspire você a criar suas próprias APIs poderosas.

Se tiver alguma dúvida ou precisar de assistência, não hesite em entrar em contato. Divirta-se desenvolvendo com o projeto Book API!