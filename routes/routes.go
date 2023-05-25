package routes

import (
	"github.com/gorilla/mux"
)

func ClientFormsRouter() *mux.Router {
	r := mux.NewRouter()

	extensions(r)
	skinTest(r)

	return r
}
