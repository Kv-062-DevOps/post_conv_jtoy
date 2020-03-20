//Service 3. Additional, get YAML.
//1. Listen incoming traffic on port http :8083
//2. Decode structure from YAML into map.
//3. Print mapped data on the screen.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
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
	log.Println(router)
	fmt.Println(router)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	team := Team{
		Employee{Name: "Vasya", Age: 25},
		Employee{Name: "Petya", Age: 30},
	}
	if err := yaml.NewEncoder(w).Encode(team); err != nil {
		panic(err)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lastname := vars["lastname"]
	var emp Employee
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &emp); err != nil {
		w.WriteHeader(400)
	}

	w.WriteHeader(200)
	fmt.Fprintln(w, "your name is:", emp.Name, lastname)
}

type Employee struct {
	Name string `json:"name"`
	Age  uint8  `json:"employee_age"`
}

type Team []Employee

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
	Route{
		Name:        "Hello",
		Method:      "POST",
		Pattern:     "/hello/{lastname}",
		HandlerFunc: HelloHandler,
	},
}
