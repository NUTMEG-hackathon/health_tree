package model

import (
  "fmt"
  "time"

  "github.com/jinzhu/gorm"
    _"github.com/go-sql-driver/mysql"
)

func Init() (db *gorm.DB) {
  DBMS := "mysql"
  USER := "go_test"
  PASS := "password"
  PROTOCOL := "tcp(db:3306)"
  DBNAME := "go_database"

  CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

  count := 0
  db, err := gorm.Open(DBMS, CONNECT)
  if err != nil {
    for {
      if err == nil {
        fmt.Println("")
        break
      }
      fmt.Print(".")
      time.Sleep(time.Second)
      count++
      if count > 180 {
        fmt.Println("")
        fmt.Println("DB connect fail")
        panic(err)
      }
      db, err = gorm.Open(DBMS, CONNECT)
    }
  }
	fmt.Println("DB connect succeed")

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Tree{})
	db.AutoMigrate(&Log{})
	db.AutoMigrate(&Question{})
	db.AutoMigrate(&Answer{})

  return db
}
