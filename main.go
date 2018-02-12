package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Category struct {
	Title string	`json:"title,omitempty"`
}

type Page struct {
    Title 		string		`json:"title,omitempty"`
    Body 		string		`json:"text,omitempty"`
    Category 	*Category	`json:"category,omitempty"`
}

var pages []Page
func main() {
	var router = mux.NewRouter()

	pages = append(pages, Page{Title: "First page", Body: "This is first page", Category: &Category{Title: "First category"}})
	pages = append(pages, Page{Title: "Second page", Body: "This is second page", Category: &Category{Title: "Second category"}})
	pages = append(pages, Page{Title: "Third page", Body: "This is third page", Category: &Category{Title: "Third category"}})

	router.HandleFunc("/pages", GetPages).Methods("GET")
	router.HandleFunc("/page/{id}", GetPage).Methods("GET")
	//router.HandleFunc("/check", save).Methods("GET")
	//router.HandleFunc("/new", load).Methods("GET")
	router.HandleFunc("/heythere/{msg}", view).Methods("GET")

	fmt.Println("Hello this is server!")
	log.Fatal(http.ListenAndServe(":8001", router))
}

func GetPages(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pages)
}

func GetPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range pages {
		if item.Title == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
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

//func load(w http.ResponseWriter, r *http.Request) {
//	vars := r.URL.Query()
//    filename := vars.Get("title") + ".txt"
//    body, _ := ioutil.ReadFile(filename)
//
//    json.NewEncoder(w).Encode(&Page{Title: filename, Body: body})
//}