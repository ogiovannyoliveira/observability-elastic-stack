package main

import (
	"fmt"
	"net/http"

	// "github.com/giovannylucas/observability-elastic-stack/golang-server/database"
	"github.com/giovannylucas/observability-elastic-stack/golang-server/routes"
	"github.com/gorilla/handlers"
)



func main() {
	// connect to the database
	// database.ConnectOnDatabase()

	// create a new router
	r := routes.HandleRequests()

	// start the server
	fmt.Println("Server started at port 3334!")
	http.ListenAndServe(":3334", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r))
}