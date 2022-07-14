package models

import (
	"github.com/amar-jay/first-go-crud-v2/packages/config"
  "gorm.io/gorm"
)

var db *gorm.DB

type Movie struct {
  gorm.Model
  Id string `json:"id"`
  Name string `gorm:"" json:"name"`
  Director Director `json:"director"`
}

type Director struct {
  gorm.Model
  FirstName string `json:"firstname"`
  Surname string `json:"surname"`
}

func init()  {
  config.Connection()
  db = config.GetDB()
  db.AutoMigrate(&Movie{})
}

func CreateMovie(movie *Movie) *Movie {
 // db.NewRecord(movie)
  db.Create(&movie)
  return movie 
}

func GetAllMovies() []Movie {
  var Movies []Movie
  db.Find(&Movies)
  return Movies
}

func GetMovie(id int64) (*Movie,*gorm.DB) {
  var m Movie
  db:= db.Where("ID=?", id).Find(&m)

  
  return &m, db
}

func DeleteMovie(id int64) (*Movie,*gorm.DB) {
  var m Movie
  db:= db.Where("ID=?", id).Delete(m)

  return &m, db
}

func UpdateMovie(id int64) {

}
