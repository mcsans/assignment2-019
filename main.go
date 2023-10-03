package main

import (
	"github.com/mcsans/assignment2-019/models"
	"github.com/mcsans/assignment2-019/routers"
)

func main() {
	models.ConnectDatabase()
	routers.StartServer().Run(":8080")
}