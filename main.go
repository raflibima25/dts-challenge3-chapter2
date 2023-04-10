package main

import (
	"challenge-3-chapter-2/database"
	"challenge-3-chapter-2/routers"
)

func main() {
	// var PORT = ":8080"

	database.StartDB()

	routers.StartServer().Run(":8080")

}
