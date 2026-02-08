package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Jacksondprice/river-workshop-service/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/classrooms", handlers.ListClassrooms)
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Fatal(http.ListenAndServe(addr, mux))
}
