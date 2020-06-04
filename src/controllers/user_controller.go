package controllers

import (
	"net/http"
	"fmt"
)

// "/hello"に対しての処理
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// w に書き込み
	fmt.Fprint(w, "hello world!\n")
}