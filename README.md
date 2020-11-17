# node-go-react-challenge

## A1
### Desenvolver um projeto React-JS simples, que entregue as seguintes funcionalidades:
* Integrar com uma api de Empregos: https://api.graphql.jobs/
* Criar uma página exibindo todas as vagas disponÍveis (Jobs).
* Escolher os campos mais relevantes a serem exibidos
* Paginar o resultado (10 elementos por página)
* Ao clicar em uma das vagas exibidas, mostrar um modal com detalhes dessa vaga
* Fazer uma nova chamada para a query Job, passando a chave da vaga escolhida
* Exibir os campos que achar necessário
* Mostrar os botões de fechar e confirmar (confirmar candidatura) ao clicar em Confirmar
* Chamar uma mutation para a candidatura (subscribe) de acordo com o schema dessa mutation
* Exibir sucesso ou erro de acordo com a chamada da mutation

Entrega:
* O projeto precisa ser feito em React-JS usando Typescript como linguagem
* Não precisa ter tela de login
* O layout é livre



## A2 
### Criar um backend GraphQL, expondo uma query que retorne o resultado de uma API rest
* utilizar a api https://viacep.com.br
* disponibilizar uma query que retorne a consulta de um CEP específico
* o parâmetro de entrada será o número do cep a ser consultado
* a saída deverá ser uma estrutura com as informações que a API do viacep já retorna para a consulta do CEP

Entrega:
* O projeto pode ser feito em NodeJS utilizando o Apollo como servidor GraphQL, pontos extras se fizer o backend em Go.
* Considerar os projetos A1 e A2 como sendo projetos separados
