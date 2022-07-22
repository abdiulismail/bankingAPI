package app

import (
	"banking/handlers"
	"log"
	"net/http"
)

func Start() {

	mux := http.NewServeMux()

	//define routes
	mux.HandleFunc("/greet", handlers.Greet)
	mux.HandleFunc("/customers", handlers.GetAllCustomers)

	//starting server
	log.Print("started application at port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
