package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	addr := os.Getenv("ADDR")

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/hello", HelloHandler)
	mux.HandleFunc("/v1/time", CurrentTimeHandler)

	log.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func CurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("current time is %v", currentTime)))
}
