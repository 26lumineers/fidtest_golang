package route

import (
	"fidtest_golang/controller"
	"github.com/gorilla/mux"
)

func GetFindValueRoute(r *mux.Router) *mux.Router{
	router := r.PathPrefix("/find").Subrouter()
	router.HandleFunc("/xyz",controller.FindValue).Methods("GET")

	return router
}