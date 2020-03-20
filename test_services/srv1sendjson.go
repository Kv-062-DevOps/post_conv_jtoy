//Service 1. Additional, send JSON.
//1. Create data array.
//2. Print data on the screen.
//2. Encode data array into JSON.
//3. Send JSON to service 2 on port http :8082

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Kv-062-DevOps/post_conv_jtoy/handlers"
)

func main() {
	MakeRequest()
}

func MakeRequest() {

	message := handlers.Employ{
		E_Id:       "  1",
		F_Name:     "Pablo",
		S_Name:     "Picasso",
		Type:       "designer",
		Exper:      "5",
		Def_Salary: "2000",
	}

	fmt.Println(message)

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://127.0.0.1:8082", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}
