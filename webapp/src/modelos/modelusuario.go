package modelos

import (
	"net/http"
	"time"
)

// Usuario representa uma pessoa utilizando a rede social
type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Nick        string       `json:"nick"`
	Email       string       `json:"email"`
	CriadoEm    time.Time    `json:"criadoEm"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}

// BuscarUsuarioCompleto faz 4 requisições na API para montar o usuário
func BuscarUsuarioCompleto(usuarioID uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarDadosDoUsuario(canalUsuario, usuarioID, r)
	go BuscarSeguidores(canalSeguidores, usuarioID, r)
	go BuscarSeguindo(canalSeguindo, usuarioID, r)
	go BuscarPublicacoes(canalPublicacoes, usuarioID, r)

}

func BuscarDadosDoUsuario(canal <-chan Usuario, usuarioID uint64, r *http.Request) {

}

func BuscarSeguidores(canal <-chan []Usuario, usuarioID uint64, r *http.Request) {

}

func BuscarSeguindo(canal <-chan []Usuario, usuarioID uint64, r *http.Request) {

}

func BuscarPublicacoes(canal <-chan []Publicacao, usuarioID uint64, r *http.Request) {

}
