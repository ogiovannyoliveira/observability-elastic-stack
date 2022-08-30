package routes

import (
	"net/http"

	"github.com/giovannylucas/observability-elastic-stack/golang-server/controllers"
	"github.com/gorilla/mux"

	"go.elastic.co/apm/module/apmgorilla/v2"
)

func HandleRequests() *mux.Router {
	// create a new router
	router := mux.NewRouter()
	 
	// apmgorilla.Instrument(router)

	router.Use(apmgorilla.Middleware())

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
		return
	}).Methods("GET")

	router.HandleFunc("/events", controllers.CreateEvent).Methods("POST")

	return router
}