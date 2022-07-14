package controllers

import (
  "net/http"
_  "encoding/json"
_  "github.com/gorilla/mux"
_  "strconv"
_  "github.com/amar-jay/first-go-crud-v2/packages/models"
_  "github.com/amar-jay/first-go-crud-v2/packages/utils"
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

