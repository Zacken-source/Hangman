package main

import (
	"log"
	"main/src/game"
	"net/http"
	_ "net/http/pprof"
)


func main() {
	var g game.Game

	//g.Init()
	//g.Load()
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	//g.Run()
	//g.Unload()
	g.Close()
}

