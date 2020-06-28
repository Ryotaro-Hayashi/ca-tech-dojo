package controllers

import (
	"net/http"
	"fmt"
	"ca-tech-dojo/pkg/database"
	"ca-tech-dojo/pkg/jwt"
	"ca-tech-dojo/internal/game_api/user/models"
	"encoding/json"
	"log"
	"io/ioutil"
)

type UserController struct {
	database.UserRepository
	*jwt.JwtHandler
}

func NewUserController(sqlHandler *database.SqlHandler, jwtHandler *jwt.JwtHandler) *UserController {
	return &UserController {
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
func (controller *UserController) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { // GETリクエストのみ許可
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "The request is limited to the GET method")
		return
	}

	// ヘッダーのContent-Typeを検証
	// if r.Header.Get("Content-Type") != "application/json" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	fmt.Fprint(w, "The ContentーType is limited to the application/json")
	// 	return
	// }

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

// トークンを生成してユーザーを保存して、保存したユーザーidを返す
func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { // POSTリクエストのみ許可
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "The Request is limited to the POST method")
		return
	}

	// ヘッダーのContent-Typeを検証
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "The ContentーType is limited to the application/json")
		return
	}

	// リクエストボディを読み取って[]byte型の変数に格納
	body, err := ioutil.ReadAll(r.Body) 
    if err != nil { 
      panic(err) 
	} 

	user := models.User{}
    err = json.Unmarshal(body, &user)  // []byte型を構造体に変換
    if err != nil { 
     panic(err) 
    } 

	log.Printf("The request body is %+v", user)

	user, err = controller.JwtHandler.Create(user) // トークンの生成
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	id, err := controller.UserRepository.Create(user) // ユーザーを保存
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	log.Print("The stored user is %+v", user)

	fmt.Fprint(w, id)
}