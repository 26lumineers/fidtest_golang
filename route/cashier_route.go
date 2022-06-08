package route

import (
	"fidtest_golang/controller"
	"github.com/gorilla/mux"
)

func GetCashierRoute(r *mux.Router) *mux.Router{
	router := r.PathPrefix("/cash").Subrouter()
	router.HandleFunc("/cashier",controller.Cashier).Methods("POST")
	router.HandleFunc("/cashier-inf",controller.Cashier_infinity).Methods("POST")

	return router
}