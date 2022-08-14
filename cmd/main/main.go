package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/maksiecluster/booking/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoomsRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Starting server at port 8099\n")
	log.Fatal(http.ListenAndServe("localhost:8099", r))
}
