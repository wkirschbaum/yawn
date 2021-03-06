package yawn

import (
	"log"
	"net/http"
	"os"
)

func ListenAndServe(handler http.Handler) {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	addr := ":" + port
	log.Println("Starting server on " + addr)
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		panic(err)
	}
}
