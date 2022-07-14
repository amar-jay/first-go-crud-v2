package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)



var db *gorm.DB


func Connection() {
  dns := "host=localhost user=manan password=manan dbname=gomoviesapi port=5432 sslmode=disable"
  dbInst, err := gorm.Open(postgres.New(postgres.Config{
    DSN: dns,
    PreferSimpleProtocol: true,
  }), &gorm.Config{})
  if err != nil {
    panic(err)
  }

  db= dbInst
}

func GetDB() *gorm.DB {
  return db
}
