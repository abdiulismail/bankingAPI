package main

import (
	"banking/handlers"
	"log"
	"net/http"
)

func main(){
	//define routes
	http.HandleFunc("/greet",handlers.Greet)
	http.HandleFunc("/customers",handlers.GetAllCustomers)


	//starting server
	log.Print("started application at port 8080")
	log.Fatal(	http.ListenAndServe(":8080",nil))
}
