package controllers

import (
	"net/http"
	"fmt"
	"ca-tech-dojo/pkg/database"
	"ca-tech-dojo/pkg/jwt"
	"encoding/json"
	"log"
)

type UserController struct {
	UserRepository database.UserRepository
	jwt.JwtHandler
}

func NewUserController(sqlHandler database.SqlHandler, jwtHandler jwt.JwtHandler) UserController {
	return UserController {
		UserRepository: database.UserRepository {
			SqlHandler: sqlHandler,
		},
		JwtHandler: jwtHandler,
	}
}

// "/hello"に対しての処理
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// w に書き込み
	fmt.Fprint(w, "hello world!\n")
}

func (controller UserController) GoodnightHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "good night!\n")
}


// ユーザー一覧をJSONで返す
func (controller UserController) Index(w http.ResponseWriter, r *http.Request) {
	// ユーザーをDBから取得
	users, err := controller.UserRepository.GetAll()
	if err != nil {
		fmt.Fprint(w, err)
	}
	log.Print("The users struct is ", users)

	usersByte, err := json.Marshal(users) // 構造体を []byte へ変換
    if err != nil {
        fmt.Fprint(w, err)
    }
    usersJson := string(usersByte) // []byte をJSON文字列に変換

	fmt.Fprint(w, usersJson)
}

// トークンを生成してユーザーを保存する
func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {
	tokenString, err := controller.JwtHandler.Create() // トークンの生成
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	id, err := controller.UserRepository.Create() // ユーザーを保存
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	log.Print("The token is ", tokenString)

	fmt.Fprint(w, id)
}