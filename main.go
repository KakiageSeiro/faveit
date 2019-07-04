package main

import (
	"KakiageSeiro/faveit/handler"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", handler.IndexHandler) //一覧

	http.ListenAndServe(":80", r)
}

