package app

import (
	"banking/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	//define routes
	router.HandleFunc("/greet", handlers.Greet)
	router.HandleFunc("/customers", handlers.GetAllCustomers)

	//starting server
	log.Print("started application at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
