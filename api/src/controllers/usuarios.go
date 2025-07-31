package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
)

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro) //recebe 3 parâmetros: resposta, código de status e mensagem de erro
		return
	}

	var usuario modelos.Usuario //criando usuario que está no pacote de modelos
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro) //recebe 3 parâmetros: resposta, código de status e mensagem de erro
		return
	}

	db, erro := banco.Conectar() //criando conexao com o banco de dados
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro) //recebe 3 parâmetros: resposta, código de status e mensagem de erro
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db) //criando um novo repositorio de usuarios
	usuario.ID, erro = repositorio.Criar(usuario)             //chamando o metodo criar do repositorio de usuarios
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro) //recebe 3 parâmetros: resposta, código de status e mensagem de erro
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario) //enviando resposta para o usuario
}

// BuscarUsuarios busca todos os usuários salvos no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários!"))
}

// BuscarUsuario busca um usuário salvo no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário!"))
}

// AtualizarUsuario altera as informações de um usuário no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário!"))
}

// DeletarUsuario exclui as informações de um usuário do banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}
