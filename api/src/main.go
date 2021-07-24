package main

import (
  "github.com/NUTMEG-hackathon/echo-tree-api/src/router"
)

func main() {
  router := router.NewRouter()
  router.Logger.Fatal(router.Start(":8888"))
}
