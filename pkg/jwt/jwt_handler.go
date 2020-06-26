package jwt

import (
	"io/ioutil"
	// "crypto/rsa"
	"log"
	jwt "github.com/dgrijalva/jwt-go"
	// request "github.com/dgrijalva/jwt-go/request"
	"ca-tech-dojo/internal/game_api/user/models"
)

type JwtHandler struct {
	Token *jwt.Token
}

func NewJwtHandler() (jwtHandler *JwtHandler){
	jwtHandler = new(JwtHandler)
	return
}

// var varifyKey *rsa.PublicKey  // 公開鍵
// var signKey   *rsa.PrivateKey // 秘密鍵

// トークンを生成して返す
func (handler *JwtHandler) Create(u models.User) (user models.User, err error) {
	signBytes, err := ioutil.ReadFile("./docs/key/private-key.pem") // 秘密鍵の読み込み
	if err != nil {
		panic(err)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes) // 秘密鍵をパース
    if err != nil {
		log.Print("error is ", err.Error())
        panic(err)
	}

	// 署名アルゴリズムを指定してトークン型を生成
	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims) // クレームを設定
	claims["name"] = u.Name

	tokenString, err := token.SignedString(signKey) // 署名
	if err != nil {
		log.Print("error is ", err.Error())
		return
	}

	u.Token = tokenString
	user = u

	return
}