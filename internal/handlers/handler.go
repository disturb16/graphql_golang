package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/disturb16/graphql_golang/internal/services"
)

// handler main structure with handlers functionality
type handler struct {
	Service *services.Service
}

// healthcheck model definition
type healthcheck struct {
	Host     string    `json:"host"`
	Datetime time.Time `json:"datetime"`
}

// New initialize main handler
func New(s *services.Service) *handler {
	return &handler{
		Service: s,
	}
}

// Home handler example
func (h *handler) home(w http.ResponseWriter, req *http.Request) {

	output := "Hello World"
	jsonResponse(w, output)
}

func (h *handler) healthCheck(w http.ResponseWriter, r *http.Request) {
	// local properties
	var out healthcheck
	var err error
	out.Host, err = os.Hostname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(out.Host)
	out.Datetime = time.Now()

	jsonResponse(w, out)
}
