//Additional service for test sending JSON (such as Python Frontend works).
//1. Create data array.
//2. Print data on the screen.
//2. Encode data array into JSON.
//3. Send JSON to MAIN service on port http :8082

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
		E_Id:       "1",
		F_Name:     "Pablo",
		S_Name:     "Picasso",
		Type:       "designer",
		Exper:      "5",
		Def_Salary: "2000",
	}

	fmt.Println(message)

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println("Error 422 Unprocessable Entity in data")
		fmt.Println(err)
		return
	}

	resp, err := http.Post("http://127.0.0.1:8082", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		//log.Fatalln(err)
		fmt.Println("Error 503 Service POST Converter Unavailable at port 8082")
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println("Error 400 Bad Request recieved from POST converting service")
		fmt.Println(err)
		return
	}
	log.Println(string(body))
}
