package model

import (
  "fmt"
)

type User struct {
  ID       int    `json:"id" gorm:"praimaly_key"`
  Name     string `json:"name"`
  Password string `json:"password"`
}

type Users []User

func UserIndex() Users {
  db := Init()
  defer db.Close()
  var users Users
  db.Find(&users)
  return users
}

func CreateUser(user *User) {
  db := Init()
  defer db.Close()
  db.Create(user)
}

func FindUser(u *User) User {
  db := Init()
  defer db.Close()
  var user User
  db.Where(u).First(&user)
  return user
}

func UserSeed(){
  db := Init()
  defer db.Close()

  users := Users{
    { ID:1, Name:"sam1", Password:"samsam" },
    { ID:2, Name:"sam2", Password:"samsam" },
    { ID:3, Name:"sam3", Password:"samsam" },
  }

  for i,user := range users{
    fmt.Println(i,user)
    db.Create(user)
  }

}
