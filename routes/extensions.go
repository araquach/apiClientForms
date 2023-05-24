package routes

import "github.com/araquach/apiClientForms/handlers"

func extensions() {
	s := R.PathPrefix("/api/client-forms").Subrouter()

	s.HandleFunc("/extensions/{salon}/{first_name}/{last_name}", handlers.ApiGetExtensionClients).Methods("GET")
	s.HandleFunc("/extensions", handlers.ApiExtensionsCreate).Methods("POST")
	s.HandleFunc("/extensions-details/{id}", handlers.ApiExtensionsDetails).Methods("GET")
}
