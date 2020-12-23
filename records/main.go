package main

import "github.com/timoth-y/scrapnote-api/records/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}