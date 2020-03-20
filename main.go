//SERVICE. ENCODING JSON->YAML.
//1. Listen incoming traffic on port http :8082
//2. Decode structure from JSON into map.
//3. Print mapped data on the screen.
//4. Encode data map into YAML.
//5. Send YAML to service 3 on port http :8083

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ghodss/yaml"
	//"github.com/ghodss/yaml"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		y1, err := yaml.JSONToYAML(body)

		resp, err := http.Post("http://127.0.0.1:8083", "application/yaml", bytes.NewBuffer(y1))
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(resp.Status)

		fmt.Fprintf(w, resp.Status)

	})
	http.ListenAndServe(":8082", nil)

}
