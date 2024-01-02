package donation

import (
	"github.com/go-chi/chi/v5"
)

func V1DonationRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/donation", func(r chi.Router) {
		r.Get("/{id}", GetDonationByIdHandler)
		r.Post("/", CreateDonationHandler)
	})

	return r
}
