package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/check", check).Methods("GET")
	router.HandleFunc("/new", handle).Methods("GET")
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