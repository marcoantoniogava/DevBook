package rotas

import (
	"net/http"
)

// Rota representa todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request) //Função que irá lidar com a requisição
	RequerAutenticacao bool
}
