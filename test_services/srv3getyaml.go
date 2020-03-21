//Service 3. Additional for test, get YAML (such as Python DB module works).
//1. Listen requests on port http :8083 path /add
//2. Decode body from YAML.
//3. Print structure on the screen.

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kv-062-DevOps/post_conv_jtoy/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
)

func main() {
	router := NewRouter()
	//log.Fatal(http.ListenAndServe(":8083", router))
	log.Fatal(http.ListenAndServe("127.0.0.1:8083", router))

}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	var message handlers.Employ
	if err := yaml.NewDecoder(r.Body).Decode(&message); err != nil {
		panic(err)
	}
	fmt.Println(message)
	w.WriteHeader(200)
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routesArray {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}

var routesArray = []Route{
	Route{
		Name:        "Root",
		Method:      "POST",
		Pattern:     "/add",
		HandlerFunc: RootHandler,
	},
}
