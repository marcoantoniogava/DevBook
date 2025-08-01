package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{ //Definindo rotas
	{
		URI:                "/usuarios", //Endereço
		Metodo:             http.MethodPost, //POST
		Funcao:             controllers.CriarUsuario, //Função executada
		RequerAutenticacao: false, //Não precisa login
	},
	{
		URI:    "/usuarios",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI:    "/usuarios/{usuarioId}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:    "/usuarios/{usuarioId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:    "/usuarios/{usuarioId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
}
