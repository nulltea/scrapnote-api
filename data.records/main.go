package main

import "github.com/timoth-y/scrapnote-api/data.records/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}