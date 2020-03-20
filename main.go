//SERVICE. ENCODING JSON->YAML.
//1. Listen incoming traffic on port http :8082
//2. Decode structure from JSON into map.
//3. Print mapped data on the screen.
//4. Encode data map into YAML.
//5. Send YAML to service 3 on port http :8083

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ghodss/yaml"
)

type Item struct {
	Emp_ID         uint   `json:"emp_id"`
	First_Name     string `json:"first_name"`
	Second_Name    string `json:"second_name"`
	Types          string `json:"types"` //developer, designer, manager
	Default_Salary int    `json:"default_salary"`
	Experience     uint   `json:"experience"`
}

func main() {

	item := Item{Emp_ID: 2, First_Name: "Bill", Second_Name: "Gates", Types: "developer"}

	jitem, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(jitem))

	y, err := yaml.JSONToYAML(jitem)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//var item1 Item
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		y1, err := yaml.JSONToYAML(body)

		//log.Println(y1)
		//fmt.Println(y1)

		//if err := json.Unmarshal(body, &item1); err != nil {
		//	w.WriteHeader(400)
		//}

		resp, err := http.Post("http://127.0.0.1:8083", "application/yaml", bytes.NewBuffer(y1))
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Fprintf(w, string(resp.Status))
		log.Println(resp)
		fmt.Println(resp)

	})
	http.ListenAndServe(":8082", nil)

}
