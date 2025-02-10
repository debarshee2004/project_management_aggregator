package main

import (
	"server/project_management/config"
)

func main() {
	// Connect to database
	config.ConnectDatabase()
}
