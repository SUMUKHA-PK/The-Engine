package routing

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRouting adds all routes to the server
func SetupRouting(r *mux.Router) *mux.Router {
	r.HandleFunc("/query", IncomingQuery).Methods(http.MethodPost, http.MethodOptions)
	return r
}
