package config

import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)



var db *gorm.DB


func Connection() {
  dbInst, err := gorm.Open("postgres", "postgres:Abdel-manan1978@/gomovieapi?charset=utf8&parseTime=True&loc=Local",)
  if err != nil {
    panic(err)
  }

  db= dbInst
}

func GetDB() *gorm.DB {
  return db
}
