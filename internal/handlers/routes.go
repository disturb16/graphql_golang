package handlers

import (
	"log"
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/disturb16/graphql_golang/settings"
	"github.com/disturb16/graphql_golang/schema"
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

// Router returns api router
func (h *Handler) Router() *mux.Router {
	router := mux.NewRouter()

	// set default prefix for Service
	api := router.PathPrefix("/blog-service").Subrouter()
	graphqlHandler, err := schema.New(h.Service)

	if err != nil{
		log.Fatal(err)
	}

	// set endpoints
	api.HandleFunc("/home", h.home).Methods("GET")
	api.HandleFunc("/healthcheck", h.healthCheck).Methods("GET")
	api.Handle("/graphql", graphqlHandler)

	// sets base context data for all handlers
	api.Use(initBaseContext)

	return router
}
