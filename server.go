package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	// "/hello"に対しての処理を設定
	http.HandleFunc("/hello", helloHandler)

	// 第二引数でhandlerにnilを渡して、DefaultServeMuxを使用
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// "/hello"に対しての処理
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// w に書き込み
	fmt.Fprint(w, "hello world!\n")
}