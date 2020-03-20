//Service 1. Additional, send JSON.
//1. Create data array.
//2. Print data on the screen.
//2. Encode data array into JSON.
//3. Send JSON to service 2 on port http :8082

package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Item struct {
	Emp_ID         uint   `json:"emp_id"`
	First_Name     string `json:"first_name"`
	Second_Name    string `json:"second_name"`
	Types          string `json:"types"` //developer, designer, manager
	Experience     uint   `json:"experience"`
	Default_Salary int32  `json:"default_salary"`
}

func main() {
	MakeRequest()
}

func MakeRequest() {

	message := map[string]interface{}{
		"Emp_ID":         1,
		"First_Name":     "Pablo",
		"Second_Name":    "Picasso",
		"Types":          "designer",
		"Experience":     5,
		"Default_Salary": 2000,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://127.0.0.1:8082", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	log.Println(result["data"])
}
