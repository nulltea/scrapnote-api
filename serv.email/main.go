package main

import "github.com/timoth-y/scrapnote-api/serv.email/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}