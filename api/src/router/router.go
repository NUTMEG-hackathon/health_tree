package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
  "github.com/NUTMEG-hackathon/echo-tree-api/src/controller"
)

func NewRouter() *echo.Echo {
  e := echo.New()

  e.Use(middleware.CORS())
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // application
  e.GET("/", controller.Hello)
  e.GET("/seed", controller.Seed)

  // user
  e.GET("/user_index", controller.UserIndex)
  e.POST("/signup", controller.Signup)
  e.POST("/signin", controller.Signin)

  api := e.Group("/api")
  api.Use(middleware.JWTWithConfig(controller.Config)) // /api 下はJWTの認証が必要
  api.GET("/tree", controller.GetTree) // GET /api/tree
  api.POST("/tree", controller.AddTree) // POST /api/tree
  api.PUT("/tree", controller.TreePoint) // PUT /api/tree

  // tree
  e.GET("/tree_index", controller.TreeIndex)

  // log
  e.GET("/log_index", controller.LogIndex)
  e.POST("/logs", controller.AddLog)

  // question
  e.GET("/question_index", controller.QuestionIndex)
  e.GET("/get_question", controller.GetQuestion)

  // answer
  e.GET("/answer_index", controller.AnswerIndex)

  return e
}

