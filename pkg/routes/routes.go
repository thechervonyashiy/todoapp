package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thechervonyashiy/todoapp/pkg/handlers"
)

func InitRoutes(handler *handlers.Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", handler.SignUp) // POST /auth/sign-up
		r.Post("/sign-in", handler.SignIn) // POST /auth/sign-ip
	})

	r.Route("/lists", func(r chi.Router) {
		r.Get("/", handler.GetAllLists) // GET /lists
		r.Post("/", handler.CreateList) // POST /lists

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.GetListById)   // GET /lists/{id}
			r.Put("/", handler.UpdateList)    // PUT /lists/{id}
			r.Delete("/", handler.DeleteList) // DELETE /lists/{id}

			r.Get("/items", handler.GetAllItems) // GET /lists/{id}/items
			r.Post("/items", handler.CreateItem) // POST /lists/{id}/items

			r.Route("/{item_id}", func(r chi.Router) {
				r.Get("/", handler.GetItemById)   // GET /lists/{id}/items/{item_id}
				r.Put("/", handler.UpdateItem)    // PUT /lists/{id}/items/{item_id}
				r.Delete("/", handler.DeleteItem) // DELETE /lists/{id}/items/{item_id}
			})
		})
	})

	return r
}
