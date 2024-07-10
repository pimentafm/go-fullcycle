package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("handler started")
	defer log.Println("handler ended")

	select {
	case <-time.After(5 * time.Second):
		log.Println("Request processed successfully")
		w.Write([]byte("Request processed successfully"))
	case <-ctx.Done():
		log.Println("request cancelled")
	}
}
