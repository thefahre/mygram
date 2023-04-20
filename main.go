package main

import (
	"mygram/database"
	"mygram/routers"
	// "github.com/swaggo/gin-swagger" // gin-swagger middleware
	// "github.com/swaggo/files" // swagger embed files
)

func init() {

}

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(":8080")
}
