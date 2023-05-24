package routes

import (
	"github.com/gorilla/mux"
)

var R mux.Router

func ClientFormsRouter() {
	extensions()
	skinTest()

	return
}
