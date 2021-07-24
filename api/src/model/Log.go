package model

import (
  "fmt"
)

type Log struct {
  ID        int    `json:"id" gorm:"praimaly_key"`
  Win       bool   `json:"win"`
  TreeID    int    `json:"tree_id"`
}

type Logs []Log

func LogIndex() Logs {
  db := Init()
  defer db.Close()

  var logs Logs
  db.Find(&logs)
  return logs
}


func CreateLog(log *Log) {
  db := Init()
  defer db.Close()

  db.Create(log)
}

func FindLogs(t *Log) Logs {
  db := Init()
  defer db.Close()

  var logs Logs
  db.Where(t).Find(&logs)
  return logs
}

func LogSeed(){
  db := Init()
  defer db.Close()

  logs := Logs{
    { ID:1, Win:true,  TreeID:1 },
    { ID:2, Win:false, TreeID:1 },
    { ID:3, Win:true,  TreeID:2 },
    { ID:4, Win:false, TreeID:2 },
    { ID:5, Win:true,  TreeID:3 },
    { ID:6, Win:false, TreeID:3 },
  }

  for i,log := range logs{
    fmt.Println(i,log)
    db.Create(log)
  }

}
