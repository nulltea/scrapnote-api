package main

import "github.com/timoth-y/scrapnote-api/serv.auth/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}