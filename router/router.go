package router

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/pjol/bera-bd-backend/handlers/auth"
	"github.com/pjol/bera-bd-backend/handlers/cfi"
	"github.com/pjol/bera-bd-backend/handlers/ramps"
	m "github.com/pjol/bera-bd-backend/utils/middleware"
)

func New(a *auth.Service, r *ramps.Service, c *cfi.Service) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Token", "X-Admin-Key"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Use(m.AuthMiddleware)

	return router
}

func withAdmin(handlerFunc http.HandlerFunc, s *auth.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqKey := r.Header.Get("X-Admin-Key")
		envKey := os.Getenv("ADMIN_KEY")
		if reqKey == envKey && envKey != "" {
			handlerFunc(w, r)
			return
		}

		id, ok := r.Context().Value("userDid").(string)
		if !ok {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		isAdmin := s.IsAdmin(r.Context(), id)
		if !isAdmin {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		handlerFunc(w, r)
	}
}
