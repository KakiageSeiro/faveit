package handler

import (
	"github.com/unrolled/render"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request){


	rendering(w, "index", nil)
}

//httpステータス200でHTMLレンダリングする
func rendering(w http.ResponseWriter, templateName string, data interface{}) {
	re := render.New(render.Options{
		Extensions: []string{".html"},
		Charset:    "UTF-8",
	})
	re.HTML(w, http.StatusOK, templateName, data)
}
