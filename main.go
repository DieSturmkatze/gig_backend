package main

import (
	"backend/gig/db"
	"backend/gig/server"
)

func main() {
	db.Init()
	defer db.UnInit()

	server.Serve()
}
