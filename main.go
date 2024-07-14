package main

import (
	"fmt"
	"image-processing/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/", routes.IndexRoute)

	port := ":3000"
	fmt.Println("Listening on port", port)
	http.ListenAndServe(port, r)
}
