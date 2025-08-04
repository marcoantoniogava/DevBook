package autenticacao

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//CriarToken retorna um token assinado com as permissões do usuário
func CriarToken(usuarioID uint64) (string, error) { //retorna o token e um erro
	permissoes := jwt.MapClaims{} //map que vai ter as permissoes dentro do token
	permissoes["authorized"] = true //quem tiver esse campo está autorizado
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix() //token expira depois de 6h, .unix converte para a anotação unix de data, devolve a quantidade de segundos que passaram desde o dia 01/01/1970
	permissoes["usuarioId"] = usuarioID //usuario que é dono do token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes) //criando token novo a partir dos claims que foram criados, passa o método de signing HS256, e as permissoes que ele vai ter
	return token.SignedString([]byte("Secret")) //secret, chave que vai ser usada para fazer a assinatura e garantir a autenticidade do token
}
