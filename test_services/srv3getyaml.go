//Service 3. Additional, get YAML.
//1. Listen incoming traffic on port http :8083
//2. Decode structure from YAML into map.
//3. Print mapped data on the screen.

package main

import (
	"fmt"
	"log"
	"net/http"

	//"github.com/gorilla/mux"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	//"gopkg.in/yaml.v2"
)

type Item struct {
	Emp_ID         uint   `yaml:"emp_id"`
	First_Name     string `yaml:"first_name"`
	Second_Name    string `yaml:"second_name"`
	Types          string `yaml:"types"` //developer, designer, manager
	Default_Salary int    `yaml:"default_salary"`
	Experience     uint   `yaml:"experience"`
}

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8083", router))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	var item Item
	if err := yaml.NewDecoder(r.Body).Decode(&item); err != nil {
		panic(err)
	}
	fmt.Println(item)
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
		Pattern:     "/",
		HandlerFunc: RootHandler,
	},
}
