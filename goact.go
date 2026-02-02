package goact

import (
	"html/template"
	"net/http"
)

type DataFunc func(*http.Request) any

func MustTemplate(name, html string) *template.Template {
	return template.Must(
		template.New(name).Parse(html),
	)
}

func Handle(
	mux *http.ServeMux,
	path string,
	tmpl *template.Template,
	dataFn DataFunc,
) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		data := dataFn(r)

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
		}
	})
}
