package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	const addr = "localhost:8080"

	fmt.Print("LOADING SERVER... ")

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte("{}"))
	})

	fmt.Print("OK\nBUILDING SERVER... ")

	srv := http.Server{
		Handler:      serveMux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("OK\nRUNNING SERVER. ")
	srv.ListenAndServe()

	fmt.Println("SERVER DIED.")
}
