package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kitabisa/kibitalk/api/readiness"
	"github.com/kitabisa/kibitalk/api/v1/donation"
)

func ApplyRoutes(r *chi.Mux) {
	r.With(middleware.RequestID).Mount("/v1", donation.V1DonationRoutes())
	r.With(middleware.RequestID).Mount("/health_check", readiness.Routes())
}
