package handler

import (
	"net/http"
	"html/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request){
	// テンプレート読み込み
	t := template.Must(template.ParseFiles("templates/Index.html"))
	// テンプレートを描画
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
