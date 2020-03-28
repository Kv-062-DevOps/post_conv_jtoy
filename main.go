//POST Data Converting Service JSON to YAML:
//1. Listen requests from frontend on port http :8082
//2. Decode body from JSON.
//3. Encode data into YAML.
//4. Print structure on the screen.
//5. Send YAML to DB service on port http :8083 path /add.

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/ghodss/yaml"
)

func main() {

	postport := "0.0.0.0:" + os.Getenv("POSTPORT")                                       //default is 8082
	backlink := "http://" + os.Getenv("BACKADDR") + ":" + os.Getenv("BACKPORT") + "/add" //defaults are "127.0.0.1" and 8083

	fmt.Println("POSTPORT =", postport) //default is "0.0.0.0:8082"
	fmt.Println("BACKLINK =", backlink) //default is "http://127.0.0.1:8083/add"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "400 Bad Request from Frontend", 400)
			fmt.Println(err)
			return
		}

		converted, err := yaml.JSONToYAML(body)
		if err != nil {
			http.Error(w, "422 Unprocessable Entity recieved", 422)
			fmt.Println(err)
			return
		}

		//resp, err := http.Post("http://127.0.0.1:8083/add", "application/yaml", bytes.NewBuffer(converted))
		resp, err := http.Post(backlink, "application/yaml", bytes.NewBuffer(converted))
		if err != nil {
			http.Error(w, "503 Service DB Unavailable at this link", 503)
			fmt.Println(err)
			return
		}
		fmt.Println(resp.Status)
		fmt.Fprintf(w, resp.Status)
		fmt.Println(bytes.NewBuffer(converted))
		fmt.Println("POSTPORT =", postport) //default is "0.0.0.0:8082"
		fmt.Println("BACKLINK =", backlink) //default is "http://127.0.0.1:8083/add"

	})

	//http.ListenAndServe("0.0.0.0:8082", nil)
	http.ListenAndServe(postport, nil)

}
