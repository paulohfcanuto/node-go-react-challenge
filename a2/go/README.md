# Backend go and GraphQL to Wrapping a REST API

## Observações:
* Ficou faltando a parte de testes.

## Como rodar local:
```shell script
  go get -t github.com/graphql-go/graphql-go-handler
  go get -t github.com/graphql-go/graphql
  go run main.go
```

## Exemplo de query para consulta de cep com curl:
```shell script
curl --request POST \
  --url http://localhost:3000/graphql \
  --header 'Content-Type: application/json' \
  --data '{
	"operationName": null,
	"variables": {},
	"query": "{\n  viacepAPI(cep: \"88807010\") {\n    cep\n    logradouro\n    complemento\n    bairro\n    localidade\n    uf\n    ibge\n    gia\n    ddd\n    siafi\n  }\n}\n"
}'
```