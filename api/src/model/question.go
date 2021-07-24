package model

import (
  "fmt"
)

type Question struct {
  ID        int    `json:"id" gorm:"praimaly_key"`
  Title     string `json:"title"`
  Answers []Answer
}

type Questions []Question

func QuestionIndex() Questions {
  db := Init()
  defer db.Close()

  var questions Questions
  db.Find(&questions)
  return questions
}


func FindQuestion(number int) Question {
  db := Init()
  defer db.Close()

  fmt.Println(number)

  var question Question

  var answers Answers
  db.First(&question, number)
  db.Where(Answer{QuestionID:question.ID}).Find(&answers)
  fmt.Println(answers)
  question.Answers = answers
  return question
}


func QuestionSeed(){
  db := Init()
  defer db.Close()

  questions := Questions{
    { ID:1,  Title:"question1" },
    { ID:2,  Title:"question2" },
    { ID:3,  Title:"question3" },
    { ID:4,  Title:"question4" },
    { ID:5,  Title:"question5" },
    { ID:6,  Title:"question6" },
    { ID:7,  Title:"question7" },
    { ID:8,  Title:"question8" },
    { ID:9,  Title:"question9" },
    { ID:10, Title:"question10" },
    { ID:11, Title:"question11" },
    { ID:12, Title:"question12" },
    { ID:13, Title:"question13" },
    { ID:14, Title:"question14" },
    { ID:15, Title:"question15" },
  }

  for i,question := range questions{
    fmt.Println(i,question)
    db.Create(question)
  }

}

