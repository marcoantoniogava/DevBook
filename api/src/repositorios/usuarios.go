package repositorios

import (
	"api/src/modelos"
	"database/sql"
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
