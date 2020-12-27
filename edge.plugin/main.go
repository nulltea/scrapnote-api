package main

import "github.com/timoth-y/scrapnote-api/edge.plugin/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}