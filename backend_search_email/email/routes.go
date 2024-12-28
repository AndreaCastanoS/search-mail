package email

import (
	"github.com/go-chi/chi/v5"
)

func InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/search", SearchEmails) 
	r.Get("/email", GetEmailByID)


	return r
}
