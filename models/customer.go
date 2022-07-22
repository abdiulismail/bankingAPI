package models


type Customer struct {
	Name string `json:"full_name"`
	City string `json:"city"`
	ZipCode string `json:"zip_code"`
}

