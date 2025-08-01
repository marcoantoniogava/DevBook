package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body) //Lê o corpo da requisição
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro) //recebe 3 parâmetros: resposta, código de status e mensagem de erro
		return
	}

	var usuario modelos.Usuario //criando usuario que está no pacote de modelos
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro) //recebe 3 parâmetros: resposta, código de status e mensagem de erro
		return
	}

	if erro = usuario.Preparar(); erro != nil {
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
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario")) //converte a string para minuscula, vai trazer tudo que tiver na query (rota) e vai pegar o valor do campo usuario
	db, erro := banco.Conectar() //conecta ao banco de dados
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro) //chama a função erro lá do respostas.go e envia o erro para o usuario
		return
	}
	defer db.Close() //fecha a conexão com o banco de dados

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.Buscar(nomeOuNick)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro) //chama a função erro lá do respostas.go e envia o erro para o usuario
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
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
