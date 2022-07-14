package routers

import (
  "github.com/gorilla/mux"
  "github.com/amar-jay/first-go-crud-v2/packages/controllers"
)
func HandleRoutes(router *mux.Router) {

  router.HandleFunc("/getMovies", controllers.GetAllMovies).Methods("GET")
  router.HandleFunc("/getMovies/{id}", controllers.GetMovie).Methods("GET")
  router.HandleFunc("/createMovie/{id}", controllers.CreateMovie).Methods("POST")
  router.HandleFunc("/updateMovie/{id}", controllers.UpdateMovie).Methods("PUT")
  router.HandleFunc("/deleteMovie/{id}", controllers.DeleteMovie).Methods("DELETE")

}
