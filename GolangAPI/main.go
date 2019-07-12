package main

import (
	"github.com/s14228so/WilicoExtra/GolangAPI/db"
	"github.com/s14228so/WilicoExtra/GolangAPI/server"
)

func main() {
	db.Init()
	server.Init()
}
