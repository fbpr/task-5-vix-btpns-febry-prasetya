package main

import (
	"github.com/fbpr/task-5-vix-btpns-febry-prasetya/database"
  "github.com/fbpr/task-5-vix-btpns-febry-prasetya/router"
)

func main() {
  database.StartDB()
  r := router.StartServer()

  r.Run()
}
