package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"io/ioutil"
)

type Page struct {
    Title string
    Body []byte
}

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/check", save).Methods("GET")
	router.HandleFunc("/new", load).Methods("GET")
	router.HandleFunc("/heythere/{msg}", view).Methods("GET")

	fmt.Println("Hello this is server!")
	log.Fatal(http.ListenAndServe(":8001", router))
}

func check(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Yes it is it!")
}

func handle(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	message := vars.Get("msg")

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func view(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message := vars["msg"]

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func save(w http.ResponseWriter, r *http.Request) {
    p := &Page{Title: "Hello world", Body: []byte("This is our body-text")}
    filename := p.Title + ".txt"
    ioutil.WriteFile(filename, p.Body, 0600)
}

func load(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
    filename := vars.Get("title") + ".txt"
    body, _ := ioutil.ReadFile(filename)

    json.NewEncoder(w).Encode(&Page{Title: filename, Body: body})
}