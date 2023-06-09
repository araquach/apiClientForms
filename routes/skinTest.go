package routes

import (
	"github.com/araquach/apiClientForms/handlers"
	"github.com/gorilla/mux"
)

func skinTest(r *mux.Router) {
	s := r.PathPrefix("/api/client-forms").Subrouter()

	s.HandleFunc("/skintests/{salon}/{first_name}/{last_name}", handlers.ApiGetTestedClients).Methods("GET")
	s.HandleFunc("/skintest", handlers.ApiSkinTestCreate).Methods("POST")
	s.HandleFunc("/skintest-details/{id}", handlers.ApiSkinTestDetails).Methods("GET")
}
