package controllers

import (
	"net/http"
	"webapp/src/utils"
)

//CarregarTelaDeLogin vai carregar a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

//CarregarPaginaDeCadastroDeUsuario vai carregar a página de cadastro de usuário
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request)  {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}
