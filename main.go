package main

import (
	"praktikum/configs"
	"praktikum/routes"
)

func main() {
	// configs.LoadEnv()
	configs.LoadDatabase()
	routes.Route()
}
