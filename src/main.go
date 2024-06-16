package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

type Logger struct {
	handler http.Handler
}

type ResponseHeader struct {
	handler     http.Handler
	headerName  string
	headerValue string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	addr := os.Getenv("ADDR")

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/hello", HelloHandler)
	mux.HandleFunc("/v1/time", CurrentTimeHandler)

	wrappedMux := NewLogger(NewResponseHeader(mux, "X-My-Header", "my header value"))

	log.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, wrappedMux))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func CurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("current time is %v", currentTime)))
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

// NewResponseHeader constructs a new ResponseHeader middleware handler
func NewResponseHeader(handlerToWrap http.Handler, headerName string, headerValue string) *ResponseHeader {
	return &ResponseHeader{handlerToWrap, headerName, headerValue}
}

// ServeHTTP handles the request by adding the response header
func (rh *ResponseHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//add the header
	w.Header().Add(rh.headerName, rh.headerValue)
	//call the wrapped handler
	rh.handler.ServeHTTP(w, r)
}
