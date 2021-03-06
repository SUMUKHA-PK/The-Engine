package main

import (
	"log"

	"github.com/SUMUKHA-PK/Basic-Golang-Server/server"
	"github.com/SUMUKHA-PK/The-Engine/internal/routing"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	m := make(map[string]int)
	r = routing.SetupRouting(r)
	serverData := server.Data{
		Router:        r,
		Port:          "33330",
		IP:            "127.0.0.1",
		HTTPS:         false,
		ConnectionMap: m,
	}

	err := server.Server(&serverData)
	if err != nil {
		log.Fatalf("Could not start server : %v", err)
	}
}
