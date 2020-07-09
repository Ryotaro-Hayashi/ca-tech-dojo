package main

import (
	"net/http"
	"log"
	"ca-tech-dojo/internal/game_api/user/controllers"
	"ca-tech-dojo/pkg/database"
	"ca-tech-dojo/pkg/jwt"
)

func main() {
	sqlHandler := database.NewSqlHandler()
	jwtHandler := jwt.NewJwtHandler()

	// コントローラーの初期化
	userController := controllers.NewUserController(sqlHandler, jwtHandler)

	http.HandleFunc("/users/get", userController.Index)

	http.HandleFunc("/user/create", userController.Create)

	// 第二引数でhandlerにnilを渡して、DefaultServeMuxを使用
	log.Fatal(http.ListenAndServe(":8080", nil))
}