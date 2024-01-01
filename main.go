package main

import (
	"github.com/Danazzz/task-5-pbi-btpns-INymGdeArtadanaMahaputraW/database"
	"github.com/Danazzz/task-5-pbi-btpns-INymGdeArtadanaMahaputraW/router"
)

func main() {
	database.InitDB()
	database.MigrateDB()
	r := router.RouteInit()
	r.Run(":8080")
}