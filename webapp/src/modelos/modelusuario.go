package modelos

import "time"

// Usuario representa uma pessoa utilizando a rede social
type Usuario struct {
	ID       uint64 `json:"id"`
	Nome     string `json:"nome"`
	Nick     string `json:"nick"`
	Email    string `json:"email"`
	CriadoEm time.Time `json:"criadoEm"`
	Seguidores []Usuario `json:"seguidores"`
	Seguindo   []Usuario `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}
