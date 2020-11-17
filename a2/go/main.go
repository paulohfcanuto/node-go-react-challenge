package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"io/ioutil"
	"log"
	"net/http"
)

type Cep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(
			createQueryType(createCepType()),
		),
	})
	if err != nil {
		log.Fatalf("Falha ao criar um novo schema, erro: %v", err)
	}
	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &schema,
	})
	http.Handle("/graphql", handler)
	log.Println("Iniciando o servidor em: http://localhost:3000/graphql")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func createQueryType(cepType *graphql.Object) graphql.ObjectConfig {
	return graphql.ObjectConfig{Name: "QueryType", Fields: graphql.Fields{
		"viacepAPI": &graphql.Field{
			Type: cepType,
			Args: graphql.FieldConfigArgument{
				"cep": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				cep := p.Args["cep"]
				v, _ := cep.(string)
				log.Printf("fetching cep: %s", v)
				return fetchCepByiD(v)
			},
		},
	}}
}

func createCepType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Cep",
		Fields: graphql.Fields{
			"cep": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"logradouro": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"complemento": &graphql.Field{
				Type: graphql.String,
			},
			"bairro": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"localidade": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"uf": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"ibge": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"gia": &graphql.Field{
				Type: graphql.String,
			},
			"ddd": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"siafi": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	})
}

func fetchCepByiD(cep string) (*Cep, error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s: %s", "Não foi possível realizar a requisição", resp.Status)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Não foi possível ler os dados")
	}
	result := Cep{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, errors.New("Não foi possível desserializar os dados")
	}
	return &result, nil
}
