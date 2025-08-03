package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da API
type Rota struct {
	URI                string //Endereço da rota
	Metodo             string //Método HTTP
	Funcao             func(http.ResponseWriter, *http.Request) //Função que irá lidar com a requisição
	RequerAutenticacao bool //Se precisa estar logado
}

//Configurar coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios //Pega as rotas definidas
	rotas = append(rotas, rotaLogin) //Adiciona a rota de login

	for _, rota := range rotas { //Para cada rota...
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo) //Registra no router: "quando vier requisição X, chame a função Y"
	}

	return r //Retorna o router configurado
}
