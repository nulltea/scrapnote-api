package main

import "github.com/timoth-y/scrapnote-api/record/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}