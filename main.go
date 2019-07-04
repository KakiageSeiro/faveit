package main

import (
	"KakiageSeiro/faveit/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.IndexHandler)
	http.ListenAndServe(":80", nil)
}

