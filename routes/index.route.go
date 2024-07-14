package routes

import (
	"image-processing/controllers"

	"github.com/go-chi/chi/v5"
)

func IndexRoute(r chi.Router) {
	r.Get("/", controllers.IndexHandler)
	r.Post("/upload", controllers.UploadHandler)
}
