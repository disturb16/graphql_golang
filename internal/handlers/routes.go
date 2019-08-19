package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/disturb16/graphql_golang/schema"
	"github.com/disturb16/graphql_golang/settings"
	"github.com/gorilla/mux"
)

// Middleware that sets requestId and apiKey variables in context
func initBaseContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		requestID := req.Header.Get("Kong-Request-ID")
		apiKey := req.Header.Get("api-key")

		reqTrack := map[string]string{
			"Kong-Request-ID": requestID,
			"api-key":         apiKey,
		}

		ctx := context.WithValue(req.Context(), settings.RequestTracking, reqTrack)
		req = req.WithContext(ctx)

		//call hanlder
		next.ServeHTTP(w, req)
	})
}

func (h *handler) initServiceInContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := context.WithValue(req.Context(), "service", h.Service)
		req = req.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

// Router returns api router
func (h *handler) Router() *mux.Router {
	router := mux.NewRouter()

	// set default prefix for Service
	api := router.PathPrefix("/blog-service").Subrouter()
	graphqlHandler, err := schema.NewHandler()

	if err != nil {
		log.Fatal(err)
	}
	// sets base context data for all handlers
	api.Use(initBaseContext)
	api.Use(h.initServiceInContext)

	// set endpoints
	api.HandleFunc("/healthcheck", h.healthCheck).Methods("GET")
	api.HandleFunc("/graphql", graphqlHandler.ServeHTTP)

	return router
}
