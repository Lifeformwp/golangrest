package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/user/rest-api/lib/persistence"
)

func ServeAPI(endpoint string, dbHandler persistence.DatabaseHandler) error {
	handler := NewEventHandler(dbHandler)
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{searchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)

	return http.ListenAndServe(endpoint, r)
}
