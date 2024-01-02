package readiness

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/kitabisa/kibitalk/config/database"
	"net/http"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/db", ReadinessHandler)
	return r
}

func ReadinessHandler(rw http.ResponseWriter, r *http.Request) {
	err := database.MySqlDB.Ping(context.Background())
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(200)
}
