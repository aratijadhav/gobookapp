package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aratijadhav1610/bookapp/routesutil"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func init() {

	router = mux.NewRouter()

	createroute := routesutil.CreateRoute()

	router.HandleFunc(createroute.Url, createroute.HandlerName).Methods(createroute.Methods...)

	fmt.Println("This is init function")
}

func main() {

	fmt.Println("Starting on 8080 port")
	log.Fatal(http.ListenAndServe(":8080", router))

}
