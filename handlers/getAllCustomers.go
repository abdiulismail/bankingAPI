package handlers

import (
	"banking/models"
	"encoding/json"
	"net/http"
)

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []models.Customer{
		{"John","Nairobi","1234"},
		{"Alex","Lisbon","345"},
	}
	w.Header().Add("Content-Type","application/json")
	json.NewEncoder(w).Encode(customers)
	//res, _ := json.Marshal(customers)
	//w.Write(res)
}
