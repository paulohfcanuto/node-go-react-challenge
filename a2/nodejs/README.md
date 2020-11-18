# Backend node.js and GraphQL to Wrapping a REST API

## Observações:
 * Foi utilizado um cache de cada chamada para inserir o retorno em um redis e posteriormente retornar esses dados.
 * Utilizei o client http fetch por ser algo simples, em um projeto de produção eu utilizaria o axios ou algum client com mais recursos
 * Algumas dependencias do graphql estão deprecadas, em caso de projeto de produção eu migraria para o graphql-tools
 * Ficou faltando um tratamento para erros na api de cep, caso seja informado um cep que não existe, por exemplo.
 
## Como rodar local:
```shell script
 docker-compose up redis
 npm i
 npm run start:dev
```

## Playground:
 * http://localhost:8080/graphql

## Exemplo de query para consulta de cep:
```text
query {
  viacepAPI(cep: "88807010") {
    cep
    logradouro
    complemento
    bairro
    localidade
    uf
    ibge
    gia
    ddd
    siafi
  }
}
```

## Como rodar os testes:
 ```shell script  
  npm run test
 ```
