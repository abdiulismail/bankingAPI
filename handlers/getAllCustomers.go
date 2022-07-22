package handlers

import (
	"banking/models"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []models.Customer{
		{"John", "Nairobi", "1234"},
		{"Alex", "Lisbon", "345"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		_ = xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(customers)
	}
	//res, _ := json.Marshal(customers)
	//w.Write(res)
}
