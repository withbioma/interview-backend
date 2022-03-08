package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/withbioma/interview-backend/controllers"
)

type Router struct {
	ProductController             *controllers.ProductController
	ProductVariantGroupController *controllers.ProductVariantGroupController
}

func main() {
	productController := controllers.GetProductControllerInstance()
	productVariantGroupController := controllers.GetProductVariantGroupControllerInstance()

	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Route("/products", func(r chi.Router) {
			r.Get("/", productController.GetAll)
			r.Post("/", productController.Create)
			r.Route("/{product_id}", func(r chi.Router) {
				r.Get("/", productController.Get)
				r.Delete("/", productController.Delete)
			})
		})

		r.Route("/product_variant_groups", func(r chi.Router) {
			r.Get("/", productVariantGroupController.GetAll)
			r.Post("/", productController.Create)
			r.Route("/{product_variant_group_id}", func(r chi.Router) {
				r.Delete("/", productController.Delete)
			})
		})
	})

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err, -1)
	}
}
