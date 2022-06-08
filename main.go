package main

import (
	"fidtest_golang/route"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hi fid")
	RouterApp()
}
func RouterApp(){
	router := mux.NewRouter().StrictSlash(true)
	path := router.PathPrefix("/api").Subrouter()
	route.GetCashierRoute(path)
	route.GetFindValueRoute(path)
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
