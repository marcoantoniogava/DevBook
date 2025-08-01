package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

// Usuarios representa um repositório de usuarios
type Usuarios struct {
	db *sql.DB //struct que vai receber o banco de dados
}

// NovoRepositorioDeUsuarios é uma função que recebe o banco de dados e retorna um novo repositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuario no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) { //Método vai estar dentro do repositório de usuários, vai criar um usuário, vai receber um parâmetro (um modelo de usuario) e vai retornar um id e um erro
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

// Buscar traz todos os usuarios que atendem um filtro de nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) { //recebe um nome ou nick e retorna uma lista de usuarios e um erro
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?", //vai procurar por um cara que tenha o nome ou nick igual ao nomeOuNick
		nomeOuNick, nomeOuNick, //são as duas "?"
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close() //fecha a conexão com as linhas

	var usuarios []modelos.Usuario //cria uma lista de usuarios

	for linhas.Next() { //itera sobre as linhas
		var usuario modelos.Usuario //cria um usuario para cada linha

		if erro = linhas.Scan(// vai ler cada linha e atribuir os valores a cada campo do usuario
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario) //adiciona à lista de usuarios, o usuario que acabou de ser lido
	}

	return usuarios, nil
}
