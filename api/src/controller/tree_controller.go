package controller

import (
	"net/http"
	"fmt"
//    "strconv"
	"github.com/labstack/echo"
  "github.com/NUTMEG-hackathon/echo-tree-api/src/model"
)


func TreeIndex(c echo.Context) error {
  trees := model.TreeIndex()
  return c.JSON(http.StatusOK, trees)
}

func AddTree(c echo.Context) error {
	tree := new(model.Tree)
	if err := c.Bind(tree); err != nil {
			return err
	}

	if tree.Name == "" {
			return &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: "invalid to or message fields",
			}
	}

	user_id := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: user_id}); user.ID == 0 {
			return echo.ErrNotFound
	}

	tree.UserID = user_id
	model.CreateTree(tree)

	return c.JSON(http.StatusCreated, tree)
}

func GetTree(c echo.Context) error {
    user_id := userIDFromToken(c)
    if user := model.FindUser(&model.User{ID: user_id}); user.ID == 0 {
        return echo.ErrNotFound
    }

    tree := model.FindTree(&model.Tree{UserID: user_id})
    return c.JSON(http.StatusOK, tree)
}

func TreePoint(c echo.Context) error {
  user_id := userIDFromToken(c)
  if user := model.FindUser(&model.User{ID: user_id}); user.ID == 0 {
      return echo.ErrNotFound
  }
  tree := model.FindTree(&model.Tree{UserID: user_id})

	tree_param := new(model.Tree)
	if err := c.Bind(tree_param); err != nil {
			return err
	}

  tree.Point = tree.Point + tree_param.Point
  if err := model.UpdateTree(&tree); err != nil {
    return echo.ErrNotFound
  }

  fmt.Println(tree_param.Point)
  return c.JSON(http.StatusOK, tree)
}
