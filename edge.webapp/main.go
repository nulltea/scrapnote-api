package main

import "github.com/timoth-y/scrapnote-api/edge.webapp/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}