package controller

import (
  "net/http"
  "github.com/labstack/echo"
  "github.com/NUTMEG-hackathon/echo-tree-api/src/model"
)


func AnswerIndex(c echo.Context) error {
  answers := model.AnswerIndex()
  return c.JSON(http.StatusOK, answers)
}

