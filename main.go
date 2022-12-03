package main

import (
	"my-gram-1/configs"
	"my-gram-1/routes"
)

func main() {
	configs.StartDB()
	routes.MainRouter()
}
