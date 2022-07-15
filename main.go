package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/amar-jay/first-go-crud-v2/packages/routers"
	"github.com/gorilla/mux"
)
func main(){
  router := mux.NewRouter()

  routers.HandleRoutes(router)
  
  http.Handle("/", router)
  fmt.Printf("Starting Server at http://localhost:4000\n")
  log.Fatal(http.ListenAndServe(":4000", router))

}
