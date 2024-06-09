package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	certFile := os.Getenv("CERT_FILE")
	keyFile := os.Getenv("KEY_FILE")

	if certFile == "" || keyFile == "" {
		fmt.Println("SERVER_CERT_FILE and SERVER_KEY_FILE must be set")
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})

	err := http.ListenAndServeTLS(":8443", certFile, keyFile, nil)
	if err != nil {
		panic(err)
	}
}
