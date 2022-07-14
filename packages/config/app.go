package config

import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)



var db *gorm.DB


func Connection() {
  config := gorm.Config{}
  dbInst, err := gorm.Open("postgres", postgres.Open("host=localhost user=manan password=manan dbname=gomoviesapi port=5432 sslmode=disable", &config)
  if err != nil {
    panic(err)
  }

  db= dbInst
}

func GetDB() *gorm.DB {
  return db
}
