# Sistema de marcação de consultas
Desafio final da disciplina Backend 3 do curso Certified Tech Develpor da Digital House (2° ano -  especialização em Backend).
Implementação de uma API com objetivo de administrar a marcação de consultas para uma clínica odontológica.

## Especificações

 - Administração de dados de dentistas: listar, adicionar, alterar ou excluir dentistas. Registrar os seus sobrenome, nome e matrícula. Desenvolvimento de um CRUD para a entidade Dentista. 
    - POST: adicionar dentista. 
    - GET: trazer dentista pelo seu ID.
    - PUT: atualizar dentista.
    - PATCH: atualizar um dentista através de algum dos seus campos. 
    - DELETE: excluir dentista.

- Administração de dados dos pacientes: listar, adicionar, alterar ou excluir pacientes. Registrar os seus sobrenome, RG, nome e data de cadastro. Desenvolvimento de um CRUD para a entidade Paciente.
    - POST: adicionar dentista. 
    - GET: trazer paciente pelo seu ID.
    - PUT: atualizar paciente.
    - PATCH: atualizar um paciente através de algum dos seus campos. 
    - DELETE: excluir paciente.

- Marcação de consulta: deve ser possível atribuir a um paciente uma consulta com um dentista em uma determinada data e hora, e também adicionar uma descrição à consulta.Desenvolvimento de um CRUD para a entidade Consulta.
    - POST: adicionar consulta.
    - GET: trazer consulta pelo seu ID.
    - PUT: atualizar consulta.
    - PATCH: atualizar consulta por algum dos seus campos.
    - DELETE: excluir consulta.
    - GET: trazer consulta pelo RG do paciente. Deve conter o detalhamento da consulta (Data-Hora, descrição, Paciente e Dentista).
    
## Requerimentos técnicos
Desenvolvimento em design orientado a pacotes
- Camada/domínio de entidades de negócio.
- Camada/domínio de acesso a dados (Repository).
- Camada de acesso a dados (banco de dados).
- Camada/domínio service.
- Camada/domínio handler.

## Tecnologias utilizadas
- Golang
- MySql
- Docker
- Postman

## Como executar o projeto

Pré-requisitos: Golang e MySql

```bash
# clonar repositório
git clone git@github.com:vanessa-maganhoto/desafio-final-esp-backend3.git

# executar o projeto
go run cmd/server/main.go
```

## Autores
[Mateus Benites](https://www.linkedin.com/in/mateus-benites-dias/),
[Vanessa Matos](https://www.linkedin.com/in/vanessaammatos) e 
[Wirley Almeida](https://www.linkedin.com/in/wirley-almeida-dev/)
