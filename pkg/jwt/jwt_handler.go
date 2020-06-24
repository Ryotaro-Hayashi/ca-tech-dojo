package jwt

import (
	"io/ioutil"
	// "crypto/rsa"
	"log"
	jwt "github.com/dgrijalva/jwt-go"
	// request "github.com/dgrijalva/jwt-go/request"
)

type JwtHandler struct {
	Token *jwt.Token
}

func NewJwtHandler() (jwtHandler JwtHandler){
	jwtHandler = JwtHandler{}
	return
}

// var varifyKey *rsa.PublicKey  // 公開鍵
// var signKey   *rsa.PrivateKey // 秘密鍵

// トークンを生成して返す
func (handler *JwtHandler) Create() (tokenString string, err error) {
	userName := "testman"

	signBytes, err := ioutil.ReadFile("./docs/key/private-key.pem") // 秘密鍵の読み込み
	if err != nil {
		panic(err)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes) // 秘密鍵をパース
    if err != nil {
		log.Print("error is ", err.Error())
        panic(err)
	}

	// 署名アルゴリズムを指定してトークンを生成
	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims) // クレームを設定
	claims["name"] = userName

	tokenString, err = token.SignedString(signKey) // 署名
	if err != nil {
		log.Print("error is ", err.Error())
		return
	}

	return
}