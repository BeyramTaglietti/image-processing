package utils

import (
	"net/http"

	"github.com/a-h/templ"
)

func SendComponent(w http.ResponseWriter, r *http.Request, component templ.Component) {
	templ.Handler(component).ServeHTTP(w, r)
}
