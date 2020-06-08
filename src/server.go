package main

import (
	"net/http"
	"log"
	"ca-tech-dojo/controllers"
	"ca-tech-dojo/database"
)

func main() {
	// コントローラーの初期化とデータベースの初期化
	userController := controllers.NewUserController(database.NewSqlHandler())

	// "/hello"に対しての処理を設定
	http.HandleFunc("/hello", controllers.HelloHandler)

	http.HandleFunc("/user/create", userController.GoodnightHandler)

	// 第二引数でhandlerにnilを渡して、DefaultServeMuxを使用
	log.Fatal(http.ListenAndServe(":8080", nil))
}