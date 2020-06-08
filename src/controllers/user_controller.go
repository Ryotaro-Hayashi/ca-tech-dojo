package controllers

import (
	"net/http"
	"fmt"
	"ca-tech-dojo/database"
)

type UserController struct {
	SqlHandler database.SqlHandler
}

func NewUserController(sqlHandler database.SqlHandler) UserController {
	return UserController {
		SqlHandler: sqlHandler,
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