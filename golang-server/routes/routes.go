package routes

import (
	"github.com/giovannylucas/observability-elastic-stack/golang-server/controllers"
	"github.com/gorilla/mux"
)

func HandleRequests() *mux.Router {
	// create a new router
	router := mux.NewRouter()

	router.HandleFunc("/events", controllers.CreateEvent).Methods("GET")

	return router
}