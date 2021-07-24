package controller

import (
  "net/http"
  "math/rand"
 //   "fmt"

  "github.com/labstack/echo"
  "github.com/NUTMEG-hackathon/echo-tree-api/src/model"
)


func QuestionIndex(c echo.Context) error {
  questions := model.QuestionIndex()
  return c.JSON(http.StatusOK, questions)
}

func GetQuestion(c echo.Context) error {
  number := []int{}
  number = append(number,rand.Intn(15)+1)
  for len(number) < 5 {
    var flag = 0
    add_num := rand.Intn(15)+1
    for i := 0; i < len(number); i++{
      if number[i] == add_num {
        flag = 1
      }
    }
    if flag == 0{
      number = append(number,add_num)
    }
  }

  question1 := model.FindQuestion(number[0])
  question2 := model.FindQuestion(number[1])
  question3 := model.FindQuestion(number[2])
  question4 := model.FindQuestion(number[3])
  question5 := model.FindQuestion(number[4])

  questions := []model.Question{
    question1,
    question2,
    question3,
    question4,
    question5,
  }
  return c.JSON(http.StatusOK, questions)
}
