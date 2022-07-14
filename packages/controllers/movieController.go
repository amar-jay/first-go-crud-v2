package controllers

import (
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "strconv"
  "github.com/amar-jay/first-go-crud-v2/packages/models"
  "github.com/amar-jay/first-go-crud-v2/packages/utils"
)
func GetAllMovies(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
}


func GetMovie(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
}


func CreateMovie(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
}

func UpdateMovie(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
}


func DeleteMovie(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
}

