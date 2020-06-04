package main

import (
	"net/http"
	"log"
	"ca-tech-dojo/controllers"
)

func main() {
	// "/hello"に対しての処理を設定
	http.HandleFunc("/hello", controllers.HelloHandler)

	// 第二引数でhandlerにnilを渡して、DefaultServeMuxを使用
	log.Fatal(http.ListenAndServe(":8080", nil))
}