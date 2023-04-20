package main

import (
	"fmt"
	"mygram/database"
	"mygram/routers"
)

func init() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(":8080")
}

func main() {
	fmt.Println("Init complete.")
}
