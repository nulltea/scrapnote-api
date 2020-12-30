package main

import "github.com/timoth-y/scrapnote-api/data.users/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}