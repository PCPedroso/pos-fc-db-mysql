# Conectar ao banco de dados
docker-compose exec mysql bash

# Criando a tabela
create table produto (id varchar(255), nome varchar(80), preco decimal(10,2), primary key (id));