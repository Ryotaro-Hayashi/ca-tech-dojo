package controllers

import (
	"net/http"
	"fmt"
	"ca-tech-dojo/pkg/database"
	"ca-tech-dojo/pkg/jwt"
	"ca-tech-dojo/internal/game_api/user/models"
	"encoding/json"
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

// ユーザー一覧をJSONで返す
func (controller *UserController) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { // GETリクエストのみ許可
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "The request is limited to the GET method")
		return
	}

	// ユーザーをDBから取得
	users, err := controller.UserRepository.GetAll()
	if err != nil {
		fmt.Fprint(w, err)
	}

	usersByte, err := json.Marshal(users) // 構造体を []byte へ変換
    if err != nil {
        fmt.Fprint(w, err)
    }
    usersJson := string(usersByte) // []byte をJSON文字列に変換

	fmt.Fprint(w, usersJson)
}

// トークンを生成してユーザーを保存して、保存したユーザーのトークンをJSONで返す
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

	userCreateRequest := models.UserCreateRequest{}
    err = json.Unmarshal(body, &userCreateRequest)  // []byte型を構造体に変換
    if err != nil { 
     panic(err) 
    }

	tokenString, err := controller.JwtHandler.Create(userCreateRequest.Name) // トークンの生成
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	id, err := controller.UserRepository.Create(userCreateRequest.Name, tokenString) // ユーザーを保存
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	tokenString, err = controller.UserRepository.FindTokenById(id) // idでユーザ検索
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	userCreateResponse := models.UserCreateResponse{
		Token: tokenString,
	}

	userCreateResponseByte, err := json.Marshal(userCreateResponse) // 構造体を []byte へ変換
    if err != nil {
        fmt.Fprint(w, err)
    }
    tokenJson := string(userCreateResponseByte) // []byte をJSON文字列に変換

	fmt.Fprint(w, tokenJson)
}