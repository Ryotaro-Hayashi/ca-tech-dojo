package main

import (
	"net/http"
	"log"
	"ca-tech-dojo/internal/game_api/user/controllers"
	"ca-tech-dojo/pkg/database"
)

func main() {
	// コントローラーの初期化とデータベースの初期化
	userController := controllers.NewUserController(database.NewSqlHandler())

	// "/hello"に対しての処理を設定
	http.HandleFunc("/hello", controllers.HelloHandler)
	http.HandleFunc("/good-night", userController.GoodnightHandler)
	http.HandleFunc("/users/get", userController.Index)

	// 第二引数でhandlerにnilを渡して、DefaultServeMuxを使用
	log.Fatal(http.ListenAndServe(":8080", nil))
}