package middlewares

import (
	"fmt"
	"log"
	"net/http"
)
//middleware -> camada entre a requisição e a resposta
//Logger escreve informações da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host) //exibindo informações da requisição no terminal
		next(w, r)
	}
}

//Autenticar verifica se o usuário fazendo a requisição está autenticando
func Autenticar(next http.HandlerFunc) http.HandlerFunc { //handlerfunc recebe um w responsewriter e um r *request
	return func(w http.ResponseWriter, r *http.Request) { //vai chamar a função que vai validar o token, e vai executar o q veio no parametro (next)
		fmt.Println("Validando...")
		next(w, r)
	}
}
