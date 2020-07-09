package jwt

import (
	"log"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"github.com/google/uuid"
)

type JwtHandler struct {
	Token *jwt.Token
}

func NewJwtHandler() (jwtHandler *JwtHandler){
	jwtHandler = new(JwtHandler)
	return
}

const (
	signKey = "ca-tech-dojo-key" // 共通鍵
)

// トークンを生成して、ユーザーの構造体に格納して返す
func (handler *JwtHandler) Create(name string) (tokenString string, err error) {

	// 署名アルゴリズムを指定してトークン型を生成
	token := jwt.New(jwt.SigningMethodHS256)

	// uuid を生成
	uu, err := uuid.NewRandom()
    if err != nil {
      return
    }

	claims := token.Claims.(jwt.MapClaims) // クレームを設定
	claims["name"] = name
	claims["iat"] = time.Now().Unix() // トークンの発行時間
	claims["sub"] = uu.String() // 一意な識別子

	log.Print("the uuid is ", uu.String())

	tokenString, err = token.SignedString([]byte(signKey)) // 署名
	if err != nil {
		log.Print("error is ", err.Error())
		return
	}
	
	return
}