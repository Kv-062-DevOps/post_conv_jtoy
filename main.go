//MAIN SERVICE. ENCODING JSON->YAML.
//1. Listen incoming requests on port http :8082
//2. Decode body from JSON.
//3. Encode data into YAML.
//4. Print structure on the screen.
//5. Send YAML to service 3 on port http :8083 path /add

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ghodss/yaml"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		converted, err := yaml.JSONToYAML(body)
		if err != nil {
			log.Fatalln(err)
		}

		resp, err := http.Post("http://127.0.0.1:8083/add", "application/yaml", bytes.NewBuffer(converted))
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(resp.Status)
		fmt.Fprintf(w, resp.Status)
		fmt.Println(bytes.NewBuffer(converted))

	})

	http.ListenAndServe(":8082", nil)

}
